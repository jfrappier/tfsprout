# Removed checks

Nine check IDs still exist but no longer report anything. They were removed when the Terraform Plugin SDK dropped the v1 APIs they detected.

**Check IDs are never reused.** A removed check keeps its ID permanently, and its analyzer stays registered as a no-op. This means old `//lintignore:` comments and old CI flags referencing these IDs continue to work instead of breaking your build.

## The removed checks

| Check | Detected | Replacement |
|---|---|---|
| [R008](../checks/R008.md) | Deprecated `(schema.ResourceData).SetPartial()` | None — delete the call |
| [S034](../checks/S034.md) | `Schema` configuring `PromoteSingle` | None — invalid after Terraform 0.12, delete the field |
| [V002](../checks/V002.md) | `validation.CIDRNetwork()` | `validation.IsCIDRNetwork()` |
| [V003](../checks/V003.md) | `validation.IPRange()` | `validation.IsIPv4Range` |
| [V004](../checks/V004.md) | `validation.SingleIP()` | `validation.IsIPAddress` |
| [V005](../checks/V005.md) | `validation.ValidateJsonString` | `validation.StringIsJSON` |
| [V006](../checks/V006.md) | `validation.ValidateListUniqueStrings` | `validation.ListOfUniqueStrings` |
| [V007](../checks/V007.md) | `validation.ValidateRegexp` | `validation.StringIsValidRegExp` |
| [V008](../checks/V008.md) | `validation.ValidateRFC3339TimeString` | `validation.IsRFC3339Time` |

Note that `IsIPv4Range`, `IsIPAddress`, `StringIsJSON`, `ListOfUniqueStrings`, `StringIsValidRegExp`, and `IsRFC3339Time` are **values, not calls** — write `ValidateFunc: validation.IsIPAddress`, without parentheses. `IsCIDRNetwork` is the exception and still takes arguments.

## Why they were removed

All nine targeted `terraform-plugin-sdk` **v1** APIs. Those functions do not exist in v2, so a provider that compiles against v2 cannot contain the patterns these checks looked for. Keeping them would cost analysis time on every run and never produce a report.

The removal happened upstream in `tfproviderlint` v0.30.0, and tfsprout inherited it. That version number belongs to the upstream lineage, not to a tfsprout release — tfsprout's own history starts at v0.1.0. See [Migrating from tfproviderlint](../migrating-from-tfproviderlint.md).

## What happens if you reference one

**Enabling it is harmless.** `tfsprout -V002 ./...` runs and reports nothing. The analyzer is registered with the documentation string `REMOVED check`, so `tfsprout help V002` confirms its status.

**Disabling it is harmless.** `tfsprout -V002=false ./...` works. You can leave stale disable flags in CI indefinitely.

**Its per-check flags are gone.** Removed analyzers register no flags of their own, so a flag like `-V002.some-option` will fail with an unknown-flag error even though `-V002` itself is accepted.

**Old ignore comments are inert.** A leftover `//lintignore:V002` suppresses nothing because nothing reports. It is safe to leave, and safe to clean up.

## Cleaning up

None of this is urgent. If you want to tidy a provider that migrated from SDK v1:

```shell
grep -rn "lintignore:\(R008\|S034\|V00[2-8]\)" --include="*.go" .
```

Any hit is a suppression for a check that can no longer fire, and can be deleted along with whatever v1 API call it was guarding — if that call still compiles, you are still on SDK v1.
