# V005

_This terraform-plugin-sdk (v1) analyzer was removed upstream in [tfproviderlint](https://github.com/bflad/tfproviderlint) v0.30.0 and reports nothing in tfsprout. Its ID is retained so existing `//lintignore:` comments and CI flags keep working._

The V005 analyzer reports usage of the deprecated [ValidateJsonString](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation#ValidateJsonString) validation function that should be replaced with [StringIsJSON](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation#StringIsJSON).

## Flagged Code

```go
ValidateFunc: validation.ValidateJsonString,
```

## Passing Code

```go
ValidateFunc: validation.StringIsJSON,
```

## Ignoring Reports

Singular reports can be ignored by adding a `//lintignore:V005` Go code comment at the end of the offending line or on the line immediately preceding, e.g.

```go
//lintignore:V005
ValidateFunc: validation.ValidateJsonString,
```
