set -e

a=$({ sleep 2; exit 12; } | { exit 34; } | { exit 56; }; exit 0)
echo 'ret: '$?

a=$({ sleep 2; exit 12; } | { exit 34; } | { exit 56; }; exit ${PIPESTATUS[0]}) || ret=$?
echo '(Using PIPESTATUS) ret: '$ret
