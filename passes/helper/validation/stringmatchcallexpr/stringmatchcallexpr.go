package stringmatchcallexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"stringmatchcallexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameStringMatch,
)
