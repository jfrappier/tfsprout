# Building from source

## Requirements

- **Go 1.25 or later.** This project follows the [Go support policy](https://golang.org/doc/devel/release.html#policy) and supports the two latest major releases. The 1.25 floor comes from `golang.org/x/tools`, whose Go 1.27-compatible release requires it.
- **Go Modules.** Dependency management is via `go.mod`; there is no vendor directory.

CI builds and tests against both `1.25.x` and `1.27.x`. The 1.27 job exists specifically to guard the `package without types was imported` regression — see [Troubleshooting](../usage/troubleshooting.md).

## Building

```shell
git clone https://github.com/jfrappier/tfsprout.git
cd tfsprout
go build ./...
```

## Installing your build

```shell
go install ./cmd/tfsprout
go install ./cmd/tfsproutx
```

This puts the binaries in `$GOBIN` (typically `$GOPATH/bin`), shadowing any released version on your `PATH`. Confirm which one you are running:

```shell
tfsprout -V=full
```

A locally built binary reports the version baked into `version/version.go` rather than a release tag, since the release version is injected at link time by GoReleaser. See [Releasing](releasing.md).

## Trying it against a real provider

The fastest way to sanity-check a change is to run your build against an actual provider:

```shell
go install ./cmd/tfsprout
cd /path/to/terraform-provider-example
tfsprout ./...
```

To exercise one check in isolation:

```shell
tfsprout -AT001 ./...
```

## Project layout

| Path | Contents |
|---|---|
| `cmd/tfsprout/` | Standard command entry point |
| `cmd/tfsproutx/` | Extended command entry point |
| `passes/` | Standard checks, plus the information-gathering analyzers they depend on |
| `xpasses/` | Extra checks |
| `helper/analysisutils/` | Constructors and runners for common analyzer shapes |
| `helper/astutils/` | Generic Go AST utilities |
| `helper/terraformtype/` | Terraform Plugin SDK type models and package path constants |
| `helper/cmdflags/` | Shared command flags, currently the version flag |
| `version/` | Version constants, overridden at link time on release |

Note that `passes/` holds both checks and non-check analyzers. Directories named for a check ID (`AT001`, `S013`) are checks; the others (`commentignore`, `helper`, `stdlib`, `terraform`, `testaccfuncdecl`, `testfuncdecl`) gather information for them. See [How tfsprout works](../concepts/how-it-works.md).

## Helpful tooling

Writing an analyzer means reasoning about AST shapes. These help:

- [`astdump`](https://github.com/wingyplus/astdump) — display the AST of a Go file
- [`ssadump`](https://pkg.go.dev/golang.org/x/tools/cmd/ssadump) — display and interpret the SSA form of a Go program

Dumping the AST of a small file containing the pattern you want to match is usually the quickest way to work out what to write.

## Dependencies

Updates are managed by [Dependabot](https://docs.github.com/en/code-security/dependabot). CI enforces a tidy module:

```shell
go mod tidy
git diff --exit-code -- go.mod go.sum
```

Run both before opening a pull request.

## Building the documentation site

The site at <https://jfrappier.github.io/tfsprout/> is built with
[MkDocs](https://www.mkdocs.org/) and the
[Cinder](https://github.com/chrissimpkins/cinder) theme. Every dependency is
pinned in `requirements-docs.txt`, which needs **Python 3.10 or later** —
`pymdown-extensions` 11 dropped 3.9, and on an older interpreter pip reports
only that no matching distribution was found:

```shell
python3 -m venv .venv
.venv/bin/pip install -r requirements-docs.txt
.venv/bin/mkdocs serve
```

That serves the site on <http://127.0.0.1:8000> and rebuilds on save. CI builds
with `--strict`, which turns broken internal links and unresolved anchors into
errors, so run it that way before opening a pull request:

```shell
.venv/bin/mkdocs build --strict
```

### What is generated rather than written

`docs-theme/hooks.py` builds two kinds of page at build time, so do not look for
them in `docs/`:

| Page | Built from |
|---|---|
| `checks/<ID>` | `passes/<ID>/README.md` or `xpasses/<ID>/README.md` |
| `changelog` | `CHANGELOG.md` |

The practical consequence is that a check's documentation lives beside its
analyzer. Adding a check directory with a `README.md` is all it takes for the
check to get a page, a sidebar entry, and a search index entry — see
[Adding an analyzer](adding-an-analyzer.md).

Check READMEs are read both on GitHub and on the site, so they use GitHub's
`> [!NOTE]` alert syntax; the hook rewrites it into a MkDocs admonition when it
builds the page.

## Next steps

- [Adding an analyzer](adding-an-analyzer.md)
- [Testing](testing.md)
