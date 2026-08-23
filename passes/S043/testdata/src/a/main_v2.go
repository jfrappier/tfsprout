package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	// A set block containing a write-only attribute.
	_ = schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{ // want "schema of TypeSet should not contain WriteOnly attributes"
			Schema: map[string]*schema.Schema{
				"token": {
					Type:      schema.TypeString,
					Optional:  true,
					WriteOnly: true,
				},
			},
		},
	}

	// A computed block containing a write-only attribute.
	_ = schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{ // want "Computed schema should not contain WriteOnly attributes"
			Schema: map[string]*schema.Schema{
				"token": {
					Type:      schema.TypeString,
					Optional:  true,
					WriteOnly: true,
				},
			},
		},
	}

	// Nested one level deeper.
	_ = schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{ // want "schema of TypeSet should not contain WriteOnly attributes"
			Schema: map[string]*schema.Schema{
				"inner": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"token": {
								Type:      schema.TypeString,
								Optional:  true,
								WriteOnly: true,
							},
						},
					},
				},
			},
		},
	}

	// A set block with no write-only attribute anywhere.
	_ = schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	}

	// A list block that is neither a set nor computed may contain write-only
	// attributes.
	_ = schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"token": {
					Type:      schema.TypeString,
					Optional:  true,
					WriteOnly: true,
				},
			},
		},
	}

	// Elem of *schema.Schema is not a block and cannot carry attributes.
	_ = schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	}
}
