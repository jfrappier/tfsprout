package testcheckresourceattrcallexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/resource"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"testcheckresourceattrcallexpr",
	resource.IsFunc,
	resource.PackagePath,
	resource.FuncNameTestCheckResourceAttr,
)
