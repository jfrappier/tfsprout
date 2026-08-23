package a

import (
	"testdata/src/a/schema"
)

func foutside() {
	_ = schema.Schema{
		Type:         schema.TypeList,
		ValidateFunc: func() {},
	}

	_ = schema.Schema{
		Type:         schema.TypeSet,
		ValidateFunc: func() {},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type:         schema.TypeList,
			ValidateFunc: func() {},
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type:         schema.TypeSet,
			ValidateFunc: func() {},
		},
	}
}

func foutside_diag() {
	_ = schema.Schema{
		Type:             schema.TypeList,
		ValidateDiagFunc: func() {},
	}

	_ = schema.Schema{
		Type:             schema.TypeSet,
		ValidateDiagFunc: func() {},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type:             schema.TypeList,
			ValidateDiagFunc: func() {},
		},
	}
}
