package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Touch struct {
	offset int
	modulo int
	cost int
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}

	filters := map[int]int {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' ' || r == ':'
		})
		nums := make([]int, len(parts))
		for i, p := range(parts) {
			v, e := strconv.Atoi(p)
			if e != nil {
				panic(e)
			}
			nums[i] = v
		}
		filters[nums[0]] = nums[1]
	}

	touches := map[Touch]bool {}
	for depth, ran := range filters {
		touches[Touch {
			offset: depth,
			modulo: ran * 2 - 2,
			cost: depth * ran,
		}] = true
	}

	cost0, _ := check(0, touches)
	fmt.Printf("Cost: %d\n", cost0)

	for i := 0; ; i++ {
		if _, caught := check(i, touches); !caught {
			fmt.Printf("Found gap: %d\n", i)
			break
		} else {
		//	fmt.Printf("No gap: %d\n", i)
		}
	}
}

func check(delay int, touches map[Touch]bool) (int, bool) {
	cost := 0
	caught := false
	for touch := range touches {
		if (delay + touch.offset) % touch.modulo == 0 {
			cost += touch.cost
			caught = true
		}
	}
	return cost, caught
}
