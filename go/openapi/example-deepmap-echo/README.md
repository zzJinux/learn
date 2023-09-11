## Getting started
Install `oapi-codegen`
```
go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
```

Generate
```
oapi-codegen  -package petstore petstore-expanded.yaml > petstore.gen.go
```
