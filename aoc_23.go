package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func aoc_23_run() {
	file, err := os.Open("aoc_23_in.txt")
	if err != nil {
		panic(err)
	}

	ins := []string {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ins = append(ins, line)
	}

	a := 1
	b := 0
	ptr := 0

	for ptr >= 0 && ptr < len(ins) {
		i := ins[ptr]
		fmt.Printf("[%d] %5d, %5d; %s\n", ptr, a, b, i)
		ptr, a, b = applyIns(i, ptr, a, b)
	}

	fmt.Printf("[%d] %5d, %5d\n", ptr, a, b)
}

func applyIns(line string, ptr int, a int, b int) (int, int, int) {
	parts := strings.FieldsFunc(line, splitTokens)
	ins := parts[0]

	val1 := parseValue(parts[1], a, b)
	var val2 int
	if len(parts) > 2 {
		val2 = parseValue(parts[2], a, b)
	}
	isA := parts[1] == "a"

	switch ins {
	case "hlf":
		if isA {
			return ptr+1, a/2, b 
		} else {
			return ptr+1, a, b/2
		}
	case "tpl":
		if isA {
			return ptr+1, a*3, b
		} else {
			return ptr+1, a, b*3
		}
	case "inc":
		if isA {
			return ptr+1, a+1, b
		} else {
			return ptr+1, a, b+1
		}
	case "jmp":
		return ptr + val1, a, b
	case "jie":
		if val1 % 2 == 0 {
			return ptr+val2, a, b
		} else {
			return ptr+1, a, b
		}
	case "jio":
		if val1 == 1 {
			return ptr+val2, a, b
		} else {
			return ptr+1, a, b
		}
	default:
		panic("Unknown instruction")
	}

}

func parseValue(val string, a int, b int) int {
	var val1 int
	switch val {
	case "a":
		val1 = int(a)
	case "b":
		val1 = int(b)
	default:
		val1_i, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		val1 = int(val1_i)
	}
	return val1
}

func splitTokens(c rune) bool {
	return c == ' ' || c == ','
}
