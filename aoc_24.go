package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

func aoc_24_run() {
	file, err := os.Open("aoc_24_in.txt")
	if err != nil {
		panic(err)
	}
	weights := []int {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		w, _ := strconv.Atoi(line)
		weights = append(weights, w)
	}


	total := sumWeights(weights)
	target := total / 4

	fmt.Printf("Total: %d, Target:  %d\n", total, target)
	
	pm := perms(weights, make([]int, 0), target)
	fmt.Printf("Found %d permutations\n", len(pm))

	var bestEntangle int64 = math.MaxInt64
	bestPackage := math.MaxInt
	bestPack := []int {}
	for _, perm := range(pm) {
		pack := len(perm)
		var entangle int64 = 1
		for _, w := range(perm) {
			entangle *= int64(w)
		}
		
		if bestPackage > pack {
			bestPackage = pack
			bestEntangle = entangle
			bestPack = perm
		} else if bestPackage < pack {
			continue
		}
		if entangle < bestEntangle {
			bestEntangle = entangle
			bestPack = perm
		}
	}
	fmt.Printf("Best: %+v\n\nPacks: %d,\nEntange: %d\n", bestPack, bestPackage, bestEntangle)

}

func perms(weights []int, existing []int, target int) [][]int {
	pr := [][]int {}
	//fmt.Printf("Perms: %+v w, %+v e\n", weights, existing)
	for i := len(weights)-1; i >=0; i-- {
		//fmt.Printf(" Weight (%d) to %+v\n", weights[i], existing)
		newExisting := slices.Clone(existing)
		newExisting = append(newExisting, weights[i])
		total := sumWeights(newExisting)
		if total > target {
			//fmt.Printf("  Discarded\n")
			continue
		}
		if total == target {
			pr = append(pr, newExisting)
			//fmt.Printf("  Perfect %+v\n", newExisting)
		} else {
			extraPerms := perms(weights[i+1:], newExisting, target)
			pr = append(pr, extraPerms...)
			//fmt.Printf("  Extras %+v\n", extraPerms)
		}
		
	}
	// fmt.Printf("Combined to %+v\n", pr)
	return pr
}

func sumWeights(weights []int) int {
	weight := 0
	for _, w := range(weights) {
		weight += w
	}
	return weight
}

