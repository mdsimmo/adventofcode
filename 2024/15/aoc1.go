package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	WALL  rune = '#'
	ROBOT rune = '@'
	BOX   rune = 'O'
	SPACE rune = '.'
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
	grid := [][]rune{}
	robotPos := Point{}
	scanner := bufio.NewScanner(file)
	// Read the grid
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		runes := []rune(line)
		for i, r := range runes {
			if r == ROBOT {
				robotPos.x = i
				robotPos.y = len(grid)
				runes[i] = SPACE
			}
		}
		grid = append(grid, runes)
	}

	seq := []Point{}
	// Read the sequence
	for scanner.Scan() {
		line := scanner.Text()
		for _, r := range line {
			var dir Point
			switch r {
			case '^':
				dir = Point{x: 0, y: -1}
			case 'v':
				dir = Point{x: 0, y: 1}
			case '<':
				dir = Point{x: -1, y: 0}
			case '>':
				dir = Point{x: 1, y: 0}
			default:
				panic("Unknown sequence character")
			}
			seq = append(seq, dir)
		}
	}

outer:
	for _, dir := range seq {
		// check for obstruction
		pNext := robotPos
		fmt.Printf("Process: %+v\n", dir)
	inner:
		for {
			pNext.x += dir.x
			pNext.y += dir.y
			switch grid[pNext.y][pNext.x] {
			case WALL:
				// Blocked, move to next sequece
				fmt.Printf("  Hit Wall\n")
				continue outer
			case BOX:
				// contine looping until space found
				fmt.Printf("  Found Box\n")
				continue inner
			case SPACE:
				break inner
			default:
				panic("Missed a case?")
			}
		}
		// Move boxes across
		for pNext != robotPos {
			pPre := pNext
			pPre.x -= dir.x
			pPre.y -= dir.y
			grid[pNext.y][pNext.x] = grid[pPre.y][pPre.x]
			fmt.Printf("  Moving Item: %+v->%+v\n", pNext, pPre)
			pNext = pPre
		}
		robotPos.x += dir.x
		robotPos.y += dir.y

		// printGrid(robotPos, grid)
	}

	sum := 0
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			switch grid[j][i] {
			case BOX:
				sum += i + j*100
			}
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

func printGrid(robot Point, grid [][]rune) {
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			if robot.x == i && robot.y == j {
				fmt.Printf("%c", ROBOT)
			} else {
				fmt.Printf("%c", grid[j][i])
			}
		}
		fmt.Println()
	}
}
