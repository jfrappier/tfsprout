package resourcedatagetchangeassignstmt

import (
	"github.com/jfrappier/tfsprout/helper/analysisutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/schema"
)

var Analyzer = analysisutils.ReceiverMethodAssignStmtAnalyzer(
	"resourcedatagetchangeassignstmt",
	schema.IsReceiverMethod,
	schema.PackagePath,
	schema.TypeNameResourceData,
	"GetChange",
)
