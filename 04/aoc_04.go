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
	//  Remove first 9 characters and split
	// card_sp := card[9:len(card)]
	// fmt.Printf("--- After pruning: %s\n", card_sp)
	splits_1 := strings.Split(card[9:], "|")
	// fmt.Printf("--- Split 1 (winning numbers): %s\n", splits_1[0])
	// fmt.Printf("--- Split 2 (drawn numbers): %s\n", splits_1[1])
	winning_numbers := parse_seq(splits_1[0])
	drawn_numbers := parse_seq(splits_1[1])

	return winning_numbers, drawn_numbers
}

// Calculate integer powers
func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

// Scores the card.
// In the above example, card 1 has five winning numbers (41, 48, 83, 86, and 17) and
// eight numbers you have (83, 86, 6, 31, 17, 9, 48, and 53). Of the numbers you have,
// four of them (48, 83, 17, and 86) are winning numbers! That means card 1 is worth 8
// points (1 for the first match, then doubled three times for each of the three matches
// after the first).
func score_card(winning_numbers []int, drawn_numbers []int) (int, int) {
	// For each winning number, search if it matches a drawn number.
	// Then count the matches
	num_matches := 0
	// fmt.Printf("len(w) = %d, len(d) = %d\n", len(winning_numbers), len(drawn_numbers))
	for _, w := range winning_numbers {
		for _, d := range drawn_numbers {
			// fmt.Printf("[%d - %d] and  [%d - %d]\n", ix_w, w, ix_d, d)
			if d == w {
				// fmt.Printf("==================== Match: [%d - %d] and [%d - %d]\n", ix_w, w, ix_d, d)
				num_matches += 1
			}
		}
		// fmt.Printf("\n")
	}
	if num_matches > 0 {
		return num_matches, IntPow(2, num_matches-1)
	} else {
		return num_matches, 0
	}
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

	total_score := 0 // This accumulates the score reported for part 1.
	ix_card := 0

	card_counts := make([]int, 192, 200) // For part 2: Counts the copies of cards that are won
	card_scores := make([]int, 192, 200)
	// Initially, we have 1 copy of each card
	for ix_c := range card_counts {
		card_counts[ix_c] = 1
	}

	for scanner.Scan() {
		fmt.Printf("%d\n", ix_card)
		current_card := scanner.Text()
		// Parse the current card
		winning_numbers, drawn_numbers := parse_card(current_card)

		// Now score the card
		num_matches, this_score := score_card(winning_numbers, drawn_numbers)
		card_scores[ix_card] = num_matches
		total_score += this_score

		// fmt.Printf("card %d: %d matches\n", ix_card+1, num_matches)

		// For part two, we add one copy to each of the next s cards, where s is the number of winning numbers
		for s := 1; s <= num_matches; s++ {
			if (ix_card + s) < len(card_counts) {
				card_counts[ix_card+s] += card_counts[ix_card]
			}
		}

		ix_card += 1
		if ix_card < 5 {
			for ix_s := 0; ix_s < len(card_counts); ix_s++ {
				fmt.Printf("[%d] : %03d\n", ix_s, card_counts[ix_s])
			}
		}
	}
	fmt.Printf("Total score from all cards: %d\n", total_score)
	// Correct answer: 21213

	for ix_s := 0; ix_s < len(card_counts); ix_s++ {
		fmt.Printf("[%d] : %03d\n", ix_s, card_scores[ix_s])
	}
	sum_cards := 0
	for _, c := range card_counts {
		sum_cards += c
	}

	fmt.Printf("Total number of cards won: %d\n", sum_cards)
	// Correct answer: 8549735
}
