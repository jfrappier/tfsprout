package S028_test

import (
	"testing"

	"github.com/jfrappier/tfsprout/passes/S028"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS028(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S028.Analyzer, "testdata/src/a")
}
