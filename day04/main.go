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
	}
	return false
}

func shiftBy(input string, counter int) string {
	for i := 0; i < len(input); i++ {
		char := rune(input[i])
		for j := 0; j < counter; j++ {
			if int(char) == 122 {
				char = 'a'
			} else if char == ' ' {
				char = '-'
			} else if char == '-' {
				char = ' '
			} else {
				char = rune(int(char) + 1)
			}
		}
		input = input[:i] + string(char) + input[i+1:]
	}
	return input
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
		sectorID, _ := strconv.Atoi(input[sep+1 : bracketPos])

		if verifyChecksum(room, checksum) {
			sum += sectorID
		}
	}
	fmt.Println(sum)

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		bracketPos := strings.Index(input, "[")
		input = input[:bracketPos]
		sep := strings.LastIndex(input, "-")
		seq, _ := strconv.Atoi(input[sep+1:])
		input = input[:sep]

		input = shiftBy(input, seq)

		if strings.HasPrefix(input, "north") {
			fmt.Println(seq)
		}
	}
}
