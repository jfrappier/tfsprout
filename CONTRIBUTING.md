# Contributing to tfsprout

Thanks for your interest in improving tfsprout.

## Getting oriented

Start with [How tfsprout works](docs/concepts/how-it-works.md). tfsprout is a [`go/analysis`](https://pkg.go.dev/golang.org/x/tools/go/analysis) driver with a two-layer architecture — most analyzers in the repository gather information and never report — and that structure determines how nearly every change is written.

## Guides

| Guide | When you need it |
|---|---|
| [Building from source](docs/contributing/building.md) | Toolchain requirements, project layout, useful AST tooling |
| [Adding an analyzer](docs/contributing/adding-an-analyzer.md) | Writing a new check end to end |
| [Testing](docs/contributing/testing.md) | `analysistest`, testdata modules, `// want` comments, golden files |
| [Implementing a custom lint tool](docs/contributing/custom-lint-tool.md) | Reusing these analyzers in your own binary |
| [Releasing](docs/contributing/releasing.md) | Maintainers only |

## Quick reference

```shell
go build ./...                                # build
go test ./...                                 # test
go install ./cmd/tfsprout                     # install your build
./scripts/check-docs-sync.sh                  # verify the check index is current
```

Before opening a pull request:

```shell
go mod tidy
git diff --exit-code -- go.mod go.sum
go test ./...
```

Then run your build against a real Terraform Provider. Fixtures do not catch false positives; idiomatic code in a well-maintained provider does.

## Adding a check, in brief

1. Create `passes/<ID>/` — or `xpasses/<ID>/` if a reasonable provider author could disagree with the finding.
2. Write the analyzer, requiring `commentignore` and consulting the ignorer before every report.
3. Register it in `passes/checks.go` (or `xpasses/checks.go`), keeping the list alphabetical.
4. Add `testdata/` with failing, passing, and `//lintignore:` fixtures.
5. Add `README.md` to the check directory.
6. Add a row to [`docs/reference/checks.md`](docs/reference/checks.md) — CI fails if a check directory has no row.

The full walkthrough is in [Adding an analyzer](docs/contributing/adding-an-analyzer.md).

## Conventions worth knowing

- **Check IDs are permanent.** They are never renumbered and never reused. Removed checks stay registered as no-ops so existing `//lintignore:` comments and CI flags keep working. See [Removed checks](docs/reference/removed-checks.md).
- **The first line of a check's `Doc` constant is its description** in `tfsprout help` and in the check index. Write it as `check for ...`.
- **Report messages are prefixed with the analyzer name**, as `"%s: message"`, so findings stay greppable by check ID.
- **Go 1.25 is the floor**, and CI tests against both 1.25.x and 1.27.x. Do not drop either job.

## Reporting bugs

Open an issue using the [bug report template](.github/ISSUE_TEMPLATE/Bug_Report.md). Include:

- `tfsprout -V=full` output
- `go version` output
- The exact command you ran
- A minimal code sample that reproduces the behavior

A false positive — a check firing on correct code — is a bug, and a minimal reproducing sample is the most useful thing you can attach.

## Proposing a check

Open an issue using the [feature request template](.github/ISSUE_TEMPLATE/Feature_Request.md) before writing code. Worth addressing up front:

- What incorrect or risky pattern does it catch?
- What should authors write instead?
- Would every provider agree this is a problem? If not, it belongs in `xpasses/`.

## Dependencies

Updates are handled by [Dependabot](https://docs.github.com/en/code-security/dependabot). CI enforces a tidy module.

## License

Contributions are made under the [Mozilla Public License 2.0](LICENSE).
