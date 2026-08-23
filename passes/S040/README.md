# S040

The S040 analyzer reports cases of schemas which only enable `Computed` and
configure `ValidateDiagFunc`, which will fail provider schema validation. There
is no practitioner input to validate on a computed-only attribute.

This is the `ValidateDiagFunc` counterpart to `S010`.

## Flagged Code

```go
&schema.Schema{
    Computed:         true,
    Type:             schema.TypeString,
    ValidateDiagFunc: /* ... */,
}
```

## Passing Code

```go
&schema.Schema{
    Computed: true,
    Type:     schema.TypeString,
}

# OR

&schema.Schema{
    Optional:         true,
    Type:             schema.TypeString,
    ValidateDiagFunc: /* ... */,
}
```

## Ignoring Reports

Singular reports can be ignored by adding a `//lintignore:S040` Go code comment at the end of the offending line or on the line immediately preceding, e.g.

```go
//lintignore:S040
&schema.Schema{
    Computed:         true,
    Type:             schema.TypeString,
    ValidateDiagFunc: /* ... */,
}
```
