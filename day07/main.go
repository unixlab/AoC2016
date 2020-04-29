package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkABBA(input string) bool {
	for i := 97; i < 123; i++ {
		doubleCharPos := strings.Index(input, (string(i) + string(i)))
		if doubleCharPos > 0 && doubleCharPos < len(input)-2 {
			if input[doubleCharPos-1] == input[doubleCharPos+2] && input[doubleCharPos] != input[doubleCharPos+2] {
				return true
			}
		}
	}
	return false
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		for strings.Index(line, "[") > -1 {
			startABBA := strings.Index(line, "[")
			endABBA := strings.Index(line, "]")
			abba := line[startABBA+1 : endABBA]
			if checkABBA(abba) {
				break
			}
			line = line[:startABBA] + "|" + line[endABBA+1:]
		}
		if strings.Index(line, "[") > -1 {
			continue
		}
		for _, v := range strings.Split(line, "|") {
			if checkABBA(v) {
				counter++
				break
			}
		}
	}
	fmt.Println(counter)
}
