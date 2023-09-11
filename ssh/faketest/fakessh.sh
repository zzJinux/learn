#!/usr/bin/env bash

echo "$#" > args
printf '%s\n' "$@" >> args

read -d '' -r p <<'EOF'
Warning: Permanently added '[127.0.0.1]:2222' (ED25519) to the list of known hosts.
test@127.0.0.1's password: \
EOF
echo -n "${p::-1}"

read -r REPLY
echo "$REPLY" > reply

