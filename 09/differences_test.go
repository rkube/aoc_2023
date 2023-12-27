package main

import (
	"fmt"
	"testing"
)

/*
 * Run with  'go test'
 */

func Test_part1_differences(t *testing.T) {
	fmt.Printf("-------- Test cases part 1 ---------\n")
	test_cases := []string{"0 3 6 9 12 15", "1 3 6 10 15 21", "10 13 16 21 30 45"}
	for _, line_str := range test_cases {

		line_int := parse_line(line_str)
		ip_val := calc_differences(line_int)
		fmt.Printf("%s - %d\n", line_str, ip_val)
	}
}

func Test_part2_differences(t *testing.T) {
	fmt.Printf("-------- Test cases part 2 ---------\n")
	test_cases := []string{"0 3 6 9 12 15", "1 3 6 10 15 21", "10 13 16 21 30 45"}
	for _, line_str := range test_cases {
		line_int := parse_line(line_str)
		ip_val := calc_differences_2(line_int)
		fmt.Printf("%s - Extrapolated: %d\n", line_str, ip_val)
	}
}
