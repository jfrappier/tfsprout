package singleipselectorexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"singleipselectorexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameSingleIP,
)
