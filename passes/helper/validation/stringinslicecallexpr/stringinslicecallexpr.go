package stringinslicecallexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"stringinslicecallexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameStringInSlice,
)
