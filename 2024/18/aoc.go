package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func (p Point) Add(p2 Point) Point {
	return Point{
		x: p.x + p2.x,
		y: p.y + p2.y,
	}
}

func (p Point) Sub(p2 Point) Point {
	return Point{
		x: p.x - p2.x,
		y: p.y - p2.y,
	}
}

func (p Point) Mult(s int) Point {
	return Point{
		x: p.x * s,
		y: p.y * s,
	}
}

func (p Point) RotateLeft() Point {
	return p.MultComplex(Point{0, 1})
}

func (p Point) RotateRight() Point {
	return p.MultComplex(Point{0, -1})
}

func (p Point) MultComplex(p2 Point) Point {
	// (a+bi) * (c+di)
	// = ac + adi + cbi - bd
	// = ac - bd + i(ad + cb)
	return Point{
		x: (p.x * p2.x) - (p.y * p2.y),
		y: (p.x*p2.y + p.y*p2.x),
	}
}

func (p Point) Adj() []Point {
	list := make([]Point, 4)
	for i, dir := range []Point{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
		list[i] = p.Add(dir)
	}
	return list
}

func (p Point) Diag() []Point {
	list := make([]Point, 8)
	for i, dir := range []Point{
		{0, 1}, {0, -1}, {-1, 0}, {1, 0},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	} {
		list[i] = p.Add(dir)
	}
	return list
}

func (p Point) Dist2(p2 Point) int {
	x := p2.x - p.x
	y := p2.y - p.y
	return x*x + y*y
}

func (p Point) DistGrid(p2 Point) int {
	x := p2.x - p.x
	y := p2.y - p.y
	return Abs(x) + Abs(y)
}

func Min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func Abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}

const (
	width  int = 70
	height int = 70
)

func main() {
	// Read all the points
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	points := []Point{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := ExtractInts(scanner.Text())
		points = append(points, Point{nums[0], nums[1]})
	}

	// Binary search until we find the min/max pass/block point
	grid := map[Point]bool{}
	bestPath, _ := findPath(grid)
	iMin := 0
	iMax := len(points)
	iPre := 0
	pathFinds := 1
	for i := (iMax + iMin) / 2; iMin < iMax-1; i = (iMin + iMax) / 2 {
		// update the grid with the changed points from the last iteration
		if iPre > i {
			for j := i; j < iPre; j++ {
				grid[points[j]] = false
			}
			bestPath = map[Point]bool{}
		} else if iPre < i {
			// When adding points, optimise by checking if any
			// of the new points have blocked the previous path.
			// If not, just keep adding points until we do block.
			allOk := true
			for j := iPre; j < i || allOk; j++ {
				p := points[j]
				grid[p] = true
				if _, blocked := bestPath[p]; blocked {
					allOk = false
				}
				i = Max(i, j+1)
			}
		} else {
			panic("Should not happen")
		}

		newPath, passable := findPath(grid)
		bestPath = newPath
		pathFinds++

		if !passable {
			// make easier
			iMax = i
		} else {
			// make harder
			iMin = i
		}
		iPre = i
	}
	blocker := points[iMin]
	fmt.Printf("Blocking: %d,%d\n", blocker.x, blocker.y)
	fmt.Printf("Path find runs: %d\n", pathFinds)
}

// Does A* search for the exit
func findPath(grid map[Point]bool) (map[Point]bool, bool) {
	start := Point{
		x: 0,
		y: 0,
	}
	end := Point{
		x: width,
		y: height,
	}

	type Entry struct {
		cost int
		path map[Point]bool
		est  int
		loc  Point
	}
	search := []Entry{{
		loc:  start,
		path: map[Point]bool{},
	}}
	explored := map[Point]bool{}
	for len(search) > 0 {
		best := search[len(search)-1]
		search = search[:len(search)-1]
		// fmt.Printf("Test: %+v\n", best)

		if best.loc == end {
			// fmt.Printf("Best: %d\n", best.len)
			return best.path, true
		}
		if explored[best.loc] {
			continue
		}
		explored[best.loc] = true

		for _, pNew := range best.loc.Adj() {
			// fmt.Printf("  Branch: %+v\n", pNew)
			if grid[pNew] {
				// fmt.Printf("  Wall\n")
				continue
			}
			if pNew.x < 0 || pNew.y < 0 || pNew.x > width || pNew.y > height {
				// fmt.Printf("  OOB\n")
				continue
			}
			if explored[pNew] {
				// fmt.Printf("  Explored\n")
				continue
			}

			costNew := best.cost + 1
			estNew := costNew + pNew.DistGrid(end)
			newPath := maps.Clone(best.path)
			newPath[pNew] = true
			search = append(search, Entry{
				cost: costNew,
				est:  estNew,
				path: newPath,
				loc:  pNew,
			})
		}

		slices.SortFunc(search, func(a, b Entry) int {
			return b.est - a.est
		})
	}

	return nil, false
}

func ExtractInts(line string) []int {
	parts := strings.FieldsFunc(line, func(r rune) bool {
		return r < '0' || r > '9'
	})
	nums := make([]int, len(parts))
	for i, p := range parts {
		var e error
		nums[i], e = strconv.Atoi(p)
		if e != nil {
			panic(e)
		}
	}
	return nums
}
