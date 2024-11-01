package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SearchIndex struct {
	been []string
	togo []string
}

func aoc_9_run() {
	file, err := os.Open("aoc_9_in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	places := make(map[string] map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		from := parts[0]
		to := parts[2]
		dist, e := strconv.ParseInt(parts[4], 10, 32)
		if e != nil {
			panic(e)
		}
		if places[from] == nil {
			places[from] = make(map[string]int)
		}
		if places[to] == nil {
			places[to] = make(map[string]int)
		}
		places[from][to] = int(dist)
		places[to][from] = int(dist)	
	}

	str, _ := json.MarshalIndent(places, "", " ")
	fmt.Println(string(str))

	searches := []SearchIndex {}
	for key, value := range(places) {
		been := key
		needtogo := []string {}
		for place := range(value) {
			needtogo = append(needtogo, place)
		}
		searches = append(searches, SearchIndex{
			been: []string { been },
			togo: needtogo,
		})
	}

	fmt.Printf("%+v\n", searches)


	best := 10000000
	bestpath := make([]string, 0)
	worst := 0
	for len(searches) > 0 { 
		search := searches[len(searches)-1]
		searches = searches[:len(searches)-1]
		if len(search.togo) != 0 {
			for i := 0; i < len(search.togo); i++ {
				// copy 'been' and append the indexed place
				been := make([]string, len(search.been))
				copy(been, search.been)
				been = append(been, search.togo[i])

				// copy togo and remove the indexed place
				togo := make([]string, len(search.togo))
				copy(togo, search.togo)
				togo[i] = togo[len(togo)-1]
				togo = togo[:len(togo)-1]

				searches = append(searches, SearchIndex{
					been,
					togo,
				})
			}
		} else {
			totcost := 0
			for i := 0; i < len(search.been)-1; i++ {
				a := search.been[i]
				b := search.been[i+1]
				cost := places[a][b]
				totcost += cost
			}
			fmt.Printf("%+v: %d\n", search, totcost)

			if best > totcost {
				best = totcost
				bestpath = search.been
			}
			if worst < totcost {
				worst = totcost
			}
		}
		fmt.Printf("%+v\n", searches)
	}
	
	fmt.Printf("Best: %+v: %d\n", bestpath, best)
	fmt.Printf("Worst: %+v: %d\n", bestpath, worst)
}
