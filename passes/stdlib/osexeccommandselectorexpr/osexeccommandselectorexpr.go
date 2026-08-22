package osexeccommandselectorexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
)

var Analyzer = analysisutils.StdlibFunctionSelectorExprAnalyzer(
	"osexeccommandselectorexpr",
	"os/exec",
	"Command",
)
