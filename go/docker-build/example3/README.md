# Caching 2: Cache mounts

The `go` command stores downloaded modules files to the `$GOMODCACHE` (default: `$GOPATH/pkg/mod`) directory,
and caches build outputs the $GOCACHE (default: `<platform-dependent user cache directory>/go-build`) directory.
For Linux, `GOMODCACHE` is usually `/root/go/pkg/mod` and `GOCACHE` is usually `/root/.cache/go-build`.

Cache mounts let you specifiy a persistent package cache to be used during builds.
Having a persistent cache means that even if you rebuild a layer, you only download new or changed packages.
Contents of the cache directories persists between builder invocations.

Let's verify cache mounts work as expected. First, we need a clean slate:

```
docker builder prune -af
```

Build first time with the log output:

```
docker build --target=bin --progress=plain . 2> log1.txt
```

Modify the modules file to upgrade an existing module or add a new module. For example: 

```
go get github.com/go-chi/chi/v5@latest
```

Run another build, and again with the log output:

```
docker build --target=bin --progress=plain . 2> log2.txt
```

Compare two logs.