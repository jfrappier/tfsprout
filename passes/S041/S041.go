// Package S041 defines an Analyzer that checks for
// Schema with WriteOnly and an incompatible field configured
package S041

import (
	"golang.org/x/tools/go/analysis"

	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/schema"
	"github.com/jfrappier/tfsprout/passes/commentignore"
	"github.com/jfrappier/tfsprout/passes/helper/schema/schemainfo"
)

const Doc = `check for Schema with WriteOnly and an incompatible field configured

The S041 analyzer reports cases of schemas which enable WriteOnly alongside
Computed, ForceNew, Default, or DefaultFunc, which will fail provider schema
validation.

A write-only attribute is never persisted to state, so there is nothing for
Terraform to compare against or default from.`

const analyzerName = "S041"

// writeOnlyIncompatibleFields are rejected by (schemaMap).internalValidate when
// WriteOnly is enabled, in the order reported.
var writeOnlyIncompatibleFields = []string{
	schema.SchemaFieldComputed,
	schema.SchemaFieldDefault,
	schema.SchemaFieldDefaultFunc,
	schema.SchemaFieldForceNew,
}

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

		for _, fieldName := range writeOnlyIncompatibleFields {
			if !schemaInfo.DeclaresField(fieldName) {
				continue
			}

			pass.Reportf(schemaInfo.Fields[fieldName].Pos(), "%s: schema should not configure %s with WriteOnly", analyzerName, fieldName)
		}
	}

	return nil, nil
}
