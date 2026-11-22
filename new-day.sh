#!/usr/bin/env bash
# Usage: ./new-day.sh <day>     e.g. ./new-day.sh 3  -> creates day03/
#
# Optional: export AOC_SESSION=<your session cookie> to also download input.txt.
set -euo pipefail

YEAR=2026
cd "$(dirname "$0")"

if [[ $# -ne 1 || ! "$1" =~ ^[0-9]+$ ]]; then
  echo "usage: $0 <day>" >&2
  exit 1
fi

day=$((10#$1))
dir=$(printf "day%02d" "$day")

if [[ -e "$dir" ]]; then
  echo "$dir already exists" >&2
  exit 1
fi

mkdir "$dir"
cp _template/main.go _template/main_test.go "$dir/"
touch "$dir/test.txt"

if [[ -n "${AOC_SESSION:-}" ]]; then
  curl -fsS -A "${AOC_USER_AGENT:-advent-of-code-2026 go solutions}" \
    -b "session=$AOC_SESSION" \
    "https://adventofcode.com/$YEAR/day/$day/input" -o "$dir/input.txt" \
    && echo "Downloaded input for day $day"
else
  touch "$dir/input.txt"
fi

echo "Created $dir/ — paste the example into $dir/test.txt"
