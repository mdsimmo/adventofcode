package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main1() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	grid := map[complex128]bool {}
	scanner := bufio.NewScanner(file)
	in_y := 0
	for scanner.Scan() {
		line := scanner.Text()
		offset := len(line)/2
		for i, r := range line {
			grid[complex(float64(i-offset), float64(in_y-offset))] = r == '#'
		}
		in_y++
	}


	infect := 0
	pos := complex(0, 0)
	dir := complex(0, -1)
	
	printGrid(grid, pos)
	
	for i := 0; i < 10000; i++ {
		if grid[pos] {
			dir *= complex(0, 1)
		} else {
			dir *= complex(0, -1)
			infect++
		}
		grid[pos] = !grid[pos]
		pos += dir

		//printGrid(grid, pos)
	}
	fmt.Printf("Infected: %d\n", infect)
}

func printGrid(in map[complex128]bool, pos complex128) {
	size := 0
	for v := range in {
		dist := int(math.Abs(real(v)) + math.Abs(imag(v)))
		if dist > size {
			size = dist
		}
	}

	for j := -size-1; j < size+1; j++ {
		for i := -size-1; i < size+1; i++ {
			p := complex(float64(i), float64(j))
			if p == pos {
				fmt.Print("@")
			} else if in[p] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
