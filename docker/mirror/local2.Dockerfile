FROM registry-test:5000/mirror/library/alpine:3.15.10
RUN ["cat", "/etc/alpine-release"]
