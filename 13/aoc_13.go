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
	new_matrix := make([]byte, 200)
	for scanner.Scan() {
		current_line := scanner.Text()
		if len(current_line) == 0 {
			// Append the current matrix
			arrays = append(arrays, matrix_t{data: new_matrix, nrows: row_count, ncols: this_N})

			this_N = 0
			row_count = 0
			fmt.Printf("Blank line\n")
			continue
		}
		if len(current_line) != this_N {
			append_row = true
			this_N = len(current_line)
			fmt.Printf("Setting append = %v, this_N=%d\n", append_row, this_N)
		}

		if append_row {
			fmt.Printf("%02d: %s\n", row_count, current_line)
			bb := ([]byte)(current_line)
			for _, b := range bb {
				new_matrix = append(new_matrix, b)
			}
			row_count += 1
		}

		// if row_count == this_N-1 {
		// 	this_N = 0
		// 	row_count = 0
		// }

		// if new_matrix {
		// 	current_line := scanner.Text()
		// 	N := len(current_line)
		// 	for row_count := 1; row_count < N; row_count++ {
		// 		current_line = scanner.Text()
		// 		fmt.Printf("%d: %s\n", row_count, current_line)
		// 	}
		// 	new_matrix = false
		// }

		// if new_matrix {
		// 	current_line := scanner.Text()
		// 	fmt.Printf("%s\n", current_line)
		// 	N := len(current_line)
		// 	matrix := make([]byte, N)
		// 	for row_count := 1; row_count < N; row_count++ {
		// 		// bb := ([]byte)(current_line)
		// 		for _, bb := range ([]byte)(current_line) {
		// 			matrix = append(matrix, bb)
		// 		}
		// 		current_line = scanner.Text()
		// 		fmt.Printf("%s\n", current_line)
		// 	}
		// 	new_matrix = false
		// 	fmt.Println(matrix)
		// 	arrays = append(arrays, matrix)
		// }
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
