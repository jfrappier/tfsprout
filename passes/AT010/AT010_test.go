package AT010_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jfrappier/tfsprout/passes/AT010"
)

func TestAT010(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, AT010.Analyzer, "testdata/src/a")
}
