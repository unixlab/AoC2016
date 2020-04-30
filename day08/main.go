package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/fatih/color"
)

func main() {
	var grid [50][6]bool
	counter := 0
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		a := 0
		b := 0
		mode := 0
		line := scanner.Text()
		regex := regexp.MustCompile("rect ([0-9]+)x([0-9]+)")
		if regex.MatchString(line) {
			mode = 1
			tempRegexMatch := regex.FindStringSubmatch(line)
			a, _ = strconv.Atoi(tempRegexMatch[1])
			b, _ = strconv.Atoi(tempRegexMatch[2])
		}
		regex = regexp.MustCompile("rotate row y=([0-9]+) by ([0-9]+)")
		if regex.MatchString(line) {
			mode = 2
			tempRegexMatch := regex.FindStringSubmatch(line)
			a, _ = strconv.Atoi(tempRegexMatch[1])
			b, _ = strconv.Atoi(tempRegexMatch[2])

		}
		regex = regexp.MustCompile("rotate column x=([0-9]+) by ([0-9]+)")
		if regex.MatchString(line) {
			mode = 3
			tempRegexMatch := regex.FindStringSubmatch(line)
			a, _ = strconv.Atoi(tempRegexMatch[1])
			b, _ = strconv.Atoi(tempRegexMatch[2])
		}

		switch mode {
		case 1:
			for x := 0; x < a; x++ {
				for y := 0; y < b; y++ {
					grid[x][y] = true
				}
			}
		case 2:
			for i := 0; i < b; i++ {
				var store bool
				y := a
				for x := 0; x < len(grid)-1; x++ {
					store = grid[x+1][y]
					grid[x+1][y] = grid[0][y]
					grid[0][y] = store
				}
			}
		case 3:
			for i := 0; i < b; i++ {
				var store bool
				x := a
				for y := 0; y < len(grid[x])-1; y++ {
					store = grid[x][y+1]
					grid[x][y+1] = grid[x][0]
					grid[x][0] = store
				}
			}
		}
	}

	for y := 0; y < len(grid[0]); y++ {
		for x := 0; x < len(grid); x++ {
			if grid[x][y] {
				counter++
			}
		}
	}

	fmt.Println(counter)

	for y := 0; y < len(grid[0]); y++ {
		for x := 0; x < len(grid); x++ {
			if grid[x][y] {
				color.New(color.BgWhite).Printf(" ")
			} else {
				color.New(color.BgBlack).Printf(" ")
			}
		}
		fmt.Println()
	}
}
