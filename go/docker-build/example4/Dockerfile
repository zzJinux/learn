# syntax=docker/dockerfile:1

FROM --platform=${BUILDPLATFORM} golang:1.21.1-alpine AS base
WORKDIR /src
ENV CGO_ENABLED=0
# -x flag is to show the download process
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=bind,source=go.sum,target=go.sum \
  --mount=type=bind,source=go.mod,target=go.mod \
  go mod download -x

FROM base AS build
ARG TARGETOS
ARG TARGETARCH
RUN --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/go/pkg/mod \
  --mount=type=bind,target=. \
  GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/example .

FROM base AS unittest
# "go test -v .", running in package list mode, caches successful package test results
# so we mount the cache directory so that Go only run tests if there have been changes.
RUN --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/go/pkg/mod \
  --mount=type=bind,target=. \
  go test -v .

FROM golangci/golangci-lint:v1.54-alpine AS lint-base

FROM base AS lint
RUN --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/root/.cache/golangci-lint \
  --mount=type=bind,target=. \
  --mount=from=lint-base,src=/usr/bin/golangci-lint,target=/usr/bin/golangci-lint \
  golangci-lint run --timeout 10m0s ./...

FROM scratch AS bin-unix
COPY --from=build /out/example /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /out/example /example.exe

FROM bin-${TARGETOS} AS bin
