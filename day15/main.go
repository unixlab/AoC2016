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

	time := 0
	for {
		for i := 0; i < len(discs); i++ {
			discs[i].currentPosition = (discs[i].currentPosition + 1) % discs[i].Positions
		}
		if checkAllPos(discs) {
			break
		}
		time++
	}
	fmt.Println(time)
}
