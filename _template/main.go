package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Input file from the command line, defaulting to input.txt
	path := "input.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	input, err := readInput(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	start := time.Now()
	fmt.Println("Part 1:", part1(input), time.Since(start))

	start = time.Now()
	fmt.Println("Part 2:", part2(input), time.Since(start))
}

// readInput reads the whole file and strips the trailing newline(s).
func readInput(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(data), "\r\n"), nil
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	_ = lines
	return 0
}

func part2(input string) int {
	return 0
}
