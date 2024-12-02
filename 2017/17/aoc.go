package main

import (
	"fmt"
)

func main() {
	part1(2017)
	part2()
}

func part2() {
	steps := 337
	position := 0
	pos1Val := 0
	for bufferSize := 1; bufferSize <= 50_000_000; bufferSize++ {
		position = (position + steps) % bufferSize + 1
		if position == 1 {
			pos1Val = bufferSize
		}
		if position == 0 {
			panic("Should never happen")
		}
	}
	fmt.Printf("Zero Pos: %d\n", pos1Val)
}

func part1(iters int) {
	steps := 337
	buffer := []int { 0 }
	position := 0
	for i := 1; i <= iters; i++ {
		position = (position + steps) % len(buffer) + 1
		buffer = append(buffer, 0)
		for j := len(buffer)-1; j > position; j-- {
			buffer[j] = buffer[j-1]
		}
		buffer[position] = i
	}
	fmt.Printf("Last position: %d\n", buffer[(position+1)%len(buffer)])
}
