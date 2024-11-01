package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func aoc_8_run() {
	file, _ := os.Open("aoc_8_in.txt")
	scanner := bufio.NewScanner(file)
	sum := 0
	sum2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		chars := len(line)
		memchars := (count(line)-2) // 2 for quotes
		sum += chars - memchars
		fmt.Printf("%s: %d, %d\n", line, chars, memchars)

		sum2 += strings.Count(line, "\"") + strings.Count(line, "\\") + 2 

	}
	fmt.Printf("Result 1: %d\n", sum)
	fmt.Printf("result 2: %d\n", sum2)
}

func count(input string) int {
	count := 0
	for i := 0; i < len(input); i++ {
		char := input[i]
		if char == '\\' {
			switch input[i+1] {
			case 'x':
				i += 3
			default:
				i += 1
			}
		}
		count++
	}
	return count
}

