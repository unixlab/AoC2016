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

func getAbas(input string) []string {
	var abas []string
	for i := 1; i < len(input)-1; i++ {
		if input[i-1] == input[i+1] && input[i] != input[i+1] {
			abas = append(abas, input[i-1:i+2])
		}
	}
	return abas
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

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	counter = 0
LINES:
	for scanner.Scan() {
		var hyperNets []string
		line := scanner.Text()
		lineSupernet := line
		for strings.Index(lineSupernet, "[") > -1 {
			startSN := strings.Index(lineSupernet, "[")
			endSN := strings.Index(lineSupernet, "]")
			hyperNets = append(hyperNets, lineSupernet[startSN+1:endSN])
			lineSupernet = lineSupernet[:startSN] + "|" + lineSupernet[endSN+1:]
		}
		var abas []string
		for _, v := range strings.Split(lineSupernet, "|") {
			abas = append(abas, getAbas(v)...)
		}
		for _, v := range abas {
			reverseAba := v[1:2] + v[0:1] + v[1:2]
			for _, net := range hyperNets {
				if strings.Index(net, reverseAba) > -1 {
					counter++
					continue LINES
				}
			}
		}
	}
	fmt.Println(counter)
}
