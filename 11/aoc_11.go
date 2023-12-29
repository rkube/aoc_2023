package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Returns a number of rows that don't contain a galaxy (#)
func find_empty_rows(chart []byte, nrow int, ncol int) []int {
	empty_rows := []int{}

	found_hash := false
	for ix_row := 0; ix_row < nrow; ix_row++ {
		found_hash = false
		for ix_col := 0; ix_col < ncol; ix_col++ {
			if chart[ix_row*ncol+ix_col] == '#' {
				found_hash = true
			}
		}
		if found_hash == false {
			empty_rows = append(empty_rows, ix_row)
		}
	}
	return empty_rows
}

// Returns a number of columns that don't contain a galaxy (#)
func find_empty_cols(chart []byte, nrow int, ncol int) []int {
	empty_cols := []int{}
	found_hash := false

	for ix_col := 0; ix_col < ncol; ix_col++ {
		found_hash = false
		for ix_row := 0; ix_row < nrow; ix_row++ {
			if chart[ix_row*ncol+ix_col] == '#' {
				found_hash = true
			}
		}
		if found_hash == false {
			empty_cols = append(empty_cols, ix_col)
		}
	}
	return empty_cols
}

// Duplicate empty rows
func duplicate_row(chart []byte, nrow int, ncol int, duplicate_rows []int) []byte {
	new_nrow := nrow + len(duplicate_rows)
	new_chart := make([]byte, new_nrow*ncol)

	ix_new_row := 0 // Used to index the new array
	for ix_row := 0; ix_row < nrow; ix_row++ {
		// Copy current line from original chart to new chart
		for ix_col := 0; ix_col < ncol; ix_col++ {
			new_chart[ix_new_row*ncol+ix_col] = chart[ix_row*ncol+ix_col]
		}
		// If we have to duplicate the row, add it to the new array
		for _, v := range duplicate_rows {
			// We have to check if the current original row needs to be duplicated.
			if ix_row == v {
				// fmt.Printf("Duplicating: ix_row = %d -> ix_new_row = %d\n", ix_row, ix_new_row)
				ix_new_row += 1
				for ix_col := 0; ix_col < ncol; ix_col++ {
					new_chart[ix_new_row*ncol+ix_col] = chart[ix_row*ncol+ix_col]
				}

			}
		}
		ix_new_row += 1
	}
	return new_chart
}

func duplicate_col(chart []byte, nrow int, ncol int, duplicate_cols []int) []byte {
	new_ncol := ncol + len(duplicate_cols)
	new_chart := make([]byte, nrow*new_ncol)

	ix_new_col := 0
	for ix_row := 0; ix_row < nrow; ix_row++ {
		ix_new_col = 0
		for ix_col := 0; ix_col < ncol; ix_col++ {
			new_chart[ix_row*new_ncol+ix_new_col] = chart[ix_row*ncol+ix_col]
			// Check if current column needs to be duplicated
			for _, v := range duplicate_cols {
				if ix_col == v {
					ix_new_col += 1
					new_chart[ix_row*new_ncol+ix_new_col] = chart[ix_row*ncol+ix_col]
				}
			}
			ix_new_col += 1
		}
	}
	return new_chart
}

func load_map(filename string, nrow int, ncol int) []byte {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	//Chart is a len*len byte array
	chart := make([]byte, nrow*ncol)
	ix_row := 0

	for scanner.Scan() {
		current_line := scanner.Text()
		for ix_col := 0; ix_col < ncol; ix_col++ {
			chart[ix_row*nrow+ix_col] = current_line[ix_col]
		}
		ix_row += 1
	}

	return chart
}

func main() {
	fmt.Printf("Advent of Code 2023 - Day 11\n")

	chart := load_map("input_11", 141, 141)
	fmt.Printf("%s\n", chart[0:141])
}
