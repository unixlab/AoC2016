package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	pad := [3][3]int{[3]int{1, 4, 7}, [3]int{2, 5, 8}, [3]int{3, 6, 9}}
	x := 1
	y := 1

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case 'U':
				if y > 0 {
					y--
				}
			case 'D':
				if y < 2 {
					y++
				}
			case 'L':
				if x > 0 {
					x--
				}
			case 'R':
				if x < 2 {
					x++
				}
			}
		}
		fmt.Printf("%d", pad[x][y])
	}
	fmt.Println()
}
