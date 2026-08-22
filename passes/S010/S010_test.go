package S010_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jfrappier/tfsprout/passes/S010"
)

func TestS010(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S010.Analyzer, "testdata/src/a")
}
