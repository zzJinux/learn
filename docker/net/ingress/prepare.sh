#!/usr/bin/env bash

# Listen on 0.0.0.0:9001
docker run -d --rm --name echo1 -p 9001:5678 local/requestecho
# Listen on 127.0.0.1:9002
docker run -d --rm --name echo2 -p 127.0.0.1:9002:5678 local/requestecho
