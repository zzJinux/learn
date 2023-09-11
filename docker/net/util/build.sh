#!/usr/bin/env bash

docker build -f requestecho/Dockerfile -t local/requestecho requestecho
docker build -f alpinecurl/Dockerfile -t local/alpinecurl alpinecurl
