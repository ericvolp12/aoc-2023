package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	part2(input)
}

func part1(input string) {
	maxscores := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	possibleGames := []int{}
	possibleSum := 0

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		line = strings.Split(line, ":")[1]
		rounds := strings.Split(line, ";")
		highscores := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, round := range rounds {
			moves := strings.Split(round, ",")
			for _, move := range moves {
				move = strings.Trim(move, " ")
				parts := strings.Split(move, " ")
				num, err := strconv.Atoi(parts[0])
				if err != nil {
					log.Fatalf("failed to parse number (%s): %+v", parts[0], err)
				}
				color := parts[1]
				if highscores[color] < num {
					highscores[color] = num
				}
			}
		}
		possible := true
		for color, score := range highscores {
			if score > maxscores[color] {
				possible = false
			}
		}
		if possible {
			possibleGames = append(possibleGames, i+1)
			possibleSum += i + 1
		}
	}

	log.Printf("Possible sum: %d, possible games: %+v", possibleSum, possibleGames)
}

func part2(input string) {
	powers := []int{}
	powerSum := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.Split(line, ":")[1]
		rounds := strings.Split(line, ";")
		highscores := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, round := range rounds {
			moves := strings.Split(round, ",")
			for _, move := range moves {
				move = strings.Trim(move, " ")
				parts := strings.Split(move, " ")
				num, err := strconv.Atoi(parts[0])
				if err != nil {
					log.Fatalf("failed to parse number (%s): %+v", parts[0], err)
				}
				color := parts[1]
				if highscores[color] < num {
					highscores[color] = num
				}
			}
		}
		power := 1
		for _, num := range highscores {
			power *= num
		}
		powerSum += power
		powers = append(powers, power)
	}

	log.Printf("Power sum: %d, powers: %+v", powerSum, powers)
}
