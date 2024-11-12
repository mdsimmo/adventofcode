package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Disc struct {
	slots int
	pos int
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}

	discs := []Disc {}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		slots, _ := strconv.Atoi(parts[3])
		pos, _ := strconv.Atoi(parts[11][:len(parts[11])-1])
		
		discs = append(discs, Disc{ 
			slots: slots,
			pos: pos,
		})
	}

	// part 2
	discs = append(discs, Disc{
		slots: 11,
		pos: 0,
	})

	Outer:
	for i := 0; ; i++ {
		fmt.Printf("Checking turn %d\n", i)
		for j := 0; j < len(discs); j++ {
			pos := (discs[j].pos + j + i + 1) % discs[j].slots
			fmt.Printf("  [%d]Pos: %d\n", j, pos)
			if pos != 0 {
				fmt.Printf("  [%d]Hits\n", j)
				continue Outer 
			}
		}
		fmt.Printf("FOUND CLEAR: %d\n", i)
		return
	}
}
