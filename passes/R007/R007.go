package R007

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/schema"
	"github.com/jfrappier/tfsprout/passes/helper/schema/resourcedatapartialcallexpr"
	"github.com/jfrappier/tfsprout/passes/helper/schema/resourcedatapartialselectorexpr"
)

var Analyzer = analysisutils.DeprecatedReceiverMethodSelectorExprAnalyzer(
	"R007",
	resourcedatapartialcallexpr.Analyzer,
	resourcedatapartialselectorexpr.Analyzer,
	schema.PackagePath,
	schema.TypeNameResourceData,
	"Partial",
)
