package main

import (
	"fmt"
	"testing"
)

func Test_parse_input(t *testing.T) {
	fmt.Println("-------- Test input parsing --------")
	m := parse_input("input_14_test")
	for ix_row := 0; ix_row < m.nrows; ix_row++ {
		for ix_col := 0; ix_col < m.ncols; ix_col++ {
			fmt.Printf("%c", m.data[ix_row*m.ncols+ix_col])
		}
		fmt.Printf("\n")
	}
}

func Test_rotate_cw(t *testing.T) {
	fmt.Println("-------- Rotating clockwise --------")
	m := matrix_t{nrows: 3, ncols: 3, data: []byte{0, 1, 2, 3, 4, 5, 6, 7, 8}}
	m_t := rotate_cw(m)
	// fmt.Printf("Original: \n\t")
	fmt.Println(m.data)
	// fmt.Printf("Rotate clockwise: \n\t")
	fmt.Println(m_t.data)
}

// Test shaking the O's to top in a single row
func Test_shake_row(t *testing.T) {
	fmt.Println("------- Test shaking north --------")
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

func Test_rotate_cycle(t *testing.T) {
	fmt.Println("-------- Cycle rotation --------")
	// To shake towards west, rotate the board clockwise, so that west is the new north.
	// Then apply sort_Os, to get the round boulders on top.
	m := parse_input("input_14_test") // north-ward board
	m_w := rotate_cw(m)               // west-ward rotated board
	m_s := rotate_cw(m_w)             // south-ward rotated board
	m_e := rotate_cw(m_s)             // east-ward rotated board
	m_n := rotate_cw(m_e)             // north-ward board, should be identical to origina

	fmt.Println("----- Original board")
	for ix_row := 0; ix_row < m_n.nrows; ix_row++ {
		for ix_col := 0; ix_col < m_n.ncols; ix_col++ {
			fmt.Printf("%c", m_n.data[ix_row*m.ncols+ix_col])
		}
		fmt.Printf("\n")
	}

	fmt.Println("----- Board after 4 rotations")
	for ix_row := 0; ix_row < m_n.nrows; ix_row++ {
		for ix_col := 0; ix_col < m.ncols; ix_col++ {
			fmt.Printf("%c", m_n.data[ix_row*m.ncols+ix_col])
		}
		fmt.Printf("\n")
	}

}

func Test_shake_cycle(t *testing.T) {
	fmt.Println("-------- Test clockwise rotation --------")
	// To shake towards west, rotate the board clockwise, so that west is the new north.
	// Then apply sort_Os, to get the round boulders on top.
	m_n := parse_input("input_14_test")
	m_w := rotate_cw(m_n) // west-ward rotated board
	m_s := rotate_cw(m_w) // south-ward rotated board
	m_e := rotate_cw(m_s) // east-ward rotated board

	pos_S_n := make(map[int][]int) // Gives search intervals for original board
	pos_S_w := make(map[int][]int) // Gives search intervals for west-ward shake
	pos_S_s := make(map[int][]int) // Gives search intervals for south-ward shake
	pos_S_e := make(map[int][]int) // Gives search intervals for east-ward shake

	for ix_col := 0; ix_col < m_n.ncols; ix_col++ {
		pos_S_n[ix_col] = search_col(m_n, '#', ix_col)
		pos_S_w[ix_col] = search_col(m_w, '#', ix_col)
		pos_S_s[ix_col] = search_col(m_s, '#', ix_col)
		pos_S_e[ix_col] = search_col(m_e, '#', ix_col)
	}

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
