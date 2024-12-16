package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
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
	var start Point
	var end Point
	grid := [][]bool{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, make([]bool, len(line)))
		for i, r := range line {
			switch r {
			case '.':
			case '#':
				grid[len(grid)-1][i] = true
			case 'S':
				start = Point{x: i, y: len(grid) - 1}
			case 'E':
				end = Point{x: i, y: len(grid) - 1}
			default:
				panic("Unknown char")
			}
		}
	}

	// A* search through the grid
	type Entry struct {
		cost      int
		estimated int
		loc       Point
		dir       Point
		path      map[Point]bool
	}
	search := []Entry{{
		cost:      0,
		estimated: 0,
		loc:       start,
		dir:       Point{x: 1, y: 0},
		path:      map[Point]bool{start: true},
	}}
	searched := map[Point]map[Point]int{}
	seats := map[Point]bool{}
	for len(search) > 0 {
		best := search[len(search)-1]
		search = search[:len(search)-1]
		// fmt.Printf("Test: %+v\n", best)

		// Ignore paths that are known to not be the fastest
		if searched[best.loc] != nil {
			prevCost := searched[best.loc][best.dir]
			if prevCost != 0 && prevCost < best.cost {
				// fmt.Printf("  Discard %d\n", prevCost)
				continue
			}
		} else {
			searched[best.loc] = make(map[Point]int)
		}
		searched[best.loc][best.dir] = best.cost

		// Check if the path is the end
		if best.loc == end {
			for p := range best.path {
				seats[p] = true
			}
		}

		// find available other paths
		dirs := []Point{{x: 0, y: 1}, {x: 0, y: -1}, {x: -1, y: 0}, {x: 1, y: 0}}
		for _, d := range dirs {
			next := Point{
				x: best.loc.x + d.x,
				y: best.loc.y + d.y,
			}
			if next.x < 0 || next.y < 0 || next.x >= len(grid[0]) || next.y > len(grid) {
				continue
			}
			if grid[next.y][next.x] {
				continue
			}
			// never do u-turn, except at very start
			if d.x == best.dir.x*-1 && d.y == best.dir.y*-1 {
				continue
			}

			// Calculate new costs
			newPath := maps.Clone(best.path)
			newPath[next] = true
			stepCost := 1
			if d != best.dir {
				stepCost += 1000
			}
			newCost := best.cost + stepCost
			newEst := newCost + abs(end.x-next.x) + abs(end.y-next.y)

			newTest := Entry{
				cost:      newCost,
				estimated: newEst,
				loc:       next,
				dir:       d,
				path:      newPath,
			}
			search = append(search, newTest)
			// fmt.Printf("  Append: %+v\n", newTest)
		}
		slices.SortFunc(search, func(a, b Entry) int {
			return b.estimated - a.estimated
		})
	}
	for dir := range searched[end] {
		fmt.Printf("Score: %d\n", searched[end][dir])
		break
	}
	fmt.Printf("Seats: %d\n", len(seats))
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
