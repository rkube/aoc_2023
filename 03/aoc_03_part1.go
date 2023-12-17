package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Stores location of nunmbers in the array.
// F.ex the sequence below (row and column indices are written out here)
//
//	0123456789
//
// 0 .......$..
// 1 ..223.....
// 2 ..........
// [...]
// would have
// row=0, col_start=7, col_end=7 for "$"
// row=1, col_start=2, col_end=4 for "223"
type text_loc struct {
	row       int // row index where the number is located
	col_start int // column index where the number starts
	col_end   int // column index where the number ends
}

func char_in_string(r rune, s string) bool {
	for _, cmp := range s {
		if r == cmp {
			return true
		}
	}
	return false
}

// Extracts location of symbols # $ % & * + -/ = @
// for a line
func parse_symbols(linenr int, line string) []text_loc {
	symbol_locs := []text_loc{}

	for ix := 0; ix < len(line); ix++ {
		if char_in_string(rune(line[ix]), "#$%&*+-/=@") {
			fmt.Printf("ix %d: Special char: %c\n", ix, line[ix])
			symbol_locs = append(symbol_locs, text_loc{linenr, ix, ix})
		}
	}
	fmt.Printf("\n")
	return symbol_locs
}

// // Extract location of numbers in a line
func parse_numbers(linenr int, line string) ([]text_loc, bool) {
	number_locs := []text_loc{}
	// state == true if we are currently parsing a number
	c_state := false
	for ix := 0; ix < len(line); ix++ {
		// fmt.Printf("%d: %c\n", ix, line[ix])
		if char_in_string(rune(line[ix]), "0123456789") {
			// Check if the previous token is a '.'. In that case we are starting to parse a number
			if (ix > 0) && char_in_string(rune(line[ix-1]), "#$%&*+-/=@.") {
				c_state = true

				fmt.Printf("before: len(number_locs) = %d\n", len(number_locs))
				number_locs = append(number_locs, text_loc{linenr, ix, -1})
				fmt.Printf("after: len(number_locs) = %d\n", len(number_locs))
			}
			if (ix < 139) && char_in_string(rune(line[ix+1]), "#$%&*+-/=@.") {
				c_state = false
				ix_s := len(number_locs)
				fmt.Printf("ix_s = %d\n", ix_s)
				number_locs[ix_s-1].col_end = ix + 1

				// We just closed a number. Print it to terminal
				for cix := number_locs[ix_s-1].col_start; cix < number_locs[ix_s-1].col_end; cix++ {
					fmt.Printf("%d: %c\n", cix, line[cix])
				}
			}
		}
	}
	return number_locs, c_state
}

func main() {
	fmt.Println("Advent of code 2023 - Project 2")
	// Input has 140 lines, each 140 characters.
	f, err := os.Open("input_03")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	linenr := 0
	for scanner.Scan() {
		current_line := scanner.Text()
		fmt.Printf("%s\n", current_line)

		symbol_loc := parse_symbols(linenr, current_line)
		number_loc, s := parse_numbers(linenr, current_line)
		fmt.Printf("%d: len(symbols)=%d, len(numbers)=%d, s=%d\n", linenr, len(symbol_loc), len(number_loc), s)
		for ix_n, n := range number_loc {
			fmt.Printf("%d: %s\n", ix_n, current_line[n.col_start:n.col_end])
		}

		linenr += 1
		if linenr > 1 {
			break
		}
	}

	// remember to close the file at the end of the program
	defer f.Close()
}
