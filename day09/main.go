package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func decompressedLength(input string) int {
	pos := strings.Index(input, "(")
	if pos > 0 {
		return pos + decompressedLength(input[pos:])
	}
	if pos == 0 {
		stringNumbers := strings.Split(input[pos+1:strings.Index(input, ")")], "x")
		length, _ := strconv.Atoi(stringNumbers[0])
		start := pos + 1 + strings.Index(input, ")")
		realLength := decompressedLength(input[start : start+length])
		repeat, _ := strconv.Atoi(stringNumbers[1])
		return pos + realLength*repeat + decompressedLength(input[pos+strings.Index(input, ")")+1+length:])
	}
	return len(input)
}

func main() {
	var line, input, output string

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
	}

	pos := 0
	input = line
	for pos < len(input) {
		if strings.Index(input[pos:], "(") > -1 {
			startPos := strings.Index(input[pos:], "(")
			endPos := strings.Index(input[pos:], ")")
			output += input[pos : pos+startPos]
			formular := input[pos+startPos+1 : pos+endPos]
			xPos := strings.Index(formular, "x")
			number, _ := strconv.Atoi(formular[:xPos])
			repeat, _ := strconv.Atoi(formular[xPos+1:])
			for i := 0; i < repeat; i++ {
				output += input[pos+endPos+1 : pos+endPos+1+number]
			}
			pos += endPos + 1 + number
		} else {
			output += input[pos:]
			pos = len(input)
		}
	}
	fmt.Println(len(output))
	fmt.Println(decompressedLength(line))
}
