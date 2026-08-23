package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Schema{ // want "schema should not configure both ValidateFunc and ValidateDiagFunc"
		Type:             s.TypeString,
		ValidateFunc:     validateFunc_v2,
		ValidateDiagFunc: validateDiagFunc_v2,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema should not configure both ValidateFunc and ValidateDiagFunc"
			Type:             s.TypeString,
			ValidateFunc:     validateFunc_v2,
			ValidateDiagFunc: validateDiagFunc_v2,
		},
	}
}
