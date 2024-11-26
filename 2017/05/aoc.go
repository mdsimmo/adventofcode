package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part2 := true
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}

	jumps := []int {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		value, e := strconv.Atoi(line)
		if e != nil {
			panic(e)
		}
		jumps = append(jumps, value)
	}

	index := 0
	steps := 0
	for index >= 0 && index < len(jumps) {
		jump := jumps[index]
		if !part2 {
			jumps[index]++
		} else {
			if jump >= 3 {
				jumps[index]--
			} else {
				jumps[index]++
			}
		}
		index += jump
		steps++
		//fmt.Printf("[%d] %d => %d\n", steps, jump, index)
	}

	fmt.Printf("Steps: %d\n", steps)
}
