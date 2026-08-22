package R003_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jfrappier/tfsprout/passes/R003"
)

func TestR003(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, R003.Analyzer, "testdata/src/a")
}
