package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`

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
