package main

import (
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
