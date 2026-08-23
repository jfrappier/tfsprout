package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Schema{ // want "schema of TypeList or TypeSet should not include top level ValidateFunc or ValidateDiagFunc"
		Type:         s.TypeList,
		ValidateFunc: validateFunc_v2,
	}

	_ = s.Schema{ // want "schema of TypeList or TypeSet should not include top level ValidateFunc or ValidateDiagFunc"
		Type:         s.TypeSet,
		ValidateFunc: validateFunc_v2,
	}

	_ = s.Schema{ // want "schema of TypeList or TypeSet should not include top level ValidateFunc or ValidateDiagFunc"
		Type:             s.TypeList,
		ValidateDiagFunc: validateDiagFunc_v2,
	}

	_ = s.Schema{ // want "schema of TypeList or TypeSet should not include top level ValidateFunc or ValidateDiagFunc"
		Type:             s.TypeSet,
		ValidateDiagFunc: validateDiagFunc_v2,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema of TypeList or TypeSet should not include top level ValidateFunc or ValidateDiagFunc"
			Type:         s.TypeList,
			ValidateFunc: validateFunc_v2,
		},
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema of TypeList or TypeSet should not include top level ValidateFunc or ValidateDiagFunc"
			Type:         s.TypeSet,
			ValidateFunc: validateFunc_v2,
		},
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema of TypeList or TypeSet should not include top level ValidateFunc or ValidateDiagFunc"
			Type:             s.TypeList,
			ValidateDiagFunc: validateDiagFunc_v2,
		},
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema of TypeList or TypeSet should not include top level ValidateFunc or ValidateDiagFunc"
			Type:             s.TypeSet,
			ValidateDiagFunc: validateDiagFunc_v2,
		},
	}
}
