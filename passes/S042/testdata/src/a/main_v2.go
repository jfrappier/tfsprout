package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Schema{
		Type:      schema.TypeList,
		Optional:  true,
		Elem:      &schema.Schema{Type: schema.TypeString},
		WriteOnly: true, // want "schema of TypeList, TypeMap, or TypeSet should not enable WriteOnly"
	}

	_ = schema.Schema{
		Type:      schema.TypeSet,
		Optional:  true,
		Elem:      &schema.Schema{Type: schema.TypeString},
		WriteOnly: true, // want "schema of TypeList, TypeMap, or TypeSet should not enable WriteOnly"
	}

	_ = schema.Schema{
		Type:      schema.TypeMap,
		Optional:  true,
		Elem:      &schema.Schema{Type: schema.TypeString},
		WriteOnly: true, // want "schema of TypeList, TypeMap, or TypeSet should not enable WriteOnly"
	}

	// WriteOnly is supported on primitives.
	_ = schema.Schema{
		Type:      schema.TypeString,
		Optional:  true,
		WriteOnly: true,
	}

	_ = schema.Schema{
		Type:      schema.TypeInt,
		Optional:  true,
		WriteOnly: true,
	}

	_ = schema.Schema{
		Type:      schema.TypeBool,
		Optional:  true,
		WriteOnly: true,
	}

	// Collections without WriteOnly are fine.
	_ = schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	}

	_ = map[string]*schema.Schema{
		"values": {
			Type:      schema.TypeList,
			Optional:  true,
			Elem:      &schema.Schema{Type: schema.TypeString},
			WriteOnly: true, // want "schema of TypeList, TypeMap, or TypeSet should not enable WriteOnly"
		},
	}
}
