package V013_test

import (
	"testing"

	"github.com/jfrappier/tfsprout/passes/V013"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestV013(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V013.Analyzer, "testdata/src/a")
}
