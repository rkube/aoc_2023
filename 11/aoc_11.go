package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coords_t struct {
	row int
	col int
	id  int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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
func duplicate_row(chart []byte, nrow int, ncol int, duplicate_rows []int, n_dup int) []byte {
	new_nrow := nrow + n_dup*len(duplicate_rows)
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
				// Insert n_dup new rows at this position. Start iteration variable
				// at the next line, ix_new_row+1.
				// for ; ix_new_row < ix_new_row+n_dup; ix_new_row++ {
				for ix_this := ix_new_row + 1; ix_this < ix_new_row+n_dup+1; ix_this++ {
					// fmt.Printf("Duplicating: (n_dup = %d) ix_row = %d -> ix_new_row = %d\n", n_dup, ix_row, ix_this)
					for ix_col := 0; ix_col < ncol; ix_col++ {
						new_chart[ix_this*ncol+ix_col] = chart[ix_row*ncol+ix_col]
					}
					// fmt.Printf("-done\n")
				}
				ix_new_row += n_dup
			}
		}
		ix_new_row += 1
	}
	return new_chart
}

func duplicate_col(chart []byte, nrow int, ncol int, duplicate_cols []int, n_dup int) []byte {
	new_ncol := ncol + n_dup*len(duplicate_cols)
	new_chart := make([]byte, nrow*new_ncol)

	ix_new_col := 0
	for ix_row := 0; ix_row < nrow; ix_row++ {
		ix_new_col = 0
		for ix_col := 0; ix_col < ncol; ix_col++ {
			new_chart[ix_row*new_ncol+ix_new_col] = chart[ix_row*ncol+ix_col]
			// Check if current column needs to be duplicated
			for _, v := range duplicate_cols {
				if ix_col == v {
					for ix_this := ix_new_col + 1; ix_this < ix_new_col+n_dup+1; ix_this++ {
						new_chart[ix_row*new_ncol+ix_this] = chart[ix_row*ncol+ix_col]
					}
					ix_new_col += n_dup
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
			chart[ix_row*ncol+ix_col] = current_line[ix_col]
		}
		ix_row += 1
	}

	return chart
}

func galaxy_pos(chart []byte, nrow int, ncol int) []coords_t {
	galaxy_coords := []coords_t{}
	galaxy_count := 0
	for ix_row := 0; ix_row < nrow; ix_row++ {
		for ix_col := 0; ix_col < ncol; ix_col++ {
			if chart[ix_row*ncol+ix_col] == '#' {
				galaxy_coords = append(galaxy_coords, coords_t{row: ix_row, col: ix_col, id: galaxy_count})
				galaxy_count += 1
			}
		}
	}
	return galaxy_coords
}

// Count how many values in the array skips are smaller than coord
func smaller_in_array(coord int, skips []int) int {
	num_smaller := 0
	for _, v := range skips {
		if v < coord {
			num_smaller += 1
		}
	}
	return num_smaller
}

// Calculate the shifted position
// Shift the row/col coordinate of a galaxy by n_dup
// for each row/col that precedes the galaxy position
func shift_position(galaxy_coords coords_t, duplicate_rows []int, duplicate_cols []int, n_dup int) coords_t {
	// fmt.Printf("Original coordinate: [%03d,%03d]\n", galaxy_coords.row, galaxy_coords.col)
	num_shift_rows := smaller_in_array(galaxy_coords.row, duplicate_rows)
	num_shift_cols := smaller_in_array(galaxy_coords.col, duplicate_cols)
	// fmt.Printf("\tShift rows: %d\n", num_shift_rows*n_dup)
	// fmt.Printf("\tShift cols: %d\n", num_shift_cols*n_dup)

	return coords_t{id: galaxy_coords.id, row: galaxy_coords.row + n_dup*num_shift_rows,
		col: galaxy_coords.col + n_dup*num_shift_cols}
}

func mahattan_dist(g1 coords_t, g2 coords_t) int {
	return Abs(g1.col-g2.col) + Abs(g1.row-g2.row)
}

func main() {
	fmt.Printf("Advent of Code 2023 - Day 11\n")

	nrow, ncol := 140, 140
	chart := load_map("input_11", nrow, ncol)
	empty_rows := find_empty_rows(chart, nrow, ncol)
	empty_cols := find_empty_cols(chart, nrow, ncol)
	n_dup := 1
	// new_nrow := nrow + n_dup*len(empty_rows)
	// new_ncol := ncol + n_dup*len(empty_cols)

	// chart_exp_rows := duplicate_row(chart, nrow, ncol, empty_rows, n_dup)
	// chart_exp := duplicate_col(chart_exp_rows, new_nrow, ncol, empty_cols, n_dup)

	galaxy_coords := galaxy_pos(chart, nrow, ncol)
	galaxy_coords_shift := []coords_t{}
	fmt.Printf("%d galaxies\n", len(galaxy_coords))
	for _, g := range galaxy_coords {
		galaxy_coords_shift = append(galaxy_coords_shift, shift_position(g, empty_rows, empty_cols, n_dup))
	}

	sum_dist := 0
	for ix_p1 := 0; ix_p1 < len(galaxy_coords)-1; ix_p1++ {
		for ix_p2 := ix_p1 + 1; ix_p2 < len(galaxy_coords); ix_p2++ {
			// sum_dist += mahattan_dist(galaxy_coords[ix_p1], galaxy_coords[ix_p2])
			// fmt.Printf("%d %d - sum_dist = %d\n", ix_p1, ix_p2, sum_dist)
			sum_dist += mahattan_dist(galaxy_coords_shift[ix_p1], galaxy_coords_shift[ix_p2])
		}
	}

	fmt.Printf("Part 1: Sum of distances: %d (correct: 9565386)\n", sum_dist)

	// Work with sparse representation here. Otherwise we run out of memory
	n_dup = 1_000_000
	galaxy_coords_shift = []coords_t{}

	for _, g := range galaxy_coords {
		galaxy_coords_shift = append(galaxy_coords_shift, shift_position(g, empty_rows, empty_cols, n_dup))
	}

	sum_dist = 0
	for ix_p1 := 0; ix_p1 < len(galaxy_coords_shift)-1; ix_p1++ {
		for ix_p2 := ix_p1 + 1; ix_p2 < len(galaxy_coords_shift); ix_p2++ {
			sum_dist += mahattan_dist(galaxy_coords_shift[ix_p1], galaxy_coords_shift[ix_p2])
			// fmt.Printf("%d %d - sum_dist = %d\n", ix_p1, ix_p2, sum_dist)
		}
	}
	// Guess: 874799690595 - wrong (too high)
	fmt.Printf("Part 2: Sum of distances: %d (correct: )\n", sum_dist)

}
