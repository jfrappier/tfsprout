package osexeccommandcontextcallexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
)

var Analyzer = analysisutils.StdlibFunctionCallExprAnalyzer(
	"osexeccommandcontextcallexpr",
	"os/exec",
	"CommandContext",
)
