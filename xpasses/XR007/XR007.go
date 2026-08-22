package XR007

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/passes/stdlib/osexeccommandcallexpr"
	"github.com/jfrappier/tfsprout/passes/stdlib/osexeccommandselectorexpr"
)

var Analyzer = analysisutils.AvoidSelectorExprAnalyzer(
	"XR007",
	osexeccommandcallexpr.Analyzer,
	osexeccommandselectorexpr.Analyzer,
	"os/exec",
	"Command",
)
