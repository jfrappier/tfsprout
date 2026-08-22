package S032_test

import (
	"testing"

	"github.com/jfrappier/tfsprout/passes/S032"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS032(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S032.Analyzer, "testdata/src/a")
}
