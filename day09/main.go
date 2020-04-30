package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input, output string

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = scanner.Text()
	}

	pos := 0
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
}
