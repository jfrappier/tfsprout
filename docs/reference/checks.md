# Check index

Every check tfsprout ships, grouped by category. Each check ID links to its own page with flagged code, passing code, and suppression examples.

- **Standard checks** run in both `tfsprout` and `tfsproutx`.
- **Extra checks** — every ID beginning with `X` — run only in `tfsproutx`. See [Standard vs extra checks](../concepts/standard-vs-extra.md).

Prefix meanings are covered in [Checks and categories](../concepts/checks-and-categories.md). The `Type` column records the analysis technique; every check today is AST-based.

Three checks can rewrite your code under `-fix`: **R007**, **XR007**, and **XR008**. See [Automated fixes](../usage/automated-fixes.md).

Rows marked **REMOVED in v0.30.0** no longer report anything. Their IDs are retained permanently so existing ignore comments and CI flags keep working — see [Removed checks](removed-checks.md).

For the authoritative description of any check as the running binary sees it:

```shell
tfsprout help AT001
```

## Standard checks

Enabled by default in both commands.

### Acceptance test checks

Patterns in `_test.go` files that use `helper/resource` test harnesses. These are silent under `-test=false`.

| Check | Description | Type |
|---|---|---|
| [AT001](../../passes/AT001) | check for `TestCase` missing `CheckDestroy` | AST |
| [AT002](../../passes/AT002) | check for acceptance test function names including the word import | AST |
| [AT003](../../passes/AT003) | check for acceptance test function names missing an underscore | AST |
| [AT004](../../passes/AT004) | check for `TestStep` `Config` containing provider configuration | AST |
| [AT005](../../passes/AT005) | check for acceptance test function names missing `TestAcc` prefix | AST |
| [AT006](../../passes/AT006) | check for acceptance test functions containing multiple `resource.Test()` invocations | AST |
| [AT007](../../passes/AT007) | check for acceptance test functions containing multiple `resource.ParallelTest()` invocations | AST |
| [AT008](../../passes/AT008) | check for acceptance test function declaration `*testing.T` parameter naming | AST |
| [AT009](../../passes/AT009) | check for `acctest.RandStringFromCharSet()` calls that can be simplified to `acctest.RandString()` | AST |
| [AT010](../../passes/AT010) | check for `TestCase` including `IDRefreshName` implementation | AST |
| [AT011](../../passes/AT011) | check for `TestCase` including `IDRefreshIgnore` implementation without `IDRefreshName` | AST |
| [AT012](../../passes/AT012) | check for files containing multiple acceptance test function name prefixes | AST |

### Resource checks

Patterns in `schema.Resource` definitions and CRUD functions.

| Check | Description | Type |
|---|---|---|
| [R001](../../passes/R001) | check for `ResourceData.Set()` calls using complex key argument | AST |
| [R002](../../passes/R002) | check for `ResourceData.Set()` calls using `*` dereferences | AST |
| [R003](../../passes/R003) | check for `Resource` having `Exists` functions | AST |
| [R004](../../passes/R004) | check for `ResourceData.Set()` calls using incompatible value types | AST |
| [R005](../../passes/R005) | check for `ResourceData.HasChange()` calls that can be combined into one `HasChanges()` call | AST |
| [R006](../../passes/R006) | check for `RetryFunc` that omit retryable errors | AST |
| [R007](../../passes/R007) | check for deprecated `(schema.ResourceData).Partial` usage | AST |
| [R008](../../passes/R008) | **REMOVED in v0.30.0** check for deprecated `(schema.ResourceData).SetPartial` usage | AST |
| [R009](../../passes/R009) | check for Go panic usage | AST |
| [R010](../../passes/R010) | check for `(schema.ResourceData).GetChange` assignment which should use `(schema.ResourceData).Get` | AST |
| [R011](../../passes/R011) | check for `Resource` that configure `MigrateState` | AST |
| [R012](../../passes/R012) | check for data source `Resource` that configure `CustomizeDiff` | AST |
| [R013](../../passes/R013) | check for `map[string]*Resource` that resource names contain at least one underscore | AST |
| [R014](../../passes/R014) | check for `CreateFunc`, `CreateContextFunc`, `DeleteFunc`, `DeleteContextFunc`, `ReadFunc`, `ReadContextFunc`, `UpdateFunc`, and `UpdateContextFunc` parameter naming | AST |
| [R015](../../passes/R015) | check for `(*schema.ResourceData).SetId()` receiver method usage with unstable `resource.UniqueId()` value | AST |
| [R016](../../passes/R016) | check for `(*schema.ResourceData).SetId()` receiver method usage with unstable `resource.PrefixedUniqueId()` value | AST |
| [R017](../../passes/R017) | check for `(*schema.ResourceData).SetId()` receiver method usage with unstable `time.Now()` value | AST |
| [R018](../../passes/R018) | check for `time.Sleep()` function usage | AST |
| [R019](../../passes/R019) | check for `(*schema.ResourceData).HasChanges()` receiver method usage with many arguments | AST |

