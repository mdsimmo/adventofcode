package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}

	itemsLeft := []int {}
	itemsRight := []int {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' '
		})
		val1, e := strconv.Atoi(parts[0])
		if e != nil {
			panic(e)
		}
		val2, e := strconv.Atoi(parts[1])
		if e != nil {
			panic(e)
		}
		
		itemsLeft = append(itemsLeft, val1)
		itemsRight = append(itemsRight, val2)
	}

	slices.Sort(itemsLeft)
	slices.Sort(itemsRight)

	sum := 0
	for i := 0; i < len(itemsLeft); i++ {
		sum += Abs(itemsRight[i] - itemsLeft[i])
	}
	fmt.Printf("Sum: %d\n", sum)

	similarity := 0
	for i := 0; i < len(itemsLeft); i++ {
		for j := 0; j < len(itemsRight); j++ {
			if itemsLeft[i] == itemsRight[j] {
				similarity += itemsLeft[i]
			}
		}
	}
	fmt.Printf("Similarity: %d\n", similarity)
}

func Abs(n int) int {
	if n > 0 {
		return n
	} else {
		return -n
	}
}
