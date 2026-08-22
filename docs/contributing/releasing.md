# Releasing

For maintainers. Releases are built by [GoReleaser](https://goreleaser.com/) from `.goreleaser.yml` and published by the `release` workflow.

## What gets built

Two binaries, `tfsprout` and `tfsproutx`, each cross-compiled with `CGO_ENABLED=0` for:

| GOOS | GOARCH |
|---|---|
| `darwin` | `amd64`, `arm64`, `386` |
| `linux` | `amd64`, `arm64`, `386` |
| `windows` | `amd64`, `arm64`, `386` |

Archives are `.tar.gz`, except Windows which ships `.zip`.

## Version injection

Version information is baked in at link time, overriding the defaults in `version/version.go`:

```
-X github.com/jfrappier/tfsprout/version.Version={{.Version}}
-X github.com/jfrappier/tfsprout/version.VersionPrerelease=
-X github.com/jfrappier/tfsprout/version.GitCommit={{.ShortCommit}}
```

`VersionPrerelease` is deliberately cleared, so a released binary reports a clean version while a locally built one reports the in-repo default with its `dev` marker. That difference is how you tell the two apart with `tfsprout -V=full`.

Because the version comes from the tag, **the constants in `version/version.go` are not the source of truth for a release** — do not bump them expecting the release to follow.

## Release steps

1. **Update `CHANGELOG.md`.** Add the new version's entry **at the top of the file**, under a `# vX.Y.Z` header. The release workflow extracts everything between the first `# v` header and the next one into `RELEASE-NOTES.md`, so the topmost section becomes the GitHub release body verbatim. Get it wrong and the release ships with the previous version's notes.
2. **Commit** the changelog on `main`.
3. **Tag** the release:
   ```shell
   git tag v0.1.2
   git push origin v0.1.2
   ```
4. **Watch the workflow.** Pushing the tag triggers `.github/workflows/release.yaml`. It re-runs `go mod tidy` verification and the full test suite, and only then runs GoReleaser — a failing test blocks the release.

   The workflow matches tags of the form `v[0-9]+.[0-9]+.[0-9]+` only. Pre-release tags such as `v0.2.0-rc1` will **not** trigger it.
5. **Verify** the release page has archives for every platform, and that the milestone was closed — `milestones.close` is enabled in the config.

## Changelog

GoReleaser's automatic changelog is disabled (`changelog.disable: true`). Release notes are generated from `CHANGELOG.md` by this step in the workflow:

```shell
awk '/^# v/{c++; next} c==1{print} c>=2{exit}' CHANGELOG.md > RELEASE-NOTES.md
```

That is, the body of the **first** `# v...` section, up to the next one. Two consequences:

- The changelog must be written and committed **before** tagging. There is no way to fix the release body afterwards except by editing the GitHub release by hand.
- New entries go at the top. Appending to the bottom silently publishes the wrong notes.

## Pre-release validation

The `goreleaser` job on every pull request already runs:

```shell
goreleaser check
goreleaser build --snapshot
```

So configuration errors and cross-compilation failures surface before you tag. To reproduce locally:

```shell
goreleaser build --snapshot --clean
```

## Docker

The `Dockerfile` copies a prebuilt `tfsprout` binary from the build context rather than compiling one, so it expects a binary to exist beside it at build time.

Note its base image is `golang:1.22-bookworm`, which predates the Go 1.25 floor in `go.mod`. This is harmless today — the image only carries an already-compiled static binary and never invokes the toolchain — but it is misleading, and worth aligning the next time the file is touched.
