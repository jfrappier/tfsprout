package osexeccommandcallexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
)

var Analyzer = analysisutils.StdlibFunctionCallExprAnalyzer(
	"osexeccommandcallexpr",
	"os/exec",
	"Command",
)
