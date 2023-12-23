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
		fmt.Printf("Full House for %s: %v\n", h, is_full_house(hand, false))
	}
}

func Test_three_kind(t *testing.T) {
	fmt.Printf("===== Testing Three of a kind\n")
	test_cases := []string{"23456", "J89QA", "JA8J2", "JJJ21", "229T2"}
	fmt.Printf("\tJokers=true\n")
	for _, h := range test_cases {
		hand := build_hand(h)
		fmt.Printf("Full House for %s: %v\n", h, is_three_of_kind(hand, true))
	}
	fmt.Printf("\tJokers=false\n")
	for _, h := range test_cases {
		hand := build_hand(h)
		fmt.Printf("Full House for %s: %v\n", h, is_three_of_kind(hand, false))
	}
}

func Test_two_pair(t *testing.T) {
	fmt.Printf("===== Testing Two Pairs\n")
	test_cases := []string{"99J22", "77877", "JJ113", "AA442", "9K92K"}
	fmt.Printf("\tJokers=true\n")
	for _, h := range test_cases {
		hand := build_hand(h)
		fmt.Printf("Full House for %s: %v\n", h, is_two_pair(hand, true))
	}
	fmt.Printf("\tJokers=false\n")
	for _, h := range test_cases {
		hand := build_hand(h)
		fmt.Printf("Full House for %s: %v\n", h, is_two_pair(hand, false))
	}
}

func Test_one_pair(t *testing.T) {
	fmt.Printf("===== Testing Two Pairs\n")
	test_cases := []string{"2468T", "JKAJK", "TJK72"}
	fmt.Printf("\tJokers=true\n")
	for _, h := range test_cases {
		hand := build_hand(h)
		fmt.Printf("Full House for %s: %v\n", h, is_one_pair(hand, true))
	}
	fmt.Printf("\tJokers=false\n")
	for _, h := range test_cases {
		hand := build_hand(h)
		fmt.Printf("Full House for %s: %v\n", h, is_one_pair(hand, false))
	}
}

func Test_high_card(t *testing.T) {
	fmt.Printf("===== Testing Two Pairs\n")
	test_cases := []string{"2468T", "JKAJK", "TJK72", "AA29T"}
	fmt.Printf("\tJokers=true\n")
	for _, h := range test_cases {
		hand := build_hand(h)
		fmt.Printf("Full House for %s: %v\n", h, is_high_card(hand, true))
	}
	fmt.Printf("\tJokers=false\n")
	for _, h := range test_cases {
		hand := build_hand(h)
		fmt.Printf("Full House for %s: %v\n", h, is_high_card(hand, false))
	}
}
