package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	
	var data [](map[rune]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line:= scanner.Text()
		
		if len(data) == 0 {
			data = make([](map[rune]int), len(line))
			for i := 0; i < len(data); i++ {
				data[i] = make(map[rune]int)
			}
		}

		for i, r := range(line) {
			data[i][r]++
		}
	}

	for i := 0; i < len(data); i++ {
		best := '-'
		bestCnt := 0
		for r, cnt := range(data[i]) {
			if cnt > bestCnt {
				bestCnt = cnt
				best = r
			}
		}
		fmt.Printf("%c", best)
	}
	
	fmt.Println()

	for i := 0; i < len(data); i++ {
		best := '-'
		bestCnt := math.MaxInt
		for r, cnt := range(data[i]) {
			if cnt < bestCnt {
				bestCnt = cnt
				best = r
			}
		}
		fmt.Printf("%c", best)
	}
}
