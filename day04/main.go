package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func verifyChecksum(room string, checksum string) bool {
	characters := make(map[string]int, 26)
	var calcCheckSum []string

	for len(room) > 0 {
		char := room[0:1]
		characters[char] = strings.Count(room, char)
		room = strings.ReplaceAll(room, char, "")
	}

	for i := 0; i < 5; i++ {
		max := 0
		for _, v := range characters {
			if v > max {
				max = v
			}
		}

		counter := 0
		for _, v := range characters {
			if v == max {
				counter++
			}
		}

		if counter == 1 {
			for k, v := range characters {
				if v == max {
					calcCheckSum = append(calcCheckSum, k)
					delete(characters, k)
				}
			}
		} else {
			var chars []string
			for k, v := range characters {
				if v == max {
					chars = append(chars, k)
				}
			}
			sort.Strings(chars)
			calcCheckSum = append(calcCheckSum, chars[0])
			delete(characters, chars[0])
		}
	}

	if strings.Join(calcCheckSum, "") == checksum {
		return true
	} else {
		return false
	}
}

func main() {
	sum := 0
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		sep := strings.LastIndex(input, "-")
		room := strings.ReplaceAll(input[:sep], "-", "")
		bracketPos := strings.Index(input, "[")
		checksum := input[bracketPos+1 : len(input)-1]
		sectorId, _ := strconv.Atoi(input[sep+1 : bracketPos])

		if verifyChecksum(room, checksum) {
			sum += sectorId
		}
	}
	fmt.Println(sum)
}
