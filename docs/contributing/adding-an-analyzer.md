# Adding an analyzer

This walks through adding a new check end to end. Read [How tfsprout works](../concepts/how-it-works.md) first — the two-layer architecture determines most of what you write.

## Before you start

**Pick the right category and the next free number.** Prefixes are `AT`, `R`, `S`, `V`, with an `X` prefix for extra checks. Take the next unused number in that range; never reuse a [removed](../reference/removed-checks.md) ID.

**Decide standard or extra.** Standard checks flag things that are wrong. Extra checks flag things that are absent or debatable, and go in `xpasses/`. The test is whether a reasonable provider author could disagree with the finding. See [Standard vs extra checks](../concepts/standard-vs-extra.md).

**Check whether the information already exists.** Look through `passes/helper/...` before writing any AST traversal. If an analyzer already collects the construct you care about, your check is a filter over its results and should contain almost no traversal of its own.

## 1. Create the package

Create `passes/S038/` (or `xpasses/XS003/` for an extra check), containing `S038.go`:

```go
// Package S038 defines an Analyzer that checks for
// Schema that configure both Sensitive and Computed
package S038

import (
	"golang.org/x/tools/go/analysis"

	"github.com/jfrappier/tfsprout/passes/commentignore"
	"github.com/jfrappier/tfsprout/passes/helper/schema/schemainfocompositelit"
)

const Doc = `check for Schema that configure both Sensitive and Computed

The S038 analyzer reports cases of schemas which configure both Sensitive
and Computed, where the sensitivity has no effect.`

const analyzerName = "S038"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		commentignore.Analyzer,
		schemainfocompositelit.Analyzer,
	},
	Run: run,
}
```

Conventions that matter:

- The package name **is** the check ID.
- `Doc`'s **first line becomes the description** in `tfsprout help` and in the [check index](../reference/checks.md). Write it as `check for ...`, matching the existing checks, and keep it to one line.
- Use `analyzerName` as a constant; you need it for both suppression lookups and report messages.

## 2. Write the run function

Read your dependencies out of `pass.ResultOf`, skip suppressed nodes, and report:

```go
func run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemaInfos := pass.ResultOf[schemainfocompositelit.Analyzer].([]*schema.SchemaInfo)

	for _, schemaInfo := range schemaInfos {
		if ignorer.ShouldIgnore(analyzerName, schemaInfo.AstCompositeLit) {
			continue
		}

		if !schemaInfo.DeclaresBoolFieldWithZeroValue(schema.SchemaFieldSensitive) {
			continue
		}

		pass.Reportf(schemaInfo.AstCompositeLit.Lbrace, "%s: schema should not configure Sensitive with Computed", analyzerName)
	}

	return nil, nil
}
```

Three things are non-negotiable:

- **Always consult the ignorer** before reporting, or `//lintignore:` will not work for your check.
- **Prefix the message with the analyzer name**, as `"%s: message"`. Findings are greppable by check ID because every check does this.
- **Report a precise position.** Point at the offending field or literal, not the enclosing function.

## 3. Register it

Add the import and the analyzer to `AllChecks` in `passes/checks.go` (or `xpasses/checks.go`). Both lists are alphabetically ordered — keep them that way.

Only add checks that report. Information-gathering analyzers stay out of `AllChecks`; they are pulled in automatically through `Requires`.

`TestValidateAllChecks` in `passes/checks_test.go` calls `analysis.Validate` over the list and will catch a malformed analyzer or a dependency cycle.

## 4. Add test data

Create `testdata/` inside your check's directory. Each check's testdata is **its own Go module** with a real `terraform-plugin-sdk` dependency, so the analyzer resolves genuine SDK types rather than stubs:

```
passes/S038/testdata/
├── go.mod
├── go.sum
└── src/
    └── a/
        └── main_v2.go
```

Copy `go.mod` and `go.sum` from a neighbouring check. In the source files, mark each expected finding with a `// want` comment carrying a regular expression matched against the report message:

```go
_ = schema.Schema{
	Computed:  true,
	Sensitive: true,
} // want "schema should not configure Sensitive with Computed"
```

Cover the passing cases too — a check that only has failing fixtures will not catch false positives. Include a file exercising `//lintignore:S038` to prove suppression works.

## 5. Write the test

```go
package S038_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jfrappier/tfsprout/passes/S038"
)

func TestS038(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S038.Analyzer, "testdata/src/a")
}
```

See [Testing](testing.md) for per-check flags, suggested fixes, and golden files.

## 6. Document it

Add `README.md` to the check directory, following the structure every other check uses:

````markdown
# S038

The S038 analyzer reports cases of schemas which configure both
Sensitive and Computed, where the sensitivity has no effect.

## Flagged Code

```go
&schema.Schema{
    Computed:  true,
    Sensitive: true,
}
```

## Passing Code

```go
&schema.Schema{
    Computed: true,
}
```

## Ignoring Reports

Singular reports can be ignored by adding the a `//lintignore:S038` Go code
comment at the end of the offending line or on the line immediately
proceding, e.g.

```go
//lintignore:S038
&schema.Schema{
    Computed:  true,
    Sensitive: true,
}
```
````

Then add a row to the appropriate table in [`docs/reference/checks.md`](../reference/checks.md), using the first line of your `Doc` string as the description. CI fails if a check directory has no corresponding row.

## 7. Verify

```shell
go test ./...
go install ./cmd/tfsprout
cd /path/to/a/real/provider && tfsprout -S038 ./...
```

Running against a real provider is the step that catches false positives. A check that fires on idiomatic code in a well-maintained provider needs narrowing before it ships.

## Adding an information-gathering analyzer

If no existing analyzer surfaces what you need, add one under `passes/helper/...` mirroring the SDK package structure. These analyzers set a `ResultType`, return data, and never report.

Before writing one by hand, check `helper/analysisutils/` — `SelectorExprAnalyzer`, `FunctionCallExprAnalyzer`, `ReceiverMethodCallExprAnalyzer` and friends construct a complete analyzer from a package path and a name, and cover most cases in a single line.

Model any new SDK type in `helper/terraformtype/`, alongside the existing type models and package path constants.
