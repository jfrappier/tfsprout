package XR005_test

import (
	"testing"

	"github.com/jfrappier/tfsprout/xpasses/XR005"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestXR005(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, XR005.Analyzer, "testdata/src/a")
}
