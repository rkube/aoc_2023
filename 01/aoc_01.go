package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
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
		num_lines += 1
		// fmt.Printf("line %d (length: %d): %s\n", num_lines, len(line_str), line_str)

		// Conver to list of runes
		runes := []rune(line_str)

		first_int := 0
		last_int := 0
		// Iterate from start, find first integer and store.
		for i := 0; i < len(runes); i++ {
			first_int, err = letter_to_int(runes[i])
			if err == nil {
				break
			}
		}

		// Reverse iterate from end, find first integer and store.
		for i := len(runes) - 1; i >= 0; i-- {
			last_int, err = letter_to_int(runes[i])
			if err == nil {
				break
			}
		}

		// fmt.Printf("first_int = %d, last_int = %d, total = %d\n", first_int, last_int, 10*first_int+last_int)
		total_sum += 10*first_int + last_int
	}

	fmt.Printf("total_sum = %d\n", total_sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
