package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Entry struct {
	r rune
	count int
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Inspect: %s\n", line)
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == '-' || r == '[' || r == ']'
		})
		
		var combine strings.Builder
		for i := 0; i < len(parts)-2; i++ {
			combine.WriteString(parts[i])
		}

		sector, err := strconv.Atoi(parts[len(parts)-2])
		if err != nil {
			panic(err)
		}
		expect := parts[len(parts)-1]

		counts := make([]Entry, 26)
		for i := 0; i < 26; i++ {
			counts[i] = Entry{
				r: 'a' + rune(i),
				count: 0,
			}
		}

		for _, r := range(combine.String()) {
			counts[r-'a'].count++
		}

		slices.SortFunc(counts, func(a Entry, b Entry) int {
			if a.count == b.count {
				return int(a.r) - int(b.r)
			} else {
				return b.count-a.count
			}
		})

		var result strings.Builder
		for i := 0; i < 5; i++ {
			result.WriteRune(counts[i].r)
		}
		rtxt := result.String()

		fmt.Printf("  Exp: %s, res: %s\n", expect, rtxt)

		if rtxt == expect {
			sum += sector
		}

		var rotText strings.Builder
		for _, r := range(combine.String()) {
			rotText.WriteRune(rune((int(r-'a')+sector) % 26) + 'a')
		}
		fmt.Printf("  Decrypt: %s\n", rotText.String())
	}

	fmt.Printf("Sum: %d\n", sum)
}
