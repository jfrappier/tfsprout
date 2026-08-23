# S043

The S043 analyzer reports configuration blocks that contain a `WriteOnly`
attribute at any depth while the block itself is either a `TypeSet` or
`Computed`, which will fail provider schema validation.

A set's element identity is derived from its values, and a write-only value is
never persisted, so a set containing one cannot be matched against prior state.
A `Computed` block is likewise never configured by the practitioner.

A `TypeList` block that is not `Computed` may contain write-only attributes.

## Flagged Code

```go
&schema.Schema{
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

&schema.Schema{
    Type:     schema.TypeList,
    Computed: true,
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
```

## Passing Code

```go
&schema.Schema{
    Type:     schema.TypeList,
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
```

## Ignoring Reports

Singular reports can be ignored by adding a `//lintignore:S043` Go code comment at the end of the offending line or on the line immediately preceding, e.g.

```go
//lintignore:S043
&schema.Schema{
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
```
