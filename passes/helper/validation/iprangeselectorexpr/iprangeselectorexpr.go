package iprangeselectorexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"iprangeselectorexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameIPRange,
)
