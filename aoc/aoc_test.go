package aoc

import "testing"

func TestMod(t *testing.T) {
	tests := []struct {
		a, m, want int
	}{
		{5, 100, 5},
		{105, 100, 5},
		{-5, 100, 95},
		{-100, 100, 0},
		{-205, 100, 95},
		{0, 100, 0},
	}
	for _, tt := range tests {
		if got := Mod(tt.a, tt.m); got != tt.want {
			t.Errorf("Mod(%d, %d) = %d, want %d", tt.a, tt.m, got, tt.want)
		}
	}
}
