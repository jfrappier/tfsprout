# Check index

Every check tfsprout ships, grouped by category. Each check ID links to its own page with flagged code, passing code, and suppression examples.

- **Standard checks** run in both `tfsprout` and `tfsproutx`.
- **Extra checks** — every ID beginning with `X` — run only in `tfsproutx`. See [Standard vs extra checks](../concepts/standard-vs-extra.md).

Prefix meanings are covered in [Checks and categories](../concepts/checks-and-categories.md). The `Type` column records the analysis technique; every check today is AST-based.

Three checks can rewrite your code under `-fix`: **R007**, **XR007**, and **XR008**. See [Automated fixes](../usage/automated-fixes.md).

Rows marked **REMOVED** (tfproviderlint v0.30.0) no longer report anything. Their IDs are retained permanently so existing ignore comments and CI flags keep working — see [Removed checks](removed-checks.md).

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
| [AT001](../checks/AT001.md) | check for `TestCase` missing `CheckDestroy` | AST |
| [AT002](../checks/AT002.md) | check for acceptance test function names including the word import | AST |
| [AT003](../checks/AT003.md) | check for acceptance test function names missing an underscore | AST |
| [AT004](../checks/AT004.md) | check for `TestStep` `Config` containing provider configuration | AST |
| [AT005](../checks/AT005.md) | check for acceptance test function names missing `TestAcc` prefix | AST |
| [AT006](../checks/AT006.md) | check for acceptance test functions containing multiple `resource.Test()` invocations | AST |
| [AT007](../checks/AT007.md) | check for acceptance test functions containing multiple `resource.ParallelTest()` invocations | AST |
| [AT008](../checks/AT008.md) | check for acceptance test function declaration `*testing.T` parameter naming | AST |
| [AT009](../checks/AT009.md) | check for `acctest.RandStringFromCharSet()` calls that can be simplified to `acctest.RandString()` | AST |
| [AT010](../checks/AT010.md) | check for `TestCase` including `IDRefreshName` implementation | AST |
| [AT011](../checks/AT011.md) | check for `TestCase` including `IDRefreshIgnore` implementation without `IDRefreshName` | AST |
| [AT012](../checks/AT012.md) | check for files containing multiple acceptance test function name prefixes | AST |

### Resource checks

Patterns in `schema.Resource` definitions and CRUD functions.

| Check | Description | Type |
|---|---|---|
| [R001](../checks/R001.md) | check for `ResourceData.Set()` calls using complex key argument | AST |
| [R002](../checks/R002.md) | check for `ResourceData.Set()` calls using `*` dereferences | AST |
| [R003](../checks/R003.md) | check for `Resource` having `Exists` functions | AST |
| [R004](../checks/R004.md) | check for `ResourceData.Set()` calls using incompatible value types | AST |
| [R005](../checks/R005.md) | check for `ResourceData.HasChange()` calls that can be combined into one `HasChanges()` call | AST |
| [R006](../checks/R006.md) | check for `RetryFunc` that omit retryable errors | AST |
| [R007](../checks/R007.md) | check for deprecated `(schema.ResourceData).Partial` usage | AST |
| [R008](../checks/R008.md) | **REMOVED** (tfproviderlint v0.30.0) check for deprecated `(schema.ResourceData).SetPartial` usage | AST |
| [R009](../checks/R009.md) | check for Go panic usage | AST |
| [R010](../checks/R010.md) | check for `(schema.ResourceData).GetChange` assignment which should use `(schema.ResourceData).Get` | AST |
| [R011](../checks/R011.md) | check for `Resource` that configure `MigrateState` | AST |
| [R012](../checks/R012.md) | check for data source `Resource` that configure `CustomizeDiff` | AST |
| [R013](../checks/R013.md) | check for `map[string]*Resource` that resource names contain at least one underscore | AST |
| [R014](../checks/R014.md) | check for `CreateFunc`, `CreateContextFunc`, `DeleteFunc`, `DeleteContextFunc`, `ReadFunc`, `ReadContextFunc`, `UpdateFunc`, and `UpdateContextFunc` parameter naming | AST |
| [R015](../checks/R015.md) | check for `(*schema.ResourceData).SetId()` receiver method usage with unstable `resource.UniqueId()` value | AST |
| [R016](../checks/R016.md) | check for `(*schema.ResourceData).SetId()` receiver method usage with unstable `resource.PrefixedUniqueId()` value | AST |
| [R017](../checks/R017.md) | check for `(*schema.ResourceData).SetId()` receiver method usage with unstable `time.Now()` value | AST |
| [R018](../checks/R018.md) | check for `time.Sleep()` function usage | AST |
| [R019](../checks/R019.md) | check for `(*schema.ResourceData).HasChanges()` receiver method usage with many arguments | AST |

