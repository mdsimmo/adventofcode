package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type State struct {
	ex int
	ey int
	full [][]bool
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	state := State {
		full: make([][]bool, 0),
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r < '0' || r > '9'
		})
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		size, _ := strconv.Atoi(parts[2])
		used, _ := strconv.Atoi(parts[3])
		
		if used == 0 {
			state.ex = x
			state.ey = y
		} 
		full := size > 200 
		if (size > 0 && size < 80) || (size > 100 && size < 500) || size > 510 {
			panic("Found different size")
		}
		if (used > 0 && used < 60) || (used > 80 && used < 490) || used > 510 {
			panic("Found different used")
		}
		if size > 200 && used < 200 {
			panic("Found  large drive with low usage")
		}
		for i := len(state.full); i <= x; i++ {
			state.full = append(state.full, []bool {})
		}
		for j := len(state.full[x]); j <= y; j++ {
			state.full[x] = append(state.full[x], false)
		}
		state.full[x][y] = full
	}

	// Print the grid and find the one empty computer
	for j := 0; j < len(state.full[0]); j++ {
		for i := 0; i < len(state.full); i++ {
			if state.ex == i && state.ey == j {
				fmt.Print("_")
			} else if state.full[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	
	// Part 1 Answer
	fmt.Printf("Pairs: %d\n", calcPairs(state.full))
	
	// Look at the data arrangement printed - it is not random
	// there is a wall, with a gap at the left
	best := 0
	best += state.ex // move all the way to the left
	best += state.ey // move all the way to the top
	best += len(state.full)-1 // move all the way to the right
	best += (len(state.full)-2) * 5 // shuffle the data across to the goal
	fmt.Printf("Best: %d\n", best)
}
func calcPairs(full [][]bool) int {
	data := 0
	for j := 0; j < len(full[0]); j++ {
		for i := 0; i < len(full); i++ {
			if !full[i][j] {
				data++
			}
		}
	}
	return data - 1 // minus one for the one empty cell
}
