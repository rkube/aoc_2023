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

	// Expand rows
	chart_exp_rows := duplicate_row(chart, nrow, ncol, empty_rows)

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

	chart_exp := duplicate_col(chart_exp_rows, new_nrow, ncol, empty_cols)
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

func Test_galaxy_dist(t *testing.T) {
	fmt.Printf("\n\n--------- Expanding arrays --------\n\n")
	nrow, ncol := 10, 10
	chart := load_map("input_11_test1", nrow, ncol)
	empty_rows := find_empty_rows(chart, nrow, ncol)
	empty_cols := find_empty_cols(chart, nrow, ncol)
	new_nrow := nrow + len(empty_rows)
	new_ncol := ncol + len(empty_cols)

	chart_exp_rows := duplicate_row(chart, nrow, ncol, empty_rows)
	chart_exp := duplicate_col(chart_exp_rows, new_nrow, ncol, empty_cols)

	galaxy_coords := galaxy_pos(chart_exp, new_nrow, new_ncol)

	fmt.Printf("Found %d galaxies:\n", len(galaxy_coords))
	for _, g := range galaxy_coords {
		fmt.Printf("%d: [%02d,%02d]\n", g.id, g.row, g.col)
	}

	ix1, ix2 := 4, 8
	fmt.Printf("Distance %d [%02d,%02d] - %d [%02d,%02d] %d\n", galaxy_coords[ix1].id, galaxy_coords[ix1].row, galaxy_coords[ix1].col, galaxy_coords[ix2].id, galaxy_coords[ix2].row, galaxy_coords[ix2].col, mahattan_dist(galaxy_coords[ix1], galaxy_coords[ix2]))
	ix1, ix2 = 0, 6
	fmt.Printf("Distance %d [%02d,%02d] - %d [%02d,%02d] %d\n", galaxy_coords[ix1].id, galaxy_coords[ix1].row, galaxy_coords[ix1].col, galaxy_coords[ix2].id, galaxy_coords[ix2].row, galaxy_coords[ix2].col, mahattan_dist(galaxy_coords[ix1], galaxy_coords[ix2]))
	ix1, ix2 = 2, 5
	fmt.Printf("Distance %d [%02d,%02d] - %d [%02d,%02d] %d\n", galaxy_coords[ix1].id, galaxy_coords[ix1].row, galaxy_coords[ix1].col, galaxy_coords[ix2].id, galaxy_coords[ix2].row, galaxy_coords[ix2].col, mahattan_dist(galaxy_coords[ix1], galaxy_coords[ix2]))
	ix1, ix2 = 7, 8
	fmt.Printf("Distance %d [%02d,%02d] - %d [%02d,%02d] %d\n", galaxy_coords[ix1].id, galaxy_coords[ix1].row, galaxy_coords[ix1].col, galaxy_coords[ix2].id, galaxy_coords[ix2].row, galaxy_coords[ix2].col, mahattan_dist(galaxy_coords[ix1], galaxy_coords[ix2]))

}
