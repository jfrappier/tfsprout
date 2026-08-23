# R019

The R019 analyzer reports when there are a large number of arguments being passed to [`(*schema.ResourceData).HasChanges()`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema#ResourceData.HasChanges), which it may be preferable to use [`(*schema.ResourceData).HasChangesExcept()`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema#ResourceData.HasChangesExcept) instead.

## Options

Namespaced under the check ID on the command line, as `-R019.threshold=VALUE`.

| Flag | Default | Effect |
|---|---|---|
| `threshold` | `5` | Number of arguments before a report is emitted |

## Flagged Code

```go
d.HasChanges("attr1", "attr2", "attr3", "attr4", "attr5")
```

## Passing Code

```go
d.HasChangesExcept("metadata_attr")
```

## Ignoring Reports

Singular reports can be ignored by adding a `//lintignore:R019` Go code comment at the end of the offending line or on the line immediately preceding, e.g.

```go
//lintignore:R019
d.HasChanges("attr1", "attr2", "attr3", "attr4", "attr5")
```
