package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	rules := [][]int {}
	orders := [][]int {}
	section2 := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			section2 = true
			continue
		}
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
		
		
		if !section2 {
			rules = append(rules, nums)
		} else {
			orders = append(orders, nums)
		}
	}

	sumGood := 0
	sumBad := 0
	for _, order := range orders {
		if isGood(order, rules) {
			sumGood += order[len(order)/2]
		} else {
			// Order the list
			for !isGood(order, rules) {
				for _, rule := range rules {
					i0 := slices.Index(order, rule[0])
					i1 := slices.Index(order, rule[1])
					if i0 < 0 || i1 < 0 {
						continue
					}
					if i1 < i0 {
						// swap the elements
						order[i0] = rule[1]
						order[i1] = rule[0]
					}
				}
			}
			sumBad += order[len(order)/2]
		}
	}
	fmt.Printf("Sum Good: %d\n", sumGood)
	fmt.Printf("Sum Bad: %d\n", sumBad)
}

func isGood(order []int, rules [][]int) bool {
	for _, rule := range rules {
		i0 := slices.Index(order, rule[0])
		i1 := slices.Index(order, rule[1])
		if i0 < 0 || i1 < 0 {
			continue
		}
		if i1 < i0 {
			return false
		}
	}
	return true
}
