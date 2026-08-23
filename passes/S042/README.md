# S042

The S042 analyzer reports cases of `TypeList`, `TypeMap`, or `TypeSet` schemas
which enable `WriteOnly`, which will fail provider schema validation.
`WriteOnly` is only supported on primitive types.

## Flagged Code

```go
&schema.Schema{
    Type:      schema.TypeList,
    Optional:  true,
    Elem:      &schema.Schema{Type: schema.TypeString},
    WriteOnly: true,
}

&schema.Schema{
    Type:      schema.TypeMap,
    Optional:  true,
    Elem:      &schema.Schema{Type: schema.TypeString},
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
    Type:     schema.TypeList,
    Optional: true,
    Elem:     &schema.Schema{Type: schema.TypeString},
}
```

## Ignoring Reports

Singular reports can be ignored by adding a `//lintignore:S042` Go code comment at the end of the offending line or on the line immediately preceding, e.g.

```go
//lintignore:S042
&schema.Schema{
    Type:      schema.TypeList,
    Optional:  true,
    Elem:      &schema.Schema{Type: schema.TypeString},
    WriteOnly: true,
}
```
