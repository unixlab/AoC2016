package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func runInstructions(register map[string]int, instructions []string) int {
	wordRegex, _ := regexp.Compile("([\\-A-Za-z0-9]+)")
	for i := 0; i < len(instructions); i++ {
		instruction := wordRegex.FindAllStringSubmatch(instructions[i], -1)

		switch instruction[0][0] {
		case "cpy":
			_, exists := register[instruction[1][0]]
			var copySource int
			if exists {
				copySource = register[instruction[1][0]]
			} else {
				copySource, _ = strconv.Atoi(instruction[1][0])
			}
			register[instruction[2][0]] = copySource
		case "inc":
			register[instruction[1][0]]++
		case "dec":
			register[instruction[1][0]]--
		case "jnz":
			_, exists := register[instruction[1][0]]
			var value int
			if exists {
				value = register[instruction[1][0]]
			} else {
				value, _ = strconv.Atoi(instruction[1][0])
			}
			if value != 0 {
				intValue, _ := strconv.Atoi(instruction[2][0])
				i += intValue - 1
			}
		default:
			panic("unknown instruction")
		}
	}
	return register["a"]
}
func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var instructions []string
	register := make(map[string]int)

	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	fmt.Printf("part 1 => %d\n", runInstructions(register, instructions))
	register["c"] = 1
	fmt.Printf("part 2 => %d\n", runInstructions(register, instructions))
}
