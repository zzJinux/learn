# name1 and name2 arguments

name1=("aaaa:desc_aaaa" "bbbb:desc_bbbb")
_tttt() {
  name2=(1111 2222)
  _describe qwqwqw name1 name2
}
compdef _tttt tttt

# `tttt <tab>` displays:
# 
# -- qwqwqw --
# aaaa  -- desc_aaaa
# bbbb  -- desc_bbbb
#
# where selecting aaaa or bbbb puts 1111 or 2222 on the prompt.