### Schema checks

Patterns in `schema.Schema` definitions and attribute maps. Several of these catch configurations that fail provider schema validation at runtime.

| Check | Description | Type |
|---|---|---|
| [S001](../checks/S001.md) | check for `Schema` of `TypeList` or `TypeSet` missing `Elem` | AST |
| [S002](../checks/S002.md) | check for `Schema` with both `Required` and `Optional` enabled | AST |
| [S003](../checks/S003.md) | check for `Schema` with both `Required` and `Computed` enabled | AST |
| [S004](../checks/S004.md) | check for `Schema` with both `Required` and `Default` configured | AST |
| [S005](../checks/S005.md) | check for `Schema` with both `Computed` and `Default` configured | AST |
| [S006](../checks/S006.md) | check for `Schema` of `TypeMap` missing `Elem` | AST |
| [S007](../checks/S007.md) | check for `Schema` with both `Required` and `ConflictsWith` configured | AST |
| [S008](../checks/S008.md) | check for `Schema` of `TypeList` or `TypeSet` with `Default` configured | AST |
| [S009](../checks/S009.md) | check for `Schema` of `TypeList` or `TypeSet` with `ValidateFunc` or `ValidateDiagFunc` configured | AST |
| [S010](../checks/S010.md) | check for `Schema` of `Computed` only with `ValidateFunc` configured | AST |
| [S011](../checks/S011.md) | check for `Schema` of `Computed` only with `DiffSuppressFunc` configured | AST |
| [S012](../checks/S012.md) | check for `Schema` that `Type` is configured | AST |
| [S013](../checks/S013.md) | check for `map[string]*Schema` that one of `Computed`, `Optional`, or `Required` is configured | AST |
| [S014](../checks/S014.md) | check for `Schema` within `Elem` that `Computed`, `Optional`, and `Required` are not configured | AST |
| [S015](../checks/S015.md) | check for `map[string]*Schema` that attribute names are valid | AST |
| [S016](../checks/S016.md) | check for `Schema` that `Set` is only configured for `TypeSet` | AST |
| [S017](../checks/S017.md) | check for `Schema` that `MaxItems` and `MinItems` are only configured for `TypeList`, `TypeMap`, or `TypeSet` | AST |
| [S018](../checks/S018.md) | check for `Schema` that should use `TypeList` with `MaxItems: 1` | AST |
| [S019](../checks/S019.md) | check for `Schema` that should omit `Computed`, `Optional`, or `Required` set to `false` | AST |
| [S020](../checks/S020.md) | check for `Schema` of `Computed` only with `ForceNew` enabled | AST |
| [S021](../checks/S021.md) | check for `Schema` that should omit `ComputedWhen` | AST |
| [S022](../checks/S022.md) | check for `Schema` of `TypeMap` with invalid `Elem` of `*schema.Resource` | AST |
| [S023](../checks/S023.md) | check for `Schema` that should omit `Elem` with incompatible `Type` | AST |
| [S024](../checks/S024.md) | check for `Schema` that should omit `ForceNew` in data source schema attributes | AST |
| [S025](../checks/S025.md) | check for `Schema` of `Computed` only with `AtLeastOneOf` configured | AST |
| [S026](../checks/S026.md) | check for `Schema` of `Computed` only with `ConflictsWith` configured | AST |
| [S027](../checks/S027.md) | check for `Schema` of `Computed` only with `Default` configured | AST |
| [S028](../checks/S028.md) | check for `Schema` of `Computed` only with `DefaultFunc` configured | AST |
| [S029](../checks/S029.md) | check for `Schema` of `Computed` only with `ExactlyOneOf` configured | AST |
| [S030](../checks/S030.md) | check for `Schema` of `Computed` only with `InputDefault` configured | AST |
| [S031](../checks/S031.md) | check for `Schema` of `Computed` only with `MaxItems` configured | AST |
| [S032](../checks/S032.md) | check for `Schema` of `Computed` only with `MinItems` configured | AST |
| [S033](../checks/S033.md) | check for `Schema` of `Computed` only with `StateFunc` configured | AST |
| [S034](../checks/S034.md) | **REMOVED** (tfproviderlint v0.30.0) check for `Schema` that configure `PromoteSingle` | AST |
| [S035](../checks/S035.md) | check for `Schema` with invalid `AtLeastOneOf` attribute references | AST |
| [S036](../checks/S036.md) | check for `Schema` with invalid `ConflictsWith` attribute references | AST |
| [S037](../checks/S037.md) | check for `Schema` with invalid `ExactlyOneOf` attribute references | AST |
| [S038](../checks/S038.md) | check for `Schema` with both `ValidateFunc` and `ValidateDiagFunc` configured | AST |
| [S039](../checks/S039.md) | check for `Schema` with invalid resource identity configuration | AST |

