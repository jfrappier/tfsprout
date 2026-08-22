package timesleepcallexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
)

var Analyzer = analysisutils.StdlibFunctionCallExprAnalyzer(
	"timesleepcallexpr",
	"time",
	"Sleep",
)
