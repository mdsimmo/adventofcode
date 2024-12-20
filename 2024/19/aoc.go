package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	// Read available Patterns
	scanner.Scan()
	patterns := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
		return r == ' ' || r == ','
	})
	scanner.Scan() // blank line

	valid := 0
	possibilities := 0
	for scanner.Scan() {
		requiredPattern := scanner.Text()

		// map of sub pattern, to how many way we can go to the required pattern
		toScan := map[string]int{requiredPattern: 1}
		for len(toScan) > 0 {
			// Select the longest sub pattern
			longest := ""
			longestPossibilities := -1
			for key, val := range toScan {
				if len(key) >= len(longest) {
					longest = key
					longestPossibilities = val
				}
			}
			delete(toScan, longest)

			// We are complete
			if longest == "" {
				possibilities += longestPossibilities
				valid++
			}

			// Try and reduce the pattern down
			for _, pattern := range patterns {
				if !strings.HasPrefix(longest, pattern) {
					continue
				}
				newReq := longest[len(pattern):]
				toScan[newReq] += longestPossibilities
			}
		}
	}

	fmt.Printf("Valid: %d\n", valid)
	fmt.Printf("Possibilities: %d\n", possibilities)
}
