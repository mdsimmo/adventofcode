package main

import "fmt"
import "os"

func aoc_1_run() {
	data, err := os.ReadFile("aoc_1_input.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)
	
	level := 0
	for index := 0; index < len(data); index++ {
		char := input[index:index+1]
		if char == "(" {
			level++
		} else if char == ")" {
			level--
		} else {
			fmt.Printf("unknown: %s\n", char)
		}
	}
	
	fmt.Printf("Result: %d", level)
}

