FROM --platform=${BUILDPLATFORM} golang:1.21.1-alpine AS build
# TARGETPLATFORM == TARGETOS/TARGETARCH/TARGETVARIANT
# BUILDPLATFORM == BUILDOS/BUILDARCH/BUILDVARIANT
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/example .

FROM scratch AS bin-unix
ARG TARGETOS
ARG TARGETARCH
COPY --from=build /out/example /example-${TARGETOS}-${TARGETARCH}

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
ARG TARGETARCH
COPY --from=build /out/example /example-${TARGETARCH}.exe

FROM bin-${TARGETOS} AS bin