package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Disc struct {
	Positions       int
	currentPosition int
}

func checkAllPos(discs []Disc) bool {
	time := 0
	for _, disc := range discs {
		if (disc.currentPosition+time)%disc.Positions != 0 {
			return false
		}
		time++
	}
	return true
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	discRegex, _ := regexp.Compile("^Disc #([0-9]+) has ([0-9]+) positions; at time=([0-9]+), it is at position ([0-9]+).$")

	var discs []Disc

	for scanner.Scan() {
		var currentDisc Disc
		regexResult := discRegex.FindAllStringSubmatch(scanner.Text(), -1)
		currentDisc.Positions, _ = strconv.Atoi(regexResult[0][2])
		currentDisc.currentPosition, _ = strconv.Atoi(regexResult[0][4])
		discs = append(discs, currentDisc)
	}

	part1Discs := make([]Disc, len(discs))
	copy(part1Discs, discs)

	part2Discs := make([]Disc, len(discs))
	copy(part2Discs, discs)
	part2Discs = append(part2Discs, Disc{11, 0})

	time := 0
	for {
		for i := 0; i < len(part1Discs); i++ {
			part1Discs[i].currentPosition = (part1Discs[i].currentPosition + 1) % part1Discs[i].Positions
		}
		if checkAllPos(part1Discs) {
			break
		}
		time++
	}
	fmt.Printf("part 1 => %d\n", time)

	time = 0
	for {
		for i := 0; i < len(part2Discs); i++ {
			part2Discs[i].currentPosition = (part2Discs[i].currentPosition + 1) % part2Discs[i].Positions
		}
		if checkAllPos(part2Discs) {
			break
		}
		time++
	}
	fmt.Printf("part 2 => %d\n", time)
}