### Schema checks

Patterns in `schema.Schema` definitions and attribute maps. Several of these catch configurations that fail provider schema validation at runtime.

| Check | Description | Type |
|---|---|---|
| [S001](../../passes/S001) | check for `Schema` of `TypeList` or `TypeSet` missing `Elem` | AST |
| [S002](../../passes/S002) | check for `Schema` with both `Required` and `Optional` enabled | AST |
| [S003](../../passes/S003) | check for `Schema` with both `Required` and `Computed` enabled | AST |
| [S004](../../passes/S004) | check for `Schema` with both `Required` and `Default` configured | AST |
| [S005](../../passes/S005) | check for `Schema` with both `Computed` and `Default` configured | AST |
| [S006](../../passes/S006) | check for `Schema` of `TypeMap` missing `Elem` | AST |
| [S007](../../passes/S007) | check for `Schema` with both `Required` and `ConflictsWith` configured | AST |
| [S008](../../passes/S008) | check for `Schema` of `TypeList` or `TypeSet` with `Default` configured | AST |
| [S009](../../passes/S009) | check for `Schema` of `TypeList` or `TypeSet` with `ValidateFunc` configured | AST |
| [S010](../../passes/S010) | check for `Schema` of `Computed` only with `ValidateFunc` configured | AST |
| [S011](../../passes/S011) | check for `Schema` of `Computed` only with `DiffSuppressFunc` configured | AST |
| [S012](../../passes/S012) | check for `Schema` that `Type` is configured | AST |
| [S013](../../passes/S013) | check for `map[string]*Schema` that one of `Computed`, `Optional`, or `Required` is configured | AST |
| [S014](../../passes/S014) | check for `Schema` within `Elem` that `Computed`, `Optional`, and `Required` are not configured | AST |
| [S015](../../passes/S015) | check for `map[string]*Schema` that attribute names are valid | AST |
| [S016](../../passes/S016) | check for `Schema` that `Set` is only configured for `TypeSet` | AST |
| [S017](../../passes/S017) | check for `Schema` that `MaxItems` and `MinItems` are only configured for `TypeList`, `TypeMap`, or `TypeSet` | AST |
| [S018](../../passes/S018) | check for `Schema` that should use `TypeList` with `MaxItems: 1` | AST |
| [S019](../../passes/S019) | check for `Schema` that should omit `Computed`, `Optional`, or `Required` set to `false` | AST |
| [S020](../../passes/S020) | check for `Schema` of `Computed` only with `ForceNew` enabled | AST |
| [S021](../../passes/S021) | check for `Schema` that should omit `ComputedWhen` | AST |
| [S022](../../passes/S022) | check for `Schema` of `TypeMap` with invalid `Elem` of `*schema.Resource` | AST |
| [S023](../../passes/S023) | check for `Schema` that should omit `Elem` with incompatible `Type` | AST |
| [S024](../../passes/S024) | check for `Schema` that should omit `ForceNew` in data source schema attributes | AST |
| [S025](../../passes/S025) | check for `Schema` of `Computed` only with `AtLeastOneOf` configured | AST |
| [S026](../../passes/S026) | check for `Schema` of `Computed` only with `ConflictsWith` configured | AST |
| [S027](../../passes/S027) | check for `Schema` of `Computed` only with `Default` configured | AST |
| [S028](../../passes/S028) | check for `Schema` of `Computed` only with `DefaultFunc` configured | AST |
| [S029](../../passes/S029) | check for `Schema` of `Computed` only with `ExactlyOneOf` configured | AST |
| [S030](../../passes/S030) | check for `Schema` of `Computed` only with `InputDefault` configured | AST |
| [S031](../../passes/S031) | check for `Schema` of `Computed` only with `MaxItems` configured | AST |
| [S032](../../passes/S032) | check for `Schema` of `Computed` only with `MinItems` configured | AST |
| [S033](../../passes/S033) | check for `Schema` of `Computed` only with `StateFunc` configured | AST |
| [S034](../../passes/S034) | **REMOVED in v0.30.0** check for `Schema` that configure `PromoteSingle` | AST |
| [S035](../../passes/S035) | check for `Schema` with invalid `AtLeastOneOf` attribute references | AST |
| [S036](../../passes/S036) | check for `Schema` with invalid `ConflictsWith` attribute references | AST |
| [S037](../../passes/S037) | check for `Schema` with invalid `ExactlyOneOf` attribute references | AST |

