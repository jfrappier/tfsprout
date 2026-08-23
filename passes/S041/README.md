# S041

The S041 analyzer reports cases of schemas which enable `WriteOnly` alongside
`Computed`, `ForceNew`, `Default`, or `DefaultFunc`, which will fail provider
schema validation.

A write-only attribute is never persisted to state, so there is nothing for
Terraform to compare against or default from. One report is emitted per
offending field.

`WriteOnly` requires `Optional` or `Required`. That constraint is enforced by
`S013` together with this check: a schema with none of `Computed`, `Optional`,
or `Required` is reported by `S013`, and `WriteOnly` with `Computed` is reported
here.

## Flagged Code

```go
&schema.Schema{
    Type:      schema.TypeString,
    Optional:  true,
    ForceNew:  true,
    WriteOnly: true,
}

&schema.Schema{
    Type:      schema.TypeString,
    Optional:  true,
    Default:   "value",
    WriteOnly: true,
}
```

## Passing Code

```go
&schema.Schema{
    Type:      schema.TypeString,
    Optional:  true,
    WriteOnly: true,
}

# OR

&schema.Schema{
    Type:      schema.TypeString,
    Required:  true,
    WriteOnly: true,
}
```

## Ignoring Reports

Singular reports can be ignored by adding a `//lintignore:S041` Go code comment at the end of the offending line or on the line immediately preceding, e.g.

```go
//lintignore:S041
&schema.Schema{
    Type:      schema.TypeString,
    Optional:  true,
    ForceNew:  true,
    WriteOnly: true,
}
```
