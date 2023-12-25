package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Split line at space and return parsed array of integers
func parse_line(line string) []int {
	splits := strings.Split(line, " ")
	int_array := make([]int, len(splits))

	for ix_split, split := range splits {
		this_val, err := strconv.Atoi(split)
		if err != nil {
			log.Fatal(err)
		}
		int_array[ix_split] = this_val
	}
	return int_array
}

// Write a recursive function that calculates the difference between arrays.
// Return if all differences are zero

func calc_differences(array []int) int {
	// Return 0 if the sum of the array is zero
	sum := 0
	for _, v := range array {
		sum += v
	}
	if sum == 0 {
		return sum
	} else {

	}

}

func main() {
	fmt.Printf("Advent of code 2023 - Part 09\n")

	// Hard-code some test example first
	//
	line_str := "0 3 6 9 12 15"
	line_int := parse_line(line_str)

}
