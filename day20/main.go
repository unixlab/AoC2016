package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var ips [4294967296]bool

	for i := 0; i < 4294967296; i++ {
		ips[i] = true
	}

	inputFile, inputFileError := os.Open("input.txt")
	if inputFileError != nil {
		panic(inputFileError)
	}

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		min, _ := strconv.Atoi(strings.Split(scanner.Text(), "-")[0])
		max, _ := strconv.Atoi(strings.Split(scanner.Text(), "-")[1])
		for min < max {
			ips[min] = false
			min++
		}
		ips[max] = false
	}

	firstIP := -1
	validIPs := 0
	for i := 0; i < 4294967296; i++ {
		if ips[i] {
			if firstIP == -1 {
				firstIP = i
			}
			validIPs++
		}
	}

	fmt.Printf("part 1 => %d\n", firstIP)
	fmt.Printf("part 2 => %d\n", validIPs)
}
