package main

import (
	"fmt"
	"testing"
)

// Parse the test matrices and confirm that we get the correct sizes
func Test_parsing(t *testing.T) {
	fmt.Printf("====== Test parsing ======\n")
	matrices := parse_file("input_13_test")

	// Test matrices have both 7 rows and 9 columns
	correct_sizes := [][]int{{7, 9}, {7, 9}}

	for ix_m, m := range matrices {
		// fmt.Printf("Correct size: %d, %d\n", correct_sizes[ix_m][0], correct_sizes[ix_m][1])
		// fmt.Printf("Our size    : %d, %d\n", m.nrows, m.ncols)

		if (m.nrows != correct_sizes[ix_m][0]) || (m.ncols != correct_sizes[ix_m][1]) {
			t.Fatalf("Correct size: %d, %d. Our size: %d, %d\n", correct_sizes[ix_m][0], correct_sizes[ix_m][1], m.nrows, m.ncols)
		}
		// fmt.Printf("Matrix %02d: nrows: %02d, ncols: %02d\n", ix_m, m.nrows, m.ncols)
	}
}

func Test_iteration(t *testing.T) {
	fmt.Printf("====== Test row iteration ======\n")
	matrices := parse_file("input_13_test")

	answer_0_col := find_vertical_reflection(matrices[0])
	answer_0_row := find_vertical_reflection(transpose(matrices[0]))
	fmt.Printf("Matrix 0: vertical: %d\n", answer_0_col)
	fmt.Printf("Matrix 0: horizontal: %d\n", answer_0_row)

	answer_1_col := find_vertical_reflection(matrices[1])
	answer_1_row := find_vertical_reflection(transpose(matrices[1]))

	fmt.Printf("Matrix 1: vertical: %d\n", answer_1_col)
	fmt.Printf("Matrix 1: horizontal: %d\n", answer_1_row)
}

func Test_transpose(t *testing.T) {
	fmt.Printf("====== Testing matrix transpose ======\n")

	m := matrix_t{nrows: 4, ncols: 2, data: []byte{1, 2, 11, 12, 21, 22, 31, 32}}
	fmt.Printf("Original matrix: nrows=%d, ncols=%d\nData:", m.nrows, m.ncols)
	fmt.Println(m.data)
	for ix_row := 0; ix_row < m.nrows; ix_row++ {
		for ix_col := 0; ix_col < m.ncols; ix_col++ {
			fmt.Printf("%d ", m.data[ix_row*m.ncols+ix_col])
		}
		fmt.Printf("\n")
	}

	m_t := transpose(m)
	fmt.Printf("Transposed matrix: nrows=%d, ncols=%d\n", m_t.nrows, m_t.ncols)

	for ix_row := 0; ix_row < m_t.nrows; ix_row++ {
		for ix_col := 0; ix_col < m_t.ncols; ix_col++ {
			fmt.Printf("%d ", m_t.data[ix_row*m_t.ncols+ix_col])
		}
		fmt.Printf("\n")
	}

}
