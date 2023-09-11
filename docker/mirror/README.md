# How registry mirrors work

## Setup

Prepare functions (guaranteed to work on bash) for testing
```
source shcmd
```

Setup your local registry for testing
```
setup_local_reg
```

Create a builder
```
create_builder buildkitd.toml
```

## Facts found

- `docker pull` can resolve `localhost` but not `registry-test`.

## Resources

https://docs.docker.com/build/buildkit/configure/
