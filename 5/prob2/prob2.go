package main

import (
	"fmt"
)

func main() {
	// Put rules in a map (rules.go)

	// Put updates in arrays (updates.go)

	// Loop through updates
	total := 0
	for _, update := range updates {
		// For each update, verify every ordering has a rule
		ignore := true
		breaktoo := false
		for i := 0; i < len(update)-1; i++ {
			if breaktoo {
				i = 0
			}
			for j := i + 1; j < len(update); j++ {
				// If one pair does not have a rule, swap values and restart the process for that update
				if !containsPair(rules, update[i], update[j]) {
					ignore = false
					update[i], update[j] = update[j], update[i]
					i = 0
					j = 1
					breaktoo = true
					break
				} else {
					breaktoo = false
				}
			}
		}
		if !ignore {
			middleIndex := len(update) / 2
			middleElement := update[middleIndex]
			total += middleElement
		}
	}

	// Return total
	fmt.Println(total)

}
