package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	value     int // Numerical value (only used for numbers, not symbols)
	v_type    int // 0: symbol, 1: gear, 2: part number
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
			if rune(line[ix]) == '*' {
				// Found a gear, use v_type = 1
				symbol_locs = append(symbol_locs, text_loc{linenr, ix, ix, 0, 1})
				fmt.Printf("ix %d: Gear: %c\n", ix, line[ix])
			} else {
				// Add an ordinary symbol
				symbol_locs = append(symbol_locs, text_loc{linenr, ix, ix, 0, 0})
				fmt.Printf("ix %d: Special char: %c\n", ix, line[ix])
			}
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
	for ix_l := 0; ix_l < len(line); ix_l++ {
		// fmt.Printf("%d: %c\n", ix, line[ix])
		// If we are not parsing a number, better start it and record.
		if char_in_string(rune(line[ix_l]), "0123456789") {
			// fmt.Printf("%d: case 1", ix)
			if c_state == false {
				// fmt.Printf(" start parsing")
				c_state = true
				// Add a part number
				number_locs = append(number_locs, text_loc{linenr, ix_l, -1, 0, 2})
			}
			// fmt.Printf("\n")
		} // If the next token is either a '.' or a special character, stop parsing
		if (ix_l < 139) && char_in_string(rune(line[ix_l+1]), "#$%&*+-/=@.") {
			if c_state == true {
				c_state = false
				ix_s := len(number_locs)
				fmt.Printf("ix_s = %d\n", ix_s)

				number_locs[ix_s-1].col_end = ix_l
				// Extract numerical value
				v, err := strconv.Atoi(line[number_locs[ix_s-1].col_start : number_locs[ix_s-1].col_end+1])
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Parsed %d\n", v)
				number_locs[ix_s-1].value = v
			}
		}
		if (ix_l == 139) && c_state == true {
			ix_s := len(number_locs)
			fmt.Printf("ix_s = %d\n", ix_s)
			number_locs[ix_s-1].col_end = ix_l
			v, err := strconv.Atoi(line[number_locs[ix_s-1].col_start : number_locs[ix_s-1].col_end+1])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Parsed %d\n", v)
			number_locs[ix_s-1].value = v
		}
	}

	return number_locs, c_state
}

func is_valid_part(n text_loc, symbols []text_loc) bool {
	// For a given text location of a number, check if any symbol is adjacent.
	for _, s := range symbols {
		// If a symbol is in the row above or beyond upper row, the symbol must be in col_start-1:col_end+1
		if ((s.row == n.row-1) || (s.row == n.row+1)) && (s.col_start >= n.col_start-1) && (s.col_end <= n.col_end+1) {
			return true
			// If the symbol is in the same row, the column must be one below start or one above end
		} else if (s.row == n.row) && ((s.col_start == n.col_start-1) || (s.col_end == n.col_end+1)) {
			return true
		}
	}
	return false
}

func match_gears(g text_loc, numbers []text_loc) int {
	// Iterate over all numbers, check if they are adjacent to the gear.
	// If exactly two gears are adjacent to this gear, return their product.
	// Else, return zero
	num_match := 0
	prod_match := 1
	for _, n := range numbers {
		// If a symbol is in the row above or beyond upper row, the symbol must be in col_start-1:col_end+1
		if ((g.row == n.row-1) || (g.row == n.row+1)) && (g.col_start >= n.col_start-1) && (g.col_end <= n.col_end+1) {
			// If the symbol is in the same row, the column must be one below start or one above end
			// fmt.Printf("Adjacent: gear [%d, %d], number %d [%d, %d:%d]\n", g.row, g.col_start, n.value, n.row, n.col_start, n.col_end)
			num_match += 1
			prod_match *= n.value
		} else if (g.row == n.row) && ((g.col_start == n.col_start-1) || (g.col_end == n.col_end+1)) {
			// fmt.Printf("Adjacent: gear [%d, %d], number %d [%d, %d:%d]\n", g.row, g.col_start, n.value, n.row, n.col_start, n.col_end)
			num_match += 1
			prod_match *= n.value
		}
	}
	// zero out the product if not exactly 2 gears are adjacent to the symbol
	if num_match != 2 {
		prod_match = 0
	}
	return prod_match
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

	symbol_locs := []text_loc{}
	number_locs := []text_loc{}

	// Parse text file and extract locations of all numbers and symbols
	for scanner.Scan() {
		current_line := scanner.Text()
		fmt.Printf("%d: %s\n", linenr, current_line)

		symbol_loc := parse_symbols(linenr, current_line)
		number_loc, _ := parse_numbers(linenr, current_line)
		// fmt.Printf("line %d: len(symbols)=%d, len(numbers)=%d, s=%d\n", linenr, len(symbol_loc), len(number_loc), s)
		for ix_n, n := range number_loc {
			fmt.Printf("%d [%d:%d]: %s\n", ix_n, n.col_start, n.col_end, current_line[n.col_start:n.col_end+1])
		}

		// Append symbol locations in current line
		for _, s := range symbol_loc {
			symbol_locs = append(symbol_locs, s)
		}
		for _, n := range number_loc {
			number_locs = append(number_locs, n)
		}
		linenr += 1
	}

	sum_part_numbers := 0
	sum_gear_ratios := 0
	for _, n := range number_locs {
		if is_valid_part(n, symbol_locs) {
			// fmt.Printf("Valid: row %d, [%d:%d]: %d\n", n.row, n.col_start, n.col_end, n.value)
			sum_part_numbers += n.value
		}
	}
	fmt.Printf("sum = %d\n", sum_part_numbers)

	// Part 2: Calculate gear ratios
	// Iterate over all gears, if we have two numbers adjacent to it, calculate their ratio
	for _, s := range symbol_locs {
		if s.v_type == 1 {
			n_match := match_gears(s, number_locs)
			// fmt.Printf("Processing gear - row %d, [%d]. %d matches\n", s.row, s.col_start, n_match)
			sum_gear_ratios += n_match
		}
	}
	fmt.Printf("sum gear_ratios = %d\n", sum_gear_ratios)

	// remember to close the file at the end of the program
	defer f.Close()
}
