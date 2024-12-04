package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}

	grid := [][]rune {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	xmas := []rune("XMAS")
	xmasCount := 0
	dirs := []complex64 {
		complex(0,  1),
		complex(0, -1),
		complex(1,  1),
		complex(1,  0),
		complex(1, -1),
		complex(-1,  1),
		complex(-1,  0),
		complex(-1, -1),
	}
	for y, line := range(grid) {
		for x := range line {
			if grid[y][x] != xmas[0] {
				continue
			}
			loopDir:
			for _, dir := range dirs {
				for i, c := range xmas {
					xx := x + int(real(dir * complex64(complex(float64(i), 0))))
					yy := y + int(imag(dir * complex64(complex(float64(i), 0))))
					if xx < 0 || yy < 0 || xx >= len(line) || yy >= len(grid) {
						continue loopDir
					}
					if grid[yy][xx] != c {
						continue loopDir
					}
				}
				xmasCount++
			}
		}
	}

	fmt.Printf("Count XMAS: %d\n", xmasCount)

	masCount := 0
	dirs2 := [][]complex64 {
		{ complex(1, 1), complex(-1,  -1) },
		{ complex(1, -1), complex(-1,  1) },
	}
	for y, line := range(grid) {
		outer:
		for x := range line {
			if grid[y][x] != 'A' {
				continue
			}
			for _, dir := range dirs2 {
				xx1 := x + int(real(dir[0]))
				yy1 := y + int(imag(dir[0]))
				xx2 := x + int(real(dir[1]))
				yy2 := y + int(imag(dir[1]))
				if xx1 < 0 || yy1 < 0 || xx1 >= len(line) || yy1 >= len(grid) {
					continue outer
				}
				if xx2 < 0 || yy2 < 0 || xx2 >= len(line) || yy2 >= len(grid) {
					continue outer
				}
				c1 := grid[yy1][xx1]
				c2 := grid[yy2][xx2]
				if c1 == c2 {
					continue outer
				}
				if (c1 != 'S' && c1 != 'M') || (c2 != 'S' && c2 != 'M') {
					continue outer
				}
			}
			masCount++
		}
	}

	fmt.Printf("Count X-MAS: %d\n", masCount)
}
