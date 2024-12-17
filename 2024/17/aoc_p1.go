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
	regs := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		regs = append(regs, ExtractInts(line)...)
	}
	scanner.Scan()
	program := ExtractInts(scanner.Text())
	// fmt.Printf("Regs: %+v\n", regs)
	fmt.Printf("Program: %+v\n", program)

	var output strings.Builder
	for ptr := 0; ptr < len(program); ptr += 2 {
		opcode := program[ptr]
		operand := program[ptr+1]
		// fmt.Printf("Process: %d, %d\n", opcode, operand)
		var combo int
		if operand <= 3 {
			combo = operand
		} else if operand < 7 {
			combo = regs[operand-4]
		} else {
			combo = -1 // should not be used
		}
		// fmt.Printf("  Combo: %d\n", combo)

		switch opcode {
		case 0: // adv: a divide
			regs[0] = regs[0] / (1 << combo)
		case 1: // bxl: bitwise or
			regs[1] = regs[1] ^ operand
		case 2: // bst: modulo
			regs[1] = combo % 8
		case 3: // jnz: jump
			if regs[0] != 0 {
				ptr = operand - 2
			}
		case 4: // bxc: bitwise or
			regs[1] = regs[1] ^ regs[2]
		case 5: // out
			res := combo % 8
			output.WriteString(fmt.Sprintf("%d,", res))
		case 6: // bdv: b divide
			regs[1] = regs[0] / (1 << combo)
		case 7: // cdv: c divide
			regs[2] = regs[0] / (1 << combo)
		default:
			panic("Invalid instruction")
		}
		// fmt.Printf("  Regs: %+v\n", regs)
	}
	fmt.Printf("%s\n", output.String())
}

func ExtractInts(line string) []int {
	num_s := strings.FieldsFunc(line, func(r rune) bool {
		return r < '0' || r > '9'
	})
	nums := make([]int, len(num_s))
	for i, s := range num_s {
		num, e := strconv.Atoi(s)
		if e != nil {
			panic(e)
		}
		nums[i] = num
	}
	return nums
}
