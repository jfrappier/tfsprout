FROM golang:1.22-bookworm
WORKDIR /src
COPY tfsprout /usr/bin/tfsprout
ENTRYPOINT ["/usr/bin/tfsprout"]
CMD ["./..."]
