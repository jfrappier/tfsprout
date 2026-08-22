package S003_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jfrappier/tfsprout/passes/S003"
)

func TestS003(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S003.Analyzer, "testdata/src/a")
}
