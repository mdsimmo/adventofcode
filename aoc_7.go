package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Op int8

const (
	AND Op = iota
	OR 
	NOT
	RSHIFT
	LSHIFT
	NONE
)

type statment struct {
	line string
	op1 string
	op2 string
	op Op
	res string
}

func aoc_7_run() {
	values := make(map[string]uint16)
	operations := make([]statment, 0)
	
	file, e := os.Open("aoc_7_in2.txt")
	if e != nil {
		panic(e)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		statement := statment {
			line: line,
			op: NONE,
		}
		assigned := false

		for _, token := range(tokens) {
			switch token {
			case "AND":
				statement.op = AND
			case "OR":
				statement.op = OR
			case "NOT":
				statement.op = NOT
			case "LSHIFT":
				statement.op = LSHIFT
			case "RSHIFT":
				statement.op = RSHIFT
			case "->":
				assigned = true
			default:
				if assigned {
					statement.res = token
				} else if statement.op1 == "" {
					statement.op1 = token
				} else {
					statement.op2 = token
				}
			}
		}
		
		operations = append(operations, statement)
		fmt.Printf("%s: %s %d %s -> %s\n", line, statement.op1, statement.op, statement.op2, statement.res)
	}

	for len(operations) > 0 {
		for index := 0; index < len(operations); index++ {
			stm := operations[index]
			fmt.Printf("Checking: %s\n", stm.line)

			temp1, err1 := strconv.ParseInt(stm.op1, 10, 32)
			val1 := uint16(temp1)
			val1_ok := true
			if err1 != nil {
				val1, val1_ok = values[stm.op1]
			}
			
			temp2, err2 := strconv.ParseInt(stm.op2, 10, 32)
			val2 := uint16(temp2)
			val2_ok := true
			if err2 != nil {
				val2, val2_ok = values[stm.op2]
			}
			if stm.op2 == "" {
				val2 = 0
				val2_ok = true
			}

			fmt.Printf("  %s: %d (%t), %s: %d (%t)\n", stm.op1, val1, val1_ok, stm.op2, val2, val2_ok)
			if !val1_ok || !val2_ok {
				continue
			}

			result := uint16(0)
			switch stm.op {
			case AND:
				result = val1 & val2
			case OR:
				result = val1 | val2
			case NOT:
				result = ^val1
			case LSHIFT:
				result = val1 << val2
			case RSHIFT:
				result = val1 >> val2
			case NONE:
				result = val1
			default:
				panic("Unknown op")
			}
			fmt.Printf("   %s = %d\n", stm.res, result)
			values[stm.res] = result
			
			operations[index] = operations[len(operations)-1]
			operations = operations[:len(operations)-1]
		}
	}
	fmt.Printf("A: %d", values["a"])
}
