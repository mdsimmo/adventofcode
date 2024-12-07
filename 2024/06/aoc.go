package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	pos := complex(0, 0)
	grid := [][]bool {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, make([]bool, len(line)))
		for i, r := range line {
			grid[len(grid)-1][i] = r == '#'
			if r == '^' {
				pos = complex(float64(i), float64(len(grid)-1))
			}
		}
	}
	
	// part 1
	steps, _ := mapRoute(grid, pos)
	fmt.Printf("Exit Steps: %d\n", len(steps))
	
	loopCount := 0
	for y, line := range grid {
		for x, wall := range line {
			if wall {
				continue
			}
			p := complex(float64(x), float64(y))
			if p == pos {
				continue
			}
			grid[y][x] = true
			_, loop := mapRoute(grid, pos)
			grid[y][x] = false
			if loop {
				loopCount++
			}
		}
	}
	fmt.Printf("Loop Positions: %d\n", loopCount)
}

func mapRoute(grid [][]bool, pos complex128) (map[complex128]bool, bool) {
	type Been struct {
		pos complex128
		dir complex128
	}
	visited := map[complex128]bool { pos: true }
	been := map[Been]bool {}

	dir := complex(0, -1)
	for !been[Been{ pos, dir }] {
		posNext := pos + dir
		if real(posNext) < 0 || imag(posNext) < 0 || 
				int(real(posNext)) >= len(grid[0]) ||
				int(imag(posNext)) >= len(grid) {
			return visited, false
		}
		if grid[int(imag(posNext))][int(real(posNext))] {
			been[Been{ pos, dir }] = true
			dir *= complex(0, 1)
		} else {
			been[Been{ pos, dir }] = true
			pos = posNext
			visited[pos] = true

		}
	}
	return visited, true
}
