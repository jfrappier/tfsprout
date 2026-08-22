package XR003_test

import (
	"testing"

	"github.com/jfrappier/tfsprout/xpasses/XR003"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestXR003(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, XR003.Analyzer, "testdata/src/a")
}
