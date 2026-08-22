package stringdoesnotmatchcallexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"stringdoesnotmatchcallexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameStringDoesNotMatch,
)
