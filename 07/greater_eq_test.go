package main

import (
	"fmt"
	"log"
	"testing"
)

func Test_greater_than_o2(t *testing.T) {
	fmt.Printf("====== Test > operator\n")
	lhs_hand_str := "T2JJJ"
	lhs_this_hand := initialize_hand_map()
	for _, card := range lhs_hand_str {
		lhs_this_hand[card] += 1
	}
	lhs := hand_bid{cards: lhs_hand_str, bid: 0, hand: lhs_this_hand}

	rhs_hand_str := "T2222"
	rhs_this_hand := initialize_hand_map()
	for _, card := range rhs_hand_str {
		rhs_this_hand[card] += 1
	}
	rhs := hand_bid{cards: rhs_hand_str, bid: 0, hand: rhs_this_hand}

	v, err := greater_than_o2(lhs, rhs, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("(jokers = false) %s > %s:  %v\n", lhs_hand_str, rhs_hand_str, v)

	v, err = greater_than_o2(lhs, rhs, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("(jokers = true) %s > %s:  %v\n", lhs_hand_str, rhs_hand_str, v)
}
