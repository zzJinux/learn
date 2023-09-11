FROM nicolaka/netshoot:v0.11
RUN echo "host.docker.internal="$(dig +retry=0 +timeout=1 +short host.docker.internal); \
  echo 'gateway.docker.internal='$(dig +retry=0 +timeout=1 +short gateway.docker.internal); \
  echo curl localhost:5678; \
  curl --no-progress-meter --connect-timeout 1 localhost:5678 2>&1; \
  echo curl host.docker.internal:5678; \
  curl --no-progress-meter --connect-timeout 1 host.docker.internal:5678 2>&1; \
  echo curl gateway.docker.internal:5678; \
  curl --no-progress-meter --connect-timeout 1 gateway.docker.internal:5678 2>&1; \
  ip link; ifconfig
CMD echo "host.docker.internal="$(dig +retry=0 +timeout=1 +short host.docker.internal); \
  echo 'gateway.docker.internal='$(dig +retry=0 +timeout=1 +short gateway.docker.internal); \
  echo curl localhost:5678; \
  curl --no-progress-meter --connect-timeout 1 localhost:5678 2>&1; \
  echo curl host.docker.internal:5678; \
  curl --no-progress-meter --connect-timeout 1 host.docker.internal:5678 2>&1; \
  echo curl gateway.docker.internal:5678; \
  curl --no-progress-meter --connect-timeout 1 gateway.docker.internal:5678 2>&1;
