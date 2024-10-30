package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Deer struct {
	name string
	speed int
	fly int
	rest int
}

func aoc_14_run() {

	file, err := os.Open("aoc_14_in.txt")
	if err != nil {
		panic(err)
	}

	deers := []Deer {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		name := parts[0]
		speed, _ := strconv.Atoi(parts[3])
		fly_time, _ := strconv.Atoi(parts[6])
		rest_time, _ := strconv.Atoi(parts[13])

		deer := Deer {
			name: name,
			speed: speed,
			fly: fly_time,
			rest: rest_time,
		}
		deers = append(deers, deer)

	}

	points := make([]int, len(deers))
	for t := 1; t <= 2503; t++ {
		dists := make([]int, len(deers))
		for i := 0; i < len(deers); i++ {
			deer := deers[i]
			dists[i] = calculate_dist(t, deer.speed, deer.fly, deer.rest)
		}
		best := slices.Max(dists)
		for i := 0; i < len(deers); i++ {
			if dists[i] == best {
				points[i]++
			}
		}
	}
		
	fmt.Printf("points: %+v\n", points)
}

func calculate_dist(time int, speed int, runTime int, restTime int) int {
	cycleTime := runTime + restTime
	cycles := time / cycleTime
	currentCycleTime := time % cycleTime
	allRuntime := cycles * runTime
	if currentCycleTime > runTime {
		allRuntime += runTime
	} else {
		allRuntime += currentCycleTime
	}
	return allRuntime * speed 
}