### Validation checks

Patterns in `SchemaValidateFunc` implementations and `helper/validation` usage. Most report hand-rolled logic that duplicates a built-in validator.

| Check | Description | Type |
|---|---|---|
| [V001](../../passes/V001) | check for custom `SchemaValidateFunc` that implement `validation.StringMatch()` or `validation.StringDoesNotMatch()` | AST |
| [V002](../../passes/V002) | **REMOVED in v0.30.0** check for deprecated `CIDRNetwork` validation function usage | AST |
| [V003](../../passes/V003) | **REMOVED in v0.30.0** check for deprecated `IPRange` validation function usage | AST |
| [V004](../../passes/V004) | **REMOVED in v0.30.0** check for deprecated `SingleIP` validation function usage | AST |
| [V005](../../passes/V005) | **REMOVED in v0.30.0** check for deprecated `ValidateJsonString` validation function usage | AST |
| [V006](../../passes/V006) | **REMOVED in v0.30.0** check for deprecated `ValidateListUniqueStrings` validation function usage | AST |
| [V007](../../passes/V007) | **REMOVED in v0.30.0** check for deprecated `ValidateRegexp` validation function usage | AST |
| [V008](../../passes/V008) | **REMOVED in v0.30.0** check for deprecated `ValidateRFC3339TimeString` validation function usage | AST |
| [V009](../../passes/V009) | check for `validation.StringMatch()` call with empty message argument | AST |
| [V010](../../passes/V010) | check for `validation.StringDoesNotMatch()` call with empty message argument | AST |
| [V011](../../passes/V011) | check for custom `SchemaValidateFunc` that implement `validation.StringLenBetween()` | AST |
| [V012](../../passes/V012) | check for custom `SchemaValidateFunc` that implement `validation.IntAtLeast()`, `validation.IntAtMost()`, or `validation.IntBetween()` | AST |
| [V013](../../passes/V013) | check for custom `SchemaValidateFunc` that implement `validation.StringInSlice()` or `validation.StringNotInSlice()` | AST |
| [V014](../../passes/V014) | check for custom `SchemaValidateFunc` that implement `validation.IntInSlice()` or `validation.IntNotInSlice()` | AST |

## Extra checks

Not included in `tfsprout`. Access them via `tfsproutx`, or by [building a custom lint tool](../contributing/custom-lint-tool.md). These generally represent advanced Terraform Plugin SDK functionality that is not appropriate for every provider.

### Extra acceptance test checks

| Check | Description | Type |
|---|---|---|
| [XAT001](../../xpasses/XAT001) | check for `TestCase` missing `ErrorCheck` | AST |

### Extra resource checks

| Check | Description | Type |
|---|---|---|
| [XR001](../../xpasses/XR001) | check for usage of `ResourceData.GetOkExists()` calls | AST |
| [XR002](../../xpasses/XR002) | check for `Resource` that should implement `Importer` | AST |
| [XR003](../../xpasses/XR003) | check for `Resource` that should implement `Timeouts` | AST |
| [XR004](../../xpasses/XR004) | check for `ResourceData.Set()` calls that should implement error checking with complex values | AST |
| [XR005](../../xpasses/XR005) | check for `Resource` that `Description` is configured | AST |
| [XR006](../../xpasses/XR006) | check for `Resource` that implements `Timeouts` for missing `Create`, `Delete`, `Read`, or `Update` implementation | AST |
| [XR007](../../xpasses/XR007) | check for `os/exec.Command` usage | AST |
| [XR008](../../xpasses/XR008) | check for `os/exec.CommandContext` usage | AST |

### Extra schema checks

| Check | Description | Type |
|---|---|---|
| [XS001](../../xpasses/XS001) | check for `map[string]*Schema` that `Description` is configured | AST |
| [XS002](../../xpasses/XS002) | check for `map[string]*Schema` that keys are in alphabetical order | AST |
