package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type directions_t struct {
	left  string
	right string
}

// Access the fields of directions_t with either 'L' or 'R'
func (d directions_t) target(pos rune) string {

	if pos == 'L' {
		return d.left
	} else if pos == 'R' {
		return d.right
	} else {
		err_str := fmt.Sprintf("Unknown directions: %x\n", pos)
		log.Fatal(errors.New(err_str))
		return ""
	}
}

func find_cycles(places map[string]directions_t, directions string, node_start string, node_end string) int {
	cycles := 0        // iterations over total directions
	curr := node_start // Starting position
	// Traverse the map using the entire sequence of iterations
	// Check if we end up in "ZZZ". If yes, we are finished.
	// If not, repeat the sequnce until we arrive at ZZZ
	for {
		for _, next := range directions {
			curr = places[curr].target(next)
		}
		cycles += 1

		if curr == node_end {
			break
		}

		if cycles > 1000 {
			break
		}
	}
	return cycles
}

// Iterate over a sequence of directions
func run_cycle(places map[string]directions_t, directions string, node_start string) string { //, node_end chan string) {
	curr := node_start
	for _, next := range directions {
		curr = places[curr].target(next)
	}
	// Put end node in channel
	// node_end <- curr
	return curr
}

func main() {
	fmt.Printf("Advent of code 2023 - Part 08\n")
	f, err := os.Open("input_08")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	ix_line := 0
	directions := ""

	places := make(map[string]directions_t, 800)

	for scanner.Scan() {
		current_line := scanner.Text()
		if ix_line == 0 {
			directions = current_line
		}
		if ix_line > 1 {
			places[current_line[0:3]] = directions_t{left: current_line[7:10], right: current_line[12:15]}
			// if current_line[2] == 'A' {
			// 	fmt.Printf("Starting node: %s\n", current_line[0:3])
			// } else if current_line[2] == 'Z' {
			// 	fmt.Printf("End node: %s\n", current_line[0:3])
			// }
		}
		ix_line += 1
	}
	defer f.Close()
	// cycles_part_1 := find_cycles(places, directions, "AAA", "ZZZ")

	// The total number of steps is the number of times we had to
	// iterate over the entire sequence of steps to come to ZZZ
	// multiplied by the length of that sequence.
	// fmt.Printf("Number of steps: %d - Correct answer: %d\n", cycles_part_1*len(directions), 19631)

	/////////////////////////////////////////////////////////////////
	// Part 2
	/////////////////////////////////////////////////////////////////

	// Relevant start and end nodes:
	// Starting node: MJA - HJZ cycle length: 71
	// Starting node: RGA - HPZ cycle length: 61
	// Starting node: JMA - FKZ cycle length: 78
	// Starting node: XHA - PKZ cycle length: 70
	// Starting node: DQA - DHZ cycle length: 58
	// Starting node: AAA - ZZZ cycle length: 67

	// List of channels that receive the current position of a mover after
	// one sequence of iterations
	// ch_list := map[int](chan string){0: make(chan string), 1: make(chan string), 2: make(chan string),
	// 	3: make(chan string), 4: make(chan string), 5: make(chan string)}

	// Current position of all movers. They start out at a position with A in the end
	start_list := map[int]string{0: "MJA", 1: "RGA", 2: "JMA", 3: "XHA", 4: "DQA", 5: "AAA"}
	// curr_list := map[int]string{0: "11A", 1: "22A"} // For the test case

	cycles := make([]int, len(start_list))
	current := start_list[0]
	i := 0
	// terminate := false
	// cycle_ctr := 1
	for ix_start, pos := range start_list {
		// current := start_list[ix_start]
		i = 0
		current = pos
		for {
			next := directions[i%len(directions)]
			current = places[current].target(rune(next))
			i += 1
			if current[2] == 'Z' {
				fmt.Printf("[%d] start %s: - After i=%d cycles at%s\n", ix_start, start_list[ix_start], i, current)
				break
			}
		}
		cycles[ix_start] = 1
	}

	// Calculate least common multiple of all cycles
	// Cycles:
	// [1] start RGA: - After i=17873 cycles atHPZ
	// [2] start JMA: - After i=23147 cycles atFKZ
	// [3] start XHA: - After i=15529 cycles atPKZ
	// [4] start DQA: - After i=17287 cycles atDHZ
	// [5] start AAA: - After i=19631 cycles atZZZ
	// [0] start MJA: - After i=20803 cycles atHJZ
	// https://www.calculatorsoup.com/calculators/math/lcm.php
	// 21003205388413 - correct answer

	// fmt.Printf("Part 2 - Total cycles = %d\n", cycle_ctr)

}
