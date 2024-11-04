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
	valids := 0
	scanner := bufio.NewScanner(file)
	index := 0;
	triangs := make([][]int, 3)
	for i := 0; i < len(triangs); i++ {
		triangs[i] = make([]int, 3)
	}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.FieldsFunc(line, splitter)
		triangs[0][index], _ = strconv.Atoi(parts[0])
		triangs[1][index], _ = strconv.Atoi(parts[1])
		triangs[2][index], _ = strconv.Atoi(parts[2])

		index++
		if index == 3 {
			for i := 0; i < index; i++ {
				a := triangs[i][0]
				b := triangs[i][1]
				c := triangs[i][2]
				fmt.Printf("Testing: %3d %3d %3d\n", a, b, c)
				if a + b > c && a + c > b && b + c > a {
					valids++
					fmt.Printf("  Valid\n")
				} else {
					fmt.Printf("  Invalid\n")
				}
			}
			index = 0
		}
	}
	
	fmt.Printf("Valids: %d\n", valids)
}

func splitter(r rune) bool {
	return r == ' '
}
