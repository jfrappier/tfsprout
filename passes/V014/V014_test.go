package V014_test

import (
	"testing"

	"github.com/jfrappier/tfsprout/passes/V014"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestV014(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V014.Analyzer, "testdata/src/a")
}
