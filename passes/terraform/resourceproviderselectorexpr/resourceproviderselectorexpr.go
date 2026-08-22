package resourceproviderselectorexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/terraform"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"resourceproviderselectorexpr",
	terraform.IsFunc,
	terraform.PackagePath,
	terraform.TypeNameResourceProvider,
)
