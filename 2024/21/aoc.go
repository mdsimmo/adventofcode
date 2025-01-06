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
	keypads := 2
	if err != nil {
		panic(err)
	}

	// Keypads (0,0) is top left
	numberKeypad := map[rune]Point{
		'!': {0, 3},
		'0': {1, 3},
		'A': {2, 3},
	}
	for i := 1; i <= 9; i++ {
		numberKeypad[rune('0'+i)] = Point{
			x: (i - 1) % 3,
			y: (9 - i) / 3,
		}
	}

	directionKeypad := map[rune]Point{
		'!': {0, 0},
		'^': {1, 0},
		'A': {2, 0},
		'<': {0, 1},
		'v': {1, 1},
		'>': {2, 1},
	}
	dirLookup := map[Point]rune{
		{-1, 0}: '<',
		{1, 0}:  '>',
		{0, 1}:  'v',
		{0, -1}: '^',
	}

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		code := []rune(scanner.Text())
		fmt.Printf("Code: %s\n", string(code))
		paths := make([]map[string]bool, keypads+1)

		for i := 0; i < len(code); i++ {
			start := 'A'
			if i > 0 {
				start = code[i-1]
			}
			end := code[i]
			charPaths := findPaths(numberKeypad, dirLookup, start, end)
			// fmt.Printf("  CP [%c -> %c]: %+v\n", start, end, charPaths)
			newPaths := map[string]bool{}
			if len(paths[0]) == 0 {
				newPaths = charPaths
			} else if len(charPaths) == 0 {
				newPaths = paths[0]
			} else {
				for cp := range charPaths {
					for op := range paths[0] {
						newPaths[op+cp] = true
					}
				}
			}
			// append press "A" after each move
			newPaths2 := make(map[string]bool)
			for p := range newPaths {
				newPaths2[p+"A"] = true
			}
			paths[0] = newPaths2
			// fmt.Printf("   [0] Paths: %v\n", paths[0])
		}
		fmt.Printf("[0] Paths: %v\n", paths[0])

		for pad := 1; pad <= keypads; pad++ {
			paths[pad] = map[string]bool{}
			for path := range paths[pad-1] {
				possiblePaths := map[string]bool{}
				// fmt.Printf("[%d] Finding paths for: %s\n", pad, path)
				for i := 0; i < len(path); i++ {
					start := 'A'
					if i > 0 {
						start = rune(path[i-1])
					}
					end := rune(path[i])
					charPaths := findPaths(directionKeypad, dirLookup, start, end)
					// fmt.Printf("  CP [%c -> %c]: %+v\n", start, end, charPaths)
					newPaths := map[string]bool{}
					if len(possiblePaths) == 0 {
						newPaths = charPaths
					} else if len(charPaths) == 0 {
						newPaths = possiblePaths
					} else {
						for cp := range charPaths {
							for op := range possiblePaths {
								newPaths[op+cp] = true
							}
						}
					}
					// press A
					newPaths2 := make(map[string]bool)
					for p := range newPaths {
						newPaths2[p+"A"] = true
					}
					possiblePaths = newPaths2
					// fmt.Printf("   [%dA] Paths: %v\n", pad, possiblePaths)
				}
				// fmt.Printf("[%d] Paths for %s: %v\n", pad, path, possiblePaths)
				for k, v := range possiblePaths {
					paths[pad][k] = v
				}
			}

			// keep only shortest of each path
			// fmt.Printf("[%d] Paths: %v\n", pad, paths[pad])
			shortest := math.MaxInt
			for k := range paths[pad] {
				if len(k) < shortest {
					shortest = len(k)
				}
			}
			for k := range paths[pad] {
				if len(k) > shortest {
					delete(paths[pad], k)
				}
			}
			fmt.Printf("[%d] Paths: (%d) %v\n", pad, shortest, paths[pad])
		}

		shortest := math.MaxInt
		for k := range paths[keypads] {
			if len(k) < shortest {
				shortest = len(k)
			}
		}
		numPart, err := strconv.Atoi(strings.FieldsFunc(string(code), func(r rune) bool {
			return r < '0' || r > '9'
		})[0])
		if err != nil {
			panic(err)
		}
		complexity := shortest * numPart
		fmt.Printf("Complexity: %d*%d=%d\n", shortest, numPart, complexity)
		sum += complexity
	}
	fmt.Printf("Sum: %d\n", sum)
}

func findPaths(keypad map[rune]Point, dirLookup map[Point]rune, start rune, end rune) map[string]bool {

	type Entry struct {
		pos  Point
		path string
	}
	startPoint := keypad[start]
	endPoint := keypad[end]
	// fmt.Printf("  Find Paths: %+v -> %+v\n", startPoint, endPoint)
	bad := keypad['!']
	step := Point{x: Signum(endPoint.x - startPoint.x), y: Signum(endPoint.y - startPoint.y)}
	stepx := Point{step.x, 0}
	stepy := Point{0, step.y}

	possiblePaths := []Entry{{
		pos:  startPoint,
		path: "",
	}}
	for i := 0; i < startPoint.DistGrid(endPoint); i++ {
		newPaths := make([]Entry, 0)
		for _, e := range possiblePaths {
			if e.pos.x != endPoint.x {
				newPos := e.pos.Add(stepx)
				if newPos != bad {
					newPaths = append(newPaths, Entry{
						pos:  newPos,
						path: e.path + string(dirLookup[stepx]),
					})
				}
			}
			if e.pos.y != endPoint.y {
				newPos := e.pos.Add(stepy)
				if newPos != bad {
					newPaths = append(newPaths, Entry{
						pos:  newPos,
						path: e.path + string(dirLookup[stepy]),
					})
				}
			}
		}
		possiblePaths = newPaths
	}
	paths := make(map[string]bool)
	for _, e := range possiblePaths {
		paths[e.path] = true
		if e.pos != endPoint {
			panic("Found path that did not finish")
		}
	}
	return paths
}
