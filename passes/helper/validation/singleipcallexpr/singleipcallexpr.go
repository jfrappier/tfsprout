package singleipcallexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"singleipcallexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameSingleIP,
)
