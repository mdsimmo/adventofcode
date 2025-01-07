package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	OR  string = "OR"
	AND string = "AND"
	XOR string = "XOR"
)

type Ins struct {
	op1 string
	op2 string
	ins string
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	values := map[string]bool{}
	scanner := bufio.NewScanner(file)
	// Read the values
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, ": ")
		values[parts[0]] = parts[1] == "1"
	}

	// Read the wires
	wires := map[string]Ins{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		ins := Ins{
			op1: parts[0],
			op2: parts[2],
			ins: parts[1],
		}
		wires[parts[4]] = ins
	}

	z := 0
	maxBits := 0
	for i := 0; ; i++ {
		tag := fmt.Sprintf("z%02d", i)
		_, ok1 := values[tag]
		_, ok2 := wires[tag]
		if !ok1 && !ok2 {
			maxBits = i
			break
		}
		if calculate(wires, tag, values) {
			z |= 1 << i
		}
	}
	fmt.Printf("z: %d\n", z)

	// Part 2 - fix the adder

	// These are the values I found by semi manual inspection
	swaps := map[string]string{
		"z08": "cdj",
		"z32": "gfm",
		"z16": "mrb",
		"dhm": "qjd",
	}
	badTags := []string{}
	for s1, s2 := range swaps {
		swap(&wires, s1, s2)
		badTags = append(badTags, s1)
		badTags = append(badTags, s2)
	}
	slices.Sort(badTags)
	fmt.Printf("Bad Tags: ")
	for i, tag := range badTags {
		if i != 0 {
			fmt.Print(",")
		}
		fmt.Printf("%s", tag)
	}

	// If the tags above are swapped correctly, then the below code will not crash

	// Each bit follows the following rules (with some special cases at the start/end)
	// bth[n] = x[n] & y[n]
	// one[n] = x[n] ^ y[n]
	// ovf[n] = bth[n] | ext[n]
	// ext[n] = ovf[n-1] & one[n]  // extended overflow
	// z[n] = one[n] ^ ovf[n-1]

	// rename the bth[n] and one[n] tags
	for i := 0; i < maxBits-1; i++ {
		xTag := fmt.Sprintf("x%02d", i)
		yTag := fmt.Sprintf("y%02d", i)
		bthTag := fmt.Sprintf("bth%02d", i)
		oneTag := fmt.Sprintf("one%02d", i)

		foundXor := false
		foundAnd := false
		for wire, ins := range wires {
			if (ins.op1 == xTag && ins.op2 == yTag) || (ins.op2 == xTag && ins.op1 == yTag) {
				if wire[0] == 'z' {
					if i == 0 {
						foundXor = true
						continue
					}
					panic(fmt.Sprintf("Direct assign to z: %d\n", i))
				}
				switch ins.ins {
				case AND:
					replace(&wires, wire, bthTag)
					foundAnd = true
				case XOR:
					replace(&wires, wire, oneTag)
					foundXor = true
				default:
					panic("Should not exist")
				}
			}
		}
		if !foundXor {
			panic(fmt.Sprintf("Did not find xor: %d", i))
		}
		if !foundAnd {
			panic(fmt.Sprintf("Did not find and: %d", i))
		}
	}

	//Rename the ovf and ext tags
	for i := 2; i < maxBits; i++ {
		zTag := fmt.Sprintf("z%02d", i)
		oneTag := fmt.Sprintf("one%02d", i)
		bthTagN1 := fmt.Sprintf("bth%02d", i-1)
		ovfTagN1 := fmt.Sprintf("ovf%02d", i-1)
		extTagN1 := fmt.Sprintf("ext%02d", i-1)

		wire, ok := wires[zTag]
		if !ok {
			panic(fmt.Sprintf("Z tag missing: %d", i))
		} else if wire.op1 == oneTag {
			replace(&wires, wire.op2, ovfTagN1)
		} else if wire.op2 == oneTag {
			replace(&wires, wire.op1, ovfTagN1)
		} else if i != maxBits-1 {
			panic(fmt.Sprintf("Z wire not right: %d", i))
		}
		if wire.ins != XOR && i != maxBits-1 {
			panic(fmt.Sprintf("Z wire not xor: %d", i))
		}

		wire, ok = wires[ovfTagN1]
		if !ok {
			if i == maxBits-1 {
				continue
			}
			panic(fmt.Sprintf("overflow missing: %d", i))
		} else if wire.op1 == bthTagN1 {
			replace(&wires, wire.op2, extTagN1)
		} else if wire.op2 == bthTagN1 {
			replace(&wires, wire.op1, extTagN1)
		} else {
			panic(fmt.Sprintf("overflow wrong: %d", i))
		}
		if wire.ins != OR {
			panic(fmt.Sprintf("overflow bad op: %d", i))
		}
	}

	// print the tags for import into excel
	// for wire, ins := range wires {
	// 	fmt.Printf("%s\t%s\t%s\t->\t%s\n", ins.op1, ins.ins, ins.op2, wire)
	// }
}

func swap(wires *map[string]Ins, w1 string, w2 string) {
	ins1 := (*wires)[w1]
	ins2 := (*wires)[w2]
	(*wires)[w2] = ins1
	(*wires)[w1] = ins2
}

func replace(wires *map[string]Ins, old string, neww string) {
	for wire, ins := range *wires {

		if ins.op1 == old {
			ins.op1 = neww
		}
		if ins.op2 == old {
			ins.op2 = neww
		}
		if wire == old {
			delete(*wires, wire)
			wire = neww
		}
		(*wires)[wire] = ins
	}
}

func calculate(wires map[string]Ins, value string, values map[string]bool) bool {
	if val, ok := values[value]; ok {
		return val
	}

	ins, ok := wires[value]
	if !ok {
		panic("Wire not found")
	}

	v1 := calculate(wires, ins.op1, values)
	v2 := calculate(wires, ins.op2, values)
	var res bool
	switch ins.ins {
	case AND:
		res = v1 && v2
	case OR:
		res = v1 || v2
	case XOR:
		res = (v1 && !v2) || (!v1 && v2)
	default:
		panic("Unknown instruction")
	}
	values[value] = res
	return res
}
