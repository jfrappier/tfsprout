# Checks and categories

Every tfsprout check has a stable identifier: a letter prefix naming its category, followed by a three-digit number.

## Prefixes

| Prefix | Category | What it covers |
|---|---|---|
| `AT` | Acceptance tests | `resource.TestCase` and `resource.TestStep` usage, acceptance test function naming |
| `R` | Resources | `schema.Resource` definitions, CRUD functions, `ResourceData` method usage |
| `S` | Schemas | `schema.Schema` definitions and `map[string]*schema.Schema` attribute maps |
| `V` | Validation | `SchemaValidateFunc` implementations and `helper/validation` usage |

An `X` prefix on any of these — `XAT001`, `XR002`, `XS001` — marks an **extra** check, available only in `tfsproutx`. See [Standard vs extra checks](standard-vs-extra.md).

The prefix is not merely descriptive. It tells you where a finding lives and roughly what fixing it will involve: `AT` findings are in `_test.go` files, `S` findings are usually a one-line schema edit, `R` findings often require reasoning about CRUD behavior.

## Numbers are stable and never reused

A check ID, once assigned, belongs to that check forever.

**Checks are never renumbered.** `AT001` has always meant "TestCase missing CheckDestroy" and always will.

**Retired IDs are not recycled.** When a check is removed, its analyzer stays registered as a no-op rather than disappearing. Nine IDs are in this state today — see [Removed checks](../reference/removed-checks.md). This is what lets an old `//lintignore:V002` comment or a stale `-V002=false` CI flag remain harmless instead of breaking a build.

**Numbers are assigned sequentially, not semantically.** There is no meaning to `S001` being lower than `S039` beyond the order they were written. Adjacent numbers are unrelated, and gaps in a range mean a check was removed.

The practical consequence: you can pin a check ID in an ignore comment, a CI flag, or a code review checklist and expect it to keep meaning the same thing across upgrades.

## What a check consists of

Each check is a Go package under `passes/` (or `xpasses/`) named for its ID, containing:

- **The analyzer** — `AT001.go`, exporting `Analyzer` and a `Doc` constant.
- **The documentation** — `README.md`, with flagged code, passing code, and a suppression example.
- **Test data** — `testdata/src/a/`, a miniature provider exercising the check.

The `Doc` constant is the authoritative description. Its first line is the one-liner in the [check index](../reference/checks.md), and the whole string is what `tfsprout help AT001` prints.

## Severity

There is none. tfsprout has no severity levels, no warning-versus-error distinction, and no configurable thresholds. A check either reports or it does not, and any report causes a non-zero exit.

This is deliberate — the `go/analysis` framework has no severity concept — but it shapes how you adopt the tool. Since you cannot downgrade a check to a warning, the ways to live with findings you are not ready to fix are to disable the check, or to suppress individual sites with `//lintignore:`. Both are covered in [Ignoring reports](../usage/ignoring-reports.md).

## Overlapping checks

Some checks look like they could be one configurable check. `S013` and `S014` both concern whether `Computed`/`Optional`/`Required` are configured, in opposite directions. `S025` through `S033` are nine near-identical checks for fields that should not accompany a `Computed`-only schema.

They are separate because checks cannot see each other's findings — only shared layer 1 results — and because separate IDs let you enable, disable, and suppress each case independently. See [How tfsprout works](how-it-works.md).

## What a check ID does not tell you

- **Whether it can autofix.** Only `R007`, `XR007`, and `XR008` can. See [Automated fixes](../usage/automated-fixes.md).
- **Whether it accepts flags.** Most do not. Run `tfsprout help NAME` to find out.
- **Whether it still reports.** Nine do not. See [Removed checks](../reference/removed-checks.md).
