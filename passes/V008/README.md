# V008

_This terraform-plugin-sdk (v1) analyzer was removed upstream in [tfproviderlint](https://github.com/bflad/tfproviderlint) v0.30.0 and reports nothing in tfsprout. Its ID is retained so existing `//lintignore:` comments and CI flags keep working._

The V008 analyzer reports usage of the deprecated [ValidateRFC3339TimeString](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation#ValidateRFC3339TimeString) validation function that should be replaced with [IsRFC3339Time](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation#IsRFC3339Time).

## Flagged Code

```go
ValidateFunc: validation.ValidateRFC3339TimeString,
```

## Passing Code

```go
ValidateFunc: validation.IsRFC3339Time,
```

## Ignoring Reports

Singular reports can be ignored by adding a `//lintignore:V008` Go code comment at the end of the offending line or on the line immediately preceding, e.g.

```go
//lintignore:V008
ValidateFunc: validation.ValidateRFC3339TimeString,
```
