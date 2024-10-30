package main

import (
	"fmt"
)

func aoc_20_run() {
	min_presents := 36000000
	for i := 50; ; i++ {
		presents := count_presents_2(i)
		fmt.Printf("%d: %d\n", i, presents)
		if presents >= min_presents {
			fmt.Printf("House %d: %d\n", i, presents)
			return
		}
	}
}

func count_presents(house int) int {
	presents := 0
	for i := 1; i <= house/2; i++ { 
		if house % i == 0 {
			presents += i * 10
		}
	}
	presents += house * 10
	return presents
}

func count_presents_2(house int) int {
	presents := 0
	for i := house/50; i <= house/2; i++ { 
		if house % i == 0 {
			if house / i <= 50 {
				presents += i * 11
			}
		}
	}
	presents += house * 11
	return presents
}
