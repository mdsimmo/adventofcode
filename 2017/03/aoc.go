package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	in := 289326
	//in := 1024

	part1Done := false
	part2Done := false

	grid := map[Point]int {}

	dirs := []Point {
		{ x:  1, y:  0 },
		{ x:  0, y: -1 },
		{ x: -1, y:  0 },
		{ x:  0, y:  1 },
	}

	loc := Point {
		x: 0,
		y: 0,
	}
	dirIndex := -1
	grid[loc] = 1
	index := 1
	for {
		// check if we can turn
		nextDir := dirs[(dirIndex + 1) % len(dirs)]
		nextCorner := Point { x: loc.x + nextDir.x, y: loc.y + nextDir.y }
		if grid[nextCorner] == 0 {
			dirIndex = (dirIndex + 1) % len(dirs)
		}

		// Continue
		dir := dirs[dirIndex]
		loc = Point { x: loc.x + dir.x, y: loc.y + dir.y }
		sum := 0
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				sum += grid[Point{ x: loc.x + i, y: loc.y + j }]
			}
		}
		grid[loc] = sum
		index++

		//fmt.Printf("[%7d] (%4d, %4d) Dist: %4d, Val: %7d \n", index, loc.x, loc.y, abs(loc.x) + abs(loc.y), grid[loc])
		if grid[loc] > in && !part2Done {
			fmt.Printf("Part 2: %d\n", grid[loc])
			part2Done = true
		}
		
		if index == in {
			fmt.Printf("Part 1: %d\n", abs(loc.x) + abs(loc.y))
			part1Done = true
		}

		if part1Done && part2Done {
			return
		}
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
