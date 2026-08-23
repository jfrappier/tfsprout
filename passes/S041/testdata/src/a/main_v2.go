package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Schema{
		Type:      schema.TypeString,
		Optional:  true,
		Computed:  true, // want "schema should not configure Computed with WriteOnly"
		WriteOnly: true,
	}

	_ = schema.Schema{
		Type:      schema.TypeString,
		Optional:  true,
		ForceNew:  true, // want "schema should not configure ForceNew with WriteOnly"
		WriteOnly: true,
	}

	_ = schema.Schema{
		Type:      schema.TypeString,
		Optional:  true,
		Default:   "value", // want "schema should not configure Default with WriteOnly"
		WriteOnly: true,
	}

	_ = schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		DefaultFunc: defaultFunc_v2, // want "schema should not configure DefaultFunc with WriteOnly"
		WriteOnly:   true,
	}

	// Reports once per offending field.
	_ = schema.Schema{
		Type:      schema.TypeString,
		Optional:  true,
		Computed:  true,    // want "schema should not configure Computed with WriteOnly"
		Default:   "value", // want "schema should not configure Default with WriteOnly"
		WriteOnly: true,
	}

	// Valid write-only attributes.
	_ = schema.Schema{
		Type:      schema.TypeString,
		Optional:  true,
		WriteOnly: true,
	}

	_ = schema.Schema{
		Type:      schema.TypeString,
		Required:  true,
		WriteOnly: true,
	}

	// The same fields without WriteOnly are other checks' concern.
	_ = schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		ForceNew: true,
		Default:  "value",
	}

	_ = map[string]*schema.Schema{
		"token": {
			Type:      schema.TypeString,
			Optional:  true,
			ForceNew:  true, // want "schema should not configure ForceNew with WriteOnly"
			WriteOnly: true,
		},
	}
}

func defaultFunc_v2() (interface{}, error) {
	return nil, nil
}
