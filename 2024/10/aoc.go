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
	grid := [][]int{}
	heads := map[complex128]bool{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, make([]int, len(line)))
		for i, r := range line {
			height := int(r - '0')
			grid[len(grid)-1][i] = height
			if height == 0 {
				heads[complex(float64(i), float64(len(grid)-1))] = true
			}
		}
	}

	dirs := []complex128{
		complex(0, -1),
		complex(0, 1),
		complex(1, 0),
		complex(-1, 0),
	}
	paths1 := 0
	paths2 := 0
	for head := range heads {
		search := []complex128{head}
		been := map[complex128]bool{}

		for len(search) > 0 {
			pos := search[0]
			search = search[1:]
			height := grid[int(imag(pos))][int(real(pos))]

			if height == 9 {
				paths2++
				if !been[pos] {
					paths1++
					been[pos] = true
				}
				continue
			}

			for _, dir := range dirs {
				p2 := pos + dir

				x := int(real(p2))
				y := int(imag(p2))
				if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[y]) {
					continue
				}
				h2 := grid[y][x]
				if h2 != height+1 {
					continue
				}

				search = append(search, p2)
			}
		}
	}

	fmt.Printf("Paths 1: %d\n", paths1)
	fmt.Printf("Paths 2: %d\n", paths2)
}
