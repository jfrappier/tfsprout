package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Schema{ // want "schema should not configure both ValidateFunc and ValidateDiagFunc"
		Type:             schema.TypeString,
		ValidateFunc:     validateFunc_v2,
		ValidateDiagFunc: validateDiagFunc_v2,
	}

	_ = schema.Schema{
		Type:         schema.TypeString,
		ValidateFunc: validateFunc_v2,
	}

	_ = schema.Schema{
		Type:             schema.TypeString,
		ValidateDiagFunc: validateDiagFunc_v2,
	}

	_ = schema.Schema{
		Type: schema.TypeString,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not configure both ValidateFunc and ValidateDiagFunc"
			Type:             schema.TypeString,
			ValidateFunc:     validateFunc_v2,
			ValidateDiagFunc: validateDiagFunc_v2,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type:             schema.TypeString,
			ValidateDiagFunc: validateDiagFunc_v2,
		},
	}
}

func validateFunc_v2(v interface{}, k string) (ws []string, errors []error) {
	return ws, errors
}

var validateDiagFunc_v2 schema.SchemaValidateDiagFunc
