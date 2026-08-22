package resourceproviderfactoryselectorexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/terraform"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"resourceproviderfactoryselectorexpr",
	terraform.IsFunc,
	terraform.PackagePath,
	terraform.TypeNameResourceProviderFactory,
)
