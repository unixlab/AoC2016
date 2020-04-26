package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getDistance(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2)))
}

func main() {
	var input string

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = scanner.Text()
	}

	d := 0
	x := 0
	y := 0

	steps := strings.Split(strings.ReplaceAll(input, " ", ""), ",")

	for _, step := range steps {
		if step[0] == 'R' {
			if d == 360 {
				d = 90
			} else {
				d += 90
			}
		}
		if step[0] == 'L' {
			if d == 0 {
				d = 270
			} else {
				d -= 90
			}
		}

		move, _ := strconv.Atoi(step[1:])

		switch d {
		case 0:
			y += move
		case 90:
			x += move
		case 180:
			y -= move
		case 270:
			x -= move
		case 360:
			y += move
		default:
			panic(d)
		}
	}

	fmt.Println(getDistance(0, 0, x, y))
}
