# How tfsprout works

tfsprout is a [`go/analysis`](https://pkg.go.dev/golang.org/x/tools/go/analysis) driver. Understanding that framework explains most of the tool's behavior — why it needs your code to compile, why it is slow on large providers, and why the codebase has twice as many analyzers as it has checks.

## The go/analysis model

An `analysis.Analyzer` is a unit of work with a name, a documentation string, a list of other analyzers it `Requires`, and a `Run` function. The framework loads and type-checks your packages, then runs each analyzer once per package, in dependency order, caching results.

Analyzers communicate through results. If analyzer B declares `Requires: []*analysis.Analyzer{A}`, then B's `Run` can read A's output from `pass.ResultOf[A]`. This is the mechanism that lets tfsprout share expensive AST traversals across dozens of checks.

The entry points are thin. `cmd/tfsprout/tfsprout.go` registers a version flag and hands a list of analyzers to `multichecker.Main`:

```go
func main() {
	cmdflags.AddVersionFlag()
	multichecker.Main(passes.AllChecks...)
}
```

`cmd/tfsproutx` is the same, appending `xpasses.AllChecks`. Everything else — flag parsing, package loading, `-fix`, `-json`, exit codes — comes from the framework.

## Two layers of analyzers

This is the key structural idea, and the thing that surprises most first-time contributors: **most analyzers in this repository are not checks.**

**Layer 1 — information gatherers.** Analyzers under `passes/helper/...`, `passes/stdlib/...`, and `passes/terraform/...` walk the AST and collect facts about Terraform Plugin SDK constructs. They never report anything. For example, `passes/helper/resource/testcaseinfo` returns every `resource.TestCase` composite literal in the package, already parsed into a struct that knows which fields were declared. `passes/helper/schema/resourcedatasetcallexpr` returns every `ResourceData.Set()` call expression.

**Layer 2 — checks.** The `AT`/`R`/`S`/`V` and `X*` analyzers consume those results and decide what is worth reporting. They contain almost no AST traversal of their own.

`AT001` is representative. It requires two analyzers and does nothing but filter:

```go
var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		commentignore.Analyzer,
		testcaseinfo.Analyzer,
	},
	Run: run,
}
```

Its `Run` reads the list of test cases, skips ignored files and ignored nodes, and reports the ones missing `CheckDestroy`.

The payoff is that the expensive work — finding every `TestCase` in a package — happens once no matter how many `AT` checks are enabled, because the framework caches `testcaseinfo`'s result.

## Recognizing SDK types

Layer 1 analyzers identify SDK constructs by resolving types against **package paths**, declared as constants under `helper/terraformtype/`. tfsprout does not import the Terraform Plugin SDK; it asks the type checker whether an expression's type resolves to a named type in the expected package path.

This is why tfsprout imposes no constraint on which SDK version your provider uses, and why a provider that does not import those packages produces no findings at all. See [Scope and SDK support](scope-and-sdk-support.md).

Two helper packages support this work:

- `helper/terraformtype/` — models SDK types: `schema.Schema`, `schema.Resource`, `resource.TestCase`, `validation` functions, `diag` types. Includes field-name constants and helpers like `DeclaresField`.
- `helper/astutils/` — generic Go AST utilities: examining basic literals, composite literals, function types, field lists, and package qualifiers.

## Shared analyzer constructors

Many checks are structurally identical — "report usage of X", "report usage of X and suggest Y instead". Rather than hand-writing each one, `helper/analysisutils/` provides constructors that build a complete analyzer from a few parameters:

```go
var Analyzer = analysisutils.RemovedAnalyzer("R008")
```

The constructors in `analyzers.go` pair with runners in `runners.go`. This is also where automated fixes come from — only the deprecation and avoidance runners attach `SuggestedFixes`, which is why so few checks support `-fix`. See [Automated fixes](../usage/automated-fixes.md).

## Suppression

`passes/commentignore` is a layer 1 analyzer with an unusual job: it builds a map from check ID to source ranges covered by a `//lintignore:` comment. Nearly every check requires it and calls `ignorer.ShouldIgnore(analyzerName, node)` before reporting.

Because it maps comments to the AST node they attach to, ignoring a composite literal ignores everything inside it. See [Ignoring reports](../usage/ignoring-reports.md).

## Consequences worth knowing

**Your code must compile.** Analysis runs after type checking. If your provider does not build, tfsprout exits `1` without reporting anything useful.

**There is no incremental cache across runs.** The framework caches within a run, not between them. Every invocation type-checks the module from scratch.

**Checks are independent.** No check can see another check's findings, only layer 1 results. That is why overlapping checks like `S013` and `S014` exist separately rather than as one configurable check.

## See also

- [Adding an analyzer](../contributing/adding-an-analyzer.md) — writing a check against this architecture
- [Implementing a custom lint tool](../contributing/custom-lint-tool.md) — reusing these analyzers in your own binary
