package printlint_test

import (
	"testing"

	"github.com/markusthoemmes/printlint"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, printlint.Analyzer, "test")
}