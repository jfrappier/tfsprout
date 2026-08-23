package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	_ = map[string]*schema.Schema{
		//lintignore:S039
		"id": {
			Type:              schema.TypeString,
			Computed:          true,
			RequiredForImport: true,
		},
	}
}
