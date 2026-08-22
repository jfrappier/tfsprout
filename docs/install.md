# Install

tfsprout ships two binaries. Install whichever matches the check set you want; installing both is fine.

| Binary | Checks |
|---|---|
| `tfsprout` | Standard checks |
| `tfsproutx` | Standard checks plus [extra checks](concepts/standard-vs-extra.md) |

## Release binaries

Prebuilt binaries for Linux, macOS, and Windows on `amd64`, `arm64`, and `386` are attached to each [release](https://github.com/jfrappier/tfsprout/releases).

Download the archive for your platform, extract it, and place the binary somewhere on your `PATH`.

## go install

To build from source into your `$GOBIN` directory (typically `$GOPATH/bin`):

```shell
go install github.com/jfrappier/tfsprout/cmd/tfsprout@latest
```

For the command that includes extra checks:

```shell
go install github.com/jfrappier/tfsprout/cmd/tfsproutx@latest
```

This requires **Go 1.25 or later** — see [Scope and SDK support](concepts/scope-and-sdk-support.md#go-version-support).

## Pinning a version

`@latest` is convenient locally and a liability in CI, where an unpinned linter turns an upstream release into an unrelated build failure. Pin explicitly:

```shell
go install github.com/jfrappier/tfsprout/cmd/tfsprout@v0.1.1
```

If your provider already pins its tooling through a `tools.go` file, add the import there instead so the version lives in `go.mod`:

```go
//go:build tools

package tools

import (
    _ "github.com/jfrappier/tfsprout/cmd/tfsprout"
)
```

Then `go mod tidy`, and install with `go install github.com/jfrappier/tfsprout/cmd/tfsprout`.

## Docker

The repository includes a `Dockerfile` that wraps a prebuilt binary:

```dockerfile
FROM golang:1.22-bookworm
WORKDIR /src
COPY tfsprout /usr/bin/tfsprout
ENTRYPOINT ["/usr/bin/tfsprout"]
CMD ["./..."]
```

It expects a `tfsprout` binary in the build context rather than compiling one, so build a binary first (or extract one from a release archive) before `docker build`. Mount your provider at `/src`:

```shell
docker run --rm -v "$PWD:/src" tfsprout ./...
```

## GitHub Action

For CI, the [tfsprout-github-action](https://github.com/jfrappier/tfsprout-github-action) handles installation for you. See [CI integration](usage/ci-integration.md).

## Verifying the install

```shell
tfsprout -V
```

Use `-V=full` for the fully qualified version string, which is what to include when filing an issue.

## Next steps

- [Running tfsprout](usage/running-locally.md)
- [Migrating from tfproviderlint](migrating-from-tfproviderlint.md) if you are replacing an existing setup
