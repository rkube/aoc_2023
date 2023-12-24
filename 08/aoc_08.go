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

func find_cucle(node_start string, node_end string) int {

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
			if current_line[2] == 'A' {
				fmt.Printf("Starting node: %s\n", current_line[0:3])
			} else if current_line[2] == 'Z' {
				fmt.Printf("End node: %s\n", current_line[0:3])
			}
		}
		ix_line += 1
	}
	defer f.Close()

	// fmt.Printf("Directions: %s, length=%d\n", directions, len(directions))
	iterations := 0 // iterations over total directions
	curr := "AAA"   // Starting position
	// Traverse the map using the entire sequence of iterations
	// Check if we end up in "ZZZ". If yes, we are finished.
	// If not, repeat the sequnce until we arrive at ZZZ
	for {
		for _, next := range directions {
			curr = places[curr].target(next)
		}
		iterations += 1

		if curr == "ZZZ" {
			break
		}

		if iterations > 1000 {
			break
		}
	}
	// The total number of steps is the number of times we had to
	// iterate over the entire sequence of steps to come to ZZZ
	// multiplied by the length of that sequence.
	fmt.Printf("Number of steps: %d\n", iterations*len(directions))
}
