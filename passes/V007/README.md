# V007

_This terraform-plugin-sdk (v1) analyzer was removed upstream in [tfproviderlint](https://github.com/bflad/tfproviderlint) v0.30.0 and reports nothing in tfsprout. Its ID is retained so existing `//lintignore:` comments and CI flags keep working._

The V007 analyzer reports usage of the deprecated [ValidateRegexp](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation#ValidateRegexp) validation function that should be replaced with [StringIsValidRegExp](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation#StringIsValidRegExp).

## Flagged Code

```go
ValidateFunc: validation.ValidateRegexp,
```

## Passing Code

```go
ValidateFunc: validation.StringIsValidRegExp,
```

## Ignoring Reports

Singular reports can be ignored by adding a `//lintignore:V007` Go code comment at the end of the offending line or on the line immediately preceding, e.g.

```go
//lintignore:V007
ValidateFunc: validation.ValidateRegexp,
```
