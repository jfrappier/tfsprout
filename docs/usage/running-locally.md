# Running tfsprout

`tfsprout` and `tfsproutx` work identically; only the registered check set differs. Everything below applies to both.

## Analyzing a provider

Change into the provider's directory and pass a package pattern:

```shell
tfsprout ./...
```

`./...` recurses through every package in the module. A bare `tfsprout` with **no** package argument prints usage and exits `1` — unlike some linters, there is no implicit default.

To analyze a subtree, which is much faster on large providers:

```shell
tfsprout ./internal/service/ec2/...
```

Findings are printed to stderr in `go vet` format, and the process exits `3` if anything was reported:

```
internal/service/example/resource_thing.go:42:3: AT001: missing CheckDestroy
```

See [Exit codes and output](../reference/exit-codes-and-output.md).

## Discovering checks

List every registered check:

```shell
tfsprout help
```

Read the full documentation for one check, including any flags it accepts:

```shell
tfsprout help AT001
```

This is the authoritative description. The [check index](../reference/checks.md) mirrors it with worked examples of flagged and passing code.

## Selecting checks

**Run only specific checks.** Naming any check disables all the others:

```shell
tfsprout -AT001 ./...
tfsprout -AT001 -AT005 -S013 ./...
```

This is the practical way to adopt tfsprout on a provider that has never been linted — start with the checks you already pass and grow the list.

**Run everything except specific checks.** Set them to `false`:

```shell
tfsprout -R009=false ./...
```

Do not mix the two styles in one invocation; pick whichever expresses your intent.

**Configure an individual check.** Per-check flags are namespaced with a dot:

```shell
tfsprout -AT001.ignored-filename-prefixes=data_source_,legacy_ ./...
```

## Test files

Test files are analyzed by default. Since the `AT` checks only match inside `_test.go` files, disabling them silences that entire category:

```shell
tfsprout -test=false ./...   # schema and resource findings only
```

## Applying fixes

```shell
tfsprout -fix ./...
```

Only three checks can rewrite code, and `-fix` edits files in place without a backup. Preview first with `-diff`, and read [Automated fixes](automated-fixes.md) before running it.

## Running via go vet

tfsprout can act as a `go vet` tool:

```shell
go vet -vettool $(which tfsprout) ./...
```

In this mode `go vet` handles package loading and build configuration. Use it when you need build tags, which the binary cannot set on its own:

```shell
go vet -vettool $(which tfsprout) -tags=integration ./...
```

## Suppressing individual findings

Add a `//lintignore:` comment naming the check ID:

```go
//lintignore:R009 // panic is unreachable, guarded above
panic("unreachable")
```

See [Ignoring reports](ignoring-reports.md) for scoping rules and adoption strategy.

## Machine-readable output

```shell
tfsprout -json ./...
```

JSON goes to stdout, and the exit code is always `0` — a CI job using `-json` must inspect the output rather than the exit status.

## See also

- [CLI reference](../reference/cli.md) — every flag
- [CI integration](ci-integration.md) — wiring this into a pipeline
- [Troubleshooting](troubleshooting.md) — when something looks wrong
