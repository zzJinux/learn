#!/usr/bin/env bash
if [ -f requestecho.pid ]; then
  pid=$(cat requestecho.pid)
  kill -- -$(ps -o pgid= $pid)
  rm requestecho.pid requestecho.log
fi
