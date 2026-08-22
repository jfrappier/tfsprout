# Testing

```shell
go test ./...
```

Tests are built on [`analysistest`](https://pkg.go.dev/golang.org/x/tools/go/analysis/analysistest), the standard harness for `go/analysis` analyzers. Each check runs against a miniature provider in its own `testdata/` directory and asserts that reports land exactly where `// want` comments say they should.

## The testdata module

Each check's `testdata/` is a **separate Go module** with a real `terraform-plugin-sdk` dependency:

```
passes/AT001/testdata/
├── go.mod
├── go.sum
└── src/
    ├── a/                  # primary fixtures
    ├── prefixes/           # fixtures for -ignored-filename-prefixes
    └── suffixes/           # fixtures for -ignored-filename-suffixes
```

This matters: the analyzers resolve SDK types through the type checker, so fixtures must import the genuine SDK rather than a stub. Copy `go.mod` and `go.sum` from a neighbouring check when adding a new one.

Because these are real modules, they need their dependencies downloaded before tests can run offline. CI does this explicitly:

```shell
for moddir in {,x}passes/*/testdata; do (cd "$moddir" && go mod download); done
```

Run that once locally if you hit `GOPROXY=off` errors from `analysistest`.

## Want comments

Mark each expected report with a `// want` comment. The string is a **regular expression** matched against the report message:

```go
_ = resource.TestCase{} // want "missing CheckDestroy"

resource.Test(t, resource.TestCase{}) // want "missing CheckDestroy"
```

A line with no `// want` comment must produce no report. This is what makes false positives visible — every passing fixture you add is an assertion that the check stays quiet.

Always include:

- **Failing cases**, one per distinct code shape the check should catch.
- **Passing cases**, especially shapes that look similar but are correct.
- **A suppression case**, proving `//lintignore:` works for your check.

## The test function

```go
package AT001_test

import (
    "testing"

    "golang.org/x/tools/go/analysis/analysistest"

    "github.com/jfrappier/tfsprout/passes/AT001"
)

func TestAT001(t *testing.T) {
    testdata := analysistest.TestData()
    analysistest.Run(t, testdata, AT001.Analyzer, "testdata/src/a")
}
```

The package is `AT001_test`, external to the analyzer package. `analysistest.TestData()` resolves the `testdata` directory relative to the test file.

## Testing per-check flags

Set the flag on the analyzer and point at a fixture directory built for it:

```go
func TestAT001CustomSuffixes(t *testing.T) {
    testdata := analysistest.TestData()
    analyzer := AT001.Analyzer
    analyzer.Flags.Set("ignored-filename-suffixes", "_data_source_test.go")
    analysistest.Run(t, testdata, analyzer, "testdata/src/suffixes")
}
```

Note that `AT001.Analyzer` is a package-level pointer, so setting a flag mutates shared state for the rest of the test binary. Give each flag configuration its own fixture directory, and be aware that test ordering within the package matters if you set the same flag twice.

## Testing suggested fixes

For checks that support `-fix`, use `RunWithSuggestedFixes`:

```go
import (
    "testing"

    "golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzerFixes(t *testing.T) {
    testdata := analysistest.TestData()
    analysistest.RunWithSuggestedFixes(t, testdata, Analyzer, "testdata/src/a")
}
```

The harness applies the fixes and compares the result against a `.golden` file sitting beside each source file — `testdata/src/a/main.go` is checked against `testdata/src/a/main.go.golden`.

Only three checks currently produce fixes. See [Automated fixes](../usage/automated-fixes.md).

## Validating the check list

`passes/checks_test.go` runs `analysis.Validate` over `AllChecks`, catching malformed analyzers, duplicate names, and dependency cycles. `xpasses` has the equivalent. Adding a check to `AllChecks` gets you this for free.

## CI

Pull requests run three jobs:

| Job | What it enforces |
|---|---|
| `go mod` | `go mod tidy` produces no diff |
| `go test` | Tests pass on Go **1.25.x** and **1.27.x** |
| `goreleaser` | `goreleaser check` and a snapshot build succeed |

The dual Go version matrix is deliberate: 1.25 is the `go.mod` minimum, and 1.27 guards the `package without types was imported` regression fixed in v0.1.1. Do not drop either.

## Before opening a pull request

```shell
go mod tidy
git diff --exit-code -- go.mod go.sum
go test ./...
```

Then run your build against a real provider — it is the only way to find false positives that fixtures miss. See [Building from source](building.md).
