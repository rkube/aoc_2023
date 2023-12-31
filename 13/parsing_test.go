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

	find_vertical_reflection(matrices[0])
}

func Test_transpose(t *testing.T) {
	fmt.Printf("====== Testing matrix transpose ======\n")

	m := matrix_t{nrows: 4, ncols: 2, data: []byte{1, 2, 11, 12, 21, 22, 31, 32}}
	fmt.Printf("Original matrix: nrows=%d, ncols=%d\n", m.nrows, m.ncols)
	fmt.Println(m.data)

	m_t := transpose(m)
	fmt.Printf("Transposed matrix: nrows=%d, ncols=%d\n", m_t.nrows, m_t.ncols)
	fmt.Printf("len: %d\n", len(m_t.data))
	fmt.Println(m_t.data)

}
