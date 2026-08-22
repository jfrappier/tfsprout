package S020_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jfrappier/tfsprout/passes/S020"
)

func TestS020(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S020.Analyzer, "testdata/src/a")
}
