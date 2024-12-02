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
	bufferSize := 1
	position := 0
	zeroPos1 := 0
	for i := 1; i <= 5_000_000; i++ {
		position = (position + steps) % bufferSize + 1
		bufferSize++
		if position == 1 {
			zeroPos1 = i
		}
		if position == 0 {
			panic("Should never happen")
		}
		//check := part1(i)
		//if check != zeroPos1 {
		//	panic("Mismatch")
		//}
		//fmt.Printf("Zero: %d\n", zeroPos1)
	}
	fmt.Printf("Zero Pos: %d\n", zeroPos1)
}

func part1(iters int) int {
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
		//fmt.Printf("[%4d] %4d: %+v\n", i, position, buffer)
	}
	//fmt.Printf("Last position: %d\n", buffer[(position+1)%len(buffer)])
	if buffer[0] != 0 {
		panic("Buffer 0 was not zero")
	}
	return buffer[1]
}
