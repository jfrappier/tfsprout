package resourcedatapartialselectorexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/schema"
)

var Analyzer = analysisutils.ReceiverMethodSelectorExprAnalyzer(
	"resourcedatapartialselectorexpr",
	schema.IsReceiverMethod,
	schema.PackagePath,
	schema.TypeNameResourceData,
	"Partial",
)
