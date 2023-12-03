package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input :=
		`................................................965..583........389.................307.................512......................395.....387
		........................#....374...382....250...*..........737*....*896.395...........*....................$.........................#......
		..494.........532-...474......*.......#....*...................522......*..........%...........................%...+................269.....
		.....*..#................506..143........375......77.....155...........400.518...64....773...718..797........694....972.603.....*...........
		....479.795...............*..........800...........*.$.......264*636.......@..............&..*...*.......499...............*...5.20.........
		515...................512.484...*....*...=......390...427...................................644.804.........*...@......-..532............28.
		..........607...........&.....105...906...910.......@............979.969...........-..=.............462....414..101.361..........283$.......
		..........*...781................................925..............=...*..........434..899....368.......*..................33...........*....
		.......850.......=........559..249...................732.....430....132........................*.....................817-.=.........613.381.
		....................157....*...&....978..............$......-...................*626.-.......297.............750............................
		..........312...........606........*....136.............593.............638...........177.....................*....672.772....998.491=......`

	part2(input)
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	lineLength := len(lines[0])

	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	expr, _ := regexp.Compile("[#\\$*+=%@/&^-]")

	input = strings.Join(lines, "")
	log.Printf("%+v", input)

	inputLen := len(input)

	expr2, _ := regexp.Compile("[0-9]+")

	nums := expr2.FindAllString(input, -1)
	locs := expr2.FindAllStringIndex(input, -1)
	valids := []bool{}
	for range locs {
		valids = append(valids, false)
	}

	log.Printf("%+s | %+v", nums, locs)

	inputArr := strings.Split(input, "")
	// Above = i - lineLength
	// Below = i + lineLength
	// UpLeft = i - lineLegth - 1
	// UpRight = i - lineLength + 1
	// DownLeft = i + lineLength - 1
	// DownRight = i + lineLength + 1
	// Right = i+1
	// Left = i-1

	for i, char := range inputArr {
		if char == "." {
			continue
		}
		var above *string
		var below *string
		var upLeft *string
		var upRight *string
		var downLeft *string
		var downRight *string
		var left *string
		var right *string

		if i > lineLength {
			above = &inputArr[i-lineLength]
		}
		if i > lineLength+1 {
			upLeft = &inputArr[i-lineLength-1]
		}
		if i > lineLength-1 {
			upRight = &inputArr[i-lineLength+1]
			// log.Printf("%d:%d | %s %+v", i, i-lineLength+1, char, *upRight)

		}
		if inputLen > i+lineLength {
			below = &inputArr[i+lineLength]
		}
		if inputLen > i+lineLength+1 {
			downRight = &inputArr[i+lineLength+1]
		}
		if inputLen > i+lineLength-1 {
			downLeft = &inputArr[i+lineLength-1]
		}
		if i > 0 {
			left = &inputArr[i-1]
		}
		if inputLen > i+1 {
			right = &inputArr[i+1]
		}

		valid := false
		if above != nil && len(expr.Find([]byte(*above))) > 0 {
			valid = true
		}
		if upLeft != nil && len(expr.Find([]byte(*upLeft))) > 0 {
			valid = true
		}
		if upRight != nil && len(expr.Find([]byte(*upRight))) > 0 {
			valid = true
		}
		if below != nil && len(expr.Find([]byte(*below))) > 0 {
			valid = true
		}
		if downLeft != nil && len(expr.Find([]byte(*downLeft))) > 0 {
			valid = true
		}
		if downRight != nil && len(expr.Find([]byte(*downRight))) > 0 {
			valid = true
		}
		if right != nil && len(expr.Find([]byte(*right))) > 0 {
			valid = true
		}
		if left != nil && len(expr.Find([]byte(*left))) > 0 {
			valid = true
		}

		if valid {
			for j, loc := range locs {
				if loc[0] <= i && loc[1] > i {
					valids[j] = true
				}
			}
		}
	}

	validNums := []int{}
	runningSum := 0
	for i, valid := range valids {
		if valid {
			asInt, _ := strconv.Atoi(nums[i])
			validNums = append(validNums, asInt)
			runningSum += asInt
		}
	}

	log.Printf("Sum: %d | Valid nums: %+v", runningSum, validNums)
}

