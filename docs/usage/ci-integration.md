# CI integration

## GitHub Actions

A maintained action is available: [tfsprout-github-action](https://github.com/jfrappier/tfsprout-github-action). It installs the binary and runs it against your provider.

To wire it up by hand instead, which gives you control over the pinned version:

```yaml
name: Lint
on: pull_request

jobs:
  tfsprout:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version-file: 'go.mod'
      - run: go install github.com/jfrappier/tfsprout/cmd/tfsprout@v0.1.1
      - run: tfsprout ./...
```

**Pin the version.** With `@latest`, a tfsprout release you did not ask for can turn a green pull request red without a code change on your side.

**Mind the Go version.** `go-version-file: 'go.mod'` uses your provider's Go version to run the linter. Analyzing with a Go 1.27 toolchain requires tfsprout v0.1.1 or later — see [Troubleshooting](troubleshooting.md).

## Gating on the result

tfsprout exits `3` when it reports findings and `1` when analysis itself fails. Test for any non-zero status rather than a specific code:

```shell
if ! tfsprout ./...; then
  echo "tfsprout reported findings"
  exit 1
fi
```

Most CI systems do this for you by failing the step on any non-zero exit. The trap is a hand-written conditional testing `-eq 1`, which will silently pass on real findings. See [Exit codes and output](../reference/exit-codes-and-output.md).

Note that `-json` forces exit code `0`. If you use JSON output for annotations, parse the document to decide pass or fail — do not rely on the exit status.

## Running through go vet

If your provider already runs `go vet` in CI, adding tfsprout as a vettool avoids a second package-loading pass over the module:

```shell
go vet -vettool $(which tfsprout) ./...
```

This is also the only way to pass build tags, which matters if your acceptance tests sit behind one.

## Adopting incrementally

A mature provider will produce hundreds of findings on its first run. Turning everything on and requiring a green build is not a realistic first step. Two strategies:

**Enable a growing subset.** Start with checks you already satisfy and add more as you fix code:

```shell
tfsprout -AT001 -AT005 -R014 -S013 ./...
```

CI is green from the first commit, and every newly enabled check is a small reviewable change. The downside is that the flag list becomes long and needs maintaining.

**Enable everything and ignore the backlog.** Turn on all checks, then add `//lintignore:` comments referencing a tracking issue. New violations fail immediately, since only pre-existing code carries a suppression. The downside is one large mechanical commit up front.

For a provider under active development the second strategy usually pays off faster. See [Ignoring reports](ignoring-reports.md).

## Large providers

Analysis type-checks every package before running, with no incremental cache, so runtime scales with module size. For providers with thousands of resources:

- Shard by service directory across parallel jobs: `tfsprout ./internal/service/ec2/...`
- Analyze only packages touched by the pull request
- Drop `-test=false` in a fast pre-check job, then run the full set nightly

## Reporting findings inline

For pull request annotations, `-json` emits diagnostics keyed by package and analyzer, with position and message fields suitable for feeding a reviewdog-style pipeline. Remember to derive pass/fail from the document rather than the exit code.
