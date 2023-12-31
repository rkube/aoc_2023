package main

import (
	"fmt"
	"testing"
)

func Test_arrangements(t *testing.T) {
	// We want to test how many arrangements the
	// pattern .??..?##? has for the configuration
	// 1, 3
	// Answer: 4
	// 1: .#...###.
	// 2: ..#..###.
	// 3: .#....###
	// 4: ..#...###
	//

	fmt.Printf("Test")
	pattern := ".??..?##?"
	counts := []int{1, 3}
	num_matches := count_arrangements(pattern, counts)
	if count_arrangements(pattern, counts) != 4 {
		t.Fatalf("Pattern: %s\nCounts: 1,3. Correct: 4, ours: %d\n", pattern, num_matches)
	}
	fmt.Printf("Matching patterns: %d\n", num_matches)
}
