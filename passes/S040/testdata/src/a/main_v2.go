package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Schema{ // want "schema should not only enable Computed and configure ValidateDiagFunc"
		Computed:         true,
		Type:             schema.TypeString,
		ValidateDiagFunc: validateDiagFunc_v2,
	}

	// Optional with Computed is configurable, so validation is meaningful.
	_ = schema.Schema{
		Computed:         true,
		Optional:         true,
		Type:             schema.TypeString,
		ValidateDiagFunc: validateDiagFunc_v2,
	}

	_ = schema.Schema{
		Required:         true,
		Type:             schema.TypeString,
		ValidateDiagFunc: validateDiagFunc_v2,
	}

	// ValidateFunc on a computed-only schema is S010's concern, not S040's.
	_ = schema.Schema{
		Computed:     true,
		Type:         schema.TypeString,
		ValidateFunc: validateFunc_v2,
	}

	_ = schema.Schema{
		Computed: true,
		Type:     schema.TypeString,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not only enable Computed and configure ValidateDiagFunc"
			Computed:         true,
			Type:             schema.TypeString,
			ValidateDiagFunc: validateDiagFunc_v2,
		},
	}
}

func validateFunc_v2(v interface{}, k string) (ws []string, errors []error) {
	return ws, errors
}

var validateDiagFunc_v2 schema.SchemaValidateDiagFunc
