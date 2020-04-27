package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	counter := 0

	var row1 []int
	var row2 []int
	var row3 []int

	for scanner.Scan() {
		var sides [3]int
		re := regexp.MustCompile(`([^ \t]+)`)
		for k, v := range re.FindAllString(scanner.Text(), -1) {
			sides[k], _ = strconv.Atoi(v)
		}
		if sides[0]+sides[1] > sides[2] && sides[0]+sides[2] > sides[1] && sides[1]+sides[2] > sides[0] {
			counter++
		}
		row1 = append(row1, sides[0])
		row2 = append(row2, sides[1])
		row3 = append(row3, sides[2])
	}
	fmt.Println(counter)

	counter = 0

	var rows []int
	rows = append(rows, row1...)
	rows = append(rows, row2...)
	rows = append(rows, row3...)

	for i := 2; i < len(rows); i = i + 3 {
		if rows[i-2]+rows[i-1] > rows[i] && rows[i-2]+rows[i] > rows[i-1] && rows[i-1]+rows[i] > rows[i-2] {
			counter++
		}
	}
	fmt.Println(counter)
}
