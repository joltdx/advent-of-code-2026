package main

import (
	"strings"
	"testing"
)

// Example input from the puzzle text. Backtick strings are raw: newlines and
// spaces are kept exactly, so don't indent the lines inside them.
const example = `
paste
example
here
`

// tidy removes the leading/trailing newlines that come from writing a raw
// string on its own lines, matching what readInput does with a file.
// Leading spaces on the first line are kept (some puzzles need them).
func tidy(s string) string {
	return strings.Trim(s, "\r\n")
}

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example", example, 0},
		// {"single line", "abc", 0},
		// {"edge case", `
		// ...
		// `, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tidy(tt.input)); got != tt.want {
				t.Errorf("part1() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example", example, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tidy(tt.input)); got != tt.want {
				t.Errorf("part2() = %d, want %d", got, tt.want)
			}
		})
	}
}

// Benchmarks run against the real input.txt: go test -bench . -run '^$'
// (add -benchmem to also see allocations). Skipped if input.txt is missing or empty.
//
// If part1/part2 take parsed data instead of a string, parse once *before*
// the loop so only the solving is timed, e.g.:
//
//	data, err := parse(benchInput(b))
//	if err != nil {
//		b.Fatal(err)
//	}
//	for b.Loop() {
//		part1(data)
//	}
func BenchmarkPart1(b *testing.B) {
	input := benchInput(b)
	for b.Loop() {
		part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := benchInput(b)
	for b.Loop() {
		part2(input)
	}
}

func benchInput(b *testing.B) string {
	b.Helper()
	input, err := readInput("input.txt")
	if err != nil {
		b.Skip("no input.txt:", err)
	}
	if input == "" {
		b.Skip("input.txt is empty")
	}
	return input
}

// Helper functions in main.go can be tested the same way, e.g.:
//
// func TestParseLine(t *testing.T) {
// 	got := parseLine("1,2,3")
// 	want := []int{1, 2, 3}
// 	if !slices.Equal(got, want) {
// 		t.Errorf("parseLine() = %v, want %v", got, want)
// 	}
// }
