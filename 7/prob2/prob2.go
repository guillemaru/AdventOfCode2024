package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)
// Get number of digits of int
func getDigits(num int64) int64 {
        return int64(math.Floor(math.Log10(float64(num)) + 1))
}

func main() {
	// Read the entire file into a byte slice
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	// remove last line as its empty
	lines = lines[0 : len(lines)-1]
	var totalSum int64 = 0
	for _, line := range lines {
		var numbers []int64
		splitline := strings.Split(line, ":")
		numbersStr := strings.Split(splitline[1], " ")
		for _, numberStr := range numbersStr {
			number, _ := strconv.ParseInt(numberStr, 10, 64)
			numbers = append(numbers, number)
		}
		numbers = numbers[1:]
		bigNumber, _ := strconv.ParseInt(splitline[0], 10, 64)
		maxSize := int(math.Pow(3, (float64)(len(numbers)-1)))
		//fmt.Println(bigNumber, maxSize)
		//fmt.Println(numbers)
		for i := 0; i < maxSize; i++ {
			j := i
			// Create a deep copy of the slice
			cpyNumbers := make([]int64, len(numbers))
			copy(cpyNumbers , numbers)
			//fmt.Println("cpyNumbers", cpyNumbers)
			//for k := 0; k < len(numbers)-1; k++ {
			//	remainder := j % 2
			//	j /= 2
			//	if remainder == 1 {
			//		cpyNumbers[k+1] *= cpyNumbers[k]
			//		cpyNumbers[k] = 0
			//	}
			//}
			//var sum int64 = 0
			//for k := 0; k < len(numbers); k++ {
			//	fmt.Print(cpyNumbers[k], " ")
			//	sum += cpyNumbers[k]
			//}
			var sum int64 = cpyNumbers[0]
			for k := 0; k < len(numbers)-1; k++ {
				remainder := j % 3
				j /= 3
				if remainder == 1 {
					sum *= cpyNumbers[k+1]
				} else if remainder == 0 {
					sum += cpyNumbers[k+1]
				} else {
					sum = sum * int64(math.Pow(10, float64(getDigits(cpyNumbers[k+1])))) + cpyNumbers[k+1]
				}
				//fmt.Print(remainder, " ")
			}
			//fmt.Println(": ", sum)

			if sum == bigNumber {
				totalSum += sum
				break
			}
		}
	}
	fmt.Println(totalSum)
}
