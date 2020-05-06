package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Bot struct {
	BusyHands int
	HandA     int
	HandB     int
}

func (b Bot) low(botNr int) int {
	checkCompare(b.HandA, b.HandB, botNr)
	if b.HandA < b.HandB {
		return b.HandA
	}
	return b.HandB
}

func (b Bot) high(botNr int) int {
	checkCompare(b.HandA, b.HandB, botNr)
	if b.HandA < b.HandB {
		return b.HandB
	}
	return b.HandA
}

func (b Bot) add(value int) Bot {
	if b.BusyHands == 0 {
		b.HandA = value
		b.BusyHands++
	} else if b.BusyHands == 1 {
		b.HandB = value
		b.BusyHands++
	} else {
		panic("bot overloaded")
	}
	return b
}

func checkCompare(a int, b int, c int) {
	if a == 61 && b == 17 {
		fmt.Println(c)
	}
	if a == 17 && b == 61 {
		fmt.Println(c)
	}
}

func main() {
	var output [100]int
	var lines []string
	bots := make(map[int]Bot, 100)

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for len(lines) > 0 {
		line := lines[0]
		lines = lines[1:]
		regexValue := regexp.MustCompile("value ([0-9]+) goes to bot ([0-9]+)")
		regexBot := regexp.MustCompile("bot ([0-9]+) gives ([a-z]+) to ([a-z]+) ([0-9]+) and ([a-z]+) to ([a-z]+) ([0-9]+)")

		if regexValue.MatchString(line) {
			tempRegexMatch := regexValue.FindStringSubmatch(line)
			value, _ := strconv.Atoi(tempRegexMatch[1])
			botNumber, _ := strconv.Atoi(tempRegexMatch[2])
			bot := bots[botNumber]
			bot = bot.add(value)
			bots[botNumber] = bot
		} else if regexBot.MatchString(line) {
			tempRegexMatch := regexBot.FindStringSubmatch(line)
			botNumber1, _ := strconv.Atoi(tempRegexMatch[1])
			mode1 := tempRegexMatch[2]
			dest1 := tempRegexMatch[3]
			dest1Number, _ := strconv.Atoi(tempRegexMatch[4])
			mode2 := tempRegexMatch[5]
			dest2 := tempRegexMatch[6]
			dest2Number, _ := strconv.Atoi(tempRegexMatch[7])
			bot := bots[botNumber1]
			if bot.BusyHands < 2 {
				lines = append(lines, line)
				continue
			}
			var value int
			if mode1 == "low" {
				value = bot.low(botNumber1)
			} else {
				value = bot.high(botNumber1)
			}
			if dest1 == "bot" {
				bot := bots[dest1Number]
				bot = bot.add(value)
				bots[dest1Number] = bot
			} else {
				output[dest1Number] = value
			}
			if mode2 == "low" {
				value = bot.low(botNumber1)
			} else {
				value = bot.high(botNumber1)
			}
			if dest2 == "bot" {
				bot := bots[dest2Number]
				bot = bot.add(value)
				bots[dest2Number] = bot
			} else {
				output[dest2Number] = value
			}
		} else {
			panic("unkown line")
		}
	}
	fmt.Println(output[0] * output[1] * output[2])
}
