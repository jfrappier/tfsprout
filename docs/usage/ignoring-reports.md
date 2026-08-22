# Ignoring reports

Not every report is worth acting on, and no provider adopts a linter with a clean slate. tfsprout suppresses individual findings through `//lintignore:` comments in your Go source.

There is no configuration file and no repository-wide ignore list. Suppression is always local to the code it applies to, which keeps the reason visible next to the exception.

## Basic syntax

Add a `//lintignore:` comment naming the check ID, either on the line immediately before the offending code or trailing it on the same line:

```go
//lintignore:AT001
resource.ParallelTest(t, resource.TestCase{
    PreCheck:  func() { testAccPreCheck(t) },
    Providers: testAccProviders,
    Steps:     []resource.TestStep{ /* ... */ },
})
```

The check ID is exactly the ID from the [check index](../reference/checks.md) — `AT001`, `R014`, `S013`, `XR002`, and so on.

## Ignoring several checks at once

Separate IDs with commas. Whitespace around the list is trimmed, but do not put spaces between the IDs:

```go
//lintignore:S013,S016
"attribute_name": {
    Type: schema.TypeString,
},
```

## Explaining yourself

Anything after a second `//` is ignored by the parser, so you can record *why* the exception exists:

```go
//lintignore:R009 // panic is unreachable, guarded by the check above
panic("unreachable")
```

This is the recommended form. A bare ignore comment is indistinguishable from an oversight six months later.

## How scoping works

The comment attaches to the **AST node it precedes or trails**, and suppresses reports anywhere inside that node's source range. That has two practical consequences.

**Ignoring a composite literal ignores everything within it.** Putting `//lintignore:S013` above a whole `map[string]*schema.Schema` suppresses S013 for every attribute in the map, not just the first:

```go
//lintignore:S013
map[string]*schema.Schema{
    "first":  {Type: schema.TypeString},  // suppressed
    "second": {Type: schema.TypeString},  // also suppressed
}
```

Place the comment as tightly as possible around the code you actually mean to exempt.

**Ignoring a function ignores its whole body.** A comment above a `func` declaration suppresses that check for every statement inside. This is occasionally what you want — for example, exempting one generated file's worth of code — but it is a blunt instrument.

## Limitations

- **Line comments only.** `//lintignore:` is recognized; `/* lintignore:AT001 */` is not.
- **No wildcards.** There is no `//lintignore:*` or `//lintignore:S*`. Each ID must be named.
- **No file-level directive.** To exempt an entire file, either place ignores on each top-level declaration, or exclude the file from the packages you pass to the tool.
- **IDs are not validated.** A typo such as `//lintignore:S0013` silently does nothing. If an ignore comment does not seem to work, check the ID first.

## Adopting tfsprout on an existing provider

Running a linter for the first time against a mature provider typically produces hundreds of reports. Blanket-ignoring them defeats the point, and fixing them all before the first green build is rarely realistic. Two approaches work well:

**Enable checks incrementally.** Start by running only the checks you already satisfy, and add more as you fix them:

```shell
tfsprout -AT001 -AT005 -S013 ./...
```

This keeps CI green from day one and makes each newly enabled check a small, reviewable change. See [Running tfsprout](running-locally.md) for check selection.

**Or enable everything and ignore the backlog.** Turn on all checks, add `//lintignore:` comments with a tracking issue in the explanatory comment, and burn them down over time:

```go
//lintignore:XR002 // TODO(#412): implement Importer
```

The advantage is that *new* violations are caught immediately, since only pre-existing code carries an ignore. The cost is a large mechanical first commit.

For a provider under active development, the second approach usually pays off faster.

## Alternatives to ignoring

Before suppressing a report, consider whether the check should be off entirely. If a check does not suit your provider at all, disabling it in your invocation is more honest than scattering ignores:

```shell
tfsprout -R009=false ./...
```

And if you want that decision to be permanent rather than a CI flag, build a custom lint tool that omits the analyzer. See [Implementing a custom lint tool](../contributing/custom-lint-tool.md).
