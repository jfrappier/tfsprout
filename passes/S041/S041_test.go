package S041_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jfrappier/tfsprout/passes/S041"
)

func TestS041(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S041.Analyzer, "testdata/src/a")
}
