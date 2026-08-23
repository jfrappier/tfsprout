# S038

The S038 analyzer reports cases of schemas which configure both `ValidateFunc`
and `ValidateDiagFunc`, which will fail provider schema validation.

The two fields are mutually exclusive. `ValidateDiagFunc` is the replacement for
the deprecated `ValidateFunc`; it receives the attribute path and can return
warnings and errors as diagnostics. Configuring both is usually a partially
finished migration from one to the other.

## Flagged Code

```go
&schema.Schema{
    Type:             schema.TypeString,
    ValidateFunc:     /* ... */,
    ValidateDiagFunc: /* ... */,
}
```

## Passing Code

```go
&schema.Schema{
    Type:             schema.TypeString,
    ValidateDiagFunc: /* ... */,
}

# OR

&schema.Schema{
    Type:         schema.TypeString,
    ValidateFunc: /* ... */,
}
```

## Ignoring Reports

Singular reports can be ignored by adding a `//lintignore:S038` Go code comment at the end of the offending line or on the line immediately preceding, e.g.

```go
//lintignore:S038
&schema.Schema{
    Type:             schema.TypeString,
    ValidateFunc:     /* ... */,
    ValidateDiagFunc: /* ... */,
}
```
