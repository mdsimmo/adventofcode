package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func main() {
	dest := Point {
		x: -1, //31,
		y: -1, //39,
	}
	input := 1350
	maxSteps := 50 // set to big number for part 1

	for y := 0; y < dest.y; y++ {
		for x := 0; x < dest.x + 5; x++ {
			if x == dest.x && y == dest.y {
				fmt.Print("X")
			} else if isWall(Point{x, y}, input) {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}

	searched := map[Point]bool {}
	toSearch := map[int]map[Point]int {
		0:  { 
			Point {
				x: 1,
				y: 1,
			}: 0},
	}

	for len(toSearch) > 0 {
		// Find best guess to take next
		bestScore := math.MaxInt
		for score := range(toSearch) {
			if score < bestScore {
				bestScore = score
			}
		}
		scanMap := toSearch[bestScore]
		var point Point
		var steps int
		for p, s := range(scanMap) {
			point = p
			steps = s
			delete(scanMap, p)
			if len(scanMap) == 0 {
				delete(toSearch, bestScore)
			}
			break
		}
		fmt.Printf("Scaning %d - %+v\n", steps, point)
		if searched[point] == true {
			fmt.Printf("  Skipped (already searched)\n")
			continue
		}
		searched[point] = true

		// Generate possible steps
		for _, delta := range([][]int {
				{ 0, 1 },
				{ 1, 0 },
				{ -1, 0 },
				{ 0, -1 },
			}) {
			next := Point {
				x: point.x + delta[0],
				y: point.y + delta[1],
			}
			fmt.Printf("  Check %+v\n", next)
			if searched[next] {
				fmt.Printf("    Discard (already searched)\n")
				continue
			}
			if isWall(next, input) {
				fmt.Printf("    Discard (wall)\n")
				continue
			}
			dist := abs(next.x - dest.x) + abs(next.y - dest.y)
			if dist == 0 {
				fmt.Printf("    SOLUTION FOUND: %d", steps+1)

				return
			}

			// Part 2
			if steps + 1 > maxSteps {
				fmt.Printf("    Discard (overstepped)\n")
				continue
			}

			score := dist + steps
			if toSearch[score] == nil {
				toSearch[score] = map[Point]int{next: steps + 1}
			} else {
				if toSearch[score][next] == 0 {
					toSearch[score][next] = steps + 1
				}
			}
			fmt.Printf("    Added (%d)\n", score)
		}
	}
	fmt.Printf("Reachable: %d\n", len(searched))
	//panic("Should never happen")
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func isWall(point Point, input int) bool {
//	fmt.Printf("Checking: %+v\n", point)
	x := point.x
	y := point.y

	if x < 0 || y < 0 {
		return true
	}

	a := x*x + 3*x + 2*x*y + y + y*y + input
//	fmt.Printf("  a: %d %x\n", a, a)
	sum := 0
	for a != 0 {
		if a % 2 != 0 {
			sum++
		}
		a = a >> 1
	}
//	fmt.Printf("  Sum: %d\n", sum)
	return sum % 2 != 0
}
