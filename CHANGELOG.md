# v0.2.0 (Unreleased)

NOTES

* This release changes which findings are reported. A provider that ran clean on v0.1.1 may report new findings without any change to its own code, and `tfsprout` exits `3` when it does. Plan to re-baseline CI when upgrading. See [Exit codes and output](docs/reference/exit-codes-and-output.md).

FEATURES

* **New Check:** `S039`: check for `Schema` with invalid resource identity configuration. An identity attribute may only configure `Type`, `Description`, `Elem`, and exactly one of `RequiredForImport` or `OptionalForImport`; anything else fails provider schema validation via `(*schema.ResourceIdentity).InternalIdentityValidate`. Because the SDK rejects the import fields outside an identity schema, this also catches the inverse mistake of setting `RequiredForImport` on an ordinary resource attribute.
* **New Check:** `S038`: check for `Schema` with both `ValidateFunc` and `ValidateDiagFunc` configured. The two are mutually exclusive and configuring both fails provider schema validation with `ValidateFunc and ValidateDiagFunc cannot both be set`.

BUG FIXES

* `S013`: no longer reports resource identity schema attributes. Identity schemas are declared as an ordinary `map[string]*schema.Schema` but configure `RequiredForImport`/`OptionalForImport` in place of `Computed`, `Optional`, or `Required`, so every attribute of every identity schema was reported. The bug is inherited from `tfproviderlint` (see [bflad/tfproviderlint#340](https://github.com/bflad/tfproviderlint/issues/340)) and affects any provider adopting Terraform 1.12 resource identity. Verified against `terraform-provider-scaleway`, where it accounted for 21 of the 22 findings tfsprout reported. Identity schemas are now validated on their own terms by the new `S039`.

ENHANCEMENTS

* `S009`: now also reports `ValidateDiagFunc` configured on a `TypeList` or `TypeSet` schema, not just `ValidateFunc`. The Terraform Plugin SDK rejects both identically (`ValidateFunc and ValidateDiagFunc are not yet supported on lists or sets`), so they are one rule and share the `S009` ID. Existing `//lintignore:S009` comments continue to suppress both, and no previously reported finding has changed position — only the report message, which now names both fields.

# v0.1.1

BUG FIXES

* Fix `internal error: package "context" without types was imported from ...` crash when analyzing providers under Go 1.27. Upgraded `golang.org/x/tools` from `v0.30.0` to `v0.49.0`, whose reworked `go/packages` loader falls back to export data instead of a hard `log.Fatalf` when a dependency package arrives without complete type information. The minimum Go version (`go` directive) is now `1.25.0`, as required by the newer `golang.org/x/tools`.
* Fix checks silently failing to match types referenced through a Go type alias (e.g. `type RetryError = resource.RetryError`) under Go 1.23+. With materialized type aliases now the default (`gotypesalias=1`), `go/types` reports a `*types.Alias` where a `*types.Named` was previously seen, so the `IsType*` matchers fell through and stopped firing. The type-matching switches in `helper/terraformtype` and `helper/astutils`, plus `R004`, now resolve aliases with `types.Unalias` before matching.

# v0.1.0

NOTES

* tfsprout is a fork of [`tfproviderlint`](https://github.com/bflad/tfproviderlint), a static analysis tool for Terraform Providers. This release renames the project and its commands (`tfproviderlint` -> `tfsprout`, `tfproviderlintx` -> `tfsproutx`) and updates the Go module path to `github.com/jfrappier/tfsprout`. No lint check functionality has changed. For the history prior to the fork, see the upstream [`tfproviderlint` CHANGELOG](https://github.com/bflad/tfproviderlint/blob/main/CHANGELOG.md).
* **v0.1.0 is intended as a drop-in replacement** for the last `tfproviderlint` release: the lint checks and their behavior are identical, and there are no new features. Only the project name, command names, and module path have changed. Migrating is a matter of swapping the binary/command names (`tfproviderlint` -> `tfsprout`, `tfproviderlintx` -> `tfsproutx`).
