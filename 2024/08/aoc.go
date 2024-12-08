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
	grid := [][]rune{}
	lookup := map[rune][]complex128{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
		for x, r := range line {
			if r != '.' {
				if lookup[r] == nil {
					lookup[r] = []complex128{}
				}
				lookup[r] = append(lookup[r], complex(float64(x), float64(len(grid)-1)))
			}
		}
	}

	nodesPart1 := map[complex128]bool{}
	nodesPart2 := map[complex128]bool{}
	for _, set := range lookup {
		for i, p1 := range set {
			for j, p2 := range set {
				if i == j {
					continue
				}
				// only need one point, as the other i/j pairing will get the other point
				node := p1 - (p2 - p1)
				if real(node) < 0 || imag(node) < 0 ||
					real(node) >= float64(len(grid[0])) ||
					imag(node) >= float64(len(grid)) {
				} else {
					nodesPart1[node] = true
				}

				// Part 2
				nodesPart2[p1] = true
				diff := p2 - p1
				p := p1
				for {
					p -= diff
					if real(p) < 0 || imag(p) < 0 ||
						real(p) >= float64(len(grid[0])) ||
						imag(p) >= float64(len(grid)) {
						break
					}
					nodesPart2[p] = true
				}
			}
		}
	}

	fmt.Printf("Nodes Part 1: %d\n", len(nodesPart1))
	fmt.Printf("Nodes Part 2: %d\n", len(nodesPart2))

}
