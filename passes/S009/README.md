# S009

The S009 analyzer reports cases of `TypeList` or `TypeSet` schemas configuring
`ValidateFunc` or `ValidateDiagFunc`, which will fail schema validation.

Neither validator is supported at the top level of a list or set. Validation
belongs on the element schema instead, where it runs against each element.

## Flagged Code

```go
&schema.Schema{
    Type:         schema.TypeList,
    Elem:         &schema.Schema{Type: schema.TypeString},
    ValidateFunc: /* ... */,
}

&schema.Schema{
    Type:             schema.TypeSet,
    Elem:             &schema.Schema{Type: schema.TypeString},
    ValidateDiagFunc: /* ... */,
}
```

## Passing Code

```go
&schema.Schema{
    Type: schema.TypeList,
    Elem: &schema.Schema{Type: schema.TypeString},
}

&schema.Schema{
    Type: schema.TypeSet,
    Elem: &schema.Schema{Type: schema.TypeString},
}

&schema.Schema{
    Type: schema.TypeList,
    Elem: &schema.Schema{
      Type:         schema.TypeString,
      ValidateFunc: /* ... */,
    },
}

&schema.Schema{
    Type: schema.TypeSet,
    Elem: &schema.Schema{
      Type:             schema.TypeString,
      ValidateDiagFunc: /* ... */,
    },
}
```

## Ignoring Reports

Singular reports can be ignored by adding a `//lintignore:S009` Go code comment at the end of the offending line or on the line immediately preceding, e.g.

```go
//lintignore:S009
&schema.Schema{
    Type:         schema.TypeList,
    Elem:         &schema.Schema{Type: schema.TypeString},
    ValidateFunc: /* ... */,
}
```
