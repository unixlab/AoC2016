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

	newPad := [7][7]string{
		[7]string{"x", "x", "x", "x", "x", "x", "x"},
		[7]string{"x", "x", "x", "5", "x", "x", "x"},
		[7]string{"x", "x", "2", "6", "A", "x", "x"},
		[7]string{"x", "1", "3", "7", "B", "D", "x"},
		[7]string{"x", "x", "4", "8", "C", "x", "x"},
		[7]string{"x", "x", "x", "9", "x", "x", "x"},
		[7]string{"x", "x", "x", "x", "x", "x", "x"}}

	x = 1
	y = 3
	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case 'U':
				if newPad[x][y-1] != "x" {
					y--
				}
			case 'D':
				if newPad[x][y+1] != "x" {
					y++
				}
			case 'L':
				if newPad[x-1][y] != "x" {
					x--
				}
			case 'R':
				if newPad[x+1][y] != "x" {
					x++
				}
			}
		}
		fmt.Printf("%s", newPad[x][y])
	}
	fmt.Println()
}
