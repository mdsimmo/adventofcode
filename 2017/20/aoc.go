package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Vec struct {
	x int
	y int
	z int
}

type Pt struct {
	p Vec
	v Vec
	a Vec
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	pts := []Pt {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return (r < '0' || r > '9') && r != '-'
		})
		var pt Pt
		pt.p.x, _ = strconv.Atoi(parts[0])
		pt.p.y, _ = strconv.Atoi(parts[1])
		pt.p.z, _ = strconv.Atoi(parts[2])
		pt.v.x, _ = strconv.Atoi(parts[3])
		pt.v.y, _ = strconv.Atoi(parts[4])
		pt.v.z, _ = strconv.Atoi(parts[5])
		pt.a.x, _ = strconv.Atoi(parts[6])
		pt.a.y, _ = strconv.Atoi(parts[7])
		pt.a.z, _ = strconv.Atoi(parts[8])
		pts = append(pts, pt)
	}

	// part 1: find closest particle
	closest := math.MaxInt
	closestIndex := -1
	for i, pt := range pts {
		t := 10000
		pos := Vec {
			x: pt.a.x * t*t + pt.v.x * t + pt.p.x,
			y: pt.a.y * t*t + pt.v.y * t + pt.p.y,
			z: pt.a.z * t*t + pt.v.z * t + pt.p.z,
		}
		dist := Abs(pos.x) + Abs(pos.y) + Abs(pos.z)
		if dist == closest {
			panic("Two particals with same speed found")
		}
		if dist < closest {
			closest = dist
			closestIndex = i
		}
	}
	fmt.Printf("Slowest: %d\n", closestIndex)


	for t := 0; ; t++ {
		for i, pt := range pts {
			pt.v.x += pt.a.x
			pt.v.y += pt.a.y
			pt.v.z += pt.a.z
			pt.p.x += pt.v.x
			pt.p.y += pt.v.y
			pt.p.z += pt.v.z
			pts[i] = pt
		}
		
		//fmt.Printf("%+v\n", pts)

		for i := 0; i < len(pts); i++ {
			p1 := pts[i]
			collision := false
			for j := i+1; j < len(pts); j++ {
				p2 := pts[j]
				if p1.p.x == p2.p.x && p1.p.y == p2.p.y && p1.p.z == p2.p.z {
					pts = append(pts[:j], pts[j+1:]...)
					j--
					collision = true
				}
			}
			if collision {
				pts = append(pts[:i], pts[i+1:]...)
				i--
				fmt.Printf("[%d] Particals: %d\n", t, len(pts))
			}
		}
	}
}

func Abs(n int) int {
	if n > 0 {
		return n
	} else {
		return -n
	}
}
