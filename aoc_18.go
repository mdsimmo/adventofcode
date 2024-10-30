package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func aoc_18_run() {
	file, err := os.Open("aoc_18_in.txt")
	if err != nil {
		panic(err)
	}
	grid := make([][]bool, 100)
	scanner := bufio.NewScanner(file)
	y := 0;
	for scanner.Scan() {
		line := scanner.Text()
		grid[y] = make([]bool, 100)
		for i, c := range(line) {
			switch c {
			case '#':
				grid[y][i] = true
			case '.':
				grid[y][i] = false
			default:
				panic("Don't know what that was")
			}
		}
		y++
	}

	for t := 0; t < 100; t++ {
		grid[0][0] = true
		grid[0][99] = true
		grid[99][0] = true
		grid[99][99] = true

		gridnew := make([][]bool, 100)
		for i := 0; i < 100; i++ {
			gridnew[i] = slices.Clone(grid[i])
		}

		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				neighs := 0
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						x := i + dx;
						y := j + dy;
						if x < 0 || x >= 100 || y < 0 || y >= 100 || (x==i && y==j) {
							// do nothing
						} else if grid[x][y] {
							neighs++
						}
					}
				}
				state := grid[i][j]
				var newstate bool
				if neighs == 3 {
					newstate = true
				} else if neighs == 2 {
					newstate = state
				} else {
					newstate = false
				}
				gridnew[i][j] = newstate
			}
		}

		grid = gridnew
	}

	grid[0][0] = true
	grid[0][99] = true
	grid[99][0] = true
	grid[99][99] = true

	numon := 0
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if grid[i][j] {
				numon++
			}
		}
	}

	fmt.Printf("Num lights on: %d", numon)
}

