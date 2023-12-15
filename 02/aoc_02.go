package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Takes the substring "game xxx" as input and returns the number xxx
// converted into an integer
func get_game_id(line string) (int, error) {
	// Split further by separating space
	splits_1 := strings.Split(line, " ")
	game_id, err := strconv.Atoi(splits_1[1])

	return game_id, err
}

type game_draws struct {
	blue  int
	green int
	red   int
}

// Takes as input a string like
// "1 blue, 11 red, 6 green;
// Find the total of red, blue, and green draws
// func parse_game(game string) (game_draws, error) {
func parse_game(game string) game_draws {
	my_game := game_draws{0, 0, 0}
	draws := strings.Split(game, ",")
	for _, val_d := range draws {
		// fmt.Printf("ix_d = %d - val_d = %s\n", ix_d, val_d)
		// Split the string by space. Second split is number of cubes. Third split is color of cube
		draw := strings.Split(val_d, " ")
		val, _ := strconv.Atoi(draw[1])

		if strings.Compare(draw[2], "blue") == 0 {
			my_game.blue = val
		} else if strings.Compare(draw[2], "green") == 0 {
			my_game.green = val
		} else if strings.Compare(draw[2], "red") == 0 {
			my_game.red = val
		}
	}
	return my_game

}

func main() {
	fmt.Println("Advent of Code 2024 - Project 1")

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// list of game ids that are compatible with the bounds
	var valid_games []int
	const max_blue int = 14
	const max_green int = 13
	const max_red int = 12

	for scanner.Scan() {
		current_game := scanner.Text()
		// fmt.Printf("line: %s\n", current_game)
		// Split of "Game ??:" and the draws
		splits_game_draws := strings.Split(current_game, ":")

		game_id, err := get_game_id(splits_game_draws[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("game_id = %d - draws = %s\n", game_id, splits_game_draws[1])

		// Split each game into a sequence of draws. These are separated by ";"
		draw_sequence := strings.Split(splits_game_draws[1], ";")
		draws_valid := true
		for ix_g, val_g := range draw_sequence {
			fmt.Printf("\t%d :\n", ix_g)
			out := parse_game(val_g)
			fmt.Printf("Draw: blue: %d green: %d, red: %d", out.blue, out.green, out.red)
			if (out.blue > max_blue) || (out.green > max_green) || (out.red > max_red) {
				draws_valid = false
				fmt.Printf("   ====Invalid sequence ====\n")
			} else {
				fmt.Printf("\n")
			}
		}
		if draws_valid {
			valid_games = append(valid_games, game_id)
		}
	}
	sum_valid := 0
	for _, valid_game_id := range valid_games {
		sum_valid += valid_game_id
	}
	fmt.Printf("Sum of valid game ids: %d\n", sum_valid)

}
