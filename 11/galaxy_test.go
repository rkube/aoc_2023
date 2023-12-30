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
}

func Test_expand_array(t *testing.T) {
	fmt.Printf("\n\n--------- Expanding arrays --------\n\n")
	nrow, ncol := 10, 10
	chart := load_map("input_11_test1", nrow, ncol)
	// fmt.Printf("Original size: %d\n", len(chart))

	// for ix_row := 0; ix_row < nrow; ix_row++ {
	// 	fmt.Printf("%d: ", ix_row)
	// 	for ix_col := 0; ix_col < ncol; ix_col++ {
	// 		fmt.Printf("%c", chart[ix_row*ncol+ix_col])
	// 	}
	// 	fmt.Printf("\n")
	// }

	empty_rows := find_empty_rows(chart, nrow, ncol)
	empty_cols := find_empty_cols(chart, nrow, ncol)
	new_nrow := nrow + len(empty_rows)
	new_ncol := ncol + len(empty_cols)
	n_dup := 1

	// Expand rows
	chart_exp_rows := duplicate_row(chart, nrow, ncol, empty_rows, n_dup)

	// fmt.Printf("\n\n---- Expanded rows chart - size: %d\n", len(chart_exp_rows))
	// for ix_row := 0; ix_row < new_nrow; ix_row++ {
	// 	fmt.Printf("%02d: ", ix_row)
	// 	for ix_col := 0; ix_col < ncol; ix_col++ {
	// 		fmt.Printf("%c", chart_exp_rows[ix_row*ncol+ix_col])
	// 	}
	// 	fmt.Printf("\n")
	// }

	// // Expand columns

	// chart_exp_cols := duplicate_col(chart, nrow, ncol, empty_cols)

	// fmt.Printf("\n\n---- Expanded cols chart - size: %d\n", len(chart_exp_cols))
	// for ix_row := 0; ix_row < nrow; ix_row++ {
	// 	fmt.Printf("%02d: ", ix_row)
	// 	for ix_col := 0; ix_col < new_ncol; ix_col++ {
	// 		fmt.Printf("%c", chart_exp_cols[ix_row*new_ncol+ix_col])
	// 	}
	// 	fmt.Printf("\n")
	// }

	chart_exp := duplicate_col(chart_exp_rows, new_nrow, ncol, empty_cols, n_dup)

	// fmt.Printf("\n\n---- Expanding rows and columns - size: %d\n\n", len(chart_exp))
	// for ix_row := 0; ix_row < new_nrow; ix_row++ {
	// 	fmt.Printf("%02d: ", ix_row)
	// 	for ix_col := 0; ix_col < new_ncol; ix_col++ {
	// 		fmt.Printf("%c", chart_exp[ix_row*new_ncol+ix_col])
	// 	}
	// 	fmt.Printf("\n")
	// }

	// fmt.Printf("new_nrow: %d, new_ncol: %d\n", new_nrow, new_ncol)

	chart_cmp := load_map("input_11_test2", new_nrow, new_ncol)
	for ix_row := 0; ix_row < new_nrow; ix_row++ {
		for ix_col := 0; ix_col < new_ncol; ix_col++ {
			if chart_cmp[ix_row*new_ncol+ix_col] != chart_exp[ix_row*new_ncol+ix_col] {
				t.Fatalf("Mismatch at [%d,%d]: ours: %c, true: %c\n", ix_row, ix_col, chart_exp[ix_row*new_ncol+ix_col], chart_cmp[ix_row*new_ncol+ix_col])
			}
		}
	}

}

// func Test_expand_array_2(t *testing.T) {
// 	fmt.Printf("\n\n--------- Expanding arrays (arbitrary size)--------\n\n")
// 	nrow, ncol := 10, 10
// 	chart := load_map("input_11_test1", nrow, ncol)

// 	empty_rows := find_empty_rows(chart, nrow, ncol)
// 	empty_cols := find_empty_cols(chart, nrow, ncol)

// 	fmt.Printf("\n\n --- Original chart ---- \n\n")
// 	for ix_row := 0; ix_row < nrow; ix_row++ {
// 		fmt.Printf("%02d: ", ix_row)
// 		for ix_col := 0; ix_col < ncol; ix_col++ {
// 			fmt.Printf("%c", chart[ix_row*ncol+ix_col])
// 		}
// 		fmt.Printf("\n")
// 	}

