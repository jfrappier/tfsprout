# What is tfsprout

tfsprout is a static analysis tool for **Terraform Provider source code**. You point it at a provider's Go packages and it reports patterns that are known to cause bugs, fail provider schema validation at runtime, or diverge from Terraform Plugin SDK conventions.

It is a linter in the same family as `go vet` — in fact it is built on the same [`go/analysis`](https://pkg.go.dev/golang.org/x/tools/go/analysis) framework and can be run *as* a `go vet` tool. What makes it provider-specific is that its checks understand Terraform Plugin SDK types: `schema.Schema`, `schema.Resource`, `resource.TestCase`, `validation.*`, and so on.

## What it catches

Four families of problem, each with its own check prefix:

- **`AT` — Acceptance tests.** Missing `CheckDestroy`, malformed test function names, provider configuration leaking into `TestStep.Config`, multiple `resource.Test()` calls in one function.
- **`R` — Resources.** `ResourceData.Set()` misuse, deprecated `Exists` and `MigrateState` implementations, `RetryFunc` that swallow retryable errors, unstable IDs from `resource.UniqueId()` or `time.Now()`, Go `panic` usage.
- **`S` — Schemas.** Contradictory field combinations (`Required` with `Optional`, `Computed` with `Default`), missing `Elem` on `TypeList`/`TypeSet`/`TypeMap`, invalid `ConflictsWith`/`AtLeastOneOf`/`ExactlyOneOf` attribute references, invalid attribute names.
- **`V` — Validation.** Hand-rolled `SchemaValidateFunc` that reimplement something already in `helper/validation`, and `validation.StringMatch()` calls with an empty message.

The complete list lives in the [check index](reference/checks.md). Every check has its own page with flagged code, passing code, and how to suppress it.

## What it does not do

tfsprout analyzes **provider source code**, not Terraform configuration. It will not tell you anything about `.tf` files — that is what `terraform validate` and `tflint` are for.

It also targets the **Terraform Plugin SDK (`helper/schema`)** specifically. Providers written against `terraform-plugin-framework` will produce no findings, because none of the types tfsprout matches on are present. That is a silent no-op rather than an error, so read [Scope and SDK support](concepts/scope-and-sdk-support.md) before concluding your provider is clean.

Finally, it is almost entirely a **reporting** tool. Only a small number of checks can rewrite your code — see [Automated fixes](usage/automated-fixes.md).

## Relationship to tfproviderlint

tfsprout is a fork of [`tfproviderlint`](https://github.com/bflad/tfproviderlint). The check IDs, their behavior, and the `//lintignore:` mechanism are unchanged. What differs is the project name, the command names, and the Go module path. See [Migrating from tfproviderlint](migrating-from-tfproviderlint.md).

## Two commands

| Command | Checks included |
|---|---|
| `tfsprout` | Standard checks only — appropriate for any provider |
| `tfsproutx` | Standard checks **plus** extra checks — opinionated, opt-in |

See [Standard vs extra checks](concepts/standard-vs-extra.md) for how to choose.

## Next steps

- [Install](install.md)
- [Running tfsprout](usage/running-locally.md)
- [How tfsprout works](concepts/how-it-works.md)
