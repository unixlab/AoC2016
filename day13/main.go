package main

import (
	"fmt"
	"strings"
)

type Position struct {
	Y int
	X int
}

const favoriteNumber = 1350
const gridSize = 50

var grid [gridSize][gridSize]bool
var distance [gridSize][gridSize]int

func getObstacle(curPos Position) bool {
	x := curPos.X
	y := curPos.Y
	value := x*x + 3*x + 2*x*y + y + y*y
	value += favoriteNumber
	binaryRepresentation := fmt.Sprintf("%b", value)
	numberOfOnes := strings.Count(binaryRepresentation, "1")
	if numberOfOnes%2 == 0 {
		return false
	}
	return true
}

func getLinearSurroundings(curPos Position) []Position {
	var positions []Position
	if curPos.Y-1 >= 0 {
		if !grid[curPos.Y-1][curPos.X] {
			positions = append(positions, Position{curPos.Y - 1, curPos.X})
		}
	}
	if curPos.Y+1 < gridSize {
		if !grid[curPos.Y+1][curPos.X] {
			positions = append(positions, Position{curPos.Y + 1, curPos.X})
		}
	}
	if curPos.X-1 >= 0 {
		if !grid[curPos.Y][curPos.X-1] {
			positions = append(positions, Position{curPos.Y, curPos.X - 1})
		}
	}
	if curPos.X+1 < gridSize {
		if !grid[curPos.Y][curPos.X+1] {
			positions = append(positions, Position{curPos.Y, curPos.X + 1})
		}
	}
	return positions
}

func calculateDistance(curPos Position, curStep int) {
	distance[curPos.Y][curPos.X] = curStep
	curStep++
	reachablePositions := getLinearSurroundings(curPos)
	for _, pos := range reachablePositions {
		if distance[pos.Y][pos.X] == 0 || distance[pos.Y][pos.X] > curStep+1 {
			calculateDistance(pos, curStep)
		}
	}
}

func main() {
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if getObstacle(Position{y, x}) {
				grid[y][x] = true
			} else {
				grid[y][x] = false
			}
		}
	}
	calculateDistance(Position{1, 1}, 0)
	fmt.Printf("part 1 => %d\n", distance[39][31])
}
