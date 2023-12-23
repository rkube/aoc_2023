package main

import (
	"errors"
	"fmt"
	"log"
)

// Returns true of lhs > rhs, using the second rule
// Parse cards in both hands sequentially. If the card in the
// first hand is greater than the card second hand, return true
func greater_than_o2(lhs hand_bid, rhs hand_bid) (bool, error) {
	card_values := get_card_values()
	for ix_c := 0; ix_c < 5; ix_c++ {
		lhs_card := rune(lhs.cards[ix_c])
		rhs_card := rune(rhs.cards[ix_c])
		// fmt.Printf("ix_c = %d, lhs_card: %#U = %d, rhs_card: %#U = %d\n", ix_c, lhs_card, card_values[lhs_card], rhs_card, card_values[rhs_card])
		if card_values[lhs_card] > card_values[rhs_card] {
			return true, nil
		} else if card_values[lhs_card] < card_values[rhs_card] {
			// fmt.Printf("%d < %d\n", card_values[lhs_card], card_values[rhs_card])
			return false, nil
		} else {
			continue
		}
	}
	error_str := fmt.Sprintf("Hands are identical (o2): %s - %s\n", lhs.cards, rhs.cards)
	return false, errors.New(error_str)
}

// Returns true of lhs > rhs
// Compare first by hand_type
// If hand types are equal, compare cards in order
func greater_than(lhs hand_bid, rhs hand_bid) (bool, error) {
	if find_hand_type(lhs) > find_hand_type(rhs) {
		return true, nil
	} else if find_hand_type(lhs) == find_hand_type(rhs) {
		geq, err := greater_than_o2(lhs, rhs)
		if err != nil {
			log.Fatal(err)
		}
		return geq, nil
	} else if find_hand_type(lhs) < find_hand_type(rhs) {
		return false, nil
	}
	error_str := fmt.Sprintf("Hands are identical: %s - %s\n", lhs.cards, rhs.cards)
	return false, errors.New(error_str)
}
