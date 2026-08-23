package S042_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jfrappier/tfsprout/passes/S042"
)

func TestS042(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S042.Analyzer, "testdata/src/a")
}
