lines=(${(@f)"$(<ll)"})

declare -p lines

_tttt() {
  _describe qwqwqw lines
}
compdef _tttt tttt
