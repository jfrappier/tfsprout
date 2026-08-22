package fmtsprintfcallexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
)

var Analyzer = analysisutils.StdlibFunctionCallExprAnalyzer(
	"fmtsprintfcallexpr",
	"fmt",
	"Sprintf",
)
