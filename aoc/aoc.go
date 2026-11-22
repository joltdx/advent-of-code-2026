// Package aoc holds small helpers shared between days.
// Use it from a day with: import "aoc2026/aoc"
package aoc

// Mod returns a mod m in the range [0, m), like Rust's rem_euclid.
// Go's % keeps the sign of a, so -5 % 100 == -5, but Mod(-5, 100) == 95.
func Mod(a, m int) int {
	return ((a % m) + m) % m
}
