package S043_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jfrappier/tfsprout/passes/S043"
)

func TestS043(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S043.Analyzer, "testdata/src/a")
}
