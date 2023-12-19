package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Range mapper struct defines a 1-1 mapping from
// src_start:src_start+len-1 => dst_start:dst_start+len-1
type src_dst_map struct {
	dst_start int
	src_start int
	len       int
}

// Implement a linked list to store a sequence of src_dst_maps
type Node struct {
	rg   src_dst_map
	next *Node
}

// Initial state of a linkedlist
type LinkedList struct {
	head *Node
}

// Insert a new element into a linked listt
func (list *LinkedList) Insert(new_map src_dst_map) {
	newNode := &Node{rg: new_map}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

// Map an index through the LinkedList
// Iterate thhrough the linked list. If ix is in any of the maps,
// map it from src to dst. Else, assume it is mapped from offset 0.
func (list *LinkedList) rg_map(ix int) int {
	current := list.head
	for current != nil {
		// Try mapping the index using the current list element.
		if (ix >= current.rg.src_start) && (ix < current.rg.src_start+current.rg.len) {
			// Calculate the offset between ix and src_start
			offset := ix - current.rg.src_start
			return current.rg.dst_start + offset
		}
		current = current.next
	}
	// If we haven't returned by now, we fell through all the ranges in the list.
	// This means, we return the identity mapping
	return ix
}

// Print the list
func (list *LinkedList) Display() {
	current := list.head

	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}

	fmt.Print("Linked list: \n")
	for current != nil {
		fmt.Printf("\tstart: [%d] -> [%d], length: %d\n", current.rg.src_start, current.rg.dst_start, current.rg.len)
		current = current.next
	}
	fmt.Println()
}

// Parse the seed list
// Format: seeds: 304740406 53203352...
func parse_seeds(line string, seeds []int) {
	line_cropped := line[7:]
	fmt.Printf("Cropped line: %s\n", line_cropped)
	splits := strings.Split(line_cropped, " ")
	for ix_s, this_split := range splits {
		this_seed, err := strconv.Atoi(this_split)
		if err != nil {
			log.Fatal(err)
		}
		seeds[ix_s] = this_seed
	}
	// return seeds
}

// Parse a string like 0 1 2 into a src_dst_map
func parse_map_str(line string) *src_dst_map {
	splits := strings.Split(line, " ")
	int_vals := make([]int, 3, 3)
	for ix_s, this_split := range splits {
		v, err := strconv.Atoi(this_split)
		if err != nil {
			log.Fatal(err)
		}
		int_vals[ix_s] = v
	}
	return &src_dst_map{dst_start: int_vals[0], src_start: int_vals[1], len: int_vals[2]}
}

func main() {
	fmt.Printf("Advent of Code 2023 - Day 05\n")
	f, err := os.Open("input_05")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the input file
	scanner := bufio.NewScanner(f)

	// Seed-to-soil map is a linked list
	var seed_to_soil_map LinkedList
	var soil_to_fertilizer_map LinkedList
	var fertilizer_to_water_map LinkedList
	var water_to_light_map LinkedList
	var light_to_temperature_map LinkedList
	var temperature_to_humidity_map LinkedList
	var humidity_to_location_map LinkedList

	// Line counter is used to find location of range maps in the text
	ix_line := 0
	seeds := make([]int, 20, 25)
	for scanner.Scan() {
		current_line := scanner.Text()
		// Seeds are listed in line 0. Store them in an integer slice
		if ix_line == 0 {
			parse_seeds(current_line, seeds)
			// seed_parsed = true
			for ix_s, s := range seeds {
				fmt.Printf("[%02d]:%d\n", ix_s, s)
			}
		} else if ix_line > 2 && ix_line < 15 {
			// Parse the range from a string into a src_to_dst_map and append in the linked list
			// which describes the mapping.
			new_map := parse_map_str(current_line)
			fmt.Printf("start: [%d] -> [%d], length: %d\n", new_map.src_start, new_map.dst_start, new_map.len)
			seed_to_soil_map.Insert(*new_map)
		} else if ix_line > 16 && ix_line < 38 {
			// soil-to-fertilizer map is on lines 17-37
			new_map := parse_map_str(current_line)
			soil_to_fertilizer_map.Insert(*new_map)
		} else if ix_line > 39 && ix_line < 56 {
			// Fertilizer to water map is lines 40-56
			new_map := parse_map_str(current_line)
			fertilizer_to_water_map.Insert(*new_map)
		} else if ix_line > 57 && ix_line < 103 {
			new_map := parse_map_str(current_line)
			water_to_light_map.Insert(*new_map)
		} else if ix_line > 104 && ix_line < 152 {
			new_map := parse_map_str(current_line)
			light_to_temperature_map.Insert(*new_map)
		} else if ix_line > 153 && ix_line < 177 {
			new_map := parse_map_str(current_line)
			temperature_to_humidity_map.Insert(*new_map)
		} else if ix_line > 178 && ix_line < 205 {
			new_map := parse_map_str(current_line)
			humidity_to_location_map.Insert(*new_map)
		}

		ix_line += 1
	}

	seed_to_soil_map.Display()

	// Map seeds to location
	locations := make([]int, 20, 25)
	for ix_s, seed := range seeds {
		ix_soil := seed_to_soil_map.rg_map(seed)
		ix_fert := soil_to_fertilizer_map.rg_map(ix_soil)
		ix_water := fertilizer_to_water_map.rg_map(ix_fert)
		ix_light := water_to_light_map.rg_map(ix_water)
		ix_temp := light_to_temperature_map.rg_map(ix_light)
		ix_humid := temperature_to_humidity_map.rg_map(ix_temp)
		ix_loc := humidity_to_location_map.rg_map(ix_humid)
		locations[ix_s] = ix_loc
		fmt.Printf("[seed ] %d\n[fert ] %d\n[water] %d\n[light] %d\n[temp ] %d\n[humid] %d\n[locat] %d\n", seed, ix_fert, ix_water, ix_light, ix_temp, ix_humid, ix_loc)
	}

	// Find the minimum of all locations
	min_loc := 4294967295
	for _, loc := range locations {
		fmt.Printf("%d\n", loc)
		if loc < min_loc {
			min_loc = loc
		}
	}
	fmt.Printf("Minimum location: %d\n", min_loc)

	defer f.Close()
}
