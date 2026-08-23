package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S043
	_ = schema.Schema{
		Type:     schema.TypeSet,
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
}
