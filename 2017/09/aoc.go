package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("in.txt")
	if err != nil {
		panic(err)
	}
	contents := string(file)

	sum := 0
	depth := 0
	junk := false
	cancelled := 0
	for i := 0; i < len(contents); i++ {
		switch contents[i] {
		case '{':
			if !junk {
				depth++
			} else {
				cancelled++
			}
		case '}':
			if !junk {
				sum += depth
				depth--
			} else {
				cancelled++
			}
		case '<':
			if junk {
				cancelled++
			} else {
				junk = true
			}
		case '>':
			junk = false
		case '!':
			if junk {
				i++
			}
		default: 
			if junk {
				cancelled++
			}
		}
	}

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Cancelled: %d\n", cancelled)
}
