# syntax=docker/dockerfile:1

FROM --platform=${BUILDPLATFORM} golang:1.21.1-alpine AS base
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
# -x flag is to show the download process
RUN --mount=type=cache,target=/go/pkg/mod \
  go mod download -x
COPY . .

FROM base AS build
ARG TARGETOS
ARG TARGETARCH
RUN --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/go/pkg/mod \
  GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/example .

FROM base AS unittest
# "go test -v .", running in package list mode, caches successful package test results
# so we mount the cache directory so that Go only run tests if there have been changes.
RUN --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/go/pkg/mod \
  go test -v .

FROM scratch AS bin-unix
COPY --from=build /out/example /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /out/example /example.exe

FROM bin-${TARGETOS} AS bin
