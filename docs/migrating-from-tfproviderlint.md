# Migrating from tfproviderlint

tfsprout is a fork of [`tfproviderlint`](https://github.com/bflad/tfproviderlint). **v0.1.0 is a drop-in replacement**: the checks, their IDs, their behavior, and their flags are unchanged, and there are no new checks. Only the names and the module path differ.

This page lists everything you have to change.

## What changes

| | tfproviderlint | tfsprout |
|---|---|---|
| Standard command | `tfproviderlint` | `tfsprout` |
| Extended command | `tfproviderlintx` | `tfsproutx` |
| Go module path | `github.com/bflad/tfproviderlint` | `github.com/jfrappier/tfsprout` |
| GitHub Action | `bflad/tfproviderlint-github-action` | [`jfrappier/tfsprout-github-action`](https://github.com/jfrappier/tfsprout-github-action) |

## What does not change

- **Check IDs.** `AT001` is still `AT001`, `S013` is still `S013`. Nothing was renumbered.
- **Check behavior.** No check reports more or less than it did before.
- **`//lintignore:` comments.** The suppression syntax and the key names are identical, so **you do not need to touch a single ignore comment in your provider**. See [Ignoring reports](usage/ignoring-reports.md).
- **Per-check flags.** Options like `-AT001.ignored-filename-prefixes` work exactly as before.
- **Removed checks.** The checks removed upstream in v0.30.0 are still removed here, and still accept their flags without reporting. See [Removed checks](reference/removed-checks.md).

## Migration steps

### 1. Replace the binary

```shell
go install github.com/jfrappier/tfsprout/cmd/tfsprout@latest
```

Or `cmd/tfsproutx` if you were using `tfproviderlintx`. See [Install](install.md) for release binaries and pinning.

### 2. Update your CI invocation

If you invoke the binary directly, change the command name:

```diff
-tfproviderlint ./...
+tfsprout ./...
```

If you use `go vet`:

```diff
-go vet -vettool $(which tfproviderlint) ./...
+go vet -vettool $(which tfsprout) ./...
```

If you use the GitHub Action, switch to [`jfrappier/tfsprout-github-action`](https://github.com/jfrappier/tfsprout-github-action).

### 3. Update `tools.go`, if you have one

Providers that pin lint tooling through a `tools.go` build-tagged file need the import path changed:

```diff
-_ "github.com/bflad/tfproviderlint/cmd/tfproviderlint"
+_ "github.com/jfrappier/tfsprout/cmd/tfsprout"
```

Then `go mod tidy`.

### 4. Update custom lint tools

If you built your own binary from the upstream analyzer packages, change the import paths. The package structure is identical:

```diff
-"github.com/bflad/tfproviderlint/passes"
-"github.com/bflad/tfproviderlint/xpasses"
+"github.com/jfrappier/tfsprout/passes"
+"github.com/jfrappier/tfsprout/xpasses"
```

`passes.AllChecks` and `xpasses.AllChecks` have the same names and the same contents. See [Implementing a custom lint tool](contributing/custom-lint-tool.md).

## Why you might want to migrate

Beyond the rename, tfsprout v0.1.1 fixes a crash that affects anyone analyzing a provider with a **Go 1.27** toolchain:

```
internal error: package "context" without types was imported from ...
```

This is a `golang.org/x/tools` incompatibility, not a bug in the checks themselves. See [Troubleshooting](usage/troubleshooting.md) for the full symptom and the fix.

## Verifying the migration

Run both tools against the same provider and compare. The findings should be identical:

```shell
tfproviderlint ./... > /tmp/before.txt 2>&1
tfsprout ./...        > /tmp/after.txt  2>&1
diff /tmp/before.txt /tmp/after.txt
```

The only expected differences are the tool name in any error messages. If a check reports differently, that is a bug worth [filing](https://github.com/jfrappier/tfsprout/issues).
