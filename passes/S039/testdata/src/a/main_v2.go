package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// A correctly configured identity schema reports nothing.

func fvalid_v2() {
	_ = map[string]*schema.Schema{
		"id": {
			Type:              schema.TypeString,
			Description:       "The id of the resource",
			RequiredForImport: true,
		},
		"zone": {
			Type:              schema.TypeString,
			OptionalForImport: true,
		},
	}

	_ = map[string]*schema.Schema{
		"names": {
			Type:              schema.TypeList,
			Elem:              &schema.Schema{Type: schema.TypeString},
			OptionalForImport: true,
		},
	}
}

// Ordinary resource attributes declare neither import field and are not the
// concern of this check, however they are configured.

func fnonIdentity_v2() {
	_ = map[string]*schema.Schema{
		"name": {
			Type:         schema.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validateFunc_v2,
		},
	}
}

// Identity attributes may not configure resource-only fields. This is also how
// the inverse mistake surfaces: an import field set on an ordinary resource
// attribute, which still carries Computed, Optional, or Required.

func finvalidFields_v2() {
	_ = map[string]*schema.Schema{
		"id": {
			Type:              schema.TypeString,
			Computed:          true, // want "schema should not configure Computed with RequiredForImport or OptionalForImport"
			RequiredForImport: true,
		},
	}

	_ = map[string]*schema.Schema{
		"id": {
			Type:              schema.TypeString,
			Optional:          true, // want "schema should not configure Optional with RequiredForImport or OptionalForImport"
			RequiredForImport: true,
		},
	}

	_ = map[string]*schema.Schema{
		"id": {
			Type:              schema.TypeString,
			ForceNew:          true, // want "schema should not configure ForceNew with RequiredForImport or OptionalForImport"
			RequiredForImport: true,
			Sensitive:         true, // want "schema should not configure Sensitive with RequiredForImport or OptionalForImport"
		},
	}

	_ = map[string]*schema.Schema{
		"id": {
			Type:              schema.TypeString,
			RequiredForImport: true,
			ValidateFunc:      validateFunc_v2, // want "schema should not configure ValidateFunc with RequiredForImport or OptionalForImport"
		},
	}
}

// WriteOnly is an SDK field tfsprout does not model. The allowlist means it is
// still reported.

func finvalidUnmodelledField_v2() {
	_ = map[string]*schema.Schema{
		"id": {
			Type:              schema.TypeString,
			RequiredForImport: true,
			WriteOnly:         true, // want "schema should not configure WriteOnly with RequiredForImport or OptionalForImport"
		},
	}
}

func fbothImportFields_v2() {
	_ = map[string]*schema.Schema{
		"id": {
			Type:              schema.TypeString,
			OptionalForImport: true, // want "schema should configure only one of RequiredForImport or OptionalForImport"
			RequiredForImport: true,
		},
	}
}

func finvalidTypes_v2() {
	_ = map[string]*schema.Schema{
		"tags": {
			Type:              schema.TypeMap, // want "schema should not configure TypeMap or TypeSet for resource identity"
			OptionalForImport: true,
		},
	}

	_ = map[string]*schema.Schema{
		"tags": {
			Type:              schema.TypeSet, // want "schema should not configure TypeMap or TypeSet for resource identity"
			OptionalForImport: true,
		},
	}
}

func validateFunc_v2(v interface{}, k string) (ws []string, errors []error) {
	return ws, errors
}
