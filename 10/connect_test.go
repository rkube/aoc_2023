package main

import (
	"fmt"
	"log"
	"testing"
)

// Display in which a single pipe accepts connections
func Test_single_pipe(t *testing.T) {
	fmt.Printf("================================================================\n")
	fmt.Printf("      Single pipe connections\n")
	fmt.Printf("================================================================\n")
	pipes := []byte{'S', '|', '-', 'L', 'J', '7', 'F'}

	for _, p := range pipes {
		fmt.Printf("Pipe: %c  North: %v  East: %v  South: %v  West: %v\n", p, connects_N(p), connects_E(p), connects_S(p), connects_W(p))
	}
}

// Test if a pair of pipes connect
func Test_connect_pair(t *testing.T) {
	fmt.Printf("================================================================\n")
	fmt.Printf("      Pipe  pairwise connections\n")
	fmt.Printf("================================================================\n")

	pipes := []byte{'S', '|', '-', 'L', 'J', '7', 'F'}
	directions := []byte{'N', 'E', 'S', 'W'}

	for _, p1 := range pipes {
		for _, p2 := range pipes {
			p1 := position{x: 0, y: 0, pipe: p1, distance: 0}
			p2 := position{x: 0, y: 0, pipe: p2, distance: 0}
			fmt.Printf("Connecting '%c' and '%c': ", p1, p2)
			for _, dir := range directions {
				c := connect(p1, p2, dir)
				fmt.Printf("[%c]: %v\t", dir, c)
			}
			fmt.Printf("\n")
		}
	}

}

// Traverse the test maze
func Test_maze_step(t *testing.T) {
	fmt.Printf("================================================================\n")
	fmt.Printf("      Maze stepping\n")
	fmt.Printf("================================================================\n")

	maze, curr_pos := setup_maze("input_10_test", 5)
	prev_pos := curr_pos
	fmt.Printf("Start node: [%02d, %02d]: %c (%d)\n", curr_pos.x, curr_pos.y, curr_pos.pipe, curr_pos.distance)

	for {
		fmt.Printf("----------- Step ----------\n")
		next_pos, err := maze_step(curr_pos, prev_pos, maze)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("New node: [%02d, %02d]: %c (%d)\n", next_pos.x, next_pos.y, next_pos.pipe, next_pos.distance)
		// Step
		if next_pos.pipe == 'S' {
			break
		}
		prev_pos = curr_pos
		curr_pos = next_pos
	}
}

func Test_print_maze(t *testing.T) {
	fmt.Printf("================================================================\n")
	fmt.Printf("      Maze filling\n")
	fmt.Printf("================================================================\n")
	fmt.Printf("\n\n")

	len := 23
	maze, curr_pos := setup_maze("input_10_ex3", len)
	trace_arr := make([]byte, len*len)

	// Initialize the tracing array with all '.' for readability
	for ix, _ := range trace_arr {
		trace_arr[ix] = '.'
	}
	// for iy := 0; iy < len; iy++ {
	// 	for ix := 0; ix < len; ix++ {
	// 		fmt.Printf("%c", trace_arr[iy*len+ix])
	// 	}
	// 	fmt.Printf("\n")
	// }

	// Traverse the maze, mark visited elements withh O
	prev_pos := curr_pos
	// next_pos := curr_pos

	trace_arr[curr_pos.y*len+curr_pos.x] = '7' // maze[curr_pos.y][curr_pos.x]
	for {
		// fmt.Printf("----------- Step ----------\n")
		next_pos, err := maze_step(curr_pos, prev_pos, maze)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("New node: [%02d, %02d]: %c (%d)\n", next_pos.x, next_pos.y, next_pos.pipe, next_pos.distance)
		// Step
		if next_pos.pipe == 'S' {
			break
		}
		prev_pos = curr_pos
		curr_pos = next_pos
		trace_arr[curr_pos.y*len+curr_pos.x] = maze[curr_pos.y][curr_pos.x]
	}

	for iy := 0; iy < len; iy++ {
		for ix := 0; ix < len; ix++ {
			fmt.Printf("%c", trace_arr[iy*len+ix])
		}
		fmt.Printf("\t%s\n", maze[iy])
		// fmt.Printf("\n")
	}

	I_total := 0
	O_total := 0
	for iy := 0; iy < len; iy++ {
		_, I_line, O_line := ray_casting(trace_arr[iy*len : (iy+1)*len-1])
		I_total += I_line
		O_total += O_line
	}

	fmt.Printf("Internal area: %d\n", I_total)

}
