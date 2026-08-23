package S039_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jfrappier/tfsprout/passes/S039"
)

func TestS039(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S039.Analyzer, "testdata/src/a")
}
