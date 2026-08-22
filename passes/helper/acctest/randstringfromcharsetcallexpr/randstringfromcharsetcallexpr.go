package randstringfromcharsetcallexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/acctest"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"randstringfromcharsetcallexpr",
	acctest.IsFunc,
	acctest.PackagePath,
	acctest.FuncNameRandStringFromCharSet,
)
