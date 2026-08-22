package XR008

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/passes/stdlib/osexeccommandcontextcallexpr"
	"github.com/jfrappier/tfsprout/passes/stdlib/osexeccommandcontextselectorexpr"
)

var Analyzer = analysisutils.AvoidSelectorExprAnalyzer(
	"XR008",
	osexeccommandcontextcallexpr.Analyzer,
	osexeccommandcontextselectorexpr.Analyzer,
	"os/exec",
	"CommandContext",
)
