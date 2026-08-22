package cidrnetworkselectorexpr

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"cidrnetworkselectorexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameCIDRNetwork,
)
