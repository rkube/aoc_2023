package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse_seeds(line string) []int {
	seeds := make([]int, 20, 25)
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
	return seeds
}

func main() {
	fmt.Printf("Advent of Code 2023 - Day 05\n")
	f, err := os.Open("input_05")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the input file
	scanner := bufio.NewScanner(f)
	seed_parsed := false
	for scanner.Scan() {
		current_line := scanner.Text()
		if seed_parsed == false && current_line[0:5] == "seeds" {
			seeds := parse_seeds(current_line)
			// fmt.Printf("%d\n", seeds[0])
			seed_parsed = true
			for ix_s, s := range seeds {
				fmt.Printf("[%02d]:%d\n", ix_s, s)
			}
		}

	}

	defer f.Close()
}
