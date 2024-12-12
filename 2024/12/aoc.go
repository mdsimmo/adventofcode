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
	grid := map[Point]rune{}
	scanner := bufio.NewScanner(file)
	maxPoint := Point{x: 0, y: 0}
	for scanner.Scan() {
		line := scanner.Text()
		for i, r := range line {
			grid[Point{x: i, y: maxPoint.y}] = r
		}
		maxPoint.x = len(line)
		maxPoint.y++
	}

	dirs := []Point{
		{x: 0, y: 1},
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: -1, y: 0},
	}
	score := 0
	score2 := 0
	complete := map[Point]bool{}
	for i := 0; i < maxPoint.x; i++ {
		for j := 0; j < maxPoint.y; j++ {
			basePoint := Point{x: i, y: j}
			baseRune := grid[basePoint]
			if complete[basePoint] {
				continue
			}

			blob := map[Point]bool{basePoint: true}
			touching := map[Point]map[Point]bool{}
			anyFound := true
			for anyFound {
				anyFound = false
				for p := range blob {
					for _, d := range dirs {
						pNew := Point{p.x + d.x, p.y + d.y}
						if grid[pNew] != baseRune {
							if touching[pNew] == nil {
								touching[pNew] = map[Point]bool{}
							}
							touching[pNew][p] = true
							continue
						}
						if blob[pNew] {
							continue
						}
						anyFound = true
						blob[pNew] = true
					}
				}
			}
			area := len(blob)
			perimiter := 0
			for _, v := range touching {
				perimiter += len(v)
			}
			// fmt.Printf("  Score: %c, %d * %d = %d\n", baseRune, area, perimiter, area*perimiter)
			score += perimiter * area

			for p := range blob {
				complete[p] = true
			}

			// part 2
			sides := 0
			sidesToWalk := touching
			for len(sidesToWalk) > 0 {
				// select a random in/out touching sequre
				var pOut Point
				var pIn Point
				for po, pi := range sidesToWalk {
					pOut = po
					if len(pi) == 0 {
						panic("Should never happen")
					}
					for p := range pi {
						pIn = p
						break
					}
					break
				}

				// Walk in each direction removing the sides from the walk list
				walk(pIn, pOut, Point{x: -1, y: 0}, &sidesToWalk)
				walk(pIn, pOut, Point{x: 1, y: 0}, &sidesToWalk)
				walk(pIn, pOut, Point{x: 0, y: 1}, &sidesToWalk)
				walk(pIn, pOut, Point{x: 0, y: -1}, &sidesToWalk)
				delete(sidesToWalk[pOut], pIn)
				if len(sidesToWalk[pOut]) == 0 {
					delete(sidesToWalk, pOut)
				}

				sides++
			}
			// fmt.Printf("  Score2: %c, %d * %d = %d\n", baseRune, area, sides, area*sides)
			score2 += sides * area
		}
	}
	fmt.Printf("Score: %d\n", score)
	fmt.Printf("Score2: %d\n", score2)
}

func walk(walkIn Point, walkOut Point, dir Point, sidesToWalk *map[Point]map[Point]bool) {
	for {
		walkOutNew := Point{x: walkOut.x + dir.x, y: walkOut.y + dir.y}
		walkInNew := Point{x: walkIn.x + dir.x, y: walkIn.y + dir.y}
		if len((*sidesToWalk)[walkOutNew]) == 0 {
			break
		}
		if !(*sidesToWalk)[walkOutNew][walkInNew] {
			break
		}
		walkOut = walkOutNew
		walkIn = walkInNew
		// remove the side
		delete((*sidesToWalk)[walkOutNew], walkInNew)
		if len((*sidesToWalk)[walkOutNew]) == 0 {
			delete((*sidesToWalk), walkOutNew)
		}
	}
}
