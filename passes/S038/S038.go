// Package S038 defines an Analyzer that checks for
// Schema with both ValidateFunc and ValidateDiagFunc configured
package S038

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/schema"
	"github.com/jfrappier/tfsprout/passes/commentignore"
	"github.com/jfrappier/tfsprout/passes/helper/schema/schemainfo"
)

const Doc = `check for Schema with both ValidateFunc and ValidateDiagFunc configured

The S038 analyzer reports cases of schemas which configure both ValidateFunc and
ValidateDiagFunc, which will fail provider schema validation. The two are mutually
exclusive; ValidateDiagFunc is the replacement for the deprecated ValidateFunc.`

const analyzerName = "S038"

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

		if !schemaInfo.DeclaresField(schema.SchemaFieldValidateFunc) || !schemaInfo.DeclaresField(schema.SchemaFieldValidateDiagFunc) {
			continue
		}

		switch t := schemaInfo.AstCompositeLit.Type.(type) {
		default:
			pass.Reportf(schemaInfo.AstCompositeLit.Lbrace, "%s: schema should not configure both ValidateFunc and ValidateDiagFunc", analyzerName)
		case *ast.SelectorExpr:
			pass.Reportf(t.Sel.Pos(), "%s: schema should not configure both ValidateFunc and ValidateDiagFunc", analyzerName)
		}
	}

	return nil, nil
}
