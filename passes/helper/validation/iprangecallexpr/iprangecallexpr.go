package iprangecallexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"iprangecallexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameIPRange,
)
