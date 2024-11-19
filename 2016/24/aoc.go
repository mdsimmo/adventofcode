package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

type Point struct {
	x int
	y int
}

type Entry struct {
	done []Point
	todo []Point
	pos Point
	cost int
}

func main() {

	part2 := true

	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	grid := [][]bool {}
	targets := []Point {}
	var start Point
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]bool, len(line))
		for i, r := range(line) {
			switch r {
			case '.':
				row[i] = false
				//fmt.Printf("%c", '.')
			case '#':
				row[i] = true
				//fmt.Printf("%c", '#')
			case '0':
				row[i] = false
				start = Point{
					x: i,
					y: len(grid),
				}
				targets = append(targets, start)
				//fmt.Printf("%c", '0')
			default:
				row[i] = false
				targets = append(targets, Point {
					x: i,
					y: len(grid),
				})
				//fmt.Printf("%c", 'x')
			}
		}
		grid = append(grid, row)
		//fmt.Println()
	}

	// Make distance map
	dists := make(map[Point]map[Point]int)
	for _, t1 := range(targets) {
		dists[t1] = distances(t1, grid, targets)
	}
	//fmt.Printf("Distance Map: %+v\n", dists)


	// Find best route
	bestEntry := Entry {
		cost: math.MaxInt,
	}
	search := []Entry {{
		done: []Point { start },
		todo: targets,
		pos: start,
		cost: 0,
	}}
	for len(search) > 0 {
		e := search[len(search)-1]
		search = search[:len(search)-1]
		
		//fmt.Printf("Check: %+v\n", e)

		if len(e.todo) == 0 {
			panic("Should never have zero todo here")
		}

		if e.cost >= bestEntry.cost {
			//fmt.Printf("  Discard\n")
			continue
		}

		for i, t := range(e.todo) {
			//fmt.Printf("  Check %+v\n", t)	
			newDone := slices.Clone(e.done)
			newDone = append(newDone, t)
			newTodo := slices.Clone(e.todo)
			newTodo = append(newTodo[:i], newTodo[i+1:]...)
			newCost := e.cost + dists[e.pos][t]

			finalCost := newCost
			if part2 {
				finalCost += dists[t][start]
			}

			newE := Entry {
				done: newDone,
				todo: newTodo,
				pos: t,
				cost: newCost,
			}

			if len(newE.todo) == 0 {
				if bestEntry.cost > finalCost {
					bestEntry = newE
					bestEntry.cost = finalCost
					//fmt.Printf("    New best: %+v\n", newE)
				} else {
					//fmt.Printf("    End Discard\n")
				}
			} else {
				if finalCost < bestEntry.cost {
					//fmt.Printf("    Added to todo: %+v\n", newE)
					search = append(search, newE)
				} else {
					//fmt.Printf("    Discard\n")
				}
			}
		}
	}
	
	fmt.Printf("Best: %d: %+v\n", bestEntry.cost, bestEntry.done)
}

func distances(a Point, grid [][]bool, targets []Point) map[Point]int {
	costGrid := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		costGrid[i] = make([]int, len(grid[i]))
	}
	toSearch := []Point { a }
	costGrid[a.y][a.x] = 1 // 1 to differentiate from 0
	costs := map[Point]int {
		a: 0,
	}
	//fmt.Printf("targets: %+v\n", targets)
	dirs := [][]int {{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(targets) > len(costs) {
		p := toSearch[0]
		toSearch = toSearch[1:]

		//fmt.Printf("Check: %+v\n", p)

		for _, dir := range(dirs) {
			np := Point {
				x: p.x + dir[0],
				y: p.y + dir[1],
			}
			//fmt.Printf("  Step to: %+v\n", np)
			if grid[np.y][np.x] {
				//fmt.Printf("    Discard (wall)\n")
				continue
			}
			if costGrid[np.y][np.x] != 0 {
				//fmt.Printf("    Discard (visited)\n")
				continue
			}
			newCost := costGrid[p.y][p.x] + 1
			costGrid[np.y][np.x] = newCost
			//fmt.Printf("    Set Cost: %d\n", newCost)
			toSearch = append(toSearch, np)

			for _, t := range(targets) {
				if t.x == np.x && t.y == np.y {
					costs[t] = newCost - 1 // minus one because we start at 1
					//fmt.Printf("    Target Found: %+v\n", t)
				}
			}
		}
	}
	return costs
}
