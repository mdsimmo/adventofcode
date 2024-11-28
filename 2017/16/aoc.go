package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("in.txt")
	if err != nil {
		panic(err)
	}
	contents := strings.Trim(string(file), "\n")
	dancers := ""
	for i := 'a'; i <= 'p'; i++ {
		dancers += string(i)
	}
	parts := strings.Split(contents, ",")

	// Part 1
	dancersLoop := dance(dancers, parts)
	fmt.Printf("Dance %10d: %s\n", 1, dancersLoop)


	// Part 2 - do it a billion times
	done := map[string]bool { dancers: true }
	loopsAt := 0
	dancersLoop = dancers
	for i := 0; ; i++ {
		dancersLoop = dance(dancersLoop, parts)
		if done[dancersLoop] && loopsAt == 0 {
			loopsAt = i
			break
		}
		done[dancersLoop] = true
	}
	
	maxLoops := 1_000_000_000
	dancersLoop = dancers
	for i := 0; i < maxLoops % (loopsAt+1); i++ {
		dancersLoop = dance(dancersLoop, parts)
	}
	fmt.Printf("Dance %10d: %s\n", maxLoops, dancersLoop)
}

func dance(dancers string, parts []string) string {
	for _, part := range parts {
		switch part[0] {
		case 's':
			val, e := strconv.Atoi(part[1:])
			if e != nil {
				panic(e)
			}
			dancers = dancers[len(dancers)-val:] + dancers[:len(dancers)-val]
		case 'x':
			numParts := strings.FieldsFunc(part, func(r rune) bool {
				return r < '0' || r > '9'
			})
			nums := make([]int, len(numParts))
			for i, np := range numParts {
				val, e := strconv.Atoi(np)
				if e != nil {
					panic(e)
				}
				nums[i] = val
			}
			if nums[1] < nums[0] {
				a := nums[1]
				nums[1] = nums[0]
				nums[0] = a
			}
			c0 := dancers[nums[0]]
			c1 := dancers[nums[1]]
			dancers = dancers[:nums[0]] + string(c1) + dancers[nums[0]+1:nums[1]] + string(c0) + dancers[nums[1]+1:]
		case 'p':
			spots := strings.Split(part[1:], "/")
			jLoop:
			for i := 0; i < len(dancers); i++ {
				for j := 0; j < 2; j++ {
					if dancers[i] == spots[j][0] {
						dancers = dancers[:i] + spots[(j+1)%2] + dancers[i+1:]
						continue jLoop
					}
				}
			}

		default:
			panic("Unknown instruction")
		}
		//fmt.Printf("%s\n", part)
		//fmt.Printf("  %s\n", dancers)
	}

	return dancers

}
