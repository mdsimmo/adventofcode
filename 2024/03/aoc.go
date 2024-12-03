package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	enabled := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		
		rx, e := regexp.Compile("(mul\\(\\d+\\,\\d+\\)|do\\(\\)|don't\\(\\))")
		if e != nil {
			panic(e)
		}
		results := rx.FindAllString(line, -1)
		for _, res := range results {
			if res == "do()" {
				enabled = true
				continue
			}
			if res == "don't()" {
				enabled = false
				continue
			}
			if !enabled {
				continue
			}
			parts := strings.FieldsFunc(res, func(r rune) bool {
				return r < '0' || r > '9'
			})
			if len(parts) != 2 {
				panic("Should always be two numbers")
			}
			mul := 1
			for _, p := range parts {
				val, e := strconv.Atoi(p)
				if e != nil {
					panic(e)
				}
				mul *= val
			}
			sum += mul
		}
	}


	fmt.Printf("Sum: %d\n", sum)
}
