package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Convert a rune to integer, for runes in 0,1,2,3,4,5,6,7,8,9
func letter_to_int(r rune) (int, error) {
	m := map[rune]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
	i, ok := m[r]
	if !ok {
		return -1, errors.New("Not a number betwween 0-9")
	}
	return i, nil
}

func main() {
	fmt.Println("Advent of Code 2024 - Project 1")
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	scanner := bufio.NewScanner(f)

	total_sum := 0
	num_lines := 0
	for scanner.Scan() {
		// do something with a line
		line_str := scanner.Text()

		// Convert all string numbers with the actual string
		// To take care of edge cases, such as oneight, which needs to read 18,
		// replace onnly thee middle letters of the spelled out number with the actual digits.
		num_lines += 1
		fmt.Printf("\nline %d (length: %d): %s\n", num_lines, len(line_str), line_str)
		str_1 := strings.Replace(line_str, "one", "o1e", -1)
		str_2 := strings.Replace(str_1, "two", "t2o", -1)
		str_3 := strings.Replace(str_2, "three", "t3e", -1)
		str_4 := strings.Replace(str_3, "four", "4", -1)
		str_5 := strings.Replace(str_4, "five", "d5e", -1)
		str_6 := strings.Replace(str_5, "six", "6", -1)
		str_7 := strings.Replace(str_6, "seven", "s7n", -1)
		str_8 := strings.Replace(str_7, "eight", "e8t", -1)
		str_9 := strings.Replace(str_8, "nine", "n9e", -1)
		fmt.Printf("old: %s\nnew: %s\n", line_str, str_9)

		// Conver to list of runes
		runes := []rune(str_9)

		first_int := 0
		last_int := 0
		// Iterate from start, find first integer and store.
		for i := 0; i < len(runes); i++ {
			// Try converting the first rune to an integer.
			// If this is successful, we found our first digit and stop the iteration.
			first_int, err = letter_to_int(runes[i])
			if err == nil {
				break
			}
		}

		// Reverse iterate from end, find first integer and store. Same logic as loop above, but
		// iteration is reversed to find the last digit in the string.
		for i := len(runes) - 1; i >= 0; i-- {
			last_int, err = letter_to_int(runes[i])
			if err == nil {
				break
			}
		}

		// Calculate the checksum by shifting decimal of first digit and adding the up.
		fmt.Printf("first_int = %d, last_int = %d, total = %d\n", first_int, last_int, 10*first_int+last_int)
		total_sum += 10*first_int + last_int
	}

	fmt.Printf("total_sum = %d\n", total_sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
