package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	var sum int64 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r < '0' || r > '9'
		})
		nums := make([]int, len(parts))
		for i, p := range parts {
			nums[i], err = strconv.Atoi(p)
			if err != nil {
				panic(err)
			}
		}

		if matches(nums[0], nums[1:]) {
			sum += int64(nums[0])
			if sum > math.MaxInt64 / 100 {
				panic("Extra big number")
			}
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

func matches(test int, vals []int) bool {
	// brute force every combination by encoding each possibility as an interger
	// fmt.Printf("Tests: %d = %+v\n", test, vals)
	outer:
	for i := 0; i < 1<<((len(vals)-1)*2); i++ {
		res := vals[0]
		// fmt.Printf("  %d", vals[0])
		for j := 0; j < len(vals) - 1; j++ {
			switch (i >> (j*2)) & 0b11 {
			case 0b00:
				res += vals[j+1]
				// fmt.Printf("+%d", vals[j+1])
			case 0b01:
				res *= vals[j+1]
				// fmt.Printf("*%d", vals[j+1])
			case 0b10:
				if vals[j+1] >= 1000 {
					panic("Cannot handle numbers >= 1000")
				} else if vals[j+1] >= 100 {
					res *= 1000
				} else if vals[j+1] >= 10 {
					res *= 100
				} else {
					res *= 10
				}
				res += vals[j+1]
				// fmt.Printf("|%d", vals[j+1])
			case 0b11:
				// invalid
				//fmt.Printf("ERR\n")
				// skip remaining errs
				i += (1 << j) - 1
				continue outer
			}
			// make sure overflows arnt happening
			if res > math.MaxInt / 1000 {
				panic("Big number")
			}
		}
		// fmt.Printf("=%d\n", res)
		if test == res {
			// fmt.Printf("  Good\n")
			return true
		}
	}
	// fmt.Printf("  Bad\n")
	return false
}

