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

	pattern := ".??..?##?"
	counts := []int{1, 3}

	num_matches := count_arrangements(pattern, counts)
	if count_arrangements(pattern, counts) != 4 {
		t.Fatalf("Pattern: %s\nCounts: 1,3. Correct: 4, ours: %d\n", pattern, num_matches)
	}
	fmt.Printf("Matching patterns: %d\n", num_matches)

	// Some more tests
	// ???.### 1,1,3 - 1 arrangement
	// .??..??...?##. 1,1,3 - 4 arrangements
	// ?#?#?#?#?#?#?#? 1,3,1,6 - 1 arrangement
	// ????.#...#... 4,1,1 - 1 arrangement
	// ????.######..#####. 1,6,5 - 4 arrangements
	// ?###???????? 3,2,1 - 10 arrangements

	pattern = "???.###"
	counts = []int{1}
	num_matches = count_arrangements(pattern, counts)
	fmt.Printf("Test 1: pattern = %s, num_matches = %d - correct:1\n", pattern, num_matches)

	pattern = ".??..??...?##."
	counts = []int{1, 1, 3}
	num_matches = count_arrangements(pattern, counts)
	fmt.Printf("Test 2: pattern = %s, num_matches = %d - correct: 4\n", pattern, num_matches)

	pattern = "?#?#?#?#?#?#?#?"
	counts = []int{1, 3, 1, 6}
	num_matches = count_arrangements(pattern, counts)
	fmt.Printf("Test 3: pattern = %s, num_matches = %d - correct: 1\n", pattern, num_matches)

	pattern = "????.#...#..."
	counts = []int{4, 1, 1}
	num_matches = count_arrangements(pattern, counts)
	fmt.Printf("Test 4: pattern = %s, num_matches = %d - correct: 1\n", pattern, num_matches)

	pattern = "????.######..#####."
	counts = []int{1, 6, 5}
	num_matches = count_arrangements(pattern, counts)
	fmt.Printf("Test 5: pattern = %s, num_matches = %d - correct: 4\n", pattern, num_matches)

	pattern = "?###????????"
	counts = []int{3, 2, 1}
	num_matches = count_arrangements(pattern, counts)
	fmt.Printf("Test 6: pattern = %s, num_matches = %d - correct: 10\n", pattern, num_matches)

}
