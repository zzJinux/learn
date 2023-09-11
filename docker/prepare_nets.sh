#!/usr/bin/env bash
if ! docker network inspect learn-docker &>/dev/null; then
  docker network create learn-docker
fi
if ! docker network inspect learn-docker-internal &>/dev/null; then
  docker network create --internal learn-docker-internal
fi
