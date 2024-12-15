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
	BOX_L rune = '['
	BOX_R rune = ']'
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
		runes := []rune{}
		for i, r := range line {
			switch r {
			case WALL:
				runes = append(runes, []rune{WALL, WALL}...)
			case ROBOT:
				runes = append(runes, []rune{SPACE, SPACE}...)
				robotPos.x = i * 2
				robotPos.y = len(grid)
			case BOX:
				runes = append(runes, []rune{BOX_L, BOX_R}...)
			case SPACE:
				runes = append(runes, []rune{SPACE, SPACE}...)
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

	// printGrid(robotPos, grid)
outer:
	for _, dir := range seq {
		// check for obstruction
		// fmt.Printf("Move: %+v\n", dir)
		pushList := map[Point]rune{}
		heads := []Point{robotPos}
		for len(heads) > 0 {
			head := heads[len(heads)-1]
			heads = heads[:len(heads)-1]
			if pushList[head] != 0 {
				// Already explored, skip
				continue
			}
			pushList[head] = grid[head.y][head.x]
			// fmt.Printf("  Head: %+v\n", head)

			head.x += dir.x
			head.y += dir.y
			switch grid[head.y][head.x] {
			case WALL:
				// Blocked, move to next sequece
				// fmt.Printf("  Hit Wall\n")
				continue outer
			case SPACE:
				// end the chain
				// fmt.Printf("  Space\n")
			case BOX_L:
				if dir.y != 0 {
					// Branch the path
					heads = append(heads, head)
					heads = append(heads, Point{x: head.x + 1, y: head.y})
					// fmt.Printf("  Added Box L: %+v\n", head)
				} else {
					heads = append(heads, head)
					// fmt.Printf("  Added Box -: %+v\n", head)
				}
			case BOX_R:
				if dir.y != 0 {
					// Branch the path
					heads = append(heads, head)
					heads = append(heads, Point{x: head.x - 1, y: head.y})
					// fmt.Printf("  Added Box R: %+v\n", head)
				} else {
					heads = append(heads, head)
					// fmt.Printf("  Added Box -: %+v\n", head)
				}
			default:
				panic("Missed a case?")
			}
		}
		// Delete the boxes
		for p := range pushList {
			grid[p.y][p.x] = SPACE
		}

		// Add the boxes back, offset by dir
		for pPre, tile := range pushList {
			pNext := Point{
				x: pPre.x + dir.x,
				y: pPre.y + dir.y,
			}
			grid[pNext.y][pNext.x] = tile
			// fmt.Printf("  Moving Item: %+v->%+v\n", pPre, pNext)
		}
		// move the robot
		robotPos.x += dir.x
		robotPos.y += dir.y

		// printGrid(robotPos, grid)
	}

	// Sum the boxes
	sum := 0
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			switch grid[j][i] {
			case BOX:
				fallthrough
			case BOX_L:
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
