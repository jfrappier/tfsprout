package R001_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jfrappier/tfsprout/passes/R001"
)

func TestR001(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, R001.Analyzer, "testdata/src/a")
}
