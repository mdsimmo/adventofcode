package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

type Scan struct {
	done []int
	toscan []int
}

func aoc_17_run() {
	file, err := os.Open("aoc_17_in.txt")
	if err != nil {
		panic(err)
	}
	containers := []int {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		volume, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		containers = append(containers, volume)
	}
	
	queue := []Scan {{
		done: []int {},
		toscan: containers,
	}}
	combos := 0
	best := math.MaxInt
	best_combos := 0
	for len(queue) > 0 {
		//fmt.Printf("Queue: %+v\n", queue)
		scan := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		
		fmt.Printf("Testing: %+v\n", scan)

		sum := 0
		for i := 0; i < len(scan.done); i++ {
			sum += scan.done[i]
		}
		if sum == 150 {
			combos++
			fmt.Printf("  Found: %+v\n", scan)
			
			if len(scan.done) < best {
				best = len(scan.done)
				best_combos = 1
			} else if len(scan.done) == best {
				best_combos++
			}
			
		} else if sum > 150 {
			// remove from list
		} else {
			if len(scan.toscan) > 0 {
				for i := 0; i < len(scan.toscan); i++ {
					doneNew := slices.Clone(scan.done)
					toscanNew := slices.Clone(scan.toscan)
					doneNew = append(doneNew, toscanNew[i])
					toscanNew = toscanNew[i+1:]
					
					newscan := Scan {
						done: doneNew,
						toscan: toscanNew,
					}
					queue = append(queue, newscan)
					fmt.Printf("  Append: %+v\n", newscan)
				}
			} else {
				// remove from list
			}
		}
	}
	fmt.Printf("Combos: %d\n", combos)
	fmt.Printf("Best Combos: %d (len: %d)", best_combos, best)
}
