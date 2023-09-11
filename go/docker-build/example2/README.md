# Caching 1: Separate steps


```Dockerfile
# A snippet from the previous exapmle;
COPY . .
ARG TARGETOS 
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/example .
```

Every time you change the source code, the layer cache of `COPY . .`  will be invalidated. Even when `go.mod` and `go.sum` are not changed, `go build` will download the dependencies again.

We can fix this by download the dependencies as a separate step:

```Dockerfile
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/example .
```

