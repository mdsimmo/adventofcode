package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func aoc_5_run() {
	file, err := os.Open("aoc_5_in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		text := scanner.Text()
		good := test2(text)
		if good {
			count++
		}
		fmt.Printf("%s, is %t\n", text, good)
	}

	fmt.Printf("Result: %d", count)	
	
}

func test(input string) bool {
	bad := []string { "ab", "cd", "pq", "xy" }
	has_double := false
	vowel_count := 0
	last := '-'
	vowels := []rune { 'a', 'e', 'i', 'o', 'u' }

	for _, c := range(input) {
		if c < 'a' || c > 'z' {
			continue
		}
		two := string(last) + string(c)
		if slices.Contains(bad, two) {
			fmt.Printf("Has bad: %s\n", two)
			return false
		}
		if slices.Contains(vowels, c) {
			vowel_count++
		}
		if last == c {
			has_double = true
		}
		last = c
	}
	if vowel_count < 3 {
		fmt.Println("Not enough vowels")
	}
	if !has_double {
		fmt.Println("No double")
	}
	return has_double && (vowel_count >= 3)
}

func test2(input string) bool {
	has_double := false
	has_repeat := false

	last := '-'
	second_last := '-'
	doubles := make(map[string]int)

	for index, c := range(input) {
		two := string(last) + string(c)
		if (doubles[two] > 0) && (doubles[two] < index-1) {
			has_double = true
		} else if doubles[two] == 0 {
			doubles[two] = index 
		}
		if second_last == c {
			has_repeat = true
		}
		second_last = last
		last = c
	}
	if !has_repeat {
		fmt.Println("No repeat")
	}
	if !has_double {
		fmt.Println("No double")
	}
	return has_repeat && (has_double)
}
