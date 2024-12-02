package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	safe := 0
	scanner := bufio.NewScanner(file)
	outer:
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Checking: %s\n", line)
		parts := strings.Split(line, " ")

		if isSafe(parts) {
			safe++
			fmt.Printf("  Safe\n")
			continue outer
		}

		// Part 2
		for i := 0; i < len(parts); i++ {
			newParts := slices.Clone(parts)
			newParts = append(newParts[:i], newParts[i+1:]...)
			fmt.Printf("  [%d]Remove check: %s\n", i, newParts)
			if isSafe(newParts) {
				safe++
				fmt.Printf("    Safe\n")
				continue outer
			}
		}
		fmt.Printf("  Unsafe\n")
	}
	fmt.Printf("Safe Count: %d\n", safe)
}

func isSafe(parts []string) bool {
	dir := 0
	last := 0
	for i, part := range parts {
		val, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		diff := val - last
		last = val
		if i == 0 {
			continue
		}
		if Abs(diff) < 1 || Abs(diff) > 3 {
			return false
		}
		if i == 1 {
			dir = diff
			continue
		}
		if (diff > 0 && dir < 0) || (diff < 0 && dir > 0) {
			return false
		}
	}
	return true
}

func Abs(n int) int {
	if n > 0 {
		return n
	} else {
		return -n
	}
}

