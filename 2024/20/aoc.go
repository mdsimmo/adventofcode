package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const skips int = 20
	const minDiff int = 100
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	var start Point
	var end Point
	grid := map[Point]bool{}
	scanner := bufio.NewScanner(file)
	height := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, r := range line {
			switch r {
			case '.':
			case '#':
				grid[Point{i, height}] = true
			case 'S':
				start = Point{x: i, y: height}
			case 'E':
				end = Point{x: i, y: height}
			default:
				panic("Unknown char")
			}
		}
		height++
	}

	// Build a distance map
	path := map[Point]int{
		start: 0,
	}
	for p := start; ; {
		found := false
		for _, next := range p.Adj() {
			if grid[next] {
				continue
			}
			if _, exists := path[next]; exists {
				continue
			}
			if found {
				panic("Two branches found")
			}
			p = next
			path[next] = len(path)
			found = true
		}
		if !found {
			panic("Deadend found")
		}
		if p == end {
			break
		}
	}
	// fmt.Printf("Orig Path Length: %d\n", len(path))

	// find all cheat paths within the specified distance
	cheatPaths := 0
	for cheatStart := range path {
		for dx := -skips; dx <= skips; dx++ {
			for dy := -(skips - Abs(dx)); dy <= skips-Abs(dx); dy++ {
				cheatEnd := Point{cheatStart.x + dx, cheatStart.y + dy}
				if grid[cheatEnd] {
					continue
				}
				if _, exists := path[cheatEnd]; !exists {
					continue
				}
				diff := path[cheatStart] - path[cheatEnd] - Abs(dx) - Abs(dy)
				if diff >= minDiff {
					// fmt.Printf("Cheat: %+v -> %+v = %d\n", p, next, diff)
					cheatPaths++
				}
			}
		}
	}

	fmt.Printf("Cheat Paths: %d\n", cheatPaths)
}