// 	n_dup := 1
// 	fmt.Printf("\n\n---- Expanded rows chart - n_dup: %d\nExpanding rows: ", n_dup)
// 	for _, v := range empty_rows {
// 		fmt.Printf("%02d ", v)
// 	}
// 	fmt.Printf("\n")
// 	new_nrow := nrow + n_dup*len(empty_rows)
// 	chart_exp_rows := duplicate_row(chart, nrow, ncol, empty_rows, n_dup)

// 	for ix_row := 0; ix_row < new_nrow; ix_row++ {
// 		fmt.Printf("%02d: ", ix_row)
// 		for ix_col := 0; ix_col < ncol; ix_col++ {
// 			fmt.Printf("%c", chart_exp_rows[ix_row*ncol+ix_col])
// 		}
// 		fmt.Printf("\n")
// 	}

// 	fmt.Printf("\n\n---- Expanded cols chart - n_dup: %d\nExpanding columns: ", n_dup)
// 	for _, v := range empty_cols {
// 		fmt.Printf("%02d ", v)
// 	}
// 	fmt.Printf("\n")
// 	new_ncol := ncol + n_dup*len(empty_cols)
// 	chart_exp_cols := duplicate_col(chart, nrow, ncol, empty_cols, n_dup)
// 	for ix_row := 0; ix_row < nrow; ix_row++ {
// 		fmt.Printf("%02d: ", ix_row)
// 		for ix_col := 0; ix_col < new_ncol; ix_col++ {
// 			fmt.Printf("%c", chart_exp_cols[ix_row*new_ncol+ix_col])
// 		}
// 		fmt.Printf("\n")
// 	}

// }

