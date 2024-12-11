package main

import (
	"bufio"
	"fmt"
	"log"
	//"math"
	"os"
	//"sort"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	// Read each line and process it
	var numbers1 []int
	var numbers2 []int
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")
		num, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		numbers1 = append(numbers1, num)
		num, err = strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		numbers2 = append(numbers2, num)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	// Problem 1
	/*sort.Ints(numbers1)
	        sort.Ints(numbers2)
		difference := 0
		for i:=0;i<len(numbers1);i++ {
			// Calculate absolute value of the ith element of numbers1 and numbers2
	                abs := int(math.Abs(float64(numbers1[i])-float64(numbers2[i])))
			difference += abs
		}
		fmt.Println(difference)*/
	// Problem 2
	// Put numbers2 slice in a hasmap
	myMap := make(map[int]int)
	for i := 0; i < len(numbers2); i++ {
		// Add number2 element to myMap with its frequency
		myMap[numbers2[i]]++
	}
	similarityScore := 0
	// Iterate over the map
	for i := 0; i < len(numbers1); i++{
		// Get the frequency of numbers1[i] in myMap
		frequency := myMap[numbers1[i]]
		similarityScore += numbers1[i]*frequency
	}
	fmt.Println(similarityScore)
}
