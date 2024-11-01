package main

import "fmt"

func aoc_25_run() {
	node := calcCodeNo(2978, 3083)
	result := calcThing(node)
	fmt.Printf("%d\n", result)
		
}

func calcCodeNo(row int, col int) int {
	num := 0
	add_row := row
	for i := 1; i <= col; i++ {
		num += add_row
		add_row++
	}
	add_row = row - 1
	for i := row - 2; i >= 1; i-- {
		num += i
	}
	return num
}

func calcThing(n int) int {
	result := 20151125
	for i := 1; i < n; i++ {
		result = result * 252533 % 33554393
	}
	return result
}