type gear struct {
	gearIndex int
	gear1     int
	gear2     int
	ratio     int
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	lineLength := len(lines[0])

	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	expr, _ := regexp.Compile("[*]")

	input = strings.Join(lines, "")
	log.Printf("%+v", input)

	inputLen := len(input)

	expr2, _ := regexp.Compile("[0-9]+")

	nums := expr2.FindAllString(input, -1)
	locs := expr2.FindAllStringIndex(input, -1)

	type validEnt struct {
		index     int
		gearIndex int
		valid     bool
	}

	valids := []validEnt{}
	for range locs {
		valids = append(valids, validEnt{})
	}

	// log.Printf("%+s | %+v", nums, locs)

	inputArr := strings.Split(input, "")
	// Above = i - lineLength
	// Below = i + lineLength
	// UpLeft = i - lineLegth - 1
	// UpRight = i - lineLength + 1
	// DownLeft = i + lineLength - 1
	// DownRight = i + lineLength + 1
	// Right = i+1
	// Left = i-1

	for i, char := range inputArr {
		if char == "." {
			continue
		}
		var above *string
		var below *string
		var upLeft *string
		var upRight *string
		var downLeft *string
		var downRight *string
		var left *string
		var right *string

		if i > lineLength {
			above = &inputArr[i-lineLength]
		}
		if i > lineLength+1 {
			upLeft = &inputArr[i-lineLength-1]
		}
		if i > lineLength-1 {
			upRight = &inputArr[i-lineLength+1]
		}
		if inputLen > i+lineLength {
			below = &inputArr[i+lineLength]
		}
		if inputLen > i+lineLength+1 {
			downRight = &inputArr[i+lineLength+1]
		}
		if inputLen > i+lineLength-1 {
			downLeft = &inputArr[i+lineLength-1]
		}
		if i > 0 {
			left = &inputArr[i-1]
		}
		if inputLen > i+1 {
			right = &inputArr[i+1]
		}

		valid := false
		gearIndex := 0
		if above != nil && len(expr.Find([]byte(*above))) > 0 {
			valid = true
			gearIndex = i - lineLength
		}
		if upLeft != nil && len(expr.Find([]byte(*upLeft))) > 0 {
			valid = true
			gearIndex = i - lineLength - 1
		}
		if upRight != nil && len(expr.Find([]byte(*upRight))) > 0 {
			valid = true
			gearIndex = i - lineLength + 1
		}
		if below != nil && len(expr.Find([]byte(*below))) > 0 {
			valid = true
			gearIndex = i + lineLength
		}
		if downLeft != nil && len(expr.Find([]byte(*downLeft))) > 0 {
			valid = true
			gearIndex = i + lineLength - 1
		}
		if downRight != nil && len(expr.Find([]byte(*downRight))) > 0 {
			valid = true
			gearIndex = i + lineLength + 1
		}
		if right != nil && len(expr.Find([]byte(*right))) > 0 {
			valid = true
			gearIndex = i + 1
		}
		if left != nil && len(expr.Find([]byte(*left))) > 0 {
			valid = true
			gearIndex = i - 1
		}

		if valid {
			for j, loc := range locs {
				if loc[0] <= i && loc[1] > i {
					valids[j] = validEnt{valid: true, index: i, gearIndex: gearIndex}
				}
			}
		}
	}

	gears := map[int]gear{}
	runningSum := 0
	for i, ent := range valids {
		if ent.valid {
			asInt, _ := strconv.Atoi(nums[i])
			g, ok := gears[ent.gearIndex]
			if ok {
				g.gear2 = asInt
				g.ratio = g.gear1 * asInt
				gears[ent.gearIndex] = g
				runningSum += g.ratio
			} else {
				gears[ent.gearIndex] = gear{
					gearIndex: ent.gearIndex,
					gear1:     asInt,
				}
			}
		}
	}

	for i, g := range gears {
		if g.ratio == 0 {
			delete(gears, i)
		}
	}

	log.Printf("Sum: %d | Gears: %+v", runningSum, gears)
}
