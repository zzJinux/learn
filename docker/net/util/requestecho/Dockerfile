FROM golang:1.21-alpine AS builder
WORKDIR /src
RUN --mount=target=. \
  go build -o /out/main main.go
FROM scratch
COPY --from=builder /out/main /main
ENTRYPOINT ["/main"]
