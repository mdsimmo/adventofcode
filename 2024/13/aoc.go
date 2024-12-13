package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

const cost_a int = 3
const cost_b int = 1

func main() {
	part2 := true
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums_a := extractNums(scanner.Text())
		scanner.Scan()
		nums_b := extractNums(scanner.Text())
		scanner.Scan()
		nums_p := extractNums(scanner.Text())
		scanner.Scan() // blank line

		a := numsToPoint(nums_a)
		b := numsToPoint(nums_b)
		p := numsToPoint(nums_p)

		if part2 {
			p.x += 10000000000000
			p.y += 10000000000000
		}

		cheap, _ := cheapest2(a, b, p)
		sum += cheap
	}
	fmt.Printf("Sum: %d\n", sum)
}

func cheapest2(a, b, p Point) (int, bool) {
	// solve linear equations
	num_a := p.x*b.y - p.y*b.x
	num_b := p.x*a.y - p.y*a.x
	den_a := (a.x * b.y) - (a.y * b.x)
	den_b := (a.y * b.x) - (a.x * b.y)
	if num_a%den_a == 0 {
		return (num_a/den_a)*cost_a + (num_b/den_b)*cost_b, true
	} else {
		return 0, false
	}
}

func numsToPoint(nums []int) Point {
	if len(nums) != 2 {
		panic("Length not correct")
	}
	return Point{
		x: nums[0],
		y: nums[1],
	}
}

func extractNums(line string) []int {
	parts := strings.FieldsFunc(line, func(r rune) bool {
		return r < '0' || r > '9'
	})
	nums := make([]int, len(parts))
	for i, p := range parts {
		var err error
		nums[i], err = strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
	}
	return nums
}

// The rest of this is an implemenation of A* search
// It is a REALLY DUMB way to to this puzzle

func estimate(start, a, b, p Point) int {
	cost_just_a := max(div_round_up(p.x-start.x, a.x)*cost_a, div_round_up(p.y-start.y, a.y)*cost_a)
	cost_just_b := max(div_round_up(p.x-start.x, b.x)*cost_a, div_round_up(p.y-start.y, b.y)*cost_b)
	best_cost := min(cost_just_a, cost_just_b)
	return best_cost
}

func cheapest(a, b, p Point) (int, bool) {
	type Entry struct {
		cost int
		est  int
		loc  Point
		hasB bool
	}

	search := []Entry{{
		cost: 0,
		est:  0,
		loc:  Point{x: 0, y: 0},
		hasB: false,
	}}

	for len(search) > 0 {
		test := search[0]
		search = search[1:]

		if test.loc.x == p.x && test.loc.y == p.y {
			return test.cost, true
		}

		if test.loc.x > p.x || test.loc.y > p.y {
			continue
		}

		if !test.hasB {
			aPoint := Point{
				x: test.loc.x + a.x,
				y: test.loc.y + a.y,
			}
			aNext := Entry{
				cost: test.cost + cost_a,
				est:  test.cost + cost_a + estimate(aPoint, a, b, p),
				loc:  aPoint,
				hasB: false,
			}
			search = append(search, aNext)
		}

		bPoint := Point{
			x: test.loc.x + b.x,
			y: test.loc.y + b.y,
		}
		bNext := Entry{
			cost: test.cost + cost_b,
			est:  test.cost + cost_b + estimate(bPoint, b, b, p),
			loc:  bPoint,
			hasB: true,
		}
		search = append(search, bNext)

		slices.SortFunc(search, func(a, b Entry) int {
			return a.est - b.est
		})
	}
	return 0, false
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func div_round_up(a, b int) int {
	return (a + (b - 1)) / b
}
