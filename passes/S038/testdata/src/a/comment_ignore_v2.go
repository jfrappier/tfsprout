package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S038
	_ = schema.Schema{
		Type:             schema.TypeString,
		ValidateFunc:     validateFunc_v2,
		ValidateDiagFunc: validateDiagFunc_v2,
	}
}
