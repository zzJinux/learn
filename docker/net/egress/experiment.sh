#!/usr/bin/env bash
. vars

print_and_exec() {
  echo "$@"
  "$@"
  echo
}

echo === From the default bridge
print_and_exec docker run --rm $IMAGE_NAME

echo === From an user bridge
print_and_exec docker run --rm --net learn-docker $IMAGE_NAME

echo === From an user internal bridge
print_and_exec docker run --rm --net learn-docker-internal $IMAGE_NAME

echo === From the \"host\" network
print_and_exec docker run --rm --net host $IMAGE_NAME

echo === From the root ns network
print_and_exec docker run --privileged --pid=host --rm $IMAGE_NAME nsenter -t 1 -u -n -i -- sh -c '
  echo "host.docker.internal="$(dig +retry=0 +timeout=1 +short host.docker.internal); \
  echo 'gateway.docker.internal='$(dig +retry=0 +timeout=1 +short gateway.docker.internal); \
  echo curl localhost:5678; \
  curl --no-progress-meter --connect-timeout 1 localhost:5678 2>&1; \
  echo curl host.docker.internal:5678; \
  curl --no-progress-meter --connect-timeout 1 host.docker.internal:5678 2>&1; \
  echo curl gateway.docker.internal:5678; \
  curl --no-progress-meter --connect-timeout 1 gateway.docker.internal:5678 2>&1;'
