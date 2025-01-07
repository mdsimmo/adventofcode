package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	comps := map[string]map[string]bool{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		c1 := parts[0]
		c2 := parts[1]
		if comps[c1] == nil {
			comps[c1] = make(map[string]bool)
		}
		if comps[c2] == nil {
			comps[c2] = make(map[string]bool)
		}

		comps[c1][c2] = true
		comps[c2][c1] = true
	}

	// Find the tri-sets (part 1)
	triSets := map[string]map[string]bool{}
	for comp, others := range comps {
		for other := range others {
			for otherComp := range comps[other] {
				if others[otherComp] {
					blob := map[string]bool{
						comp:      true,
						other:     true,
						otherComp: true,
					}
					triSets[name(blob)] = blob
				}
			}
		}
	}
	fmt.Printf("Tri-sets: %d\n", len(triSets))
	tTriSets := 0
	for set := range triSets {
		if set[0] == 't' || strings.Contains(set, ",t") {
			tTriSets++
		}
	}
	fmt.Printf("T Tri-sets: %d\n", tTriSets)

	// Find the largest set (part 2)
	// Start with the tri-sets, then keep adding computers to the blobs until no more can be grown
	// This is slow - takes a couple minutes
	blobs := maps.Clone(triSets)
	change := true
	for change {
		change = false
		for comp, connections := range comps {
		blobLoop:
			for _, blob := range blobs {
				if blob[comp] {
					continue
				}
				for blobComp := range blob {
					if !connections[blobComp] {
						continue blobLoop
					}
				}
				// Add a duplicate of the blob, as there might
				// be a way to build a bigger blob without the selected comp
				newBlob := maps.Clone(blob)
				newBlob[comp] = true
				name := name(newBlob)

				// dont flag re-added if the new blob already exists
				if _, exists := blobs[name]; exists {
					continue
				}
				blobs[name] = newBlob
				change = true
			}
		}
	}
	largest := 0
	largestBlob := ""
	for name, blob := range blobs {
		if len(blob) > largest {
			largest = len(blob)
			largestBlob = name
		}
	}
	fmt.Printf("Largest: %s\n", largestBlob)
}

func name(comps map[string]bool) string {
	compsArray := make([]string, len(comps))
	i := 0
	for c := range comps {
		compsArray[i] = c
		i++
	}
	slices.Sort(compsArray)

	var name strings.Builder
	for i, c := range compsArray {
		if i != 0 {
			name.WriteRune(',')
		}
		name.WriteString(c)
	}
	return name.String()
}