func Test_part1(t *testing.T) {
	fmt.Printf("\n\n--------- Testing part 1--------\n\n")
	nrow, ncol := 10, 10
	chart := load_map("input_11_test1", nrow, ncol)
	empty_rows := find_empty_rows(chart, nrow, ncol)
	empty_cols := find_empty_cols(chart, nrow, ncol)
	new_nrow := nrow + len(empty_rows)
	new_ncol := ncol + len(empty_cols)
	n_dup := 1

	chart_exp_rows := duplicate_row(chart, nrow, ncol, empty_rows, n_dup)
	chart_exp := duplicate_col(chart_exp_rows, new_nrow, ncol, empty_cols, n_dup)

	galaxy_coords := galaxy_pos(chart_exp, new_nrow, new_ncol)

	// fmt.Printf("Found %d galaxies:\n", len(galaxy_coords))
	// for _, g := range galaxy_coords {
	// 	fmt.Printf("%d: [%02d,%02d]\n", g.id, g.row, g.col)
	// }

	ix1, ix2, true_d := 4, 8, 9
	// fmt.Printf("Distance %d [%02d,%02d] - %d [%02d,%02d] %d\n", galaxy_coords[ix1].id, galaxy_coords[ix1].row, galaxy_coords[ix1].col, galaxy_coords[ix2].id, galaxy_coords[ix2].row, galaxy_coords[ix2].col, mahattan_dist(galaxy_coords[ix1], galaxy_coords[ix2]))
	if mahattan_dist(galaxy_coords[ix1], galaxy_coords[ix2]) != true_d {
		t.Fatalf("Mismatch: Distance(%d [%02d,%02d], %d [%02d,%02d]) = %d but should be %d\n", galaxy_coords[ix1].id, galaxy_coords[ix1].row, galaxy_coords[ix1].col, galaxy_coords[ix2].id, galaxy_coords[ix2].row, galaxy_coords[ix2].col, mahattan_dist(galaxy_coords[ix1], galaxy_coords[ix2]), true_d)
	}
	ix1, ix2, true_d = 0, 6, 15
	// fmt.Printf("Distance %d [%02d,%02d] - %d [%02d,%02d] %d\n", galaxy_coords[ix1].id, galaxy_coords[ix1].row, galaxy_coords[ix1].col, galaxy_coords[ix2].id, galaxy_coords[ix2].row, galaxy_coords[ix2].col, mahattan_dist(galaxy_coords[ix1], galaxy_coords[ix2]))
	if mahattan_dist(galaxy_coords[ix1], galaxy_coords[ix2]) != true_d {
		t.Fatalf("Mismatch: Distance(%d [%02d,%02d], %d [%02d,%02d]) = %d but should be %d\n", galaxy_coords[ix1].id, galaxy_coords[ix1].row, galaxy_coords[ix1].col, galaxy_coords[ix2].id, galaxy_coords[ix2].row, galaxy_coords[ix2].col, mahattan_dist(galaxy_coords[ix1], galaxy_coords[ix2]), true_d)
	}

	ix1, ix2, true_d = 2, 5, 17
	// fmt.Printf("Distance %d [%02d,%02d] - %d [%02d,%02d] %d\n", galaxy_coords[ix1].id, galaxy_coords[ix1].row, galaxy_coords[ix1].col, galaxy_coords[ix2].id, galaxy_coords[ix2].row, galaxy_coords[ix2].col, mahattan_dist(galaxy_coords[ix1], galaxy_coords[ix2]))
	if mahattan_dist(galaxy_coords[ix1], galaxy_coords[ix2]) != true_d {
		t.Fatalf("Mismatch: Distance(%d [%02d,%02d], %d [%02d,%02d]) = %d but should be %d\n", galaxy_coords[ix1].id, galaxy_coords[ix1].row, galaxy_coords[ix1].col, galaxy_coords[ix2].id, galaxy_coords[ix2].row, galaxy_coords[ix2].col, mahattan_dist(galaxy_coords[ix1], galaxy_coords[ix2]), true_d)
	}

	ix1, ix2, true_d = 7, 8, 5
	// fmt.Printf("Distance %d [%02d,%02d] - %d [%02d,%02d] %d\n", galaxy_coords[ix1].id, galaxy_coords[ix1].row, galaxy_coords[ix1].col, galaxy_coords[ix2].id, galaxy_coords[ix2].row, galaxy_coords[ix2].col, mahattan_dist(galaxy_coords[ix1], galaxy_coords[ix2]))
	if mahattan_dist(galaxy_coords[ix1], galaxy_coords[ix2]) != true_d {
		t.Fatalf("Mismatch: Distance(%d [%02d,%02d], %d [%02d,%02d]) = %d but should be %d\n", galaxy_coords[ix1].id, galaxy_coords[ix1].row, galaxy_coords[ix1].col, galaxy_coords[ix2].id, galaxy_coords[ix2].row, galaxy_coords[ix2].col, mahattan_dist(galaxy_coords[ix1], galaxy_coords[ix2]), true_d)
	}

	sum_dist := 0
	for ix_p1 := 0; ix_p1 < len(galaxy_coords)-1; ix_p1++ {
		for ix_p2 := ix_p1 + 1; ix_p2 < len(galaxy_coords); ix_p2++ {
			sum_dist += mahattan_dist(galaxy_coords[ix_p1], galaxy_coords[ix_p2])
		}
	}

	if sum_dist != 374 {
		t.Fatalf("Sum of pair distances = %d, but should be 374\n", sum_dist)
	}
	fmt.Printf("Part 1: Sum of pair distances = %d\n", sum_dist)

	// Approach 2: just find out by how much we have to shift each galaxy position
	// when accounting for empty rows/cols
	galaxy_coords_orig := galaxy_pos(chart, nrow, ncol)
	galaxy_coords_shift := []coords_t{}
	// fmt.Printf("Found %d galaxies:\n", len(galaxy_coords_orig))
	for _, g := range galaxy_coords_orig {
		// fmt.Printf("Original: %d - [%02d,%02d]\n", g.id, g.row, g.col)
		galaxy_coords_shift = append(galaxy_coords_shift, shift_position(g, empty_rows, empty_cols, 1))
		// fmt.Printf("Shifted: %d - [%02d,%02d]\n", galaxy_coords_shift[ix_g].id, galaxy_coords_shift[ix_g].row, galaxy_coords_shift[ix_g].col)
	}

	for ix_g, g := range galaxy_coords_shift {
		// fmt.Printf("From dense:  id=%d, row=%d, col=%d\n", galaxy_coords[ix_g].id, galaxy_coords[ix_g].row, galaxy_coords[ix_g].col)
		// fmt.Printf("From sparse: id=%d, row=%d, col=%d\n", g.id, g.row, g.col)
		if (g.id != galaxy_coords[ix_g].id) || (g.col != galaxy_coords[ix_g].col) || (g.row != galaxy_coords[ix_g].row) {
			fmt.Printf("Mismatch\n")
		}
	}

	sum_dist = 0
	for ix_p1 := 0; ix_p1 < len(galaxy_coords)-1; ix_p1++ {
		for ix_p2 := ix_p1 + 1; ix_p2 < len(galaxy_coords); ix_p2++ {
			sum_dist += mahattan_dist(galaxy_coords_shift[ix_p1], galaxy_coords_shift[ix_p2])
		}
	}
	if sum_dist != 374 {
		t.Fatalf("Sum of pair distances = %d, but should be 374\n", sum_dist)
	}
	fmt.Printf("Part 1: Sum of pair distances = %d\n", sum_dist)

}

