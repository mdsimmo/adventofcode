package main

import (
	"fmt"
	"slices"
	"strconv"
)

type Point struct {
	x int
	y int
}

func main() {
	input := "hfdlxzhv"
	//input := "flqrgnkx"

	grid := make([][]bool, 128)
	for i := 0; i < 128; i++ {
		grid[i] = make([]bool, 128)
		hashres := hash(input + "-" + strconv.Itoa(i))
		for j := 0; j < len(hashres); j++ {
			for k := 0; k < 8; k++ {
				if hashres[j] >> (7-k) & 1 == 1 {
					grid[i][j*8+k] = true
				}
			}
		}
	}
	
	// part 1 - count on cells
	sum := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell {
				sum++
			}
		}
	}
	fmt.Printf("Sum: %d\n", sum)
	
	fmt.Println()
	for j := 0; j < 20; j++ {
		for i := 0; i < 20; i++ {
			if grid[i][j] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	// part 2 - find regions
	groups := 0
	searched := map[Point]bool {}
	groupNos := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		groupNos[i] = make([]int, len(grid[0]))
	}
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid); i++ {
			if grid[i][j] && !searched[Point{ x: i, y: j }]{
				groups++
				
				// flood fill
				toFill := []Point {{ x: i, y: j }}
				for len(toFill) > 0 {
					p := toFill[len(toFill)-1]
					toFill = toFill[:len(toFill)-1]
					if searched[p] {
						continue
					}
					searched[p] = true
					groupNos[p.x][p.y] = groups
					for _, po := range []Point {
						{ x: p.x+1, y:p.y },
						{ x: p.x-1, y:p.y },
						{ x: p.x, y:p.y+1 },
						{ x: p.x, y:p.y-1 },
					} {
						if po.x >= 0 && po.y >= 0 && po.x < len(grid) && po.y < len(grid[po.x]) && grid[po.x][po.y] && !searched[po] {
							toFill = append(toFill, po)
						}
					}
				}
			}
		}
	}
	fmt.Printf("Groups: %d\n", groups)

	fmt.Println()
	for j := 0; j < 20; j++ {
		for i := 0; i < 20; i++ {
			if groupNos[i][j] != 0 {
				if !grid[i][j] {
					panic("searched a non grid")
				}
				fmt.Printf("%2d", groupNos[i][j]%100)
			} else {
				if grid[i][j] {
					panic("Missed a grid")
				}
				fmt.Printf(" .")
			}
		}
		fmt.Println()
	}
}

// knot hash
func hash(input string) []byte {
	lengths := []byte(input)
	lengths = append(lengths, []byte { 17, 31, 73, 47, 23 }...)
	
	data := make([]byte, 256)
	for i := 0; i < len(data); i++ {
		data[i] = byte(i)
	}
	
	position := 0
	skip := 0
	for i := 0; i < 64; i++ {
		for _, length := range(lengths) {
			subData := make([]byte, length)
			for i := 0; i < int(length); i++ {
				subData[i] = data[(position+i)%len(data)]
			}
			slices.Reverse(subData)
			for i := 0; i < int(length); i++ {
				data[(position+i)%len(data)] = subData[i]
			}
			position += int(length) + skip 
			skip++
		}
	}

	dense := make([]byte, len(data)/16)
	for i := 0; i < len(dense); i++ {
		dense[i] = 0
		for j := 0; j < 16; j++ {
			dense[i] ^= data[i*16+j]
		}
	}

	return dense
}
