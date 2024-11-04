package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("aoc_1_in.txt")
	if err != nil {
		panic(err)
	}
	text := string(file)
	cmds := strings.FieldsFunc(text, splitter)

	visited := make(map[complex128]bool)
	pos := complex(0, 0)
	dir := complex(0, -1)
	visited[pos] = true
	for _, cmd := range(cmds) {
		dir_c := cmd[0]
		if dir_c == 'L' {
			dir *= complex(0, 1)
		} else {
			dir *= complex(0, -1)
		}

		mag, err := strconv.Atoi(strings.Trim(cmd[1:], "\n\r "))
		if err != nil {
			panic(err)
		}
		for i := 0; i < mag; i++ {
			pos += complex(1, 0) * dir
			if visited[pos] {
				fmt.Printf("Double Visit: %+v (%f)\n", pos, math.Abs(real(pos))+math.Abs(imag(pos)))
			}
			visited[pos] = true
		}
		

		visited[pos] = true

		
	}
	fmt.Printf("Dist: %f", math.Abs(real(pos))+math.Abs(imag(pos)))
	
}

func splitter(c rune) bool {
	return c == ' ' || c == ','
}
