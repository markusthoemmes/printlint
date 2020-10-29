package analyzer_test

import (
	"testing"

	"github.com/markusthoemmes/printlint/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer.Analyzer, "test")
}
