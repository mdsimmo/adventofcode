package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}

	grid := [][]rune {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	// find the entry point
	pos := Point { x: -1, y : -1 }
	for i, r := range grid[0] {
		if r == '|' {
			pos = Point{ x: i, y: 0 }
			break
		}
	}
	if pos.x < 0 {
		panic("Did not find start")
	}
	dir := Point { x: 0, y: 1 }
	path := []rune {}

	steps := 0

	for true {
		r := grid[pos.y][pos.x]; 
		steps++
		switch r {
		case '+':
			// search for next path
			dir2 := Point { x: dir.y, y: dir.x }
			test := fetch(pos.x + dir2.x, pos.y + dir2.y, grid)
			if test == ' ' {
				dir2 = Point{ x: -dir2.x, y: -dir2.y }
				test2 := fetch(pos.x + dir2.x, pos.y + dir2.y, grid)
				if test2 == ' ' {
					panic("Turn was not on either side")
				}
			}
			dir = dir2
			fallthrough
		case '-':
			fallthrough
		case '|':
			pos.x += dir.x
			pos.y += dir.y
		case ' ':
			fmt.Printf("Path: %s\n", string(path))
			fmt.Printf("Steps: %d\n", steps-1) // minus one because last step does not count
			return
		default: 
			path = append(path, r)
			pos.x += dir.x
			pos.y += dir.y
		}
	}
}

func fetch(x, y int, grid [][]rune) rune {
	if x < 0 || y < 0 || y > len(grid) || x > len(grid[y]) {
		return ' '
	} else {
		return grid[y][x]
	}
}
