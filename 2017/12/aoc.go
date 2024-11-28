package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	conc := map[int][]int {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r < '0' || r > '9'
		})
		nums := make([]int, len(parts))
		for i, part := range(parts) {
			val, e := strconv.Atoi(part)
			if e != nil {
				panic(e)
			}
			nums[i] = val
		}
		conc[nums[0]] = nums[1:]
	}

	// part 1 answer
	fmt.Printf("Connected to zero: %d\n", len(findInGroup(0, conc)))

	// list every point
	uncon := map[int]bool {}
	for num := range conc {
		uncon[num] = true
	}

	// count all groups
	groups := 0
	for len(uncon) > 0 {
		//get first unconnected point
		var num int
		for n := range uncon {
			num = n
			break
		}

		group := findInGroup(num, conc)
		for n := range group {
			delete(uncon, n)
		}
		groups++
	}
	fmt.Printf("Groups: %d\n", groups)
}

func findInGroup(start int, conc map[int][]int) map[int]bool {
	changed := true
	conected := map[int]bool { start: true }
	for changed {
		changed = false
		for num := range conected {
			for _, other := range conc[num] {
				if conected[other] {
					continue
				}
				conected[other] = true
				changed = true
			}
		}
	}
	return conected
}
