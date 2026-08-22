package XR001_test

import (
	"testing"

	"github.com/jfrappier/tfsprout/xpasses/XR001"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestXR001(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, XR001.Analyzer, "testdata/src/a")
}
