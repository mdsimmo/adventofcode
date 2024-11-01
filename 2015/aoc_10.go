package main

import (
	"fmt"
	"strconv"
	"strings"
)

func aoc_10_run() {
	input := "1113122113"
	

	for i := 0; i < 51; i++ {
		fmt.Printf("%d: %d\n", i, len(input))
		input = iterate(input)
	}

}

func iterate(input string) string {
	
	var output strings.Builder
	output.Grow(len(input))
	last := '-'
	repeats := 0
	for _, char := range(input) {
		if char == last {
			repeats++
		} else {
			if repeats > 0 {
				output.WriteString(strconv.Itoa(repeats))
				output.WriteRune(last)
			}
			repeats = 1
			last = char
		}
	}
	output.WriteString(strconv.Itoa(repeats))
	output.WriteRune(last)
	return output.String()
}
