package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	//offset := 1
	offset := len(line)/2
	sum := 0
	for i, c := range(line) {
		if c == rune(line[(i+offset)%len(line)]) {
			sum += int(c - '0')
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}
