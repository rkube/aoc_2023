package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Type of a hand
type hand_type int

// Ordering for the first rule to compare cards
const (
	five_of_kind  hand_type = 6
	four_of_kind  hand_type = 5
	full_house    hand_type = 4
	three_of_kind hand_type = 3
	two_pair      hand_type = 2
	one_pair      hand_type = 1
	high_card     hand_type = 0
	invalid       hand_type = -1
)

// Maps card face values to numerical values for comparison
func get_card_values() map[rune]int {
	card_values := map[rune]int{'2': 0, '3': 1, '4': 2, '5': 3, '6': 4, '7': 5, '8': 6, '9': 7, 'T': 8, 'J': 9, 'Q': 10, 'K': 11, 'A': 12}
	return card_values
}

// This structure stores a hand together with a bid.
// The map[string] int is a mapping of the card type onto their count
// 2, 3, 4, ... T, J, Q, K, A => count
type hand_bid struct {
	hand  map[rune]int
	cards string
	bid   int
}

func (h hand_bid) Display() {
	fmt.Printf("%s\n", h.cards)
	fmt.Printf("[2] %d\t[3] %d\t [4] %d\t [5]%d\t [6] %d\t [7] %d\t [8] %d\t [9] %d [T] %d\t [J] %d\t [Q] %d\t [K] %d\t [A]%d\n",
		h.hand['2'], h.hand['3'], h.hand['4'], h.hand['5'], h.hand['6'], h.hand['7'],
		h.hand['8'], h.hand['9'], h.hand['T'], h.hand['J'], h.hand['Q'], h.hand['K'], h.hand['A'])
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
// and stuffes them into a hand_bid struct
func parse_splits(hand_str string, bid int) hand_bid {
	this_hand := initialize_hand_map()
	for _, card := range hand_str {
		this_hand[card] += 1
		// fmt.Printf("card = %c: %d\n", card, this_hand[card])
	}
	new_hand_bid := hand_bid{hand: this_hand, bid: bid, cards: hand_str}
	return new_hand_bid
}

// Takes a hand and finds determine its type, i.e. 5 of a kind, 4 of a kind etc.
func find_hand_type(h hand_bid) hand_type {
	max_count := 0
	for _, v := range h.hand {
		// fmt.Printf("hand_type: %c - %d\n", k, v)
		if v > max_count {
			max_count = v
		}
	}
	switch max_count {
	case 5:
		return five_of_kind
	case 4:
		return four_of_kind
	case 3:
		// If we have 3 identical card we need to re-check if it's a full house or 3-of-kind
		for _, v := range h.hand {
			if v == 2 {
				return full_house
			}
		}
		return three_of_kind
	case 2:
		// If max of type is 2, we need to check for another pair
		rune_first_two := '0'
		rune_second_two := '0'
		for r, v := range h.hand {
			if v == 2 && rune_first_two == '0' {
				rune_first_two = r
			} else if v == 2 && rune_first_two != '0' {
				rune_second_two = r
			}
		}
		// h.Display()
		// fmt.Printf("First pair: %c\tSecond pair: %c\n", rune_first_two, rune_second_two)
		if rune_second_two == '0' {
			return one_pair
		} else {
			return two_pair
		}
	case 1:
		return high_card
	}
	return invalid
}

// Basic bubble sort that works on an array of ints
func bubble_sort(hands map[int]hand_bid, order *[]int) {
	N := len(*order)
	// count_sort := 0
	for {
		swapped := false

		for ix := 1; ix < N; ix++ {
			cmp, err := greater_than(hands[(*order)[ix-1]], hands[(*order)[ix]])

			if err != nil {
				log.Fatal(err)
			}
			// fmt.Printf("%s > %s ? %v\n", hands[(*order)[ix-1]].cards, hands[(*order)[ix]].cards, cmp)
			if cmp {
				tmp := (*order)[ix]
				(*order)[ix] = (*order)[ix-1]
				(*order)[ix-1] = tmp
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
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

	all_hands_bids := make(map[int]hand_bid, 1000)
	ix_line := 0
	for scanner.Scan() {
		current_line := scanner.Text()

		// Parse the current line into the hand and the bid.
		splits := strings.Split(current_line, " ")
		// fmt.Printf("--- Hand (str): %s\n", splits[0])
		this_bid, err := strconv.Atoi(splits[1])
		if err != nil {
			log.Fatal(err)
		}
		this_hand_bid := parse_splits(splits[0], this_bid)
		// fmt.Printf("bid: %d\n hand:\n", this_hand_bid.bid)
		// this_hand_bid.Display()
		// c := find_hand_type(this_hand_bid)
		// fmt.Printf("hand_type = %d\n", c)

		all_hands_bids[ix_line] = this_hand_bid
		ix_line += 1

		// if ix_line > 10 {
		// 	break
		// }
	}

	// // After parsing all hands, try pair-wise comparison
	// for ix_card := 0; ix_card < 10; ix_card += 2 {
	// 	comparison, err := greater_than(all_hands_bids[ix_card], all_hands_bids[ix_card+1])
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("================================================\n")
	// 	all_hands_bids[ix_card].Display()
	// 	all_hands_bids[ix_card+1].Display()
	// 	fmt.Printf("hand type 1 = %d, hand_type 2 = %d, hand 1 >= hand 2 = %v\n", find_hand_type(all_hands_bids[ix_card]), find_hand_type(all_hands_bids[ix_card+1]), comparison)
	// }
	// order := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Initialize ordering of all cards
	order := make([]int, 1000)
	for ix := 0; ix < len(order); ix++ {
		order[ix] = ix
	}

	bubble_sort(all_hands_bids, &order)
	// for ix := 0; ix < len(order); ix++ {
	// 	this_hand := all_hands_bids[order[ix]]
	// 	fmt.Printf("%d - Hand: %s, hand_value = %d\n", ix, this_hand.cards, find_hand_type(this_hand))
	// }

	weighted_sum := 0
	for ix := 0; ix < len(order); ix++ {
		weighted_sum += all_hands_bids[order[ix]].bid * (ix + 1)
		fmt.Printf("Hand: %s Bid: %d, order: %d\n", all_hands_bids[order[ix]].cards, all_hands_bids[order[ix]].bid, ix)
	}
	fmt.Printf("Sum of bids, weighted by order: %d\n", weighted_sum)

	// 251362095 -- too high
	// 249638405 -- correct answer :D

}
