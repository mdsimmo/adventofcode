package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	CLEAN = '.'
	WEAK = 'W'
	IFEC = '#'
	FLAG = 'F'
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	grid := map[complex128]rune {}
	scanner := bufio.NewScanner(file)
	in_y := 0
	for scanner.Scan() {
		line := scanner.Text()
		offset := len(line)/2
		for i, r := range line {
			grid[complex(float64(i-offset), float64(in_y-offset))] = r
		}
		in_y++
	}


	infect := 0
	pos := complex(0, 0)
	dir := complex(0, -1)
	
	printGrid2(grid, pos)
	
	for i := 0; i < 10000000; i++ {
		switch grid[pos] {
		case 0:
			fallthrough
		case CLEAN:
			dir *= complex(0, -1)
			grid[pos] = WEAK
		case WEAK:
			// continue on same dir
			grid[pos] = IFEC
			infect++
		case IFEC:
			dir *= complex(0, 1)
			grid[pos] = FLAG
		case FLAG:
			dir = -dir
			grid[pos] = CLEAN
		default:
			panic(fmt.Sprintf("Unknown type %c", grid[pos]))
		}
		pos += dir

		//printGrid2(grid, pos)
	}
	printGrid2(grid, pos)

	fmt.Printf("Infected: %d\n", infect)
}

func printGrid2(in map[complex128]rune, pos complex128) {
	size := 0
	for v := range in {
		dist := int(math.Max(math.Abs(real(v)), math.Abs(imag(v))))
		if dist > size {
			size = dist
		}
	}

	for j := -size; j <= size; j++ {
		for i := -size; i <= size; i++ {
			p := complex(float64(i), float64(j))
			if p == pos {
				fmt.Print("@")
			} else if in[p] > 0 {
				fmt.Printf("%c", in[p])
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
