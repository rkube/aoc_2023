// Initialize folder with
// $ go mod init aoc_04

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Parse integers from a string like
//
//	88 38 30  3 48 17 19 68 73  2
func parse_seq(seq string) []int {
	// Assume each number takes up roughly 3 digits
	int_vals := make([]int, len(seq)/3, 100)
	// fmt.Printf("Creating slice: len(seq) = %d: %d\n", len(seq), len(seq)/3)
	splits := strings.Split(seq, " ")
	ix_s := 0 // Counts the successfully converted splits to int. Not the index of the splits, they can contain splits of zero length
	for _, s := range splits {
		// We need to skip empty splits
		if len(s) == 0 {
			continue
		}
		// fmt.Printf("%d: __%s__, len=%d\n", ix_s, s, len(s))
		v, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		int_vals[ix_s] = v
		ix_s += 1
	}
	return int_vals
}

// Return an array of the winning numbers and an array of the draws numbers on a card (line)
// Each card is represented by a line like this
// Card 164: 88 38 30  3 48 17 19 68 73  2 | 13 71 34 83 40 38 59 12 73  2 91 52 60 19 87 84  1 82 65  3  8 99 80 79 70
// To parse the card we
//  1. Remove the first 9 characters. These give information of the card numer.
//  2. Split the card at "|". The first split has information on the winning numbers.
//     The second split has the drawn numbers
//  3. Split each split from step 2 by " ", and parse each sub string as an int to get the
//     integer values of the splits.
func parse_card(card string) ([]int, []int) {
	fmt.Printf("%s\n", card)
	//  Remove first 9 characters and split
	// card_sp := card[9:len(card)]
	// fmt.Printf("--- After pruning: %s\n", card_sp)
	splits_1 := strings.Split(card[9:], "|")
	fmt.Printf("--- Split 1 (winning numbers): %s\n", splits_1[0])
	fmt.Printf("--- Split 2 (drawn numbers): %s\n", splits_1[1])
	winning_numbers := parse_seq(splits_1[0])
	drawn_numbers := parse_seq(splits_1[1])

	return winning_numbers, drawn_numbers
}

func main() {
	fmt.Printf("Advent of code 2024 - project 4\n")
	// Read in the input file
	f, err := os.Open("input_04")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		current_card := scanner.Text()
		// Parse the current card
		winning_numbers, drawn_numbers := parse_card(current_card)

		fmt.Printf("winning numbers (as integers)\n\t")
		for ix_w, _ := range winning_numbers {
			fmt.Printf("%d ", winning_numbers[ix_w])
		}
		fmt.Printf("\n")

		fmt.Printf("Drawn numbers (as integers)\n\t")
		for ix_d, _ := range drawn_numbers {
			fmt.Printf("%d ", drawn_numbers[ix_d])
		}
		fmt.Printf("\n")
	}

}
