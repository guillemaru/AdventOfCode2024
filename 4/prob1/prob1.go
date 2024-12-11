package main

import (
	"fmt"
	"os"
	"regexp"
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

func reverseString(s string) string {
	// Convert the string to a slice of runes
	runes := []rune(s)

	// Reverse the slice of runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the slice of runes back to a string
	return string(runes)
}

func getNthColumn(array [][]byte, n int) []byte {
	if len(array) == 0 || n < 0 || n >= len(array[0]) {
		return nil // Return nil if the column index is out of bounds
	}

	column := make([]byte, len(array))
	for i, row := range array {
		column[i] = row[n]
	}

	return column
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
	pattern := `XMAS`
	re := regexp.MustCompile(pattern)
	// Split the array in 8 ways:
	// all horizontal and backwards,
	for i := 0; i < len(data2DArray); i++ {
		row := data2DArray[i]
		matches := re.FindAllStringSubmatch(string(row), -1)
		total += len(matches)
		reverseRow := reverseString(string(row))
		matches = re.FindAllStringSubmatch(reverseRow, -1)
		total += len(matches)
	}

	// all vertical and backwards,
	for i := 0; i < len(data2DArray[0]); i++ {
		column := getNthColumn(data2DArray, i)
		matches := re.FindAllStringSubmatch(string(column), -1)
		total += len(matches)
		reversecolumn := reverseString(string(column))
		matches = re.FindAllStringSubmatch(reversecolumn, -1)
		total += len(matches)
	}
	// all diagonal in one direction and backwards,
	// Traverse diagonals from top-left to bottom-right
	for start := 0; start < len(data2DArray[0]); start++ {
		var diag []byte
		for i, j := 0, start; i < len(data2DArray) && j < len(data2DArray[0]); i, j = i+1, j+1 {
			// Access data2DArray[i][j]
			diag = append(diag, data2DArray[i][j])
		}
		matches := re.FindAllStringSubmatch(string(diag), -1)
		total += len(matches)
		reverse := reverseString(string(diag))
		matches = re.FindAllStringSubmatch(reverse, -1)
		total += len(matches)
	}

	for start := 1; start < len(data2DArray); start++ {
		var diag []byte
		for i, j := start, 0; i < len(data2DArray) && j < len(data2DArray[0]); i, j = i+1, j+1 {
			// Access data2DArray[i][j]
			diag = append(diag, data2DArray[i][j])
		}
		matches := re.FindAllStringSubmatch(string(diag), -1)
		total += len(matches)
		reverse := reverseString(string(diag))
		matches = re.FindAllStringSubmatch(reverse, -1)
		total += len(matches)
	}
	// and all diagonal in the opposite direction and backwards.
	// Traverse diagonals from top-right to bottom-left
	for start := len(data2DArray[0]) - 1; start >= 0; start-- {
		var diag []byte
		for i, j := 0, start; i < len(data2DArray) && j >= 0; i, j = i+1, j-1 {
			// Access data2DArray[i][j]
			diag = append(diag, data2DArray[i][j])
		}
		matches := re.FindAllStringSubmatch(string(diag), -1)
		total += len(matches)
		reverse := reverseString(string(diag))
		matches = re.FindAllStringSubmatch(reverse, -1)
		total += len(matches)
	}

	for start := 1; start < len(data2DArray); start++ {
		var diag []byte
		for i, j := start, len(data2DArray[0])-1; i < len(data2DArray) && j >= 0; i, j = i+1, j-1 {
			// Access data2DArray[i][j]
			diag = append(diag, data2DArray[i][j])
		}
		matches := re.FindAllStringSubmatch(string(diag), -1)
		total += len(matches)
		reverse := reverseString(string(diag))
		matches = re.FindAllStringSubmatch(reverse, -1)
		total += len(matches)
	}

	fmt.Println("Total:", total)
}
