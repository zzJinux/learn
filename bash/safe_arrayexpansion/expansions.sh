#!/bin/bash

echo "Bash Version: $BASH_VERSION"

expansion() {
  local incantation=$1 expected=$2 actual
  shift 2
  actual=$(
    set -u
    eval "$(printf 'copy=( %s ); printf "${#copy[@]}"' "$incantation")" 2>/dev/null
  ) || { printf '\t\e[1;33m%s\e[0m' "!"; return 0; }
  if (( actual == expected )); then
    printf '\t\e[32m%s\e[0m' "✓"
  else
    printf '\t\e[31m%s\e[0m' "✗"
  fi
}

test_expansion() {
  local arr=("$@")
  
  printf '$#:%s' "$#"
  for incantation in \
    '"${@:+"$@"}"' '"${@:+$@}"' '"${@+"$@"}"' '"${@+$@}"' '${@:+"$@"}' '${@:+$@}' \
    '${@+"$@"}' '${@+$@}' '"${@:1}"' '${@:1}'
  do
    expansion "$incantation" "$#" "$@"
  done
  printf '\n'
  
  printf '#arr:%s' "${#arr[@]}"
  for incantation in \
    '"${arr[@]:+"${arr[@]}"}"' '"${arr[@]:+${arr[@]}}"' '"${arr[@]+"${arr[@]}"}"' \
    '"${arr[@]+${arr[@]}}"' '${arr[@]:+"${arr[@]}"}' '${arr[@]:+${arr[@]}}' \
    '${arr[@]+"${arr[@]}"}' '${arr[@]+${arr[@]}}' '"${arr[@]:0}"' '${arr[@]:0}'
  do
    expansion "$incantation" "$#"
  done
  printf '\n'
}

printf '\t%s' '":+"' '":+' '"+"' '"+' ':+"' ':+' '+"' '+' '":0/1' ':0/1'
printf '\n'
test_expansion
test_expansion ''
test_expansion '' ''
test_expansion a 'b c'
