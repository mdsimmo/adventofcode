package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	const width int = 101
	const height int = 103

	// process the data
	file, _ := os.Open("in.txt")
	robots := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return (r < '0' || r > '9') && r != '-'
		})
		nums := make([]int, len(parts))
		for i, p := range parts {
			nums[i], _ = strconv.Atoi(p)
		}
		robots = append(robots, nums)
	}

	// find the first iteration with most of the robots clumped together
	for i := 0; ; i++ {
		lonelyRobots := printIter(robots, i, width, height, false)
		// divide 3 was a guess that seems to works
		if lonelyRobots < len(robots)/3 {
			printIter(robots, i, width, height, true)
			fmt.Printf("Iters: %d\n", i)
			break
		}
	}
}

func printIter(data [][]int, iter int, width, height int, print bool) int {
	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]int, width)
	}
	// simulate for the number of iterations
	for _, nums := range data {
		px := ((nums[0] + nums[2]*iter) + width*iter) % width
		py := ((nums[1] + nums[3]*iter) + height*iter) % height
		grid[py][px]++
	}

	// if requested, print the robots
	if print {
		for j := 0; j < height; j++ {
			for i := 0; i < width; i++ {
				if grid[j][i] > 0 {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

	// count the number of robots with no neighbours`	 -+
	lonlyRobots := 0
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if grid[j][i] > 0 {
				neighbours := 0
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						if j+dy < 0 || i+dx < 0 || j+dy >= height || i+dx >= width {
							continue
						}
						if grid[j+dy][i+dx] > 0 {
							neighbours++
						}
					}
				}
				if neighbours <= 1 {
					lonlyRobots++
				}
			}
		}
	}
	return lonlyRobots
}
