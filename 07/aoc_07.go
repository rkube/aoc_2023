package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// This structure stores a hand together with a bid.
// The map[string] int is a mapping of the card type onto their count
// 2, 3, 4, ... T, J, Q, K, A => count
type hand_bid struct {
	hand map[rune]int
	bid  int
}

func initialize_hand_map() map[rune]int {
	new_map := make(map[rune]int)
	valid_runes := []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
	for _, r := range valid_runes {
		new_map[r] = 0
	}
	return new_map
}

// Takes a hand, econded as a string "TQJTT" and its bid
// and stuffes them into
func parse_splits(hand_str string, bid int) hand_bid {
	this_hand := initialize_hand_map()
	for _, card := range hand_str {
		this_hand[card] += 1
	}
	new_hand_bid := hand_bid{hand: this_hand, bid: bid}
	return new_hand_bid
}

func main() {
	fmt.Printf("Advent of code 2023 - Part 07\n")

	f, err := os.Open("input_07")
	if err != nil {
		log.Fatal(err)
	}

	const num_hands = 1000 // Number of total hands
	// Parse the input file
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		current_line := scanner.Text()

		// Parse the current line into the hand and the bid.
		splits := strings.Split(current_line, " ")
		this_bid, err := strconv.Atoi(splits[1])
		if err != nil {
			log.Fatal(err)
		}
		this_hand_bid := parse_splits(splits[0], this_bid)
		fmt.Printf("bid: %d\n", this_hand_bid.bid)
		break
	}
}
