package S014_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jfrappier/tfsprout/passes/S014"
)

func TestS014(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S014.Analyzer, "testdata/src/a")
}