### Validation checks

Patterns in `SchemaValidateFunc` implementations and `helper/validation` usage. Most report hand-rolled logic that duplicates a built-in validator.

| Check | Description | Type |
|---|---|---|
| [V001](../checks/V001.md) | check for custom `SchemaValidateFunc` that implement `validation.StringMatch()` or `validation.StringDoesNotMatch()` | AST |
| [V002](../checks/V002.md) | **REMOVED** (tfproviderlint v0.30.0) check for deprecated `CIDRNetwork` validation function usage | AST |
| [V003](../checks/V003.md) | **REMOVED** (tfproviderlint v0.30.0) check for deprecated `IPRange` validation function usage | AST |
| [V004](../checks/V004.md) | **REMOVED** (tfproviderlint v0.30.0) check for deprecated `SingleIP` validation function usage | AST |
| [V005](../checks/V005.md) | **REMOVED** (tfproviderlint v0.30.0) check for deprecated `ValidateJsonString` validation function usage | AST |
| [V006](../checks/V006.md) | **REMOVED** (tfproviderlint v0.30.0) check for deprecated `ValidateListUniqueStrings` validation function usage | AST |
| [V007](../checks/V007.md) | **REMOVED** (tfproviderlint v0.30.0) check for deprecated `ValidateRegexp` validation function usage | AST |
| [V008](../checks/V008.md) | **REMOVED** (tfproviderlint v0.30.0) check for deprecated `ValidateRFC3339TimeString` validation function usage | AST |
| [V009](../checks/V009.md) | check for `validation.StringMatch()` call with empty message argument | AST |
| [V010](../checks/V010.md) | check for `validation.StringDoesNotMatch()` call with empty message argument | AST |
| [V011](../checks/V011.md) | check for custom `SchemaValidateFunc` that implement `validation.StringLenBetween()` | AST |
| [V012](../checks/V012.md) | check for custom `SchemaValidateFunc` that implement `validation.IntAtLeast()`, `validation.IntAtMost()`, or `validation.IntBetween()` | AST |
| [V013](../checks/V013.md) | check for custom `SchemaValidateFunc` that implement `validation.StringInSlice()` or `validation.StringNotInSlice()` | AST |
| [V014](../checks/V014.md) | check for custom `SchemaValidateFunc` that implement `validation.IntInSlice()` or `validation.IntNotInSlice()` | AST |

## Extra checks

Not included in `tfsprout`. Access them via `tfsproutx`, or by [building a custom lint tool](../contributing/custom-lint-tool.md). These generally represent advanced Terraform Plugin SDK functionality that is not appropriate for every provider.

### Extra acceptance test checks

| Check | Description | Type |
|---|---|---|
| [XAT001](../checks/XAT001.md) | check for `TestCase` missing `ErrorCheck` | AST |

### Extra resource checks

| Check | Description | Type |
|---|---|---|
| [XR001](../checks/XR001.md) | check for usage of `ResourceData.GetOkExists()` calls | AST |
| [XR002](../checks/XR002.md) | check for `Resource` that should implement `Importer` | AST |
| [XR003](../checks/XR003.md) | check for `Resource` that should implement `Timeouts` | AST |
| [XR004](../checks/XR004.md) | check for `ResourceData.Set()` calls that should implement error checking with complex values | AST |
| [XR005](../checks/XR005.md) | check for `Resource` that `Description` is configured | AST |
| [XR006](../checks/XR006.md) | check for `Resource` that implements `Timeouts` for missing `Create`, `Delete`, `Read`, or `Update` implementation | AST |
| [XR007](../checks/XR007.md) | check for `os/exec.Command` usage | AST |
| [XR008](../checks/XR008.md) | check for `os/exec.CommandContext` usage | AST |

### Extra schema checks

| Check | Description | Type |
|---|---|---|
| [XS001](../checks/XS001.md) | check for `map[string]*Schema` that `Description` is configured | AST |
| [XS002](../checks/XS002.md) | check for `map[string]*Schema` that keys are in alphabetical order | AST |
