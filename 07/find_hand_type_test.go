package main

import (
	"fmt"
	"testing"
)

func build_hand(hand_str string) hand_bid {
	this_hand := initialize_hand_map()
	for _, card := range hand_str {
		this_hand[card] += 1
	}
	return hand_bid{cards: hand_str, bid: 0, hand: this_hand}
}

// func Test_find_hand_type(t *testing.T) {
// 	// hand_str := "QQQQQ"
// 	// this_hand := initialize_hand_map()
// 	// for _, card := range hand_str {
// 	// 	this_hand[card] += 1
// 	// }
// 	// hand := hand_bid{cards: hand_str, bid: 0, hand: this_hand}

// 	hand_str := "QQQQQ"
// 	hand := build_hand(hand_str)
// 	hand_type, err := find_hand_type(hand, false)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("%s: Hand type = %d\n", hand_str, hand_type)

// }

func Test_five_kind(t *testing.T) {
	fmt.Printf("==== Testing for 5 of a kind\n")

	test_cases := []string{"TTTTT", "JAAAA", "JAJAA", "JAJJA", "JJJJA", "JJJJJ"}

	fmt.Printf("      Jokers=false\n")
	for _, h := range test_cases {
		hand := build_hand(h)
		fmt.Printf("5 of kind for %s: %v\n", h, is_five_of_kind(hand, false))
	}
	fmt.Printf("      Jokers=true:\n")
	for _, h := range test_cases {
		hand := build_hand(h)
		fmt.Printf("5 of kind for %s: %v\n", h, is_five_of_kind(hand, true))
	}
}

func Test_four_kind(t *testing.T) {
	fmt.Printf("==== Testing or 4 of a kind\n")
	test_cases := []string{"KKAKK", "83J33", "2JJ99", "J8JJ2", "22J34"}
	fmt.Printf("\tJokers=true\n")
	for _, h := range test_cases {
		hand := build_hand(h)
		fmt.Printf("4 of kind for %s: %v\n", h, is_four_of_kind(hand, true))
	}
	fmt.Printf("\tJokers=false\n")
	for _, h := range test_cases {
		hand := build_hand(h)
		fmt.Printf("4 of kind for %s: %v\n", h, is_four_of_kind(hand, false))
	}
}

func Test_full_house(t *testing.T) {
	fmt.Printf("===== Testing full house\n")
	test_cases := []string{"33377", "KAKAK", "58J85"}
	fmt.Printf("\tJokers=true\n")
	for _, h := range test_cases {
		hand := build_hand(h)
		fmt.Printf("Full House for %s: %v\n", h, is_full_house(hand, true))
	}
	fmt.Printf("\tJokers=false\n")
	for _, h := range test_cases {
		hand := build_hand(h)
		fmt.Printf("Full House for %s: %v\n", h, is_full_fouse(hand, false))
	}
}
