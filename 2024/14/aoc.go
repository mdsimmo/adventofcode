package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main1() {

	const width int = 101
	const height int = 103
	const iter int = 100
	// const width int = 11
	// const height int = 7
	// const iter int = 100

	file, _ := os.Open("in.txt")
	quads := make([]int, 4)
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
		px := ((nums[0] + nums[2]*iter) + width*iter) % width
		py := ((nums[1] + nums[3]*iter) + height*iter) % height
		if px < width/2 {
			if py < height/2 {
				quads[0]++
			} else if py > height/2 {
				quads[1]++
			}
		} else if px > width/2 {
			if py < height/2 {
				quads[2]++
			} else if py > height/2 {
				quads[3]++
			}
		}
		fmt.Printf("%+v -> %d,%d\n", nums, px, py)
	}
	factor := 1
	for _, v := range quads {
		factor *= v
	}
	fmt.Printf("Factor: %d\n", factor)
}
