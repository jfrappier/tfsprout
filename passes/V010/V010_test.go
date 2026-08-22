package V010_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jfrappier/tfsprout/passes/V010"
)

func TestV010(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V010.Analyzer, "testdata/src/a")
}
