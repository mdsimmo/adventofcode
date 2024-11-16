package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	ranges := [][]uint32 {}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		min, _ := strconv.ParseInt(parts[0], 10, 0)
		max, _ := strconv.ParseInt(parts[1], 10, 0)

		ranges = append(ranges, []uint32 {uint32(min), uint32(max)})
	}

	var sum uint32 = 0
	index, anyLeft := findNextAlowed(0, ranges)
	fmt.Printf("First IP: %d\n", index) 
	for anyLeft {
		disallow, anyLeft := findNextDisalowed(index, ranges)	
		if !anyLeft {
			sum += math.MaxUint32 - index
			break
		}
		sum += disallow - index

		index, anyLeft = findNextAlowed(disallow, ranges)
	}

	fmt.Printf("Available: %d\n", sum)
}

func findNextAlowed(index uint32, ranges [][]uint32) (uint32, bool) {
	best := uint32(math.MaxUint32)
	Outer:
	for _, r1 := range(ranges) {
		if r1[1]+1 >= best {
			continue
		}
		if r1[1]+1 <= index {
			continue
		}
		for _, r2 := range(ranges) {
			if r1[1]+1 >= r2[0] && r1[1]+1 <= r2[1] {
				continue Outer
			}
		}
		best = r1[1]+1
	}
	return best, best != math.MaxUint32
}

func findNextDisalowed(index uint32, ranges [][]uint32) (uint32, bool) {
	best := uint32(math.MaxUint32)
	Outer:
	for _, r1 := range(ranges) {
		if r1[0] >= best {
			continue
		}
		if r1[0] <= index {
			continue
		}
		for _, r2 := range(ranges) {
			if r1[0] > r2[0] && r1[0] <= r2[1] {
				continue Outer
			}
		}
		best = r1[0]
	}
	return best, best != math.MaxUint32
}
