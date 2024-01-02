package main

import (
	"fmt"
	"testing"
)

func Test_parse_input(t *testing.T) {
	m := parse_input("input_14_test")
	for ix_row := 0; ix_row < m.nrows; ix_row++ {
		for ix_col := 0; ix_col < m.ncols; ix_col++ {
			fmt.Printf("%c", m.data[ix_row*m.ncols+ix_col])
		}
		fmt.Printf("\n")
	}
}

// Test shaking the O's to top in a single row
func Test_shake_row(t *testing.T) {
	m := parse_input("input_14_test")

	// Positions refer to indices in a matrix and start from 0!
	ix_col := 9
	pos_S := search_col(m, '#', ix_col)
	pos_O := search_col(m, 'O', ix_col)
	fmt.Printf("Column %d, looking for '#':\n\t", ix_col)
	fmt.Println(pos_S)
	fmt.Printf("           looking for 'O':\n\t")
	fmt.Println(pos_O)

	fmt.Println("Shaking: ")
	sorted_Os := sort_Os(pos_O, pos_S, m.nrows)
	fmt.Printf("           After sorting: \n")

	fmt.Println(sorted_Os)

}

func Test_score_rows(t *testing.T) {
	// Get the updated positions for every row and calculate the score
	m := parse_input("input_14_test")
	score := 0
	for ix_col := 0; ix_col < m.ncols; ix_col++ {
		pos_S := search_col(m, '#', ix_col)
		pos_O := search_col(m, 'O', ix_col)
		sorted_Os := sort_Os(pos_O, pos_S, m.nrows)
		score += score_row(sorted_Os, m.nrows)
	}
	fmt.Printf("Score: %d\n", score)
}
