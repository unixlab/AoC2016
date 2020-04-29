package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var rows [8][]string

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for i := 0; i < len(line); i++ {
			rows[i] = append(rows[i], line[i])
		}
	}

	var minWord, maxWord string
	for _, row := range rows {
		chars := make(map[string]int, len(row))
		for _, v := range row {
			chars[v]++
		}
		minKey := ""
		maxKey := ""
		minValue := 0
		maxValue := 0
		for k, v := range chars {
			if v > maxValue {
				maxKey = k
				maxValue = v
			}
			if v < minValue || minValue == 0 {
				minKey = k
				minValue = v
			}
		}
		minWord += minKey
		maxWord += maxKey
	}
	fmt.Println(maxWord)
	fmt.Println(minWord)
}
