package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func aoc_16_run() {
	file, err := os.Open("aoc_16_in.txt")	
	if err != nil {
		panic(err)
	}

	file_match, err := os.Open("aoc_16_in2.txt")
	if err != nil {
		panic(err)
	}

	matches := make(map[string] int)
	scanner := bufio.NewScanner(file_match)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		matches[strings.ReplaceAll(parts[0], ":", "")], _ = strconv.Atoi(parts[1])
	}

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.FieldsFunc(line, Split)
		
		key := ""
		index := 2
		bad := false
		fmt.Printf("Viewing: %s\n", line)
		for index < len(parts) && !bad {
			part := parts[index]
			index++
			if len(part) == 0 {
				continue
			}
			val, err := strconv.Atoi(part)
			if err != nil {
				key = part
			} else {
				need := matches[key]
				fmt.Printf("Checking: %s, %d, %d\n", key, val, need)
				switch key {
				case "cats", "trees":
					bad = val < need
				case "pomeranians", "goldfish":
					bad = val > need
				default:
					bad = val != need
				}
				fmt.Printf("Bad: %t\n", bad)
			}
			if bad {
				break
			}
		}
		if !bad {
			fmt.Printf("Found: %s\n", line)
		}

	}
	fmt.Println("Complete")

}

func Split(r rune) bool {
	return r == ' ' || r == ',' || r == ':'
}
