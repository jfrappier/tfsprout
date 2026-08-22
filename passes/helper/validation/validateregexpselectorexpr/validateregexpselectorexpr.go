package validateregexpselectorexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validateregexpselectorexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameValidateRegexp,
)
