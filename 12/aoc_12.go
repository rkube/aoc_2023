package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	for scanner.Scan() {
		current_line := scanner.Text()
		expr, counts := parse_line(current_line)
		fmt.Printf("%s: %d", expr, counts[0])
	}

}
