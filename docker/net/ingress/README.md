# Ingress traffic from the Docker Desktop host
... via published ports

## Prepare

```sh
make prepare
```

## Tests

Watch logs from echo1 and echo2 in separate terminals

```sh
docker logs -f echo1
docker logs -f echo2
```

Execute:

```sh
make experiment
```
