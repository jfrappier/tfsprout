#!/usr/bin/env bash
#
# Verifies that docs/reference/checks.md lists exactly the checks that exist on
# disk. Catches the two ways the index drifts: a new check added without a
# documentation row, and a row left behind after a check is deleted.
#
# Also verifies every check directory has a README.md, since the index links to
# it.
#
# Usage: ./scripts/check-docs-sync.sh

set -euo pipefail

cd "$(dirname "$0")/.."

INDEX="docs/reference/checks.md"

if [[ ! -f "$INDEX" ]]; then
  echo "error: $INDEX not found" >&2
  exit 1
fi

# Check directories are named for their ID: AT001, R001, S001, V001, XAT001, ...
on_disk="$(
  find passes xpasses -mindepth 1 -maxdepth 1 -type d \
    | sed 's|.*/||' \
    | grep -E '^X?(AT|R|S|V)[0-9]{3}$' \
    | sort -u
)"

# Rows look like: | [AT001](../../passes/AT001) | description | AST |
documented="$(
  grep -oE '^\| \[X?(AT|R|S|V)[0-9]{3}\]' "$INDEX" \
    | sed -E 's/^\| \[//; s/\]$//' \
    | sort -u
)"

status=0

missing="$(comm -23 <(echo "$on_disk") <(echo "$documented"))"
if [[ -n "$missing" ]]; then
  status=1
  echo "error: check(s) exist on disk but are not listed in $INDEX:" >&2
  echo "$missing" | awk '{print "  " $0}' >&2
  echo >&2
  echo "Add a row to the appropriate table, using the first line of the check's Doc constant" >&2
  echo "as the description. See docs/contributing/adding-an-analyzer.md" >&2
  echo >&2
fi

stale="$(comm -13 <(echo "$on_disk") <(echo "$documented"))"
if [[ -n "$stale" ]]; then
  status=1
  echo "error: check(s) listed in $INDEX do not exist on disk:" >&2
  echo "$stale" | awk '{print "  " $0}' >&2
  echo >&2
  echo "Remove the stale row, or restore the check directory." >&2
  echo >&2
fi

# Every check the index links to must have a README for the link to resolve.
no_readme=""
while IFS= read -r id; do
  [[ -z "$id" ]] && continue
  if [[ "$id" == X* ]]; then
    dir="xpasses/$id"
  else
    dir="passes/$id"
  fi
  if [[ ! -f "$dir/README.md" ]]; then
    no_readme+="  $dir"$'\n'
  fi
done <<< "$on_disk"

if [[ -n "$no_readme" ]]; then
  status=1
  echo "error: check(s) missing README.md, so their index link is broken:" >&2
  printf '%s' "$no_readme" >&2
  echo >&2
fi

if [[ $status -eq 0 ]]; then
  echo "docs in sync: $(echo "$on_disk" | wc -l | tr -d ' ') checks documented in $INDEX"
fi

exit $status
