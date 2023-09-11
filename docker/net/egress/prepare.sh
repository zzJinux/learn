#!/usr/bin/env bash
. vars

make clean
go run requestecho.go &>requestecho.log &
echo $! > requestecho.pid
sleep 1

docker build --no-cache --progress plain -t $IMAGE_NAME - < Dockerfile
