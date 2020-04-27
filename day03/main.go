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

	for scanner.Scan() {
		var sides [3]int
		re := regexp.MustCompile(`([^ \t]+)`)
		for k, v := range re.FindAllString(scanner.Text(), -1) {
			sides[k], _ = strconv.Atoi(v)
		}
		if sides[0]+sides[1] > sides[2] && sides[0]+sides[2] > sides[1] && sides[1]+sides[2] > sides[0] {
			counter++
		}
	}
	fmt.Println(counter)
}
