package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func clear_map(m *map[int]int) {
	for k, _ := range *m {
		(*m)[k] = 0
	}
}

// This is blatantly copied from
// https://github.com/clrfl/AdventOfCode2023/blob/master/12/explanation.ipynb
func count_arrangements(pattern string, counts []int) int {
	// Compile the states from the expression.
	// The states s0, s1,... sn are just the characters of the string.
	// Initialize the states with a '.'
	states := []byte{'.'}
	// Now append blocks of #'s, separated by a .
	for _, nr := range counts {
		for ix := 0; ix < nr; ix++ {
			states = append(states, '#')
		}
		states = append(states, '.')
	}

	// fmt.Printf("Counts: ")
	// for ix := 0; ix < len(counts); ix++ {
	// 	fmt.Printf("%d ", counts[ix])
	// }
	// fmt.Printf("\nStates = %s, len=%d\n", states, len(states))
	// fmt.Printf("Pattern = %s\n", pattern)

	// Heads refer to the current state of the machine.
	// We track them with a dictionary, that keeps track how many
	// are in a given state.

	heads := make(map[int]int, len(states))   // This keeps track of the heads while iterating over the string
	new_pos := make(map[int]int, len(states)) // This keeps track of how all heads move for a single position

	// Initialize maps with all zeros.
	for ix := 0; ix < len(states); ix++ {
		heads[ix] = 0
		new_pos[ix] = 0
	}

	// Initialize a single head at state 0
	heads[0] = 1

	// Iterate over the tokens of the pattern we want to match
	for _, ch := range pattern {
		// fmt.Printf("\n%02d: '%c'\nheads:   ", ix_ch, ch)
		// for state := 0; state < len(states); state++ {
		// 	fmt.Printf("[s%d]=%d\t", state, heads[state])
		// }
		// fmt.Printf("\nnew_pos: ")
		// for state := 0; state < len(states); state++ {
		// 	fmt.Printf("[s%d]=%d\t", state, new_pos[state])
		// }
		// fmt.Printf("\n")
		// For each token, move each head through the state machine
		// for state, num_heads := range heads {
		for state := 0; state < len(states); state++ {
			num_heads := heads[state]
			if ch == '?' {
				// Every state advances for a '?'
				if state < len(states)-1 {
					// fmt.Printf("--- Advancing heads [%d,%d] to a '%c' state\n", state, num_heads, states[state+1])
					new_pos[state+1] += num_heads
				}
				// If we are currently at a '.', spawn a new head for each current head
				// at the current state (this is the self-loop on a '.'-state)
				if states[state] == '.' {
					new_pos[state] += num_heads
				}
			} else if ch == '.' {
				if (state < len(states)-1) && states[state+1] == '.' {
					// fmt.Printf("--- Advancing heads [%d,%d] to a '.' state\n", state, num_heads)
					// If we haven't exhausted the state machine and the next
					// state is a '.', matching the current char in the pattern,
					// move all heads at this state to the next state.
					// This is the right arrow out ot a '#' state
					new_pos[state+1] += num_heads
				}
				if states[state] == '.' {
					// This is the self-loop at at '.' state
					// Spawn a new head at the current state for each head in this state
					new_pos[state] += num_heads
				}
			} else if ch == '#' {
				// All heads in a state where the next state is a '#' need to advance
				if (state < len(states)-1) && (states[state+1] == '#') {
					// fmt.Printf("--- Advancing heads [%d,%d] to a '#' state\n", state, num_heads)
					new_pos[state+1] += num_heads
				}
			}
		}

		// fmt.Printf("After:\nnew_pos: ")
		// for state := 0; state < len(states); state++ {
		// 	fmt.Printf("[s%d]=%d\t", state, new_pos[state])
		// }
		for ix := 0; ix < len(states); ix++ {
			heads[ix] = new_pos[ix]
			new_pos[ix] = 0
		}

		// fmt.Printf("\nAfter updating:\nheads:   ")
		// for state := 0; state < len(states); state++ {
		// 	fmt.Printf("[s%d]=%d\t", state, heads[state])
		// }
		// fmt.Printf("\nnew_pos: ")
		// for state := 0; state < len(states); state++ {
		// 	fmt.Printf("[s%d]=%d\t", state, new_pos[state])
		// }
		// fmt.Printf("\n")

	}

	return heads[len(states)-1] + heads[len(states)-2]
}

func parse_line(line string) (string, []int) {
	expr, count_str, _ := strings.Cut(line, " ")
	// Convert the counts into integers
	splits := strings.Split(count_str, ",")
	counts := []int{}
	for _, s := range splits {
		val, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		counts = append(counts, val)
	}

	return expr, counts
}

func main() {
	fmt.Printf("Advent of Code 2023 - Day 12\n")

	f, err := os.Open("input_12")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		current_line := scanner.Text()
		expr, counts := parse_line(current_line)
		// fmt.Printf("%s: ", expr)
		// for _, c := range counts {
		// 	fmt.Printf("%d, ", c)
		// }
		// fmt.Printf("\n")
		sum += count_arrangements(expr, counts)
	}

	fmt.Printf("Part 1: sum = %d\n", sum)
}
