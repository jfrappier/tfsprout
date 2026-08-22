# V004

_This terraform-plugin-sdk (v1) analyzer was removed upstream in [tfproviderlint](https://github.com/bflad/tfproviderlint) v0.30.0 and reports nothing in tfsprout. Its ID is retained so existing `//lintignore:` comments and CI flags keep working._

The V004 analyzer reports usage of the deprecated [SingleIP](https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation#SingleIP) validation function that should be replaced with [IsIPAddress](https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation#IsIPAddress).

## Flagged Code

```go
ValidateFunc: validation.SingleIP(),
```

## Passing Code

```go
ValidateFunc: validation.IsIPAddress,
```

## Ignoring Reports

Singular reports can be ignored by adding the a `//lintignore:V004` Go code comment at the end of the offending line or on the line immediately proceding, e.g.

```go
//lintignore:V004
ValidateFunc: validation.SingleIP(),
```
