# v0.1.1 (Unreleased)

BUG FIXES

* Fix `internal error: package "context" without types was imported from ...` crash when analyzing providers under Go 1.27. Upgraded `golang.org/x/tools` from `v0.30.0` to `v0.49.0`, whose reworked `go/packages` loader falls back to export data instead of a hard `log.Fatalf` when a dependency package arrives without complete type information. The minimum Go version (`go` directive) is now `1.25.0`, as required by the newer `golang.org/x/tools`.

# v0.1.0

NOTES

* tfsprout is a fork of [`tfproviderlint`](https://github.com/bflad/tfproviderlint), a static analysis tool for Terraform Providers. This release renames the project and its commands (`tfproviderlint` -> `tfsprout`, `tfproviderlintx` -> `tfsproutx`) and updates the Go module path to `github.com/jfrappier/tfsprout`. No lint check functionality has changed. For the history prior to the fork, see the upstream [`tfproviderlint` CHANGELOG](https://github.com/bflad/tfproviderlint/blob/main/CHANGELOG.md).
* **v0.1.0 is intended as a drop-in replacement** for the last `tfproviderlint` release: the lint checks and their behavior are identical, and there are no new features. Only the project name, command names, and module path have changed. Migrating is a matter of swapping the binary/command names (`tfproviderlint` -> `tfsprout`, `tfproviderlintx` -> `tfsproutx`).
