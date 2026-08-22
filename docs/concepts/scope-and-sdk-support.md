# Scope and SDK support

tfsprout analyzes Go source code that uses the **Terraform Plugin SDK**. This page explains precisely what that means, because the failure mode when your provider is out of scope is silence, not an error.

## What tfsprout matches on

tfsprout does not import the Terraform Plugin SDK. Look at `go.mod` and you will find only `golang.org/x/tools` and `github.com/bflad/gopaniccheck`.

Instead, the analyzers match against **package paths and type names in your provider's AST**. The recognized paths are declared as constants in `helper/terraformtype/`, for example in `helper/terraformtype/helper/schema/package.go`. When a check looks for a `schema.Resource`, it is asking the type checker whether the composite literal's type resolves to a named type in the SDK's `helper/schema` package path.

Two things follow from this design:

1. **tfsprout does not constrain which SDK version your provider depends on.** There is no version conflict to resolve, and upgrading the SDK does not require upgrading tfsprout.
2. **A provider that does not import those packages produces no findings.** The analyzers simply never match. There is no warning, because "no SDK types in this package" is indistinguishable from "a package that does not happen to define resources."

## Supported: terraform-plugin-sdk

Providers built on `github.com/hashicorp/terraform-plugin-sdk/v2` — the `helper/schema` model, where resources are `*schema.Resource` values with `CreateContext`/`ReadContext`/`UpdateContext`/`DeleteContext` functions and `map[string]*schema.Schema` attribute maps — are what every check is written against.

This includes the surrounding SDK packages the checks reason about:

| Package | Used by |
|---|---|
| `helper/schema` | The `S` checks, most `R` checks |
| `helper/resource` | The `AT` checks, `R006`, `R015`, `R016` |
| `helper/validation` | The `V` checks |
| `helper/acctest` | `AT009` |
| `helper/customdiff` | `R012` and related schema analysis |
| `diag` | Diagnostics-returning function signatures |

## Not supported: terraform-plugin-framework

Providers written against `github.com/hashicorp/terraform-plugin-framework` are **not analyzed**. The framework has a different architecture — resources are Go types implementing `resource.Resource`, schemas are `schema.Schema` structs from a different module with a different attribute model — and none of the patterns tfsprout matches exist there.

Running tfsprout against a framework provider is not an error. It exits 0 with no output, which reads exactly like a clean bill of health. **It is not one.** If your provider is framework-based, or you are mid-migration and have moved some resources across, tfsprout only tells you about the resources still on the SDK.

There is no plan documented here to add framework support; the check set would have to be rewritten from scratch.

## Muxed and migrating providers

Providers that serve both SDK and framework resources through `terraform-plugin-mux` are a common intermediate state. tfsprout handles these correctly in the sense that it analyzes the SDK half and ignores the framework half — but be aware that your coverage shrinks as you migrate, and a falling report count may reflect migration rather than improvement.

## Go version support

This project follows the [Go support policy](https://golang.org/doc/devel/release.html#policy): the two latest major releases are supported.

Currently **Go 1.25 or later** is required, both to build tfsprout and to consume it as a dependency. That floor comes from `golang.org/x/tools`, whose Go 1.27-compatible release requires Go 1.25 as a minimum.

The Go version that matters is the one **analyzing** your provider, not the one your provider targets. Analyzing a provider with a Go 1.27 toolchain requires tfsprout v0.1.1 or later; earlier versions crash. See [Troubleshooting](../usage/troubleshooting.md).

## Non-goals

- **Terraform configuration.** tfsprout reads `.go` files. It has nothing to say about `.tf` files — use `terraform validate` or [`tflint`](https://github.com/terraform-linters/tflint).
- **Provider behavior at runtime.** These are static checks. They cannot tell you that your API client is wrong, only that your code shape is suspicious.
- **General Go linting.** Use `go vet`, `staticcheck`, or `golangci-lint` alongside tfsprout, not instead of it. The one deliberate overlap is `R009` (panic usage), which exists because panics in provider CRUD functions crash Terraform in a particularly unhelpful way.
