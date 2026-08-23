package a

import (
	"testdata/src/a/schema"
)

func foutside() {
	_ = schema.Schema{
		Type:             schema.TypeString,
		ValidateFunc:     func() {},
		ValidateDiagFunc: func() {},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type:             schema.TypeString,
			ValidateFunc:     func() {},
			ValidateDiagFunc: func() {},
		},
	}
}
