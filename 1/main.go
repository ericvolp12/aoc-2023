package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var numMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	input := `two1nine
	eightwothree
	eightwo
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`

	part2(input)
}

func part1(input string) {
	calibrationNumbers := []int{}
	runningSum := 0

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		var firstNum *int
		var lastNum *int

		chars := strings.Split(line, "")
		for _, char := range chars {
			asInt, err := strconv.Atoi(char)
			if err == nil {
				if firstNum == nil {
					firstNum = &asInt
				}
				lastNum = &asInt
			}
		}

		if firstNum == nil || lastNum == nil {
			log.Fatalf("Didn't find two nums on input line %d", i)
		}

		combinedNum, err := strconv.Atoi(fmt.Sprintf("%d%d", *firstNum, *lastNum))
		if err != nil {
			log.Fatalf("failed to build combined num for (%d, %d): %+v", *firstNum, *lastNum, err)
		}

		calibrationNumbers = append(calibrationNumbers, combinedNum)
		runningSum += combinedNum
	}

	log.Printf("Sum: %d, Nums: %+v", runningSum, calibrationNumbers)
}

func part2(input string) {
	calibrationNumbers := []int{}
	runningSum := 0

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		var firstNum *int
		var lastNum *int

		chars := strings.Split(line, "")
		buf := ""
		for _, char := range chars {
			asInt, err := strconv.Atoi(char)
			if err == nil {
				if firstNum == nil {
					firstNum = &asInt
				}
				lastNum = &asInt
				buf = ""
				continue
			}
			buf += char
			for numword := range numMap {
				num := numMap[numword]
				if strings.Contains(buf, numword) {
					if firstNum == nil {
						firstNum = &num
					}
					lastNum = &num
					buf = char
				}
			}
		}

		if firstNum == nil || lastNum == nil {
			log.Fatalf("Didn't find two nums on input line %d", i)
		}

		combinedNum, err := strconv.Atoi(fmt.Sprintf("%d%d", *firstNum, *lastNum))
		if err != nil {
			log.Fatalf("failed to build combined num for (%d, %d): %+v", *firstNum, *lastNum, err)
		}

		calibrationNumbers = append(calibrationNumbers, combinedNum)
		runningSum += combinedNum
	}

	log.Printf("Sum: %d, Nums: %+v", runningSum, calibrationNumbers)
}
