# Exit codes and output

tfsprout is built on the `go/analysis` [`multichecker`](https://pkg.go.dev/golang.org/x/tools/go/analysis/multichecker) driver, and inherits its output and exit behavior.

## Exit codes

| Code | Meaning |
|---|---|
| `0` | Analysis completed and no check reported a finding |
| `1` | Usage error — an unknown flag, a bad package pattern, or a package that failed to load or type-check |
| `3` | Analysis completed and at least one finding was reported |

**The failure code for findings is `3`, not `1`.** This trips people writing CI conditionals. Treat any non-zero status as failure rather than testing for a specific code:

```shell
if ! tfsprout ./...; then
  echo "lint failed"
  exit 1
fi
```

A code of `1` means tfsprout could not do its job — most often your provider does not compile. Fix the build first; a lint run against code that does not type-check is meaningless. See [Troubleshooting](../usage/troubleshooting.md).

## Report format

Findings go to **stderr**, one per line, in `go vet` format:

```
path/to/file.go:LINE:COLUMN: CHECKID: message
```

For example:

```
internal/service/example/resource_thing.go:42:3: AT001: missing CheckDestroy
internal/service/example/schema.go:17:5: S013: schema should configure one of Computed, Optional, or Required
```

The path is relative to the directory tfsprout was invoked from. The check ID prefix is part of the message, which makes findings greppable by check:

```shell
tfsprout ./... 2>&1 | grep -c 'S013:'
```

## JSON output

The driver supports `-json` for machine-readable output:

```shell
tfsprout -json ./...
```

The result is a nested object keyed by package ID, then by analyzer name, containing diagnostic objects with `posn`, `message`, and — where a check supports it — `suggested_fixes`. This is the format to parse if you are building a reviewdog-style annotation pipeline rather than scraping text.

Two differences from the default output are worth knowing:

- **JSON goes to stdout**, whereas text findings go to stderr.
- **The exit code is always `0` with `-json`**, even when findings exist. The driver reports diagnostics through the document instead of through exit status, so a CI job using `-json` must inspect the output to decide pass or fail — checking the exit code alone will always pass.

## See also

- [CLI reference](cli.md) for the full flag list
- [CI integration](../usage/ci-integration.md) for wiring this into a pipeline
