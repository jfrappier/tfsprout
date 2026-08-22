# Troubleshooting

## `internal error: package "context" without types was imported`

```
internal error: package "context" without types was imported from "github.com/example/terraform-provider-example/internal/service"
```

**Cause.** A `golang.org/x/tools` incompatibility with the **Go 1.27** toolchain. It is not a bug in any check — the run crashes during package loading, before analysis begins.

**Fix.** Upgrade to tfsprout **v0.1.1 or later**, which bumps `golang.org/x/tools` to a release that understands the Go 1.27 export data format.

```shell
go install github.com/jfrappier/tfsprout/cmd/tfsprout@latest
tfsprout -V
```

The Go version that matters is the one running tfsprout, not the one in your provider's `go.mod`. If your CI image was upgraded to Go 1.27 and lint started failing without a code change, this is why. See the [CHANGELOG](../../CHANGELOG.md).

## No output at all, exit code 0

tfsprout exits 0 and prints nothing. Three possible causes, in order of likelihood:

**Your provider does not use the Terraform Plugin SDK.** Providers built on `terraform-plugin-framework` produce no findings, because none of the types the checks match on are present. This is the most common cause and the most misleading, since it looks identical to a clean run. See [Scope and SDK support](../concepts/scope-and-sdk-support.md).

**You did not match any packages.** `tfsprout` with no arguments analyzes the current directory only. Use `./...` to recurse:

```shell
tfsprout ./...
```

**Your provider is genuinely clean.** Verify by enabling a check you know you violate, or by running against a single file you have inspected.

## `-fix` changed nothing

Expected. Only three checks produce suggested fixes, and two of them are extra checks unavailable in `tfsprout`. See [Automated fixes](automated-fixes.md).

## `flag provided but not defined: -XR002`

You are running `tfsprout`, which contains standard checks only. Extra checks (any ID starting with `X`) require `tfsproutx`:

```shell
go install github.com/jfrappier/tfsprout/cmd/tfsproutx@latest
tfsproutx -XR002 ./...
```

See [Standard vs extra checks](../concepts/standard-vs-extra.md).

## `flag provided but not defined: -V002.something`

The check is [removed](../reference/removed-checks.md). Removed analyzers accept `-V002` and `-V002=false` but register no per-check flags of their own. Drop the flag.

## An ignore comment is not working

Work through these in order:

1. **Check the ID spelling.** `//lintignore:S0013` is silently inert — IDs are not validated. Compare against the [check index](../reference/checks.md).
2. **Check the comment style.** Only `//` line comments are recognized. `/* lintignore:S013 */` does nothing.
3. **Check placement.** The comment must be on the line immediately preceding the code, or trailing it on the same line, so that it attaches to the right AST node.
4. **Check for spaces in a list.** `//lintignore:S013, S016` — the space makes the second ID ` S016`, which matches nothing. Write `//lintignore:S013,S016`.

See [Ignoring reports](ignoring-reports.md).

## `go vet` reports a different set of findings

Running through `go vet -vettool` applies `go vet`'s own package selection and build-tag handling, which can differ from running the binary directly. Most often this is build tags: acceptance test files behind a tag are invisible to whichever invocation does not set them.

With `go vet`, pass build tags through the `go` command, which understands them:

```shell
go vet -vettool $(which tfsprout) -tags=integration ./...
```

Running the binary directly, `-tags` is accepted but **has no effect** — it is a deprecated no-op inherited from the `go/analysis` driver. There is no way to set build tags on a direct invocation, so use `go vet` when you need them.

The other common cause is test files. The `AT` checks only ever fire in `_test.go` files, so if you passed `-test=false`, every acceptance test check goes quiet. See the [CLI reference](../reference/cli.md).

## Analysis is slow or runs out of memory

Large providers — several thousand resources — can take minutes and multiple gigabytes, because every package is type-checked before analysis. Options:

- Analyze a subtree rather than the whole module: `tfsprout ./internal/service/ec2/...`
- Run fewer checks: `tfsprout -AT001 -S013 ./...`
- Split CI into parallel jobs by service directory.

There is no incremental cache; each run is from scratch.

## Something else

Open an issue at [github.com/jfrappier/tfsprout/issues](https://github.com/jfrappier/tfsprout/issues). Include the output of `tfsprout -V=full`, your Go version (`go version`), and the exact command you ran.
