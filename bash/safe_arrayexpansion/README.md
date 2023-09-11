# Expanding Bash arrays safely with `set -u`

Prior to
[Bash 4.4](https://git.savannah.gnu.org/cgit/bash.git/tree/CHANGES?id=3ba697465bc74fab513a26dea700cc82e9f4724e#n878)
`set -u` treated empty arrays as "unset", and terminates the process.
There are a number of possible workarounds using array
[parameter expansion](https://www.gnu.org/software/bash/manual/html_node/Shell-Parameter-Expansion.html),
however almost all of them fail in certain Bash versions.

This gist is a supplement to [this StackOverflow post](https://stackoverflow.com/a/61551944/113632).

## tl;dr

The only safe option for expanding an array across Bash versions under `set -u` is:

```
${array[@]+"${array[@]}"}
```

Notice the quotes are *inside* the expansion, *not* surrounding the whole expression, and that it uses `+`, not `:+`.

## Alternatives

See `results.txt` for a detailed breakdown across several versions of Bash.
Some of these alternatives are "obviously" wrong due to incorrect quoting, but
they're included for completeness.

| shorthand | `$@` example | `${arr[@]}` example | broken in | notes |
| --- | --- | --- | --- | --- |
| ":+"  | `"${@:+"$@"}"` | `"${arr[@]:+"${arr[@]}"}"` | 4.2+ | |
| ":+   | `"${@:+$@}"` | `"${arr[@]:+${arr[@]}}"`     | 4.2+ | |
| "+"   | `"${@+"$@"}"` | `"${arr[@]+"${arr[@]}"}"`   | 4.2  | |
| "+    | `"${@+$@}"` | `"${arr[@]+${arr[@]}}"`       | 4.2  | |
| :+"   | `${@:+"$@"}` | `${arr[@]:+"${arr[@]}"}`     | *    | |
| :+    | `${@:+$@}` | `${arr[@]:+${arr[@]}}`         | *    | |
| +"    | `${@+"$@"}` | `${arr[@]+"${arr[@]}"}`       | N/A  | works in all tested versions (> 3.0) |
| +     | `${@+$@}` | `${arr[@]+${arr[@]}}`           | *    | |
| ":0/1 | `"${@:1}"` | `"${arr[@]:0}"`                | 4.2  | crashes, presumably a regression |
| :0/1  | `${@:1}` | `${arr[@]:0}`                    | *    | also crashes in 4.2 |
| ":-   | `"${@:-}"` | `"${arr[@]:-}"`                | *    | |
| :-    | `${@:-}` | `${arr[@]:-}`                    | *    | |
| "-    | `"${@-}"` | `"${arr[@]-}"`                  | *    | |
| -     | `${@-}` | `${arr[@]-}`                      | *    | |

If we exclude v4.2 several other expansions *do* work, including
`"${@+"$@"}"` and `"${@:1}"`, but so long you intend to support that
version these expansions are not safe.

## Reproduction

To reproduce the contents of `results.txt` run:

```shell
for v in 3.1 3.2 4.0 4.1 4.2 4.3 4.4 5.0; do
  docker run -v "$PWD:/mnt" "bash:$v" bash /mnt/expansions.sh
done
```

Bash 3.0 is intentionally excluded from the reported results, but you can run
the script against that version too to see what breaks (hint: it's a lot).
