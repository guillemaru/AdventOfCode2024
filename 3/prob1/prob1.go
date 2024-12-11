package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Read the entire file into a byte slice
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	// regex pattern: mul\(\d+,\d+\)
	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(string(data), -1)
	//fmt.Println(matches)

	total := 0
	for _, match := range matches {
		// match[1] and match[2] are the captured numbers
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		total += x * y
	}

	fmt.Println("Total:", total)
}
