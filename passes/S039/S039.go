// Package S039 defines an Analyzer that checks for
// Schema with invalid resource identity configuration
package S039

import (
	"sort"

	"golang.org/x/tools/go/analysis"

	"github.com/jfrappier/tfsprout/helper/terraformtype/helper/schema"
	"github.com/jfrappier/tfsprout/passes/commentignore"
	"github.com/jfrappier/tfsprout/passes/helper/schema/schemainfo"
)

const Doc = `check for Schema with invalid resource identity configuration

The S039 analyzer reports resource identity schema attributes that configure
fields the identity schema does not support, which will fail provider schema
validation. An identity attribute may only configure Type, Description, Elem,
and exactly one of RequiredForImport or OptionalForImport.

A schema declaring RequiredForImport or OptionalForImport is an identity
attribute; the Terraform Plugin SDK rejects both fields anywhere else. Reports
therefore also cover the inverse mistake of setting an import field on an
ordinary resource or data source attribute, which surfaces as that attribute
configuring resource-only fields alongside it.`

const analyzerName = "S039"

// identityAllowedFields are the only fields an identity schema attribute may
// configure. Everything else is rejected by (*schema.ResourceIdentity).
// InternalIdentityValidate. This is an allowlist so that SDK fields tfsprout
// does not model, such as WriteOnly and DiffSuppressOnRefresh, are still
// reported.
var identityAllowedFields = map[string]struct{}{
	schema.SchemaFieldDescription:       {},
	schema.SchemaFieldElem:              {},
	schema.SchemaFieldOptionalForImport: {},
	schema.SchemaFieldRequiredForImport: {},
	schema.SchemaFieldType:              {},
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

		if !schemaInfo.IsResourceIdentitySchema() {
			continue
		}

		if schemaInfo.DeclaresField(schema.SchemaFieldRequiredForImport) && schemaInfo.DeclaresField(schema.SchemaFieldOptionalForImport) {
			pass.Reportf(schemaInfo.Fields[schema.SchemaFieldOptionalForImport].Pos(), "%s: schema should configure only one of RequiredForImport or OptionalForImport", analyzerName)
		}

		if schemaInfo.IsOneOfTypes(schema.SchemaValueTypeMap, schema.SchemaValueTypeSet) {
			pass.Reportf(schemaInfo.Fields[schema.SchemaFieldType].Pos(), "%s: schema should not configure TypeMap or TypeSet for resource identity", analyzerName)
		}

		reportDisallowedFields(pass, schemaInfo)
	}

	return nil, nil
}

// reportDisallowedFields reports every field on an identity schema attribute
// that the identity schema does not support.
func reportDisallowedFields(pass *analysis.Pass, schemaInfo *schema.SchemaInfo) {
	// Map iteration order is not stable, so report in a fixed order to keep
	// output deterministic across runs.
	fieldNames := make([]string, 0, len(schemaInfo.Fields))

	for fieldName := range schemaInfo.Fields {
		if _, ok := identityAllowedFields[fieldName]; ok {
			continue
		}

		fieldNames = append(fieldNames, fieldName)
	}

	sort.Strings(fieldNames)

	for _, fieldName := range fieldNames {
		pass.Reportf(schemaInfo.Fields[fieldName].Pos(), "%s: schema should not configure %s with RequiredForImport or OptionalForImport", analyzerName, fieldName)
	}
}
