package fmterrorfcallexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
)

var Analyzer = analysisutils.StdlibFunctionCallExprAnalyzer(
	"fmterrorfcallexpr",
	"fmt",
	"Errorf",
)
