# Standard vs extra checks

tfsprout splits its checks into two sets, shipped as two binaries.

| Command | Checks | Analyzer list |
|---|---|---|
| `tfsprout` | Standard only | `passes.AllChecks` |
| `tfsproutx` | Standard **plus** extra | `passes.AllChecks` + `xpasses.AllChecks` |

Extra checks all carry an `X` prefix: `XAT001`, `XR001` through `XR008`, `XS001`, `XS002`. They live in `xpasses/` rather than `passes/`.

Note that `tfsproutx` is a superset. Switching to it never loses a check.

## The distinction

**Standard checks flag things that are wrong.** A schema with both `Required` and `Optional` fails provider schema validation. A `TestCase` without `CheckDestroy` leaves infrastructure behind. A `RetryFunc` that swallows retryable errors produces flaky provider behavior. These findings are defensible in any provider, which is why they are on by default.

**Extra checks flag things that are absent or debatable.** They generally represent advanced Terraform Plugin SDK functionality, or house style, that is not appropriate for every provider:

- `XR002` reports resources that do not implement `Importer`. Import support is good practice — but some resources genuinely cannot be imported.
- `XR003` reports resources without `Timeouts`. Valuable for slow cloud APIs, noise for a resource that completes instantly.
- `XS001` requires a `Description` on every schema attribute. A reasonable documentation standard, and a large mechanical change to adopt.
- `XS002` requires schema attribute keys in alphabetical order. Purely stylistic.
- `XR007` and `XR008` report any `os/exec` usage, which some providers legitimately need.

The dividing line is roughly: would a reasonable provider author disagree with this finding? If yes, it is extra.

## Choosing

**Use `tfsprout`** if you want a linter that only tells you about defects. This is the right default, and it is what the [GitHub Action](https://github.com/jfrappier/tfsprout-github-action) runs.

**Use `tfsproutx`** if you are establishing conventions across a provider and want the tool to enforce them — a new provider, or one undergoing a documentation or consistency push.

**Use `tfsprout` plus selected extra checks** for the middle ground. Install `tfsproutx` and name the extras you want alongside the standard checks you care about:

```shell
tfsproutx -XR002 -XR003 ./...
```

Remember that naming any check disables all the others, so this form runs *only* the two named. To run everything except the extras you dislike, disable those instead:

```shell
tfsproutx -XS002=false ./...
```

## Errors from mismatched commands

Passing an extra check to `tfsprout` fails, because the analyzer is not registered:

```
flag provided but not defined: -XR002
```

Install `tfsproutx`. See [Troubleshooting](../usage/troubleshooting.md).

## Building your own set

If you want a fixed, permanent check set rather than a flag list — for example, all standard checks plus two extras, with `R009` never enabled — build a custom binary. Both `AllChecks` slices are exported for exactly this:

```go
analyzers := append(passes.AllChecks, XR002.Analyzer, XR003.Analyzer)
multichecker.Main(analyzers...)
```

See [Implementing a custom lint tool](../contributing/custom-lint-tool.md).
