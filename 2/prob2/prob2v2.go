package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	// Create a new slice to avoid modifying the original
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return append(newSlice[:index], newSlice[index+1:]...)
}

func isSafe(numbers []int) bool {
        return (isAllDecreasing(numbers) || isAllIncreasing(numbers)) && checkAdjacentLevels(numbers)
}

func isAllDecreasing(numbers []int) bool {
	prevNumber := -1
	for _, number := range numbers {
		if prevNumber == -1 {
			prevNumber = number
			continue
		}
		if number > prevNumber {
			return false
		}
		prevNumber = number
	}
	return true
}

func isAllIncreasing(numbers []int) bool {
	prevNumber := -1
	for _, number := range numbers {
		if prevNumber == -1 {
			prevNumber = number
			continue
		}
		if number < prevNumber {
			return false
		}
		prevNumber = number
	}
	return true
}

func checkAdjacentLevels(numbers []int) bool {
	prevNumber := -1
	for _, number := range numbers {
		if prevNumber == -1 {
			prevNumber = number
			continue
		}
		diff := int(math.Abs(float64(number) - float64(prevNumber)))
		if diff < 1 || diff > 3 {
			return false
		}
		prevNumber = number
	}
	return true
}

func main() {
	// Read the entire file into a byte slice
	data, err := os.ReadFile("../numbers.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte slice to a string and split by lines
	lines := strings.Split(string(data), "\n")
	// Remove the last element
	if len(lines) > 0 {
		if lines[len(lines)-1] == "" {
			lines = lines[:len(lines)-1]
		}
	}

	safeLines := 0

	// Iterate over each line
	for _, line := range lines {
		// Prepare a slice to hold all numbers
		var numbers []int
		// Split the line by spaces
		fields := strings.Fields(line)
		for _, field := range fields {
			// Convert each field to an integer
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			numbers = append(numbers, num)
		}
		if isSafe(numbers) {
			safeLines++
		} else {
			for i:= 0; i < len(numbers); i++{
                                newNumbers := removeElement(numbers, i)
                                if isSafe(newNumbers) {
                                        safeLines++
					break
                                }
                        }
                }
	}

	fmt.Println(safeLines)
}
