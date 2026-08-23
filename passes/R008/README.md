# R008

_This terraform-plugin-sdk (v1) analyzer was removed upstream in [tfproviderlint](https://github.com/bflad/tfproviderlint) v0.30.0 and reports nothing in tfsprout. Its ID is retained so existing `//lintignore:` comments and CI flags keep working._

The R008 analyzer reports usage of the deprecated [(helper/schema.ResourceData).SetPartial()](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema#ResourceData.SetPartial) function that does not need replacement.

## Flagged Code

```go
d.SetPartial("example"),
```

## Passing Code

```go
// Not present :)
```

## Ignoring Reports

Singular reports can be ignored by adding a `//lintignore:R008` Go code comment at the end of the offending line or on the line immediately preceding, e.g.

```go
//lintignore:R008
d.SetPartial("example"),
```
