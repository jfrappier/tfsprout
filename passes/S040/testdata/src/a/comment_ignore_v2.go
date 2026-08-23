package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S040
	_ = schema.Schema{
		Computed:         true,
		Type:             schema.TypeString,
		ValidateDiagFunc: validateDiagFunc_v2,
	}
}
