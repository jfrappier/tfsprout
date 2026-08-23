// Package S043 defines an Analyzer that checks for
// Schema of TypeSet or Computed block containing WriteOnly attributes
package S043

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/jfrappier/tfsprout/helper/astutils"
	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/schema"
	"github.com/jfrappier/tfsprout/passes/commentignore"
	"github.com/jfrappier/tfsprout/passes/helper/schema/schemainfo"
)

const Doc = `check for Schema of TypeSet or Computed block containing WriteOnly attributes

The S043 analyzer reports configuration blocks that contain a WriteOnly
attribute at any depth while the block itself is either a TypeSet or Computed,
which will fail provider schema validation.

A set's element identity is derived from its values, and a write-only value is
never persisted, so a set containing one cannot be matched against prior state.
A Computed block is likewise never configured by the practitioner.`

const analyzerName = "S043"

// maxElemDepth bounds the Elem walk. Nested blocks in real providers are only a
// few levels deep, and a bound removes any chance of pathological recursion on
// a deeply nested literal.
const maxElemDepth = 10

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

		isSet := schemaInfo.IsType(schema.SchemaValueTypeSet)
		isComputed := schemaInfo.Schema.Computed

		if !isSet && !isComputed {
			continue
		}

		elemKvExpr := schemaInfo.Fields[schema.SchemaFieldElem]

		if elemKvExpr == nil {
			continue
		}

		if !elemResourceHasWriteOnly(pass, elemKvExpr.Value, 0) {
			continue
		}

		if isSet {
			pass.Reportf(elemKvExpr.Value.Pos(), "%s: schema of TypeSet should not contain WriteOnly attributes", analyzerName)
			continue
		}

		pass.Reportf(elemKvExpr.Value.Pos(), "%s: Computed schema should not contain WriteOnly attributes", analyzerName)
	}

	return nil, nil
}

// elemResourceHasWriteOnly reports whether an Elem expression holding a
// *schema.Resource declares a WriteOnly attribute at any depth, mirroring
// (schemaMap).hasWriteOnly in the Terraform Plugin SDK.
func elemResourceHasWriteOnly(pass *analysis.Pass, elem ast.Expr, depth int) bool {
	if depth > maxElemDepth {
		return false
	}

	if !schema.IsTypeResource(pass.TypesInfo.TypeOf(elem)) {
		return false
	}

	compositeLit := elemCompositeLit(elem)

	if compositeLit == nil {
		return false
	}

	schemaKvExpr := astutils.CompositeLitField(compositeLit, schema.ResourceFieldSchema)

	if schemaKvExpr == nil {
		return false
	}

	schemaMap, ok := schemaKvExpr.Value.(*ast.CompositeLit)

	if !ok {
		return false
	}

	for _, attribute := range schema.GetSchemaMapSchemas(schemaMap) {
		attributeInfo := schema.NewSchemaInfo(attribute, pass.TypesInfo)

		if attributeInfo.DeclaresField(schema.SchemaFieldWriteOnly) {
			return true
		}

		if nestedElem := attributeInfo.Fields[schema.SchemaFieldElem]; nestedElem != nil {
			if elemResourceHasWriteOnly(pass, nestedElem.Value, depth+1) {
				return true
			}
		}
	}

	return false
}

// elemCompositeLit unwraps the &schema.Resource{...} form that Elem almost
// always takes, as well as a bare schema.Resource{...}.
func elemCompositeLit(elem ast.Expr) *ast.CompositeLit {
	switch v := elem.(type) {
	case *ast.UnaryExpr:
		if compositeLit, ok := v.X.(*ast.CompositeLit); ok {
			return compositeLit
		}
	case *ast.CompositeLit:
		return v
	}

	return nil
}
