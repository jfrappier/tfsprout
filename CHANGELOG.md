# v0.1.1

BUG FIXES

* Fix `internal error: package "context" without types was imported from ...` crash when analyzing providers under Go 1.27. Upgraded `golang.org/x/tools` from `v0.30.0` to `v0.49.0`, whose reworked `go/packages` loader falls back to export data instead of a hard `log.Fatalf` when a dependency package arrives without complete type information. The minimum Go version (`go` directive) is now `1.25.0`, as required by the newer `golang.org/x/tools`.
* Fix checks silently failing to match types referenced through a Go type alias (e.g. `type RetryError = resource.RetryError`) under Go 1.23+. With materialized type aliases now the default (`gotypesalias=1`), `go/types` reports a `*types.Alias` where a `*types.Named` was previously seen, so the `IsType*` matchers fell through and stopped firing. The type-matching switches in `helper/terraformtype` and `helper/astutils`, plus `R004`, now resolve aliases with `types.Unalias` before matching.

# v0.1.0

NOTES

* tfsprout is a fork of [`tfproviderlint`](https://github.com/bflad/tfproviderlint), a static analysis tool for Terraform Providers. This release renames the project and its commands (`tfproviderlint` -> `tfsprout`, `tfproviderlintx` -> `tfsproutx`) and updates the Go module path to `github.com/jfrappier/tfsprout`. No lint check functionality has changed. For the history prior to the fork, see the upstream [`tfproviderlint` CHANGELOG](https://github.com/bflad/tfproviderlint/blob/main/CHANGELOG.md).
* **v0.1.0 is intended as a drop-in replacement** for the last `tfproviderlint` release: the lint checks and their behavior are identical, and there are no new features. Only the project name, command names, and module path have changed. Migrating is a matter of swapping the binary/command names (`tfproviderlint` -> `tfsprout`, `tfproviderlintx` -> `tfsproutx`).
