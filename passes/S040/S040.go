// Package S040 defines an Analyzer that checks for
// Schema with only Computed enabled and ValidateDiagFunc configured
package S040

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/schema"
	"github.com/jfrappier/tfsprout/passes/commentignore"
	"github.com/jfrappier/tfsprout/passes/helper/schema/schemainfocomputedonly"
)

const Doc = `check for Schema with only Computed enabled and ValidateDiagFunc configured

The S040 analyzer reports cases of schemas which only enables Computed
and configures ValidateDiagFunc, which will fail provider schema validation.
There is no practitioner input to validate on a computed-only attribute.

This is the ValidateDiagFunc counterpart to S010.`

const analyzerName = "S040"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		commentignore.Analyzer,
		schemainfocomputedonly.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemaInfos := pass.ResultOf[schemainfocomputedonly.Analyzer].([]*schema.SchemaInfo)
	for _, schemaInfo := range schemaInfos {
		if ignorer.ShouldIgnore(analyzerName, schemaInfo.AstCompositeLit) {
			continue
		}

		if !schemaInfo.DeclaresField(schema.SchemaFieldValidateDiagFunc) {
			continue
		}

		switch t := schemaInfo.AstCompositeLit.Type.(type) {
		default:
			pass.Reportf(schemaInfo.AstCompositeLit.Lbrace, "%s: schema should not only enable Computed and configure ValidateDiagFunc", analyzerName)
		case *ast.SelectorExpr:
			pass.Reportf(t.Sel.Pos(), "%s: schema should not only enable Computed and configure ValidateDiagFunc", analyzerName)
		}
	}

	return nil, nil
}
