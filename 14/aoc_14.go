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

// Load the input file and store in row-major matrix
func parse_input(filename string) matrix_t {
	array := make([]byte, 0, 10_000)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	n_row := 0
	n_col := 0
	for scanner.Scan() {
		n_col = 0
		current_line := scanner.Text()

		for _, b := range ([]byte)(current_line) {
			array = append(array, b)
			n_col += 1
		}
		n_row += 1
	}

	matrix := matrix_t{nrows: n_row, ncols: n_col, data: array}
	// fmt.Printf("n_row=%d, n_col=%d\n", n_row, n_col)
	return matrix
}

// Returns a slice with indices of c in a row of the matrix m
func search_col(m matrix_t, c byte, col int) []int {
	if col >= m.ncols {
		err_str := fmt.Sprintf("Can't access column %d of %dx%d matrix\n", col, m.nrows, m.ncols)
		log.Fatal(err_str)
	}
	c_ix := make([]int, 0, m.nrows)

	for ix_row := 0; ix_row < m.nrows; ix_row++ {
		if m.data[ix_row*m.ncols+col] == c {
			c_ix = append(c_ix, ix_row)
		}
	}
	return c_ix
}

// func search_row(m matrix_t, c byte, row int) []int {
// 	if row >= m.nrows {
// 		err_str := fmt.Sprintf("Can't access row %d of %dx%d matrix\n", row, m.nrows, m.ncols)
// 		log.Fatal(err_str)
// 	}
// 	c_ix := make([]int, 0, m.ncols)
// 	for ix_col := 0; ix_col < m.ncols; ix_col++ {
// 		if m.data[row*m.ncols+ix_col] == c {
// 			c_ix = append(c_ix, ix_col)
// 		}
// 	}
// 	return c_ix
// }

// Rotate a matrix clock-wise 90 degrees
func rotate_cw(m matrix_t) matrix_t {
	// Creat transpose matrix
	m_t := matrix_t{nrows: m.ncols, ncols: m.nrows, data: make([]byte, m.nrows*m.ncols)}

	// Rotate matrix clock-wise:
	//
	// 0 1 2    6 3 0
	// 3 4 5 => 7 4 1
	// 6 7 8    8 5 2

	for ix_row := 0; ix_row < m.nrows; ix_row++ {
		for ix_col := 0; ix_col < m.ncols; ix_col++ {
			m_t.data[ix_row*m.ncols+ix_col] = m.data[(m.nrows-1-ix_col)*m.ncols+ix_row]
		}
	}
	return m_t
}

// Rotates a map of positions clockwise
// O_pos is a map
// column -> [r0, r1, ...r_n] that indiciates a set of rows, as indexed by a column.
// The output is the same map, but valid for when the matrix was rotate clockwise
//
// .OO     O..
// .O.  => .OO
// O.O     O.O
//
// O_pos[0] = {2}         O_pos[0] = {0}
// O_pos[1] = {0, 1}   => O_pos[1] = {1}
// O_pos[2] = {0, 2}   => O_pos[2] = {1, 2}

func rotate_pos_vec(O_pos map[int][]int) map[int][]int {

}

// Sort O's in between the position of #'s within a single column.
// Returns a vector of the sorted O's
// pos_O: indices where a round boulder is located
// pos_S: indices where a square boulder is located
// nrows: Number of rows gives the lower boundary
func sort_Os(pos_O []int, pos_S_in []int, nrows int) []int {
	sorted_Os := make([]int, 0, len(pos_O))
	// Iterate over the positions of square boulders. Prepend -1 and append ncols to close the boundaries
	pos_S := make([]int, len(pos_S_in)+2, len(pos_S_in)+2)
	pos_S[0] = -1 // Lower bound is -1
	for ix := 0; ix < len(pos_S_in); ix++ {
		pos_S[ix+1] = pos_S_in[ix]
	}
	pos_S[len(pos_S)-1] = nrows // Upper bound is nrows

	// In each interval of S, we are going to find the boulders to place here
	// Since pos_O is sorted, we can continue from previous step when iterating
	// over the intervals of S.
	ix_O := 0 // Keeps track of O's that have been sorted
	for ix_S := 0; ix_S < len(pos_S)-1; ix_S++ {
		S_lower, S_upper := pos_S[ix_S], pos_S[ix_S+1]
		num_inserted := 0
		for { // Sort O's into the current [S_lower:S_upper]
			if ix_O >= len(pos_O) {
				break
			}
			if pos_O[ix_O] < S_upper {
				sorted_Os = append(sorted_Os, S_lower+num_inserted+1)
				num_inserted += 1
				ix_O++ // If we sorted this O, proceed to the next O.
			} else { // No O's left to sort
				break
			}

		}

		if ix_O >= len(pos_O) { // No O's left to sort
			break
		}
	}
	return sorted_Os
}

func score_row(pos_O []int, nrows int) int {
	sum_row := 0
	for _, p := range pos_O {
		sum_row += nrows - p
	}
	return sum_row
}

func main() {
	m := parse_input("input_14")
	fmt.Printf("%d %d", m.nrows, m.ncols)
	score := 0
	for ix_col := 0; ix_col < m.ncols; ix_col++ {
		pos_S := search_col(m, '#', ix_col)
		pos_O := search_col(m, 'O', ix_col)
		sorted_Os := sort_Os(pos_O, pos_S, m.nrows)
		score += score_row(sorted_Os, m.nrows)
	}
	fmt.Printf("Score: %d\n", score)
}
