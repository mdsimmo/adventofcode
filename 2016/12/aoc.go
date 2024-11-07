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
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	registers := make([]int, 4)
	registers[2] = 1
	for ptr := 0; ptr < len(lines); {
		line := lines[ptr]
		fmt.Printf("[%d] %s\n", ptr, line)
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "cpy":
			val1, e := strconv.Atoi(parts[1])
			if e != nil {
				val1 = registers[parts[1][0]-'a']
			}
			registers[parts[2][0]-'a'] = val1
			ptr++
		case "inc":
			registers[parts[1][0]-'a']++
			ptr++
		case "dec":
			registers[parts[1][0]-'a']--
			ptr++
		case "jnz":
			val1, e := strconv.Atoi(parts[1])
			if e != nil {
				val1 = registers[parts[1][0]-'a']
			}
			if val1 != 0 {
				val2, _ := strconv.Atoi(parts[2])
				ptr += val2
			} else {
				ptr++
			}
		}
		fmt.Printf("  Regs: %+v\n", registers)
	}
	fmt.Printf("%+v\n", registers)
}
