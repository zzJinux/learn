#!/usr/bin/env bash

# ipaddr bound to default interface
desktop_eth=$(ipconfig getifaddr $(route -n get default | grep interface: | awk '{print $2}'))

echo1_eth=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' echo1)
echo2_eth=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' echo2)
defaultbrdg_gw=$(docker network inspect bridge -f '{{range .IPAM.Config}}{{.Gateway}}{{end}}')
rootns_gw=$(docker run --privileged --pid=host -it --rm alpine:3.18.3 nsenter -t 1 -m -u -n -i -- ip route | grep default | awk '{print $3}')

echo Docker Desktop host ipaddr: $desktop_eth
echo echo1 container ipaddr: $echo1_eth
echo echo2 container ipaddr: $echo2_eth
echo 'default bridge gateway (docker0):' $defaultbrdg_gw
echo 'root namespace gateway:' $rootns_gw
echo host.docker.internal: \
  $(docker run -it --rm mbentley/bind-tools dig +short host.docker.internal)
echo gateway.docker.internal: \
  $(docker run -it --rm mbentley/bind-tools dig +short gateway.docker.internal)

curl_opts=(--no-progress-meter --connect-timeout 1)

desktop_curl() {
  curl "${curl_opts[@]}" -A "desktop_curl $*" "$@"
}

default_bridge_curl() {
  docker run -it --rm local/alpinecurl curl "${curl_opts[@]}" -A "default_bridge_curl $*" "$@"
}

user_bridge_curl() {
  docker run -it --rm --net learn-docker local/alpinecurl curl "${curl_opts[@]}" -A "user_bridge_curl $*" "$@"
}

user_intbridge_curl() {
  docker run -it --rm --net learn-docker-internal local/alpinecurl curl "${curl_opts[@]}" -A "user_intbridge_curl $*" "$@"
}

hostnet_curl() {
  docker run -it --rm --net host local/alpinecurl curl "${curl_opts[@]}" -A "hostnet_curl $*" "$@"
}

rootnet_curl() {
  docker run --privileged --pid=host -it --rm local/alpinecurl nsenter -t 1 -u -n -i -- curl "${curl_opts[@]}" -A "rootnet_curl $*" "$@"
}

print_and_exec() {
  echo "$@"
  "$@" 2>&1
  echo
}

echo

echo === From the Docker Desktop host

print_and_exec desktop_curl localhost:9001

print_and_exec desktop_curl localhost:9002

print_and_exec desktop_curl $desktop_eth:9001

print_and_exec desktop_curl $desktop_eth:9002

echo === From the default bridge
echo

print_and_exec default_bridge_curl localhost:9001

print_and_exec default_bridge_curl localhost:9002

print_and_exec default_bridge_curl $echo1_eth:5678

print_and_exec default_bridge_curl $echo2_eth:5678

print_and_exec default_bridge_curl host.docker.internal:9001

print_and_exec default_bridge_curl host.docker.internal:9002

echo === From an user bridge
echo

print_and_exec user_bridge_curl localhost:9001

print_and_exec user_bridge_curl localhost:9002

print_and_exec user_bridge_curl $echo1_eth:5678

print_and_exec user_bridge_curl $echo2_eth:5678

print_and_exec user_bridge_curl host.docker.internal:9001

print_and_exec user_bridge_curl host.docker.internal:9002

echo === From an user internal-bridge
echo

print_and_exec user_intbridge_curl localhost:9001

print_and_exec user_intbridge_curl localhost:9002

print_and_exec user_intbridge_curl $echo1_eth:5678

print_and_exec user_intbridge_curl $echo2_eth:5678

print_and_exec user_intbridge_curl host.docker.internal:9001

print_and_exec user_intbridge_curl host.docker.internal:9002

echo === From the \"host\" network
echo

print_and_exec hostnet_curl localhost:9001

print_and_exec hostnet_curl localhost:9002

print_and_exec hostnet_curl $echo1_eth:5678

print_and_exec hostnet_curl $echo2_eth:5678

print_and_exec hostnet_curl host.docker.internal:9001

print_and_exec hostnet_curl host.docker.internal:9002

echo === From the root ns network
echo

print_and_exec rootnet_curl localhost:9001

print_and_exec rootnet_curl localhost:9002

print_and_exec rootnet_curl $echo1_eth:5678

print_and_exec rootnet_curl $echo2_eth:5678

print_and_exec rootnet_curl host.docker.internal:9001

print_and_exec rootnet_curl host.docker.internal:9002
