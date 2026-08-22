# Implementing a custom lint tool

The `go/analysis` framework and this codebase are designed for flexibility. You may want to permanently disable certain default checks, ship a fixed check set rather than a long flag list, or add provider-specific checks of your own. All three are done by building your own binary.

`cmd/tfsproutx/tfsproutx.go` is the worked example — it does nothing but combine two analyzer lists:

```go
package main

import (
    "github.com/jfrappier/tfsprout/helper/cmdflags"
    "github.com/jfrappier/tfsprout/passes"
    "github.com/jfrappier/tfsprout/xpasses"
    "golang.org/x/tools/go/analysis"
    "golang.org/x/tools/go/analysis/multichecker"
)

func main() {
    cmdflags.AddVersionFlag()

    var analyzers []*analysis.Analyzer
    analyzers = append(analyzers, passes.AllChecks...)
    analyzers = append(analyzers, xpasses.AllChecks...)
    multichecker.Main(analyzers...)
}
```

Your binary inherits the entire driver — flag parsing, `help`, `-fix`, `-json`, exit codes — from `multichecker.Main`.

## Adding a few extra checks

The common case: all standard checks, plus the two extras you actually want.

```go
package main

import (
    "github.com/jfrappier/tfsprout/helper/cmdflags"
    "github.com/jfrappier/tfsprout/passes"
    "github.com/jfrappier/tfsprout/xpasses/XR002"
    "github.com/jfrappier/tfsprout/xpasses/XR003"
    "golang.org/x/tools/go/analysis"
    "golang.org/x/tools/go/analysis/multichecker"
)

func main() {
    cmdflags.AddVersionFlag()

    analyzers := append([]*analysis.Analyzer{}, passes.AllChecks...)
    analyzers = append(analyzers, XR002.Analyzer, XR003.Analyzer)
    multichecker.Main(analyzers...)
}
```

Copy `passes.AllChecks` into a fresh slice rather than appending to it directly — `append` may write into the package-level slice's backing array.

## Permanently excluding checks

There is no "all except" helper. To exclude a check, list the ones you want individually, the same way `AllChecks` is built in `passes/checks.go`:

```go
analyzers := []*analysis.Analyzer{
    AT001.Analyzer,
    AT002.Analyzer,
    // AT003 deliberately omitted: conflicts with our naming convention
    S013.Analyzer,
}
```

Verbose, but explicit — and unlike a `-AT003=false` flag in CI, the decision and its reason live in code where they get reviewed.

Filtering by name is the compact alternative:

```go
excluded := map[string]bool{"AT003": true, "R009": true}

var analyzers []*analysis.Analyzer
for _, a := range passes.AllChecks {
    if !excluded[a.Name] {
        analyzers = append(analyzers, a)
    }
}
```

This silently no-ops if a name is misspelled or a check is removed upstream, so consider asserting that every excluded name matched.

## Adding your own checks

Provider-specific conventions — internal naming rules, a required attribute, a banned helper — are worth encoding as analyzers in your own repository. Import them alongside tfsprout's:

```go
analyzers := append([]*analysis.Analyzer{}, passes.AllChecks...)
analyzers = append(analyzers, myprovider.NoDeprecatedClientAnalyzer)
multichecker.Main(analyzers...)
```

Two building blocks are exported for this, and using them means your checks behave like the built-in ones:

- **`helper/terraformtype/`** — models Terraform Plugin SDK types: `schema.Schema`, `schema.Resource`, `resource.TestCase`, `validation` functions, `diag` types, along with field name constants and helpers such as `DeclaresField`.
- **`helper/astutils/`** — generic Go AST utilities for basic literals, composite literals, function types, field lists, and package qualifiers.

The `passes/` directory also exports the information-gathering analyzers your checks can depend on. For example, `passes/helper/resource/retryfuncinfo` returns information from all named and anonymous declarations of `helper/resource.RetryFunc()`, so a check about retry behavior needs no traversal of its own. See [How tfsprout works](../concepts/how-it-works.md).

Support `//lintignore:` in your own checks by requiring `passes/commentignore` and calling `ignorer.ShouldIgnore(analyzerName, node)` before reporting. See [Adding an analyzer](adding-an-analyzer.md).

## Versioning

Your tool depends on tfsprout as a normal Go module:

```shell
go get github.com/jfrappier/tfsprout@v0.1.1
```

Pin it. `AllChecks` gaining a member is a behavior change for everyone who runs your binary, and you want that to arrive with a version bump you chose.

Check IDs are stable and never reused, so referencing `AT001.Analyzer` will not silently become a different check across upgrades. See [Checks and categories](../concepts/checks-and-categories.md).
