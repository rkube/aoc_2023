package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type matrix_t struct {
	data  []byte
	nrows int
	ncols int
}

// Tranpose memory layout of the matrix
func transpose(m matrix_t) matrix_t {

	data_tr := make([]byte, m.nrows*m.ncols)
	fmt.Printf("len(data_tr)= %d\n", len(data_tr))
	for ix_row := 0; ix_row < m.nrows; ix_row++ {
		for ix_col := 0; ix_col < m.ncols; ix_col++ {
			// fmt.Printf("Original: ix_row * m.ncols + ix_col= %d\n", ix_row*m.ncols+ix_col)
			// fmt.Printf("Transposed:ix_col * m.ncols + ix_row = %d\n", ix_col*m.nrows+ix_row)
			data_tr[ix_col*m.nrows+ix_row] = m.data[ix_row*m.ncols+ix_col]
		}
	}
	m_t := matrix_t{nrows: m.ncols, ncols: m.nrows, data: data_tr}
	// for ix_col := 0; ix_col < m_t.ncols; ix_col++ {
	// 	for ix_row := 0; ix_row < m_t.nrows; ix_row++ {
	// 		m_t.data[ix_col*m_t.nrows+ix_row] = (byte)(10*ix_col + ix_row)
	// 	}
	// }
	return m_t
}

// Returns the number of columns to the left of the reflection line.
// If there is no reflection line, return -1
func find_vertical_reflection(m matrix_t) int {
	fmt.Println(m.data)

	fmt.Printf("Vertical reflection")
	fmt.Printf("Matrix: nrows=%d, ncols=%d\n", m.nrows, m.ncols)

	// Start after the first column.
	// Iterate rows simultanously, to the left and to the right.
	// As long as they are identical, keep iterating.
	// If one of the iterators goes out-of-bounds, we have found a pattern

	ix_reflection := -1

	for col_start := 0; col_start < m.ncols-1; col_start++ {
		fmt.Printf("Starting at col %d\n", col_start)

		// If we can exhaust the iteration with all identical columns, we found a reflection:
		columns_identical := true
		for col_l, col_r := col_start, col_start+1; col_l >= 0 && col_r < m.ncols; col_l, col_r = col_l-1, col_r+1 {
			// Check if columns it_left and it_right are identical
			for ix_row := 0; ix_row < m.nrows; ix_row++ {
				columns_identical = columns_identical && (m.data[ix_row*m.ncols+col_l] == m.data[ix_row*m.ncols+col_r])
				// fmt.Printf("%c %c\n", m.data[ix_row*m.ncols+col_l], m.data[ix_row*m.ncols+col_r])
				// m[ix_row * m.ncols + ]
			}
			// fmt.Printf("Columns [%d, %d] identical: %v\n", col_l, col_r, columns_identical)
		}
		fmt.Printf("Iteration start: %d, columns identical: %v\n", col_start, columns_identical)
		if columns_identical {
			ix_reflection = col_start
		}
	}

	return ix_reflection + 1 // Move from 0-based to 1-based indexing
}

// func find_horizontal_reflection(m matrix_t) int {
// 	ix_reflection := -1

// 	for row_start := 0; row_start < m.nrows-1; row_start++ {
// 		fmt.Printf("Starting at row %d\n", row_start)

// 		// If we can exhaust the iteration with all identical columns, we found a reflection:
// 		rows_identical := true
// 		for row_l, row_r := row_start, row_start+1; row_l >= 0 && row_r < m.ncols; row_l, row_r = row_l-1, row_r+1 {
// 			// Check if columns it_left and it_right are identical
// 			for ix_row := 0; ix_row < m.nrows; ix_row++ {
// 				columns_identical = columns_identical && (m.data[ix_row*m.ncols+col_l] == m.data[ix_row*m.ncols+col_r])
// 				// fmt.Printf("%c %c\n", m.data[ix_row*m.ncols+col_l], m.data[ix_row*m.ncols+col_r])
// 				// m[ix_row * m.ncols + ]
// 			}
// 			// fmt.Printf("Columns [%d, %d] identical: %v\n", col_l, col_r, columns_identical)
// 		}
// 		fmt.Printf("Iteration start: %d, columns identical: %v\n", col_start, columns_identical)
// 		if columns_identical {
// 			ix_reflection = col_start
// 		}
// 	}

// 	return ix_reflection + 1 // Move from 0-based to 1-based indexing
// }

// Parses the input file and returns a slice of matrices
func parse_file(filename string) []matrix_t {
	arrays := []matrix_t{}
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	this_N := 0
	append_row := false
	row_count := 0
	new_matrix := []byte{}
	for scanner.Scan() {
		current_line := scanner.Text()
		if len(current_line) == 0 {
			// Append the current matrix
			arrays = append(arrays, matrix_t{data: new_matrix, nrows: row_count, ncols: this_N})

			this_N = 0
			row_count = 0
			// fmt.Printf("Blank line\n")
			continue
		}
		if len(current_line) != this_N {
			append_row = true
			this_N = len(current_line)
			// fmt.Printf("Setting append = %v, this_N=%d\n", append_row, this_N)
		}

		if append_row {
			// fmt.Printf("%02d: %s\n", row_count, current_line)
			bb := ([]byte)(current_line)
			for _, b := range bb {
				new_matrix = append(new_matrix, b)
			}
			row_count += 1
		}
	}

	return arrays
}

func main() {
	fmt.Printf("----- Advent of Code - day 13 -----\n")
	matrices := parse_file("input_13_test")

	for ix_m, m := range matrices {
		fmt.Printf("Matrix %02d: nrows: %02d, ncols: %02d\n", ix_m, m.nrows, m.ncols)
	}

}
