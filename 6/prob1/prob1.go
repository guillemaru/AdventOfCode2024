package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

var nextDirection = map[Direction]Direction{
	Up:    Right,
	Right: Down,
	Down:  Left,
	Left:  Up,
}

func main() {
	// Read the entire file into a byte slice
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	// Create a 2D slice of runes
	charGrid := make([][]rune, len(lines))

	// Fill the 2D slice with runes from each string
	for i, line := range lines {
		charGrid[i] = []rune(line)
	}
	// Create the "visited" slice with the same dimensions
	visited := make([][]bool, len(charGrid))
	for i := range visited {
		visited[i] = make([]bool, len(charGrid[i]))
	}
	currentDirection := Up
	currentPos := []int{0, 0}
	// Find the "^"
	for i, row := range charGrid {
		for j, char := range row {
			if string(char) == "^" {
				currentPos[0] = i
				currentPos[1] = j
				fmt.Println("Starting pos: ", i, j)
				visited[i][j] = true
			}
		}
	}
	for currentPos[0] >= 0 && currentPos[0] < len(charGrid) && currentPos[1] >= 0 && currentPos[1] < len(charGrid[0]) {
		nextPos := []int{currentPos[0], currentPos[1]}
		switch currentDirection {
		case Up:
			nextPos[0]--
		case Down:
			nextPos[0]++
		case Left:
			nextPos[1]--
		case Right:
			nextPos[1]++
		}
		// Is the next position a #?
		if nextPos[0] >= 0 && nextPos[0] < len(charGrid) && nextPos[1] >= 0 && nextPos[1] < len(charGrid[0]) && string(charGrid[nextPos[0]][nextPos[1]]) == "#" {
			currentDirection = nextDirection[currentDirection]
		} else {
			visited[currentPos[0]][currentPos[1]] = true
			currentPos = nextPos
		}
	}

	total := 0
	for _, row := range visited {
		for _, value := range row {
			if value {
				total++
			}
		}
	}
	fmt.Println(total)
}
