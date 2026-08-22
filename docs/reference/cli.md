# CLI reference

`tfsprout` and `tfsproutx` accept identical flags. They differ only in which checks are registered — see [Standard vs extra checks](../concepts/standard-vs-extra.md).

Both are `go/analysis` [`multichecker`](https://pkg.go.dev/golang.org/x/tools/go/analysis/multichecker) drivers, so most of the flag surface is inherited from that framework and behaves the same as it does in `go vet`.

## Synopsis

```
tfsprout [flags] package...
```

The package arguments use standard Go package patterns: `.`, `./...`, `./internal/service/ec2/...`, or explicit import paths.

**At least one package argument is required.** Invoking `tfsprout` with no arguments prints usage and exits `1`.

## Getting help

| Command | Effect |
|---|---|
| `tfsprout help` | List every registered check with its one-line description |
| `tfsprout help NAME` | Full documentation and flags for one check, e.g. `tfsprout help AT001` |
| `tfsprout -flags` | Print all analyzer flags as JSON |

`tfsprout help NAME` prints the same text as the check's `Doc` string, which is the authoritative description. The per-check pages linked from the [check index](checks.md) carry the same content with worked examples.

## Selecting checks

Each check registers a boolean flag named after its ID.

**Enable only specific checks** by naming them. Naming any check disables all others:

```shell
tfsprout -AT001 ./...
tfsprout -AT001 -AT005 -S013 ./...
```

**Disable specific checks** by setting them to `false`. All others stay enabled:

```shell
tfsprout -AT001=false ./...
```

These two forms do not mix usefully — pick one style per invocation.

**Per-check flags** are namespaced under the check ID with a dot. For example, `AT001` accepts filename filters:

```shell
tfsprout -AT001.ignored-filename-prefixes=data_source_,legacy_ ./...
```

Run `tfsprout help AT001` to see which flags a check accepts. Most accept none. [Removed checks](removed-checks.md) accept their `-ID` toggle but register no per-check flags.

## Output flags

| Flag | Default | Effect |
|---|---|---|
| `-json` | `false` | Emit JSON to stdout instead of text to stderr. Forces exit code `0` — see [Exit codes and output](exit-codes-and-output.md) |
| `-c N` | `-1` | Display the offending line with `N` lines of surrounding context |

## Fix flags

| Flag | Default | Effect |
|---|---|---|
| `-fix` | `false` | Apply suggested fixes in place |
| `-diff` | `false` | With `-fix`, print a unified diff instead of writing files |

Only three checks produce fixes. Read [Automated fixes](../usage/automated-fixes.md) before using either flag — `-diff` is the safe way to preview.

## Analysis flags

| Flag | Default | Effect |
|---|---|---|
| `-test` | `true` | Analyze test files in addition to non-test files |

`-test=false` suppresses **every `AT` check**, since acceptance test checks only match inside `_test.go` files. Use it when you want schema and resource findings only.

## Version flags

| Flag | Effect |
|---|---|
| `-V` | Print the version and exit |
| `-version` | Same as `-V` |
| `-V=full` | Print a fully specified version, used by the `go` command to identify the tool |

Include `tfsprout -V=full` output when filing an issue.

## Diagnostic flags

These are for debugging tfsprout itself rather than your provider.

| Flag | Effect |
|---|---|
| `-debug=STR` | Debug flags, any subset of `fpstv`. `-debug=tp` prints per-analyzer timings with parallelism disabled |
| `-cpuprofile=FILE` | Write a CPU profile |
| `-memprofile=FILE` | Write a memory profile |
| `-trace=FILE` | Write an execution trace |

## Deprecated no-op flags

Accepted for compatibility, but they do nothing: `-source`, `-v`, `-all`, `-tags`.

`-tags` is the one that catches people out. To analyze code behind a build tag, run through `go vet` — which parses `-tags` itself — rather than invoking the binary directly:

```shell
go vet -vettool $(which tfsprout) -tags=integration ./...
```

## Running as a go vet tool

```shell
go vet -vettool $(which tfsprout) ./...
```

In this mode `go vet` handles package loading and build configuration, and tfsprout runs as a unitchecker over each package. Flags intended for tfsprout still work, but must not collide with `go vet`'s own flags.

## See also

- [Running tfsprout](../usage/running-locally.md) for task-oriented examples
- [Exit codes and output](exit-codes-and-output.md) for CI scripting
