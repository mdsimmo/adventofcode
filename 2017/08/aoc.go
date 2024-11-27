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

	maxMidProcess := 0
	regs := map[string]int {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		reg := parts[0]
		op := parts[1]
		amount, e := strconv.Atoi(parts[2])
		if e != nil {
			panic(e)
		}
		regComp := regs[parts[4]]
		opComp := parts[5]
		valComp, e := strconv.Atoi(parts[6])
		if e != nil {
			panic(e)
		}

		var compResult bool
		switch opComp {
		case ">":
			compResult = regComp > valComp 
		case "<":
			compResult = regComp < valComp 
		case ">=":
			compResult = regComp >= valComp 
		case "<=":
			compResult = regComp <= valComp 
		case "==":
			compResult = regComp == valComp 
		case "!=":
			compResult = regComp != valComp
		default:
			panic("Unknown comp")
		}
		if !compResult {
			continue
		}

		switch op {
		case "inc":
			regs[reg] += amount
		case "dec":
			regs[reg] -= amount
		default:
			panic("unknown op")
		}

		if regs[reg] > maxMidProcess {
			maxMidProcess = regs[reg]
		}
	}

	max := 0
	for _, value := range(regs) {
		if value > max {
			max = value
		}
	}
	fmt.Printf("Largest: %d\n", max)
	fmt.Printf("Largest Mid Process: %d\n", maxMidProcess)
}
