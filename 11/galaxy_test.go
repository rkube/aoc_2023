package main

import (
	"fmt"
	"testing"
)

// Verify that we correctly find indices of empty rows
func Test_find_empty_rows_cols(t *testing.T) {
	nrow, ncol := 10, 10
	chart := load_map("input_11_test1", nrow, ncol)

	for ix_row := 0; ix_row < nrow; ix_row++ {
		fmt.Printf("%d: ", ix_row)
		for ix_col := 0; ix_col < ncol; ix_col++ {
			fmt.Printf("%c", chart[ix_row*ncol+ix_col])
		}
		fmt.Printf("\n")
	}

	empty_rows := find_empty_rows(chart, nrow, ncol)
	empty_cols := find_empty_cols(chart, nrow, ncol)

	fmt.Printf("%d empty rows: ", len(empty_rows))
	for _, r := range empty_rows {
		fmt.Printf("%d ", r)
	}
	fmt.Printf("\n")

	fmt.Printf("%d empty columns: ", len(empty_cols))
	for _, r := range empty_cols {
		fmt.Printf("%d ", r)
	}
	fmt.Printf("\n")

	fmt.Printf("\n\n---- Expanded rows chart:\n")
	new_nrow := nrow + len(empty_rows)
	new_chart := duplicate_row(chart, nrow, ncol, empty_rows)
	for ix_row := 0; ix_row < new_nrow; ix_row++ {
		fmt.Printf("%02d: ", ix_row)
		for ix_col := 0; ix_col < ncol; ix_col++ {
			fmt.Printf("%c", new_chart[ix_row*ncol+ix_col])
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n\n---- Expanded cols chart:\n")
	new_ncol := ncol + len(empty_cols)
	new_chart_cols := duplicate_col(chart, nrow, ncol, empty_cols)
	for ix_row := 0; ix_row < nrow; ix_row++ {
		fmt.Printf("%02d: ", ix_row)
		for ix_col := 0; ix_col < new_ncol; ix_col++ {
			fmt.Printf("%c", new_chart_cols[ix_row*new_ncol+ix_col])
		}
		fmt.Printf("\n")
	}

}
