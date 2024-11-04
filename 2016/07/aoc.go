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
	tslCount := 0
	sslCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Check: %s\n", line)
		hasAbba := false
		badAbba := false
		inside := false
		abas := make([]string, 0)
		babs := make([]string, 0)
		for i, r := range(line) {
			if r == '[' {
				inside = true
				continue
			}
			if r == ']' {
				inside = false
				continue
			}
			if i >= len(line)-2 {
				break
			}
			c0 := line[i]
			c1 := line[i+1]
			c2 := line[i+2]

			if i < len(line)-4 {
				c3 := line[i+3]

				abba := c0 == c3 && c1 == c2 && c0 != c1
				if abba {
					if inside {
						fmt.Printf("  Inner Abba, outer cycle\n")
						badAbba = true
					} else {
						fmt.Printf("  Good Abba\n")
						hasAbba = true	
					}
				}
			}
			aba := c0 == c2 && c0 != c1 && c1 != '[' && c1 != ']'
			if aba {
				if inside {
					babs = append(babs, line[i:i+3])
					fmt.Printf("  Found bab: %s\n", line[i:i+3])
				} else {
					abas = append(abas, line[i:i+3])
					fmt.Printf("  Found aba: %s\n", line[i:i+3])
				}
			}
		}
		if hasAbba && !badAbba {
			fmt.Printf(" Abba found\n")
			tslCount++
		}

		outer:
		for _, aba := range(abas) {
			aba_inv := string(aba[1]) + string(aba[0]) + string(aba[1])
			for _, bab := range(babs) {
				if aba_inv == bab {
					sslCount++
					fmt.Printf(" Count SSL\n")
					break outer
				}
			}
		}
	}
	fmt.Printf("TSL count: %d\n", tslCount)
	fmt.Printf("SSL count: %d\n", sslCount)
}
