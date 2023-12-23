package main

import (
	"bufio"
	"errors"
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
func get_card_values(jokers bool) map[rune]int {
	if jokers == false {
		return map[rune]int{'2': 0, '3': 1, '4': 2, '5': 3, '6': 4, '7': 5, '8': 6, '9': 7, 'T': 8, 'J': 9, 'Q': 10, 'K': 11, 'A': 12}
	} else {
		return map[rune]int{'J': 0, '2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8, 'T': 9, 'Q': 10, 'K': 11, 'A': 12}
	}
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

// Gets the highest occurance of a single card within a hand
func do_max_count(hand map[rune]int, jokers bool) int {
	max_count := 0
	for card, v := range hand {
		// Don't count Jokers for max_count
		if card == rune('J') && jokers {
			continue
		}
		if v > max_count {
			max_count = v
		}
	}
	return max_count
}

// Test if we can get five of kind
func is_five_of_kind(h hand_bid, jokers bool) bool {
	max_count := do_max_count(h.hand, jokers)
	if jokers {
		// See if we can form five of a kind with the number of jokers at hand
		num_jokers := h.hand[rune('J')]
		// Jokers and max count have to add up to 5
		if num_jokers+max_count == 5 {
			return true
		} else {
			return false
		}
	} else {
		return max_count == 5
	}
}

// Test if we can get four if kind
func is_four_of_kind(h hand_bid, jokers bool) bool {
	max_count := do_max_count(h.hand, jokers)
	if jokers {
		num_jokers := h.hand[rune('J')]
		return num_jokers+max_count == 4

	} else {
		return max_count == 4
	}
}

// Test if we can get a Full House
func is_full_house(h hand_bid, jokers bool) bool {
	max_count := do_max_count(h.hand, jokers)
	if jokers {
		num_jokers := h.hand[rune('J')]
		// 1st possiblity: 3 of kind, 2 of kind, no jokers
		if num_jokers == 0 {
			found_3 := false
			found_2 := false
			for _, val := range h.hand {
				found_3 = val == 3
				found_2 = val == 2
			}
			return found_3 && found_2
		} else if num_jokers == 1 {
			// 2nd possibility: 2 pairs and 1 joker
			num_twos := 0
			for _, val := range h.hand {
				if val == 2 {
					num_twos += 1
				}
			}
			return num_twos == 2
		}
		return false
	} else {
		if max_count == 3 {
			for _, v := range h.hand {
				if v == 2 {
					return true
				}
			}
			return false
		}
		return false
	}
}

// Takes a hand and finds determine its type, i.e. 5 of a kind, 4 of a kind etc.
func find_hand_type(h hand_bid, jokers bool) (hand_type, error) {
	max_count := do_max_count(h.hand, jokers)
	for _, v := range h.hand {
		if v > max_count {
			max_count = v
		}
	}
	switch max_count {
	case 5:
		return five_of_kind, nil
	case 4:
		return four_of_kind, nil
	case 3:
		// If we have 3 identical card we need to re-check if it's a full house or 3-of-kind
		for _, v := range h.hand {
			if v == 2 {
				return full_house, nil
			}
		}
		return three_of_kind, nil
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
			return one_pair, nil
		} else {
			return two_pair, nil
		}
	case 1:
		return high_card, nil
	}
	error_str := fmt.Sprintf("Could not find hand type for %s\n", h.cards)
	return invalid, errors.New(error_str)
}

// Basic bubble sort that works on an array of ints
func bubble_sort(hands map[int]hand_bid, order *[]int, jokers bool) {
	N := len(*order)
	// count_sort := 0
	for {
		swapped := false

		for ix := 1; ix < N; ix++ {
			cmp, err := greater_than(hands[(*order)[ix-1]], hands[(*order)[ix]], jokers)

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

	}

	// Initialize ordering of all cards
	order := make([]int, 1000)
	for ix := 0; ix < len(order); ix++ {
		order[ix] = ix
	}

	bubble_sort(all_hands_bids, &order, false)

	weighted_sum := 0
	for ix := 0; ix < len(order); ix++ {
		weighted_sum += all_hands_bids[order[ix]].bid * (ix + 1)
		fmt.Printf("Hand: %s Bid: %d, order: %d\n", all_hands_bids[order[ix]].cards, all_hands_bids[order[ix]].bid, ix)
	}
	fmt.Printf("Sum of bids, weighted by order: %d\n", weighted_sum)

	// 251362095 -- too high
	// 249638405 -- correct answer :D
}
