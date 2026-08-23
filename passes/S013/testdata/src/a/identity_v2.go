package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Identity schema attributes configure RequiredForImport or OptionalForImport
// instead of Computed, Optional, or Required, and must not be reported.

func fidentityInline_v2() {
	_ = schema.Resource{
		Identity: &schema.ResourceIdentity{
			Version: 0,
			SchemaFunc: func() map[string]*schema.Schema {
				return map[string]*schema.Schema{
					"id": {
						Type:              schema.TypeString,
						Description:       "The id of the resource",
						RequiredForImport: true,
					},
					"zone": {
						Type:              schema.TypeString,
						OptionalForImport: true,
					},
				}
			},
		},
	}
}

// Providers commonly build identity schemas through a helper that returns
// *schema.ResourceIdentity, which puts the map literal in an argument position
// where the enclosing ResourceIdentity is not visible from the AST. The
// RequiredForImport/OptionalForImport fields are the discriminator in both
// shapes.

func wrapSchemaMap_v2(m map[string]*schema.Schema) *schema.ResourceIdentity {
	return &schema.ResourceIdentity{
		Version:    0,
		SchemaFunc: func() map[string]*schema.Schema { return m },
	}
}

func fidentityWrapped_v2() {
	_ = wrapSchemaMap_v2(map[string]*schema.Schema{
		"id": {
			Type:              schema.TypeString,
			RequiredForImport: true,
		},
	})
}

// A schema map that configures none of the identity import fields is still an
// ordinary schema and must report as before.

func fidentityNegative_v2() {
	_ = map[string]*schema.Schema{
		"name": { // want "schema should configure one of Computed, Optional, or Required"
			Type:        schema.TypeString,
			Description: "not an identity attribute",
		},
	}
}