func Test_part2(t *testing.T) {
	fmt.Printf("\n\n --------- Testing Part 2 --------\n\n")

	nrow, ncol := 10, 10
	chart := load_map("input_11_test1", nrow, ncol)
	empty_rows := find_empty_rows(chart, nrow, ncol)
	empty_cols := find_empty_cols(chart, nrow, ncol)

	n_dup_list := []int{10 - 1, 100 - 1}
	answer_list := []int{1030, 8410}

	galaxy_coords := galaxy_pos(chart, nrow, ncol)
	fmt.Printf("Found %d galaxies:\n", len(galaxy_coords))

	for ix_t, _ := range n_dup_list {
		n_dup := n_dup_list[ix_t]

		galaxy_coords_shift := []coords_t{}

		for _, g := range galaxy_coords {
			galaxy_coords_shift = append(galaxy_coords_shift, shift_position(g, empty_rows, empty_cols, n_dup))
		}

		sum_dist := 0
		for ix_p1 := 0; ix_p1 < len(galaxy_coords_shift)-1; ix_p1++ {
			for ix_p2 := ix_p1 + 1; ix_p2 < len(galaxy_coords_shift); ix_p2++ {
				sum_dist += mahattan_dist(galaxy_coords_shift[ix_p1], galaxy_coords_shift[ix_p2])
			}
		}
		if sum_dist != answer_list[ix_t] {
			t.Fatalf("Sum of pair distances = %d, but should be %d\n", sum_dist, answer_list[ix_t])
		}
		fmt.Printf("n_dup = %d - Sum of pair distances = %d, should be %d\n", n_dup, sum_dist, answer_list[ix_t])
	}

}

// func Test_part2(t *testing.T) {
// 	fmt.Printf("\n\n--------- Testing Part 2 --------\n\n")
// 	nrow, ncol := 10, 10
// 	chart := load_map("input_11_test1", nrow, ncol)
// 	empty_rows := find_empty_rows(chart, nrow, ncol)
// 	empty_cols := find_empty_cols(chart, nrow, ncol)
// 	n_dup := 9
// 	new_nrow := nrow + n_dup*len(empty_rows)
// 	new_ncol := ncol + n_dup*len(empty_cols)

// 	chart_exp_rows := duplicate_row(chart, nrow, ncol, empty_rows, n_dup)
// 	chart_exp := duplicate_col(chart_exp_rows, new_nrow, ncol, empty_cols, n_dup)

// 	galaxy_coords := galaxy_pos(chart_exp, new_nrow, new_ncol)

// 	fmt.Printf("Found %d galaxies:\n", len(galaxy_coords))
// 	for _, g := range galaxy_coords {
// 		fmt.Printf("%d: [%02d,%02d]\n", g.id, g.row, g.col)
// 	}

// 	sum_dist := 0
// 	for ix_p1 := 0; ix_p1 < len(galaxy_coords)-1; ix_p1++ {
// 		for ix_p2 := ix_p1 + 1; ix_p2 < len(galaxy_coords); ix_p2++ {
// 			sum_dist += mahattan_dist(galaxy_coords[ix_p1], galaxy_coords[ix_p2])
// 		}
// 	}
// 	if sum_dist != 1030 {
// 		t.Fatalf("Sum of pair distances = %d, but should be 1030\n", sum_dist)
// 	}

// 	n_dup = 99
// 	new_nrow = nrow + n_dup*len(empty_rows)
// 	new_ncol = ncol + n_dup*len(empty_cols)

// 	chart_exp_rows = duplicate_row(chart, nrow, ncol, empty_rows, n_dup)
// 	chart_exp = duplicate_col(chart_exp_rows, new_nrow, ncol, empty_cols, n_dup)

// 	galaxy_coords = galaxy_pos(chart_exp, new_nrow, new_ncol)

// 	sum_dist = 0
// 	for ix_p1 := 0; ix_p1 < len(galaxy_coords)-1; ix_p1++ {
// 		for ix_p2 := ix_p1 + 1; ix_p2 < len(galaxy_coords); ix_p2++ {
// 			sum_dist += mahattan_dist(galaxy_coords[ix_p1], galaxy_coords[ix_p2])
// 		}
// 	}
// 	if sum_dist != 8410 {
// 		t.Fatalf("Sum of pair distances = %d, but should be 8410\n", sum_dist)
// 	}

// }
