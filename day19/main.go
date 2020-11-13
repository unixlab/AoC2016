package main

import "fmt"

type Elf struct {
	Number   int
	Presents int
	Next     *Elf
	Prev     *Elf
}

func main() {
	// input := 5
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
		curElf.Presents += curElf.Next.Presents
		curElf.Next.Next.Prev = curElf
		curElf.Next = curElf.Next.Next
		if curElf.Presents == input {
			break
		}
		curElf = curElf.Next
	}
	fmt.Printf("part 1 => %d\n", curElf.Number)

	firstElf = Elf{Number: 1, Presents: 1, Next: nil, Prev: nil}
	curElf = &firstElf

	for i := 1; i < input; i++ {
		prevElf := curElf
		curElf = &Elf{Number: i + 1, Presents: 1, Next: nil, Prev: prevElf}
		prevElf.Next = curElf
	}

	curElf.Next = &firstElf
	firstElf.Prev = curElf
	curElf = curElf.Next

	var stepForward int
	if input%2 == 0 {
		stepForward = input / 2
	} else {
		stepForward = (input - 1) / 2
	}

	nextElf := curElf
	for i := 0; i < stepForward; i++ {
		nextElf = nextElf.Next
	}

	lenght := input
	for {
		curElf.Presents += nextElf.Presents
		nextElf.Next.Prev = nextElf.Prev
		nextElf.Prev.Next = nextElf.Next

		if curElf.Presents == input {
			break
		}

		curElf = curElf.Next
		nextElf = nextElf.Next
		if lenght%2 == 1 {
			nextElf = nextElf.Next
		}
		lenght--
	}
	fmt.Printf("part 2 => %d\n", curElf.Number)
}
