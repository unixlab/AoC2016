package main

import "fmt"

type Elf struct {
	Number   int
	Presents int
	Next     *Elf
	Prev     *Elf
}

func main() {
	input := 3018458

	firstElf := Elf{Number: 1, Presents: 1, Next: nil, Prev: nil}
	curElf := &firstElf

	for i := 1; i < input; i++ {
		prevElf := curElf
		curElf = &Elf{Number: i + 1, Presents: 1, Next: nil, Prev: prevElf}
		prevElf.Next = curElf
	}

	curElf.Next = &firstElf
	firstElf.Prev = curElf
	curElf = curElf.Next

	for {
		if curElf.Presents > 0 {
			nextElf := curElf.Next
			for nextElf.Presents < 1 {
				nextElf = nextElf.Next
			}
			curElf.Presents += nextElf.Presents
			nextElf.Presents = 0
		}
		if curElf.Presents == input {
			break
		}
		curElf = curElf.Next
	}
	fmt.Println(curElf.Number)
}
