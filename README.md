# tfsprout

Static analysis libraries and tooling for [Terraform Provider](https://www.terraform.io/docs/providers/index.html) code.

<p align="center"><img width="250" height="250" alt="image" src="https://github.com/user-attachments/assets/8d14fa27-9636-4d45-9fa1-390cc33dd77a" /></p>

[![PkgGoDev](https://pkg.go.dev/badge/github.com/jfrappier/tfsprout)](https://pkg.go.dev/github.com/jfrappier/tfsprout)

tfsprout runs 85 checks over a Terraform Provider's Go source and reports patterns that cause bugs, fail provider schema validation at runtime, or diverge from Terraform Plugin SDK conventions — missing `CheckDestroy` in acceptance tests, contradictory schema fields, unstable resource IDs, hand-rolled validators that duplicate `helper/validation`.

> **Note:** tfsprout is a fork of [`tfproviderlint`](https://github.com/bflad/tfproviderlint). **v0.1.0 is a drop-in replacement** — the lint checks and their behavior are unchanged and there are no new features. Only the project name, command names (`tfproviderlint` -> `tfsprout`, `tfproviderlintx` -> `tfsproutx`), and Go module path differ. See [Migrating from tfproviderlint](docs/migrating-from-tfproviderlint.md).
>
> **v0.1.1** fixes the `internal error: package "context" without types was imported from ...` crash that occurs when analyzing providers under **Go 1.27**. See the [CHANGELOG](CHANGELOG.md) for details.

## Quickstart

Install the binary:

```shell
go install github.com/jfrappier/tfsprout/cmd/tfsprout@latest
```

Run it against your provider:

```shell
cd /path/to/terraform-provider-example
tfsprout ./...
```

Findings print to stderr in `go vet` format, and the process exits `3` if anything was reported:

```
internal/service/example/resource_thing.go:42:3: AT001: missing CheckDestroy
internal/service/example/schema.go:17:5: S013: schema should configure one of Computed, Optional, or Required
```

Read what a check means:

```shell
tfsprout help AT001
```

Suppress an individual finding with a comment:

```go
//lintignore:R009 // panic is unreachable, guarded above
panic("unreachable")
```

That is the whole tool. Everything else is detail.

## Documentation

**<https://jfrappier.github.io/tfsprout/>** — searchable, with a page per check.

The same pages live in [`docs/`](docs/) if you would rather read them here.

**Getting started** — [What is tfsprout](docs/what-is-tfsprout.md) · [Migrating from tfproviderlint](docs/migrating-from-tfproviderlint.md) · [Install](docs/install.md)

**Usage** — [Running tfsprout](docs/usage/running-locally.md) · [Ignoring reports](docs/usage/ignoring-reports.md) · [Automated fixes](docs/usage/automated-fixes.md) · [CI integration](docs/usage/ci-integration.md) · [Troubleshooting](docs/usage/troubleshooting.md)

**Concepts** — [How tfsprout works](docs/concepts/how-it-works.md) · [Checks and categories](docs/concepts/checks-and-categories.md) · [Standard vs extra checks](docs/concepts/standard-vs-extra.md) · [Scope and SDK support](docs/concepts/scope-and-sdk-support.md)

**Reference** — [Check index](docs/reference/checks.md) · [CLI reference](docs/reference/cli.md) · [Exit codes and output](docs/reference/exit-codes-and-output.md) · [Removed checks](docs/reference/removed-checks.md)

**Contributing** — [CONTRIBUTING.md](CONTRIBUTING.md)

## Checks

85 active checks in four categories. The full list, with a page per check, is in the [check index](docs/reference/checks.md) — or browse them at <https://jfrappier.github.io/tfsprout/reference/checks/>.

| Prefix | Category | Count |
|---|---|---|
| [`AT`](docs/reference/checks.md#acceptance-test-checks) | Acceptance tests — `TestCase` and `TestStep` usage, test function naming | 12 |
| [`R`](docs/reference/checks.md#resource-checks) | Resources — `Resource` definitions, CRUD functions, `ResourceData` usage | 18 |
| [`S`](docs/reference/checks.md#schema-checks) | Schemas — `Schema` definitions and attribute maps | 37 |
| [`V`](docs/reference/checks.md#validation-checks) | Validation — `SchemaValidateFunc` and `helper/validation` usage | 7 |
| [`X*`](docs/reference/checks.md#extra-checks) | Extra opt-in checks, available via `tfsproutx` | 11 |

A further 9 IDs are retained but no longer report, having targeted Terraform Plugin SDK v1 APIs. See [Removed checks](docs/reference/removed-checks.md).

Standard checks are enabled by default in `tfsprout`. Extra checks require `tfsproutx`:

```shell
go install github.com/jfrappier/tfsprout/cmd/tfsproutx@latest
```

See [Standard vs extra checks](docs/concepts/standard-vs-extra.md).

## Scope

tfsprout analyzes **Terraform Plugin SDK** (`helper/schema`) provider source code.

It does **not** analyze Terraform configuration — `.tf` files are the domain of `terraform validate` and `tflint` — and it does **not** support [`terraform-plugin-framework`](https://github.com/hashicorp/terraform-plugin-framework) providers. Running it against a framework provider exits `0` with no output, which looks like a clean run but is not one. See [Scope and SDK support](docs/concepts/scope-and-sdk-support.md).

## Go compatibility

This project follows the [Go support policy](https://golang.org/doc/devel/release.html#policy): the two latest major releases are supported. Currently **Go 1.25 or later** is required.

## GitHub Action

A [GitHub Action](https://github.com/features/actions) is available: [tfsprout-github-action](https://github.com/jfrappier/tfsprout-github-action). See [CI integration](docs/usage/ci-integration.md).

## License

[Mozilla Public License 2.0](LICENSE)
