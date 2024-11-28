package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("in.txt")
	if err != nil {
		panic(err)
	}
	contents := strings.Trim(string(file), "\n")
	//contents = "se,sw,se,sw,sw"
	parts := strings.Split(contents, ",")

	x, y := 0, 0
	max := 0
	for _, part := range(parts) {
		switch part {
		case "n":
			y -= 2
		case "ne":
			y -= 1
			x += 1
		case "nw":
			y -= 1
			x -= 1
		case "se":
			y += 1
			x += 1
		case "sw":
			y += 1
			x -= 1
		case "s":
			y += 2
		default:
			panic("Unknown direction")
		}
		dist := homeDist(x, y)
		if dist > max {
			max = dist
		}
	}

	fmt.Printf("Dist: %d\n", homeDist(x, y))
	fmt.Printf("Max Dist: %d\n", max)
}

func Abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func homeDist(x, y int) int {
	y = Abs(y)
	x = Abs(x)

	dist := x
	if (y-x) % 2 != 0 {
		panic("Should not happen?")
	}
	if y > x {
		dist += (y - x)/2
	}
	return dist
}
