package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}

	stones := map[int]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		for _, p := range parts {
			val, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}
			stones[val]++
		}
	}

	countStones(25, stones)
	countStones(75, stones)
}

func countStones(iterations int, stones map[int]int) {

	for i := 0; i < iterations; i++ {
		// fmt.Printf("%d %d\n", i, len(stones))
		newStones := map[int]int{}
		for s, n := range stones {
			// fmt.Printf("  [%d] Process: %d\n", j, s)
			if s == 0 {
				newStones[1] += n
				// fmt.Printf("    Made 1\n")
				continue
			}

			digs := digits(s)
			if digs%2 == 0 {
				split := int(math.Pow10(digs / 2))
				s1 := s / split
				s2 := s % split
				newStones[s1] += n
				newStones[s2] += n
				// fmt.Printf("    Insert %d, %d\n", s1, s2)
				continue
			}

			newStones[s*2024] += n
			// fmt.Printf("    *= 2024\n")
		}
		stones = newStones
		// fmt.Printf("Stones: %+v\n", stones)
	}

	sum := 0
	for _, n := range stones {
		sum += n
	}
	fmt.Printf("Stones: %d\n", sum)
}

// This is the nieve way to do it. It is waaayyyyy too slow for part 2
func main_slow() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}

	stones := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		for _, p := range parts {
			val, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}
			stones = append(stones, val)
		}
	}

	for i := 0; i < 75; i++ {
		fmt.Printf("%d %d\n", i, len(stones))
		newStones := make([]int, len(stones)*2)
		jNew := 0
		for j := 0; j < len(stones); j++ {
			s := stones[j]
			// fmt.Printf("  [%d] Process: %d\n", j, s)
			if s == 0 {
				newStones[jNew] = 1
				// fmt.Printf("    Made 1\n")
				jNew++
				continue
			}

			digs := digits(s)
			if digs%2 == 0 {
				split := int(math.Pow10(digs / 2))
				s1 := s / split
				s2 := s % split
				newStones[jNew] = s1
				newStones[jNew+1] = s2
				jNew += 2 // don't split the added stone
				// fmt.Printf("    Insert %d, %d\n", s1, s2)
				continue
			}

			newStones[jNew] = s * 2024
			jNew++
			// fmt.Printf("    *= 2024\n")
		}
		stones = newStones[:jNew]
		// fmt.Printf("Stones: %+v\n", stones)
	}

	fmt.Printf("Stones: %d\n", len(stones))
}

func digits(n int) int {
	return int(math.Ceil(math.Log10(float64(n + 1))))
}
