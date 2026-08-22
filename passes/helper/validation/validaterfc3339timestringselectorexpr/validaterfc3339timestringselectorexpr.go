package validaterfc3339timestringselectorexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validaterfc3339timestringselectorexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameValidateRFC3339TimeString,
)
