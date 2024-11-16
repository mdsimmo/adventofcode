package main

import (
	"fmt"
	"math"
)

func main() {
	num := 3012210
	fmt.Printf("Part 1: %d\n", calc(num, false))
	fmt.Printf("Part 2: %d\n", calc(num, true))

	// Simulating the solution takes way too long for large numbers
	// I have no idea why this formula works:
	// It is just the pattern I noticed by looking at the first few loops
	// for i := 1; i < 100; i++ {
	//	run(i, false)
	//}
}

func calc(num int, part2 bool) int {
	
	if part2 {
		// part 2 formula
		base_3 := math.Log2(float64(num)) / math.Log2(3)
		last_pow_3 := int(math.Pow(3, float64(int(base_3+0.0001))))
		if num == last_pow_3 {
			return num
		}
		if num - last_pow_3 <= last_pow_3 {
			return num - last_pow_3
		} else {
			return last_pow_3 + (num - last_pow_3*2) * 2
		}
	} else {
		// part 1 formula
		base_2 := math.Log2(float64(num))
		last_pow_2 := 1 << int(base_2)
		return (num - last_pow_2) * 2 + 1
	}
}

func run(num int, part2 bool) {
	elves := make([]int, num)
	for i := 0; i < len(elves); i++ {
		elves[i] = i+1
	}
	
	for i := 0; ; i++ {
		i = i % len(elves)
		j := 0
		if part2 {
			// part 2
			j = (i + len(elves)/2) % len(elves)
		} else {
			// part 1
			j = (i + 1) % len(elves)
		}
		//fmt.Printf("Elves: %+v\n", elves)
		//fmt.Printf("Elf %d steels %d\n", elves[i], elves[j])
		//fmt.Printf("[%d] %d\n", num, len(elves))
		if i == j {
			fmt.Printf("[%d] Winning elf: %d - calc: %d\n", num, elves[i], calc(num, part2))
			return
		}
		for k := j; k < len(elves)-1; k++ {
			elves[k] = elves[k+1]
		}
		elves = elves[:len(elves)-1]
		//elves = append(elves[:j], elves[j+1:]...)
		
		// correct index from the deletion
		if j < i {
			i--
		}
	}
}

