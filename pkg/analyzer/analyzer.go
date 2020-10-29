package analyzer

import (
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

// Analyzer analyzes Go test codes whether they use
var Analyzer = &analysis.Analyzer{
	Name: "printlint",
	Doc:  "printlint does stuff",
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

var targets = map[string]struct{}{
	"Sprintf": struct{}{},
}

func run(pass *analysis.Pass) (interface{}, error) {
	pssa, ok := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	if !ok {
		return nil, nil
	}

	for _, f := range pssa.SrcFuncs {
		for _, b := range f.Blocks {
			for _, i := range b.Instrs {
				// Skip everything that's not a function call.
				call, ok := i.(*ssa.Call)
				if !ok {
					continue
				}

				// Skip everything that's not in our allowlist.
				function := call.Call.Value.Name()
				if _, ok := targets[function]; !ok {
					continue
				}

				// Skip everything that has less than 2 args.
				if len(call.Call.Args) < 2 {
					continue
				}

				// Skip everything that doesn't have a string as the first argument
				// (probably broken anyway).
				if call.Call.Args[0].Type().String() != "string" {
					continue
				}

				// Skip everything that hasn't a constant string as a first argument.
				format, ok := call.Call.Args[0].(*ssa.Const)
				if !ok {
					continue
				}
				formatValue := strings.Trim(format.Value.ExactString(), `"`)
				formats := strings.Count(formatValue, "%")

				// Skip everything that hasn't exactly 1 format call.
				if formats != 1 {
					continue
				}

				if strings.HasSuffix(formatValue, "%d") {
					pass.Reportf(call.Pos(), "avoid tail format")
				}
			}
		}
	}
	return nil, nil
}
