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
	// Solution explaination:
	// From studying the program:
	// The jump instruction is at the end, and just causes the program to loop
	// Each loop iteration only cares about register A. B and C don't matter
	// The next value of A is always A / 8
	// The final value of A is 0

	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	// Skip over the register values
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
	}
	// Read the program
	scanner.Scan()
	program := ExtractInts(scanner.Text())
	// fmt.Printf("Program: %+v\n", program)

	// The end A register must be 0
	// The second last A value must meet:
	//   A / 8 == 0, i.e. A between 0 and 7
	//   out(A) = program[]
	// Likewise, A[n-2] = A[n-1] + 0-7
	// Iterate backwards until all possible starting values are found
	aValues := make([][]int, len(program)+1)
	aValues[len(program)] = []int{0}
	for ai := len(program) - 1; ai >= 0; ai-- {
		aValues[ai] = make([]int, 0)
		for _, aRes := range aValues[ai+1] {
			for i := aRes * 8; i < aRes*8+8; i++ {
				if runLoop(i, program) == program[ai] {
					aValues[ai] = append(aValues[ai], i)
				}
			}
		}
	}

	// find lowest possible starting value
	slices.Sort(aValues[0])
	fmt.Printf("%d\n", aValues[0][0])

}

// Runs the program, and returns the output.
// Does not continue after the output
func runLoop(a int, program []int) int {
	regs := make([]int, 3)
	regs[0] = a

	for ptr := 0; ptr < len(program); ptr += 2 {
		opcode := program[ptr]
		operand := program[ptr+1]
		var combo int
		if operand <= 3 {
			combo = operand
		} else if operand < 7 {
			combo = regs[operand-4]
		} else {
			combo = -1 // should not be used
		}

		switch opcode {
		case 0: // adv: a divide
			regs[0] = regs[0] / (1 << combo)
		case 1: // bxl: bitwise or
			regs[1] = regs[1] ^ operand
		case 2: // bst: modulo
			regs[1] = combo % 8
		case 3: // jnz: jump
			// `jnz` is always after `out`
			panic("Should not get here")
			// if regs[0] != 0 {
			// ptr = operand - 2
			// }
		case 4: // bxc: bitwise or
			regs[1] = regs[1] ^ regs[2]
		case 5: // out
			return combo % 8
		case 6: // bdv: b divide
			regs[1] = regs[0] / (1 << combo)
		case 7: // cdv: c divide
			regs[2] = regs[0] / (1 << combo)
		default:
			panic("Invalid instruction")
		}
	}
	panic("No output")
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
