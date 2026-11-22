# Advent of Code 2026 — in Go

One folder per day, each a standalone `main` package:

```text
dayNN/
  main.go        solution: part1(), part2()
  main_test.go   tests against the example input(s)
  test.txt       example input from the puzzle text
  input.txt      personal puzzle input (git-ignored)
```

## Workflow

```bash
./new-day.sh 1          # creates day01/ from _template/
cd day01
go test                 # run tests (add -v for detail, -run Part1 to filter)
go build && ./day01 test.txt
./day01                 # input file defaults to input.txt
go run .                # build + run in one step (or: go run . test.txt)
go test -bench . -run '^$'   # benchmark part1/part2 on input.txt (add -benchmem for allocations)
```

`-bench` and `-run` take a regex of names to match: `.` matches everything,
`'^$'` matches nothing (so the normal tests are skipped while benchmarking).

Set `AOC_SESSION` (the `session` cookie from adventofcode.com) to have
`new-day.sh` download `input.txt` for you.

From the repo root: `go test ./...` runs every day's tests, `go vet ./...` checks them.

## Shared helpers (`aoc/`)

Reusable helpers live in the `aoc` package, with tests in `aoc/aoc_test.go`.
Use them from any day:

```go
import "aoc2026/aoc"

pos = aoc.Mod(pos-dist, 100)
```

- `aoc.Mod(a, m)`: modulo that's never negative (like Rust's `rem_euclid`).
  Go's `%` can return a negative number: `-5 % 100 == -5`, but
  `aoc.Mod(-5, 100) == 95`.

To add one: write it in `aoc/` with an **Uppercase** name (lowercase names
aren't visible outside the package), add a test, and list it here.
