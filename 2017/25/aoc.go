package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Action struct {
	move int
	write int
	state int
}

type State struct {
	a []Action
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan() // start in state A
	scanner.Scan() // how many checksums
	checksums := extractInts(scanner.Text())[0]
	scanner.Scan() // Empty line

	// Read each state
	states := []State {}
	for {
		scanner.Scan() // In State X
		var state State
		state.a = make([]Action, 2)
		for i := 0; i <= 1; i++ {
			scanner.Scan() // if value is zero
			scanner.Scan() // write value n
			state.a[i].write = extractInts(scanner.Text())[0]
			scanner.Scan() // Move to right/left
			if strings.Contains(scanner.Text(), "right") {
				state.a[i].move = 1
			} else {
				state.a[i].move = -1
			}
			scanner.Scan() // continue with state Y.
			newStateString := scanner.Text()
			state.a[i].state = int(newStateString[len(newStateString)-2] - 'A')
		}
		states = append(states, state)
		if !scanner.Scan() { // empty line
			break
		}
	}

	data := map[int]int {}
	stateIndex := 0
	index := 0

	for i := 0; i < checksums; i++ {
		state := states[stateIndex]
		action := state.a[data[index]]
		data[index] = action.write
		stateIndex = action.state
		index += action.move
	}

	checksum := 0
	for _, v := range data {
		if v == 1 {
			checksum++
		}
	}
	fmt.Printf("Checksum: %d\n", checksum)
}

func extractInts(line string) []int {
	parts := strings.FieldsFunc(line, func(r rune) bool {
		return r < '0' || r > '9'
	})
	nums := make([]int, len(parts))
	for i, part := range parts { 
		n, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}
	return nums
}
