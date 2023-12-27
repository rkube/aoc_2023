package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func main() {
	fmt.Printf("Advent of code 2023 - Part 09\n")

	f, err := os.Open("input_09")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	sum_part1 := 0
	sum_part2 := 0
	for scanner.Scan() {
		current_line := scanner.Text()
		line_ints := parse_line(current_line)
		sum_part1 += calc_differences(line_ints)
		sum_part2 += calc_differences_2(line_ints)
	}

	fmt.Printf("Part 1 - %d\n", sum_part1)
	fmt.Printf("Part 1 - %d\n", sum_part2)

}
