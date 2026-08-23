# S039

The S039 analyzer reports resource identity schema attributes that configure
fields the identity schema does not support, which will fail provider schema
validation.

An identity attribute may only configure `Type`, `Description`, `Elem`, and
exactly one of `RequiredForImport` or `OptionalForImport`. Its `Type` may not be
`TypeMap` or `TypeSet`.

A schema declaring `RequiredForImport` or `OptionalForImport` is an identity
attribute — the Terraform Plugin SDK rejects both fields anywhere else. Reports
therefore also cover the inverse mistake of setting an import field on an
ordinary resource or data source attribute, which surfaces as that attribute
configuring resource-only fields alongside it.

This check is the counterpart to [S013](../S013/README.md), which skips identity
schemas precisely so that S039 can validate them on their own terms.

## Flagged Code

```go
map[string]*schema.Schema{
    "id": {
        Type:              schema.TypeString,
        Computed:          true,
        RequiredForImport: true,
    },
}

map[string]*schema.Schema{
    "id": {
        Type:              schema.TypeString,
        OptionalForImport: true,
        RequiredForImport: true,
    },
}

map[string]*schema.Schema{
    "tags": {
        Type:              schema.TypeMap,
        OptionalForImport: true,
    },
}
```

## Passing Code

```go
map[string]*schema.Schema{
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
```

## Ignoring Reports

Singular reports can be ignored by adding a `//lintignore:S039` Go code comment at the end of the offending line or on the line immediately preceding, e.g.

```go
map[string]*schema.Schema{
    //lintignore:S039
    "id": {
        Type:              schema.TypeString,
        Computed:          true,
        RequiredForImport: true,
    },
}
```
