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
		ignore := false
		breaktoo := false
		for i := 0; i < len(update)-1; i++ {
			if breaktoo {
				i = 0
			}
			for j := i + 1; j < len(update); j++ {
				// If one pair does not have a rule, ignore that update
				if !containsPair(rules, update[i], update[j]) {
					ignore = true
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
