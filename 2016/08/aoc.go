package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	lights := make([][]bool, 50)
	for i := 0; i < len(lights); i++ {
		lights[i] = make([]bool, 6)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Check: %s\n", line)
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' ' || r == 'x' || r == 'y' || r == '='
		})
		fmt.Printf("  Parts: %+v\n", parts)
		switch parts[0] {
		case "rect":
			w, _ := strconv.Atoi(parts[1])
			h, _ := strconv.Atoi(parts[2])
			for i := 0; i < w; i++ {
				for j := 0; j < h; j++ {
					lights[i][j] = true
				}
			}
		case "rotate":
			row := parts[1] == "row"
			sel, _ := strconv.Atoi(parts[2])
			by, _ := strconv.Atoi(parts[4])
			for j := 0; j < by; j++ {
				if row {
					last := lights[len(lights)-1][sel]
					for i := len(lights)-1; i > 0; i-- {
						lights[i][sel] = lights[i-1][sel]
					}
					lights[0][sel] = last
				} else {
					last := lights[sel][len(lights[sel])-1]
					for i := len(lights[sel])-1; i > 0; i-- {
						lights[sel][i] = lights[sel][i-1]
					}
					lights[sel][0] = last
				}
			}
		}
		print(lights)

	}

	sum := 0
	for i := 0; i < len(lights); i++ {
		for j := 0; j < len(lights[i]); j++ {
			if lights[i][j] {
				sum++
			}
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

func print(graph [][]bool) {
	for j := 0; j < len(graph[0]); j++ {
		for i := 0; i < len(graph); i++ {
			if graph[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
