# R012

The R012 analyzer reports cases of data source `Resource` which configure `CustomizeDiff`, which is not valid.

## Flagged Code

```go
&schema.Resource{
    CustomizeDiff: /* ... */,
}
```

## Passing Code

```go
&schema.Resource{
    // No CustomizeDiff: /* ... */,
}
```

## Ignoring Reports

Singular reports can be ignored by adding a `//lintignore:R012` Go code comment at the end of the offending line or on the line immediately preceding, e.g.

```go
//lintignore:R012
&schema.Resource{
    CustomizeDiff: /* ... */,
}
```
