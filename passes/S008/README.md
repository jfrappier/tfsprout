# S008

The S008 analyzer reports cases of `TypeList` or `TypeSet` schemas configuring `Default`,
which will fail schema validation.

## Flagged Code

```go
&schema.Schema{
    Type:    schema.TypeList,
    Default: /* ... */,
}

&schema.Schema{
    Type:    schema.TypeSet,
    Default: /* ... */,
}
```

## Passing Code

```go
&schema.Schema{
    Type: schema.TypeList,
}

&schema.Schema{
    Type: schema.TypeSet,
}
```

## Ignoring Reports

Singular reports can be ignored by adding a `//lintignore:S008` Go code comment at the end of the offending line or on the line immediately preceding, e.g.

```go
//lintignore:S008
&schema.Schema{
    Type:    schema.TypeList,
    Default: /* ... */,
}
```
