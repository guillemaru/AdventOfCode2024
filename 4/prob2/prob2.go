package main

import (
	"fmt"
	"os"
	"strings"
)

func stringTo2DArray(input string) [][]byte {
	// Split the string into lines
	lines := strings.Split(input, "\n")

	// Create a 2D slice to hold the characters
	var charArray [][]byte

	// Iterate over each line
	for _, line := range lines {
		// Convert the line to a slice of runes
		charLine := []byte(line)
		// Append to the 2D slice
		charArray = append(charArray, charLine)
	}

	return charArray
}

func main() {
	// Read the entire file into a byte slice
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	dataString := string(data)
	data2DArray := stringTo2DArray(dataString)
	// The last row appears to be empty... so delete it
	data2DArray = data2DArray[:len(data2DArray)-1]
	total := 0
	// Split the array in 8 ways:
	// all horizontal and backwards,
	for i := 1; i < len(data2DArray) - 1; i++ {
		for j := 1; j < len(data2DArray[0]) - 1; j++ {
			if data2DArray[i][j] == 'A' {
				pos1 := data2DArray[i-1][j-1]
				pos2 := data2DArray[i-1][j+1]
				pos3 := data2DArray[i+1][j-1]
				pos4 := data2DArray[i+1][j+1]
				amountofM := 0
				amountofS := 0
				if pos1 == 'M' {
					amountofM++
				}
				if pos2 == 'M' {
					amountofM++
				}
				if pos3 == 'M' {
					amountofM++
				}
				if pos4 == 'M' {
					amountofM++
				}
				if pos1 == 'S' {
					amountofS++
				}
				if pos2 == 'S' {
					amountofS++
				}
				if pos3 == 'S' {
					amountofS++
				}
				if pos4 == 'S' {
					amountofS++
				}
				if amountofM == 2 && amountofS == 2 {
					if pos2 != pos3 {
						total++
					}
				}
			}
		}
	}

	fmt.Println("Total:", total)
}
