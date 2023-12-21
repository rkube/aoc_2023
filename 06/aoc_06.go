package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Advent of code 2023 - Part 06\n")
	durations := []int{40, 70, 98, 79}        // Durations of the race
	distances := []int{215, 1051, 2147, 1005} // Distance record

	margin := 1
	for ix := 0; ix < 4; ix++ {
		// Calculate distance travelled for every possible duration of charging time
		this_dur := durations[ix]
		record_distance := distances[ix]
		ways_to_beat_record := 0
		for v := 0; v < this_dur; v++ {
			remaining_time := this_dur - v
			if v*remaining_time > record_distance {
				// fmt.Printf("%d * %d = %d > %d\n", v, remaining_time, v*remaining_time, record_distance)
				ways_to_beat_record += 1
			}
		}
		margin *= ways_to_beat_record
	}
	// Correct answer: 1084752
	fmt.Printf("Part 1: margin = %d\n", margin)

	// Part 2
	duration := 40709879
	record_distance := 215105121471005
	ways_to_beat_record := 0
	for v := 0; v < duration; v++ {
		remaining_time := duration - v
		if v*remaining_time > record_distance {
			// fmt.Printf("%d * %d = %d > %d\n", v, remaining_time, v*remaining_time, record_distance)
			ways_to_beat_record += 1
		}
	}
	fmt.Printf("Part 2: ways to beat record = %d\n", ways_to_beat_record)

}
