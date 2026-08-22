package S007_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jfrappier/tfsprout/passes/S007"
)

func TestS007(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S007.Analyzer, "testdata/src/a")
}
