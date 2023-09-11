FROM --platform=${BUILDPLATFORM} golang:1.19-alpine AS base
WORKDIR /src
ENV CGO_ENABLED=0
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=bind,source=go.sum,target=go.sum \
  --mount=type=bind,source=go.mod,target=go.mod \
  go mod download

FROM base AS build
ARG TARGETOS
ARG TARGETARCH
RUN --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/go/pkg/mod \
  --mount=type=bind,target=. \
  GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/echocors .

FROM scratch AS bin
COPY /assets /app/assets
COPY --from=build /out/echocors /app/echocors
WORKDIR /app
CMD ["/app/echocors"]