package printlint

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

const doc = "printlint does stuff"

// Analyzer analyzes Go test codes whether they use
var Analyzer = &analysis.Analyzer{
	Name: "tparallel",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	return nil, nil
}
