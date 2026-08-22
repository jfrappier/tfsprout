# tfsprout Documentation

Static analysis for [Terraform Provider](https://www.terraform.io/docs/providers/index.html) code, built on the Go [`go/analysis`](https://pkg.go.dev/golang.org/x/tools/go/analysis) framework.

New here? Read [What is tfsprout](what-is-tfsprout.md), then [Install](install.md), then [Running tfsprout](usage/running-locally.md).

Coming from `tfproviderlint`? Go straight to [Migrating from tfproviderlint](migrating-from-tfproviderlint.md).

## Getting started

| Page | What it covers |
|---|---|
| [What is tfsprout](what-is-tfsprout.md) | What the tool does, what it does not do, and who it is for |
| [Migrating from tfproviderlint](migrating-from-tfproviderlint.md) | Drop-in replacement steps for existing `tfproviderlint` users |
| [Install](install.md) | Release binaries, `go install`, Docker, and version pinning |

## Usage

| Page | What it covers |
|---|---|
| [Running tfsprout](usage/running-locally.md) | Running against a provider, selecting checks, `go vet` integration |
| [Ignoring reports](usage/ignoring-reports.md) | `//lintignore:` comments, scoping rules, and adoption strategy |
| [Automated fixes](usage/automated-fixes.md) | What `-fix` actually rewrites today |
| [CI integration](usage/ci-integration.md) | GitHub Actions, gating on exit status, incremental adoption |
| [Troubleshooting](usage/troubleshooting.md) | Common errors and what causes them |

## Concepts

| Page | What it covers |
|---|---|
| [How tfsprout works](concepts/how-it-works.md) | The `go/analysis` pipeline and the two layers of analyzers |
| [Checks and categories](concepts/checks-and-categories.md) | What `AT`/`R`/`S`/`V` mean and how check IDs behave over time |
| [Standard vs extra checks](concepts/standard-vs-extra.md) | Why `tfsproutx` exists and when to use it |
| [Scope and SDK support](concepts/scope-and-sdk-support.md) | Which providers tfsprout can analyze, and which it cannot |

## Reference

| Page | What it covers |
|---|---|
| [Check index](reference/checks.md) | Every check, with links to its full documentation |
| [CLI reference](reference/cli.md) | Flags accepted by `tfsprout` and `tfsproutx` |
| [Exit codes and output](reference/exit-codes-and-output.md) | What the tool prints and returns |
| [Removed checks](reference/removed-checks.md) | Checks that no longer report, and what replaced them |

## Contributing

| Page | What it covers |
|---|---|
| [Building from source](contributing/building.md) | Toolchain requirements and local builds |
| [Adding an analyzer](contributing/adding-an-analyzer.md) | Writing a new check end to end |
| [Testing](contributing/testing.md) | `analysistest`, testdata layout, golden files |
| [Implementing a custom lint tool](contributing/custom-lint-tool.md) | Building your own binary from tfsprout analyzers |
| [Releasing](contributing/releasing.md) | Maintainer release process |
