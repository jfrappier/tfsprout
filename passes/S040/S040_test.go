package S040_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jfrappier/tfsprout/passes/S040"
)

func TestS040(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S040.Analyzer, "testdata/src/a")
}
