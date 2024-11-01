package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func aoc_6_run() {
	var grid [1000][1000]int

	file, _ := os.Open("aoc_6_in.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		
		parts := strings.Split(line, " ")
		toggle := false
		on := false
		off := false
		startx := -1
		starty := -1
		endx := -1
		endy := -1
		for _, part := range(parts) {
			switch part {
			case "toggle":
				toggle = true
			case "turn": {}
			case "on":
				on = true
			case "off":
				off = true
			case "through": {}
			default:
				coords := strings.Split(part, ",")
				x, _ := strconv.Atoi(coords[0])
				y, _ := strconv.Atoi(coords[1])
				if startx == -1 {
					startx = x
					starty = y
				} else {
					endx = x
					endy = y
				}
			}
		}

		for i := startx; i <= endx; i++ {
			for j := starty; j <= endy; j++ {
				if on {
					grid[i][j]++
				} else if off {
					grid[i][j]--
				} else if toggle {
					grid[i][j] += 2
				}
				if grid[i][j] < 0 {
					grid[i][j] = 0
				}
			}
		}
	}

	count := 0
	for _, row := range(grid) {
		for _, cell := range(row) {
			count += cell
		}
	}
	fmt.Printf("Count: %d", count)
}
