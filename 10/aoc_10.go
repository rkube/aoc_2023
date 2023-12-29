package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type position struct {
	x        int
	y        int
	pipe     byte
	distance int
}

// returns true if a pipe connects north
func connects_N(p byte) bool {
	if (p == '|') || (p == 'L') || (p == 'J') || (p == 'S') {
		return true
	} else {
		return false
	}
}

// returns true if a pipe connects south
func connects_S(p byte) bool {
	if (p == '|') || (p == '7') || (p == 'F') || (p == 'S') {
		return true
	} else {
		return false
	}
}

// returns true if a pipe connects east
func connects_E(p byte) bool {
	if (p == '-') || (p == 'L') || (p == 'F') || (p == 'S') {
		return true
	} else {
		return false
	}
}

// returns true if a pipe connects west
func connects_W(p byte) bool {
	if (p == '-') || (p == '7') || (p == 'J') || (p == 'S') {
		return true
	} else {
		return false
	}
}

func connect(p1 position, p2 position, direction byte) bool {
	if direction == 'N' {
		// fmt.Printf("connect: %c, %v, %v\n", direction, connects_N(p1.pipe), connects_S(p2.pipe))
		return connects_N(p1.pipe) && connects_S(p2.pipe)
	} else if direction == 'E' {
		// fmt.Printf("connect: %c, %v, %v\n", direction, connects_E(p1.pipe), connects_W(p2.pipe))
		return connects_E(p1.pipe) && connects_W(p2.pipe)
	} else if direction == 'S' {
		// fmt.Printf("connect: %c, %v, %v\n", direction, connects_S(p1.pipe), connects_N(p2.pipe))
		return connects_S(p1.pipe) && connects_N(p2.pipe)
	} else if direction == 'W' {
		// fmt.Printf("connect: %c, %v, %v\n", direction, connects_W(p1.pipe), connects_E(p2.pipe))
		return connects_W(p1.pipe) && connects_E(p2.pipe)
	} else {
		return false
	}
}

// Trace out the maze by travelling into the direction
func maze_step(current position, prev position, maze []string) (position, error) {
	// Check in which direction we can move without returning to the previous position
	directions := []byte{'N', 'E', 'S', 'W'}
	proposal := position{x: 0, y: 0, pipe: '?', distance: 0}

	for _, dir := range directions {
		// out-of-bounds check for moving north
		if dir == 'N' {
			if current.y-1 < 0 {
				// out-of-bounds
				fmt.Printf("Can't move from [%d,%d] to [%d,%d]: Out of bounds\n", current.x, current.y, current.x, current.y-1)
			}
			proposal.x, proposal.y, proposal.pipe, proposal.distance = current.x, current.y-1, maze[current.y-1][current.x], current.distance+1
		} else if dir == 'E' {
			if current.x+1 > 140 {
				// out-of-bounds
				fmt.Printf("Can't move from [%d,%d] to [%d,%d]: Out of bounds\n", current.x, current.y, current.x+1, current.y)
			}
			proposal.x, proposal.y, proposal.pipe, proposal.distance = current.x+1, current.y, maze[current.y][current.x+1], current.distance+1
		} else if dir == 'S' {
			if current.y+1 > 139 {
				// out-of-bounds
				fmt.Printf("Can't move from [%d,%d] to [%d,%d]: Out of bounds\n", current.x, current.y, current.x, current.y+1)
			}
			proposal.x, proposal.y, proposal.pipe, proposal.distance = current.x, current.y+1, maze[current.y+1][current.x], current.distance+1
		} else if dir == 'W' {
			if current.x-1 < 0 {
				// out-of-bounds
				fmt.Printf("Can't move from [%d,%d] to [%d,%d]: Out of bounds\n", current.x, current.y, current.x-1, current.y)
			}
			proposal.x, proposal.y, proposal.pipe, proposal.distance = current.x-1, current.y, maze[current.y][current.x-1], current.distance+1
		}
		// Check if the proposal connects to the current one and is different from the previous positions
		fmt.Printf("Direction: %c\t", dir)
		fmt.Printf("Current: (%d)[%d, %d] %c ", current.distance, current.x, current.y, current.pipe)
		fmt.Printf("Propose: (%d)[%d, %d] %c ", proposal.distance, proposal.x, proposal.y, proposal.pipe)
		if connect(current, proposal, dir) {
			fmt.Printf(" Connect\n")
			fmt.Printf("Checking: %d!=%d: %v, %d!=%d: %v\n", proposal.x, prev.x, proposal.x != prev.x, proposal.y, prev.y, proposal.y != prev.y)
			if !((proposal.x == prev.x) && (proposal.y == prev.y)) {
				fmt.Printf("Accepted\n")
				return proposal, nil
			}
		} else {
			fmt.Printf("Don't connect\n")
		}
	}
	return proposal, errors.New("Can't find a connecting node that is not the previous node")
}

// Load the maze from file
func setup_maze(filename string, len int) ([]string, position) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	maze := make([]string, len)

	start_pos := position{x: -1, y: -1, pipe: '0', distance: 0}

	ix_line := 0
	for scanner.Scan() {
		current_line := scanner.Text()
		maze[ix_line] = current_line
		ix_S := strings.Index(maze[ix_line], "S")
		if ix_S > 0 {
			start_pos.x = ix_S
			start_pos.y = ix_line
			start_pos.pipe = 'S'
			start_pos.distance = 0
		}
		ix_line += 1
	}
	return maze, start_pos
}

func main() {
	fmt.Printf("Advent of code 2023 - Part 10\n")
	maze, curr_pos := setup_maze("input_10", 140)
	prev_pos := curr_pos
	next_pos := curr_pos
	fmt.Printf("Start node: [%02d, %02d]: %c (%d)\n", curr_pos.x, curr_pos.y, curr_pos.pipe, curr_pos.distance)

	for {
		fmt.Printf("----------- Step ----------\n")
		next_pos, err := maze_step(curr_pos, prev_pos, maze)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("New node: [%02d, %02d]: %c (%d)\n", next_pos.x, next_pos.y, next_pos.pipe, next_pos.distance)
		// Step through the maze until we are back at the start
		if next_pos.pipe == 'S' {
			break
		}
		prev_pos = curr_pos
		curr_pos = next_pos
	}

	// The node farthest away from the start is found by taking the length
	// of the entire round trip divided by 2.
	// Solution: 13434/2 = 6717
	fmt.Printf("Total length of the path: %d. Answer: %d\n", next_pos.distance, next_pos.distance/2)

	// fmt.Printf("[%d, %d]: %c %d\n", start_pos.x, start_pos.y, start_pos.pipe, start_pos.distance)
}
