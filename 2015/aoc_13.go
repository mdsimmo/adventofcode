package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type SearchIndex2 struct {
	done []string
	todo []string
}

func aoc_13_run() {
	// Parse the input scores
	scores := make(map[string]map[string]int)
	file, err := os.Open("aoc_13_in.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewScanner(file)

	for reader.Scan() {
		line := reader.Text()
		parts := strings.Split(line, " ")
		name1 := string(parts[0][0])
		name2 := string(parts[len(parts)-1][0])
		gain := parts[2] == "gain"
		score, err := strconv.ParseInt(parts[3], 10, 32)
		if err != nil {
			panic(err)
		}
		if !gain {
			score *= -1
		}

		if scores[name1] == nil {
			scores[name1] = make(map[string]int)
		}
		scores[name1][name2] = int(score)
	}

	// Add self
	if true {
		scores["y"] = make(map[string]int)
		for key := range(scores) {
			scores["y"][key] = 0
			scores[key]["y"] = 0
		}
		fmt.Printf("%+v\n", scores)
	}


	// Find optimal solution
	searches := []SearchIndex2 {}
	best := math.MinInt
	bestpath := []string {}
	
	firstSearch := SearchIndex2 {
		done: make([]string, 0),
		todo: make([]string, 0),
	}
	for key := range(scores) {
		firstSearch.todo = append(firstSearch.todo, key)
	}
	searches = append(searches, firstSearch)
	//fmt.Printf("Searches: %+v\n", searches)

	for len(searches) > 0 {
		search := searches[len(searches)-1]
		searches = searches[:len(searches)-1]
		if len(search.todo) > 0 {
			for i, person := range(search.todo) {
				newdone := make([]string, len(search.done))
				newtodo := make([]string, len(search.todo))
				copy(newdone, search.done)
				copy(newtodo, search.todo)
				newdone = append(newdone, person)
				newtodo[i] = newtodo[len(newtodo)-1]
				newtodo = newtodo[:len(newtodo)-1]
				searches = append(searches, SearchIndex2{
					done: newdone,
					todo: newtodo,
				})
			}
		} else {
			cost := 0
			for i := 0; i < len(search.done); i++ {
				cost += scores[search.done[i]][search.done[(i+1)%len(search.done)]]
				cost += scores[search.done[i]][search.done[(i-1+len(search.done))%len(search.done)]]
			}
			//fmt.Printf("Checking: %d, %+v\n", cost, search.done)
			if cost > best {
				best = cost
				bestpath = search.done
				fmt.Printf("Found better: %d, %+v\n", best, search.done)
			}
		}
		//fmt.Printf("Searches: %+v\n", searches)
	}

	fmt.Printf("Best: %d, %+v", best, bestpath)
}
