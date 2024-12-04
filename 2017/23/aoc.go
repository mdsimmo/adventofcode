package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2 := true
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	insList := []string {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		insList = append(insList, line)
	}

	registers := make([]int, 'h'-'a'+1)
	mulCount := 0
	if part2 {
		registers['a'-'a'] = 1
	}
	for i := 0; i < len(insList); i++ {
		// Interpret the line
		line := insList[i]
		//fmt.Printf("[%d] %s\n", i, line)
		// Optimise code
		if part2 && i == 10 {
			if line != "set e 2" {
				panic("Optimised wrong line")
			}
			i += 13
			registers['g'-'a'] = 0
			b := registers['b'-'a']
			registers['e'-'a'] = b
			registers['d'-'a'] = b
			for n := 2; n < b; n++ {
				if  b % n == 0 {
					registers['f'-'a'] = 0
					break
				}
			}
			//fmt.Printf("%+v\n", registers)
			continue
		}

		//fmt.Printf("[%d, %d] %s\n", id, i, line)
		parts := strings.Split(line, " ")
		ins := parts[0]
		idxs := make([]int, len(parts)-1)
		vals := make([]int, len(parts)-1)
		for j := 0; j < len(parts)-1; j++ {
			var err error = nil
			vals[j], err = strconv.Atoi(parts[j+1])
			if err != nil {
				idxs[j] = int(parts[j+1][0]) - 'a'
				vals[j] = registers[idxs[j]]
			} else {
				idxs[j] = -1 // not a register
			}
		}
		
		// Execute the instruction
		switch ins {
		case "set":
			registers[idxs[0]] = vals[1]
		case "sub":
			registers[idxs[0]] -= vals[1]
		case "mul":
			registers[idxs[0]] *= vals[1]
			mulCount++
		case "jnz":
			if vals[0] != 0 {
				// minus one to account for normal loop increment
				i += vals[1] - 1
			}
		default: 
			panic("Unknown instruction")
		}
	}
	if !part2 {
		fmt.Printf("Mul Count: %d\n", mulCount)
	} else {
		fmt.Printf("Reg H: %d\n", registers['h'-'a'])
	}
}
