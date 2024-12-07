package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Con struct {
	a int
	b int
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	cons := []Con {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "/")
		var con Con 
		con.a, err = strconv.Atoi(parts[0])
		con.b, err = strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		cons = append(cons, con)
	}

	type Entry struct {
		chain []Con
		available []Con
		open int
	}

	todo := []Entry {{
		chain: []Con {},
		available: cons,
		open: 0,
	}}
	bestScore := 0
	longest := 0
	longestBest := 0
	for len(todo) > 0 {
		e := todo[len(todo)-1]
		todo = todo[:len(todo)-1]

		anyAdded := false
		for i, con := range e.available {
			var newOpen int
			if con.a == e.open {
				newOpen = con.b
			} else if con.b == e.open {
				newOpen = con.a
			} else {
				continue
			}
			newE := Entry {
				chain: slices.Clone(e.chain),
				available: slices.Clone(e.available),
				open: newOpen,
			}
			newE.chain = append(newE.chain, con)
			newE.available = append(newE.available[:i], newE.available[i+1:]...)
			todo = append(todo, newE)
			anyAdded = true
		}

		if !anyAdded {
			score := 0
			length := 0
			for _, con := range e.chain {
				score += con.a
				score += con.b
				length++
			}
			if score > bestScore {
				bestScore = score
			}
			if length > longest {
				longest = length
				longestBest = score
			}
			continue
		}
	}

	fmt.Printf("Strongest: %d\n", bestScore)
	fmt.Printf("Longest: %d\n", longestBest)
}
