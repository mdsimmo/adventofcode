package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}

	grid := [][]rune {
		[]rune(".#."),
		[]rune("..#"),
		[]rune("###"),
	}
	patterns := map[string][][]rune {} 
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Printf("Process: %s\n", line)
		parts := strings.Split(line, " => ") 
		if len(parts) != 2 {
			panic("Should always be two")
		}
		pat := convert(parts[0])
		res := convert(parts[1])
		// find rotated matches

		rot := pat 
		size := len(rot)
		for i := 0; i < 4; i++ {
			// rotate the grid
			rotNew := clone(rot)
			for j := 0; j < size; j++ {
				for k := 0; k < size; k++ {
					pos := complex(float64(j) - float64(size-1)/2.0, float64(k) - float64(size-1)/2.0)
					pos *= complex(0, 1)
					x := int(real(pos) + float64(size-1)/2)
					y := int(imag(pos) + float64(size-1)/2)
					rotNew[x][y] = rot[j][k]
					//fmt.Printf("(%d, %d) -> (%d,%d) = %c\n", j, k, x, y, rot[j][k])
				}
			}
			rot = rotNew
			patterns[stringify(rot)] = res
			//fmt.Printf("  Rotated: %+v\n", stringify(rot))

			// flip
			flipH := clone(rot)
			flipV := clone(rot)
			for j := 0; j < size; j++ {
				for k := 0; k < size; k++ {
					flipH[j][k] = rot[size-1-j][k]
					flipV[j][k] = rot[j][size-1-k]
				}
			}
			patterns[stringify(flipV)] = res
			patterns[stringify(flipH)] = res
			//fmt.Printf("  FlipV: %+v\n", stringify(flipV))
			//fmt.Printf("  FlipH: %+v\n", stringify(flipH))
		}
	}

	//fmt.Printf("%+v\n", patterns)
	// part 1
	iterate(5, grid, patterns)
	// part 2
	iterate(18, grid, patterns)
}

func iterate(n int, start [][]rune, patterns map[string][][]rune) {
	grid := clone(start)
	for i := 0; i < n; i++ {
		size := len(grid)
		fsize := -1
		if size % 2 == 0 {
			fsize = 2
		} else if size % 3 == 0 {
			fsize = 3
		} else {
			panic("Cannot split")
		}

		newSize := (size / fsize) * (fsize + 1)
		newGrid := make([][]rune, newSize)
		for j := 0; j < newSize; j++ {
			newGrid[j] = make([]rune, newSize)
		}

		for j := 0; j < size/fsize; j++ {
			for k := 0; k < size/fsize; k++ {
				subGrid := make([][]rune, fsize)
				for dx := 0; dx < fsize; dx++ {
					subGrid[dx] = make([]rune, fsize)
					for dy := 0; dy < fsize; dy++ {
						subGrid[dx][dy] = grid[j*fsize+dx][k*fsize+dy]
					}
				}
				pat := stringify(subGrid)
				replace := patterns[pat]
				if replace == nil {
					panic(fmt.Sprintf("Pattern not found: %s", pat))
				}
				if len(replace) != fsize+1 {
					panic("Unexpected replace size")
				}
				for dx := 0; dx < len(replace); dx++ {
					for dy := 0; dy < len(replace); dy++ {
						newGrid[j*(fsize+1)+dx][k*(fsize+1)+dy] = replace[dx][dy]
					}
				}
			}
		}
		grid = newGrid
	}

	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == '#' {
				sum++
			}
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

func clone(in [][]rune) [][]rune {
	size := len(in)
	rotNew := make([][]rune, size)
	for j := 0; j < size; j++ {
		rotNew[j] = slices.Clone(in[j])
	}
	return rotNew
}

func stringify(in [][]rune) string {
	var result strings.Builder
	for i := 0; i < len(in); i++ {
		result.WriteString(string(in[i]))
		result.WriteRune('/')
	}
	return result.String()
}

func convert(in string) [][]rune {
	split := strings.Split(in, "/")
	result := make([][]rune, len(split))
	for i := 0; i < len(result); i++ {
		result[i] = []rune(split[i])
	}
	return result
}
