package XS001_test

import (
	"testing"

	"github.com/jfrappier/tfsprout/xpasses/XS001"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestXS001(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, XS001.Analyzer, "testdata/src/a")
}
