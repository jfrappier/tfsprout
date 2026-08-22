package S008_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jfrappier/tfsprout/passes/S008"
)

func TestS008(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S008.Analyzer, "testdata/src/a")
}
