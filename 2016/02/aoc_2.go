package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	result := process("in.txt", AddPosition1, Keymap)
	fmt.Printf("Pass: %x\n", result)
	
	result2 := process("in.txt", AddPosition2, Keymap2)
	fmt.Printf("Pass: %x\n", result2)
}

func process(filename string, adder func(complex128, complex128)complex128, keymap func(complex128)int) int {
	file, err := os.Open(filename)		
	if err != nil {
		panic(err)
	}
	
	pass := 0

	pos := complex(0, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pos = LinePosition(line, pos, adder) 
		result := keymap(pos)
		pass = pass * 16 + result
	}

	return pass
}


func LinePosition(line string, pos complex128, adder func(complex128, complex128)complex128) complex128 {
	for _, c := range(line) {
		var dir complex128
		switch c {
		case 'R':
			dir = complex(1, 0)
		case 'L':
			dir = complex(-1, 0)
		case 'U':
			dir = complex(0, 1)
		case 'D':
			dir = complex(0, -1)
		default:
			panic("Unknown instruction")
		}
		
		pos = adder(pos, dir)
	}
	return pos
}

func AddPosition1(pos complex128, dir complex128) complex128 {
	pos += dir
	if math.Abs(real(pos)) > 1.0 {
		pos = complex(real(pos)/math.Abs(real(pos)), imag(pos))
	}
	if math.Abs(imag(pos)) > 1.0 {
		pos = complex(real(pos), imag(pos)/math.Abs(imag(pos)))
	}
	return pos
}

func AddPosition2(pos complex128, dir complex128) complex128 {
	posnew := pos + dir
	rl := int(real(posnew))
	valid := false
	switch int(imag(posnew)) {
	case 2, -2: valid = rl == 2
	case 1, -1: valid = rl >= 1 && rl <= 3
	case 0: valid = rl >= 0 && rl <= 4
	default: valid = false
	}
	if valid {
		return posnew
	} else {
		return pos
	}
}

func Keymap(pos complex128) int {
	return 5 + (int(real(pos))%3) - (int(imag(pos))*3)
}

func Keymap2(pos complex128) int {
	switch int(imag(pos)) {
	case 2: return 1
	case 1: return 1 + int(real(pos))
	case 0: return 5 + int(real(pos))
	case -1: return 9 + int(real(pos))
	case -2: return 11 + int(real(pos))
	default:
		panic("Pos out of range")
	}
}
