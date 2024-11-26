package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	part2 := true
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	
	scanner := bufio.NewScanner(file)
	valid := 0
	outer:
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		// Sort the letters for part two (anagrams)
		if part2 {
			for i := 0; i < len(parts); i++ {
				byteParts := []byte(parts[i])
				slices.Sort(byteParts)
				parts[i] = string(byteParts)
			}
		}
		
		// Count no matches
		for i := 0; i < len(parts); i++ {
			for j := i + 1; j < len(parts); j++ {
				if parts[i] == parts[j] {
					continue outer
				}
			}
		}
		valid++
	}
	fmt.Printf("Valid Passes: %d\n", valid)
}
