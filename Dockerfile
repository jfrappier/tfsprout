# tfsprout loads and type-checks provider packages via go/packages, which shells
# out to "go env" and "go list" at runtime. A Go toolchain is therefore required
# to run the tool, not just to build it, and its version caps which modules can
# be analyzed: "go list" fails on a module whose go directive is newer than the
# toolchain. Track the newest Go release this project supports.
FROM golang:1.27-bookworm
WORKDIR /src
COPY tfsprout /usr/bin/tfsprout
ENTRYPOINT ["/usr/bin/tfsprout"]
CMD ["./..."]
