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
	
	for i := 0; ; i++ {
		good := is01010(i, lines)
		if good {
			fmt.Printf("FOUND: %d\n", i)
			return
		}
	}
}

func is01010(input int, lines []string) bool {
	fmt.Printf("Trial: %d:  ", input)
	registers := make([]int, 4)
	registers[0] = input
	outputs := 0
	lastOutput := -1
	for ptr := 0; ptr < len(lines); {
		// Look for shortcut instructions
		ptrNew, registersNew, taken := shortcutMultiply(lines, ptr, registers)
		if taken {
			ptr = ptrNew
			registers = registersNew
			//panic("Sub")
			continue
		}

		line := lines[ptr]
		//fmt.Printf("[%d] %s\n", ptr, line)
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "cpy":
			if parts[2][0] >= 'a' && parts[2][0] <= 'z' {
				val1, e := strconv.Atoi(parts[1])
				if e != nil {
					val1 = registers[parts[1][0]-'a']
				}
				registers[parts[2][0]-'a'] = val1
			}
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
				val2, e := strconv.Atoi(parts[2])
				if e != nil {
					val2 = registers[parts[2][0]-'a']
				}
				ptr += val2
			} else {
				ptr++
			}
		case "tgl":
			val1 := registers[parts[1][0]-'a']
			if ptr + val1 >= 0 && ptr + val1 < len(lines) {
				lineTarget := lines[ptr+val1]
				lineParts := strings.Split(lineTarget, " ")
				var newLine strings.Builder
				if len(lineParts) == 2 {
					if lineParts[0] == "inc" {
						newLine.WriteString("dec")
					} else {
						newLine.WriteString("inc")
					}
				} else {
					if lineParts[0] == "jnz" {
						newLine.WriteString("cpy")
					} else {
						newLine.WriteString("jnz")
					}
				}
				for i := 1; i < len(lineParts); i++ {
					newLine.WriteString(" ")
					newLine.WriteString(lineParts[i])
				}
				lines[ptr+val1] = newLine.String()
			}
			ptr++
		case "out":
			val1 := registers[parts[1][0]-'a']
			fmt.Printf("%d", val1)
			if val1 < 0 || val1 > 1 || val1 == lastOutput {
				fmt.Printf("\n")
				return false
			}
			lastOutput = val1
			outputs++
			if outputs == 100 {
				fmt.Printf("\n")
				return true
			}
			ptr++
		default:
			panic("Unknown instruction")
		}
		//fmt.Printf("  Regs: %+v\n", registers)
	}
	//fmt.Printf("%+v\n", registers)
	fmt.Printf("\n")
	return false
}

func shortcutMultiply(lines []string, ptr int, regs []int) (int, []int, bool) {

	// cpy b c; inc a; dec c; jnz c -2; dec d; jnz d -5;
	// => a += b * d; c = 0; d = 0;
	if ptr + 5 < len(lines) {
		cpy := strings.Split(lines[ptr+0], " ")
		inc := strings.Split(lines[ptr+1], " ")
		dec1 := strings.Split(lines[ptr+2], " ")
		jnz1 := strings.Split(lines[ptr+3], " ")
		dec2 := strings.Split(lines[ptr+4], " ")
		jnz2 := strings.Split(lines[ptr+5], " ")
		if cpy[0] == "cpy" && inc[0] == "inc" && dec1[0] == "dec" && jnz1[0] == "jnz" && dec2[0] == "dec" && jnz2[0] == "jnz" {
			if cpy[2] == dec1[1] && cpy[2] == jnz1[1] &&
				jnz1[2] == "-2" && jnz2[2] == "-5" &&
				dec2[1] == jnz2[1] {
				
				//fmt.Printf("Shortcut Multiply: %s += %s * %s; %s = 0; %s = 0;\n", inc[1], cpy[1], dec2[1], dec1[1], dec2[1])
				//for i := ptr; i < ptr + 6; i++ {
				//	fmt.Printf("[%d], %s\n", i, lines[i])
				//}
				
				ai := inc[1][0]-'a'
				ci := cpy[2][0]-'a'
				di := dec2[1][0]-'a'
				
				// b may be numeric or register
				valb, e := strconv.Atoi(cpy[1])
				if e != nil {
					valb = regs[cpy[1][0]-'a']
				}

				regs[ai] += valb * regs[di]
				regs[ci] = 0
				regs[di] = 0

				//fmt.Printf("  Regs: %+v\n", regs)
				return ptr + 6, regs, true
			}
		}
	}


	// jnz c 2; jnz 1 4; dec b; dec c; jnz 1 -4;
	// => b -= c; c = 0
	if ptr + 4 < len(lines) {
		jnz1 := strings.Split(lines[ptr+0], " ")
		jnz2 := strings.Split(lines[ptr+1], " ")
		dec1 := strings.Split(lines[ptr+2], " ")
		dec2 := strings.Split(lines[ptr+3], " ")
		jnz3 := strings.Split(lines[ptr+4], " ")
		if jnz1[0] == "jnz" && jnz2[0] == "jnz" && dec1[0] == "dec" && dec2[0] == "dec" && jnz3[0] == "jnz" {
			if jnz1[1] == dec2[1] &&
				jnz1[2] == "2" && 
				jnz2[1] == "1" && jnz2[2] == "4" &&
				jnz3[1] == "1" && jnz3[2] == "-4" {

				//fmt.Printf("Shortcut Substract: %s -= %s; %s = 0;\n", dec1[1], jnz1[1], jnz1[1])
				//for i := ptr; i < ptr + 5; i++ {
				//	fmt.Printf("[%d], %s\n", i, lines[i])
				//}
				
				bi := dec1[1][0]-'a'
				ci := dec2[1][0]-'a'
				
				regs[bi] -= regs[ci]
				regs[ci] = 0

				//fmt.Printf("  Regs: %+v\n", regs)
				return ptr + 6, regs, true
			}
		}
	}

	return ptr, regs, false
}