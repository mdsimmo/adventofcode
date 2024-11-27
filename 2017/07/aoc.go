package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tower struct {
	weight int
	above []string
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}

	towers := map[string]Tower {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return (r < 'a' || r > 'z') && (r < '0' || r > '9')
		})
		name := parts[0]
		weight, e := strconv.Atoi(parts[1])
		if e != nil {
			panic(e)
		}
		tower := Tower {
			weight: weight,
			above: parts[2:],
		}
		towers[name] = tower
	}

	// find the root tower
	var root string 
	for name := range towers {
		root = name
		break
	}
	lowerFound := true
	for lowerFound {
		lowerFound = false
		for name, tower := range(towers) {
			for _, nameAbove := range(tower.above) {
				if root == nameAbove {
					root = name
					lowerFound = true
				}
			}
		}
	}

	fmt.Printf("Bottom: %s\n", root)
	
	checkLeaf(root, towers)
}

func calcWeight(name string, towers map[string]Tower) (int, bool) {
	tower := towers[name]
	weight := tower.weight
	balanced := true
	firstWeigth := -1
	for _, name := range(tower.above) {
		leafWeight, leafBalanced := calcWeight(name, towers)
		weight += leafWeight
		if !leafBalanced {
			balanced = false
		}
		if firstWeigth == -1 {
			firstWeigth = leafWeight
		}
		if leafWeight != firstWeigth {
			balanced = false
		}
	}
	return weight, balanced
}

func checkLeaf(name string, towers map[string]Tower) {
	tower := towers[name]
	if len(tower.above) == 0 {
		//fmt.Printf("%s is leaf\n", name)
		return
	}
	_, balanced := calcWeight(name, towers)
	if balanced {
		//fmt.Printf("%s is balanced\n", name)
		return
	}
	//fmt.Printf("Digging in: %s\n", name)

	for _, leafName := range(tower.above) {
		//fmt.Printf("  Check sub leaf: %s\n", leafName)
		_, leafBalanced := calcWeight(leafName, towers)
		if leafBalanced {
			//fmt.Printf("    Balanced (%d)\n", leafWeight)
		} else {
			checkLeaf(leafName, towers)
			return
		}
	}
	fmt.Printf("%s has problem child \n", name)
	for i, sibling := range(tower.above) {
		sibWeight, _ := calcWeight(sibling, towers)
		ok := false
		for j, sib2 := range(tower.above) {
			if i == j {
				continue
			}
			sib2Weight, _ := calcWeight(sib2, towers)
			if sib2Weight == sibWeight {
				ok = true
				break
			}
		}
		fmt.Printf("Sibling: %s: %d kg, total %d kg, Good: %t\n", sibling, towers[sibling].weight, sibWeight, ok)

		if !ok {
			sib2 := tower.above[(i+1)%len(sibling)]
			sib2Weight, _ := calcWeight(sib2, towers)
			sibWeight, _ := calcWeight(sibling, towers)
			required := towers[sibling].weight + sib2Weight - sibWeight
			fmt.Printf("  REQUIRED WEIGHT: %d\n", required)
		}
	}

	// find the child that does not match
	// from problem, we know there are at least 3 siblings (or else two answers would be possible
}
