package validatelistuniquestringsselectorexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validatelistuniquestringsselectorexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameValidateListUniqueStrings,
)
