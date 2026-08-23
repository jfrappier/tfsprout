package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S042
	_ = schema.Schema{
		Type:      schema.TypeList,
		Optional:  true,
		Elem:      &schema.Schema{Type: schema.TypeString},
		WriteOnly: true,
	}
}
