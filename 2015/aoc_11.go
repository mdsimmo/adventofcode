package main

import (
	"fmt"
	"strconv"
)

func aoc_11_run() {
	input := "cqjxjnds"
	fmt.Println(input)

	for !isvalid(input) {
		input = increment(input)
	}
	
	fmt.Println(input)
		
	input = increment(input)
	for !isvalid(input) {
		input = increment(input)
	}
	fmt.Println(input)
}

func increment(input string) string {

	// convert to base 26 string (a -> 0, a -> 'z'-10)
	altered := ""
	for _, c := range(input) {
		newc := c - 10
		if newc < 'a' {
			newc = c - 'a' + '0'
		}
		altered += string(newc)
	}
	
	// convert to int and increment
	num, err := strconv.ParseInt(altered, 26, 64)
	if err != nil {
		panic(err)
	}
	num++

	// convert back
	newaltered := strconv.FormatInt(num, 26)
	output := ""
	for _, c := range(newaltered) {
		newc := c + 10
		if newc < 'a' {
			newc = c + 'a' - '0'
		}
		output += string(newc)
	}
	
	return output
}

func isvalid(input string) bool {
	last := '-'
	last2 := '-'
	hasRun := false
	doubleCount := '-'
	doubleOk := false

	for _, ch := range(input) {
		if ch == last {
			if doubleCount != ch {
				if doubleCount == '-' {
					doubleCount = ch
				} else {
					doubleOk = true
				}
			}
		}
		if ch == 'i' || ch == 'o' || ch == 'l' {
			return false
		}
		if ch == (last + 1) && ch == (last2 + 2) {
			hasRun = true
		}
		last2 = last
		last = ch
	}
	return doubleOk && hasRun
}
