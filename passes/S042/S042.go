// Package S042 defines an Analyzer that checks for
// Schema of TypeList, TypeMap, or TypeSet with WriteOnly enabled
package S042

import (
	"golang.org/x/tools/go/analysis"

	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/schema"
	"github.com/jfrappier/tfsprout/passes/commentignore"
	"github.com/jfrappier/tfsprout/passes/helper/schema/schemainfo"
)

const Doc = `check for Schema of TypeList, TypeMap, or TypeSet with WriteOnly enabled

The S042 analyzer reports cases of TypeList, TypeMap, or TypeSet schemas which
enable WriteOnly, which will fail provider schema validation. WriteOnly is only
supported on primitive types.`

const analyzerName = "S042"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		schemainfo.Analyzer,
		commentignore.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemaInfos := pass.ResultOf[schemainfo.Analyzer].([]*schema.SchemaInfo)

	for _, schemaInfo := range schemaInfos {
		if ignorer.ShouldIgnore(analyzerName, schemaInfo.AstCompositeLit) {
			continue
		}

		if !schemaInfo.DeclaresField(schema.SchemaFieldWriteOnly) {
			continue
		}

		if !schemaInfo.IsOneOfTypes(schema.SchemaValueTypeList, schema.SchemaValueTypeMap, schema.SchemaValueTypeSet) {
			continue
		}

		pass.Reportf(schemaInfo.Fields[schema.SchemaFieldWriteOnly].Pos(), "%s: schema of TypeList, TypeMap, or TypeSet should not enable WriteOnly", analyzerName)
	}

	return nil, nil
}
