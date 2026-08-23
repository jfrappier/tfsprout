---
title: Home
disable_toc: true
---

<div class="hero" markdown="1">

<img class="hero-logo" src="https://github.com/user-attachments/assets/8d14fa27-9636-4d45-9fa1-390cc33dd77a" alt="" width="140" height="140">

# tfsprout

<p class="hero-tagline">Static analysis for Terraform Provider source code.</p>

<p class="hero-sub">85 checks over a provider's Go source, reporting the patterns that cause bugs, fail schema validation at runtime, or drift from Terraform Plugin SDK conventions.</p>

<p class="hero-actions">
<a class="btn-primary" href="install/">Install</a>
<a class="btn-ghost" href="reference/checks/">Browse the checks</a>
<a class="btn-ghost" href="migrating-from-tfproviderlint/">Coming from tfproviderlint?</a>
</p>

</div>

<div class="quickstart" markdown="1">

## Sixty-second start

```shell
go install github.com/jfrappier/tfsprout/cmd/tfsprout@latest
cd /path/to/terraform-provider-example
tfsprout ./...
```

Findings print to stderr in `go vet` format, and the process exits `3` if anything was reported:

```text
internal/service/example/resource_thing.go:42:3: AT001: missing CheckDestroy
internal/service/example/schema.go:17:5: S013: schema should configure one of Computed, Optional, or Required
```

Read what a check means with `tfsprout help AT001`, and silence a single finding
with a comment:

```go
//lintignore:R009 // panic is unreachable, guarded above
panic("unreachable")
```

That is the whole tool. Everything below is detail.

</div>

## What tfsprout checks

<div class="card-grid" markdown="1">

<div class="card card--at" markdown="1">
### [Acceptance tests](reference/checks.md#acceptance-test-checks) <span class="card-count">12</span>
`TestCase` and `TestStep` usage, missing `CheckDestroy`, test function naming,
provider configuration leaking into a step's `Config`.
</div>

<div class="card card--r" markdown="1">
### [Resources](reference/checks.md#resource-checks) <span class="card-count">18</span>
`ResourceData.Set()` misuse, deprecated `Exists` and `MigrateState`, `RetryFunc`
that swallow retryable errors, unstable IDs, Go `panic` usage.
</div>

<div class="card card--s" markdown="1">
### [Schemas](reference/checks.md#schema-checks) <span class="card-count">37</span>
Contradictory field combinations, missing `Elem` on `TypeList`/`TypeSet`/`TypeMap`,
invalid `ConflictsWith` references, invalid attribute names.
</div>

<div class="card card--v" markdown="1">
### [Validation](reference/checks.md#validation-checks) <span class="card-count">7</span>
Hand-rolled `SchemaValidateFunc` that duplicate something already in
`helper/validation`, and `StringMatch()` calls with an empty message.
</div>

</div>

A further 11 [extra checks](concepts/standard-vs-extra.md) are opt-in through
`tfsproutx`, and 9 IDs are [retained but no longer report](reference/removed-checks.md).

## Find your way around

<div class="card-grid card-grid--plain" markdown="1">

<div class="card" markdown="1">
### Get started
- [What is tfsprout](what-is-tfsprout.md) — what it does, and what it deliberately does not
- [Install](install.md) — binaries, `go install`, Docker, version pinning
- [From tfproviderlint](migrating-from-tfproviderlint.md) — a drop-in replacement, in three steps
</div>

<div class="card" markdown="1">
### Usage
- [Running tfsprout](usage/running-locally.md) — selecting checks, `go vet` integration
- [Ignoring reports](usage/ignoring-reports.md) — `//lintignore:` scoping and adoption
- [Automated fixes](usage/automated-fixes.md) — what `-fix` really rewrites
- [CI integration](usage/ci-integration.md) — gating a build on exit status
- [Troubleshooting](usage/troubleshooting.md) — common errors and their causes
</div>

<div class="card" markdown="1">
### Concepts
- [How tfsprout works](concepts/how-it-works.md) — the `go/analysis` pipeline
- [Checks and categories](concepts/checks-and-categories.md) — what the ID prefixes mean
- [Standard vs extra](concepts/standard-vs-extra.md) — why `tfsproutx` exists
- [Scope and SDK support](concepts/scope-and-sdk-support.md) — which providers it can read
</div>

<div class="card" markdown="1">
### Reference
- [Check index](reference/checks.md) — all 94 IDs, each with its own page
- [CLI reference](reference/cli.md) — every flag both commands accept
- [Exit codes and output](reference/exit-codes-and-output.md) — what it prints and returns
- [Changelog](changelog.md) — release history
</div>

</div>

## Before you trust a clean run

!!! warning "tfsprout only reads Terraform Plugin SDK providers"

    Checks match on `helper/schema` types. A provider written against
    [`terraform-plugin-framework`](https://github.com/hashicorp/terraform-plugin-framework)
    contains none of them, so tfsprout exits `0` with no output — which looks
    like a clean run but is not one. It also says nothing about `.tf` files;
    that is the domain of `terraform validate` and `tflint`. See
    [Scope and SDK support](concepts/scope-and-sdk-support.md).

## Contributing

The [contributing guide](contributing/building.md) covers building from source,
[adding an analyzer](contributing/adding-an-analyzer.md) end to end,
[testing](contributing/testing.md) with `analysistest`, and
[building your own lint tool](contributing/custom-lint-tool.md) from tfsprout's
analyzers.
