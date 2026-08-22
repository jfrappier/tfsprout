package validatejsonstringselectorexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validatejsonstringselectorexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameValidateJsonString,
)
