# S013

The S013 analyzer reports cases of schemas which one of `Computed`,
`Optional`, or `Required` is not configured, which will fail provider
schema validation.

## Flagged Code

```go
map[string]*schema.Schema{
    "attribute_name": {
        Type: schema.TypeString,
    },
}
```

## Passing Code

```go
map[string]*schema.Schema{
    "attribute_name": {
        Computed: true,
        Type:     schema.TypeString,
    },
}

# OR

map[string]*schema.Schema{
    "attribute_name": {
        Optional: true,
        Type:     schema.TypeString,
    },
}

# OR

map[string]*schema.Schema{
    "attribute_name": {
        Required: true,
        Type:     schema.TypeString,
    },
}
```

## Resource Identity Schemas

Resource identity schemas are declared as an ordinary `map[string]*schema.Schema`
but configure `RequiredForImport` or `OptionalForImport` in place of `Computed`,
`Optional`, or `Required`. S013 skips any attribute declaring either field, so
identity schemas are not reported:

```go
&schema.ResourceIdentity{
    Version: 0,
    SchemaFunc: func() map[string]*schema.Schema {
        return map[string]*schema.Schema{
            "id": {
                Type:              schema.TypeString,
                RequiredForImport: true,
            },
        }
    },
}
```

## Ignoring Reports

Singular reports can be ignored by adding a `//lintignore:S013` Go code comment at the end of the offending line or on the line immediately preceding, e.g.

```go
//lintignore:S013
map[string]*schema.Schema{
    "attribute_name": {
        Type: schema.TypeString,
    },
}
```
