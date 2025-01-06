package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cache struct {
	pads int
	path string
}

func main() {
	file, err := os.Open("in.txt")
	keypads := 25
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

	cache := map[Cache]int{}
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Find possible paths to write number
		code := []rune(scanner.Text())
		fmt.Printf("Code: %s\n", string(code))
		paths := make([]map[string]bool, len(code))
		for i := 0; i < len(code); i++ {
			start := 'A'
			if i > 0 {
				start = code[i-1]
			}
			end := code[i]
			charPaths := keypadPath(numberKeypad, dirLookup, start, end)
			paths[i] = charPaths
		}
		// fmt.Printf("Paths: %v\n", paths)

		// Find shortest controlling path
		lengthSum := 0
		for _, options := range paths {
			shortest := math.MaxInt
			for option := range options {
				length := recursivePath(directionKeypad, dirLookup, option, keypads, cache)
				if length < shortest {
					shortest = length
				}
			}
			lengthSum += shortest
		}

		// Complexity score
		numPart, err := strconv.Atoi(strings.FieldsFunc(string(code), func(r rune) bool {
			return r < '0' || r > '9'
		})[0])
		if err != nil {
			panic(err)
		}
		complexity := lengthSum * numPart
		fmt.Printf("Complexity: %d*%d=%d\n", lengthSum, numPart, complexity)
		sum += complexity
	}
	fmt.Printf("Sum: %d\n", sum)
}

// Gets the shortest path to control the given path using the set number of keyboards
func recursivePath(keypad map[rune]Point, dirLookup map[Point]rune, path string, keypads int, cache map[Cache]int) int {
	// fmt.Printf("%s Inspect %s\n", indent(keypads), path)
	// No more keypad indirections, directly type the path
	if keypads == 0 {
		// fmt.Printf("%s  Length: %d\n", indent(keypads), len(path))
		return len(path)
	}

	// Cache results (very much needed for performance)
	cacheResult, cached := cache[Cache{path: path, pads: keypads}]
	if cached {
		// fmt.Printf("%s  Length: %d (Cahced)\n", indent(keypads), cacheResult)
		return cacheResult
	}

	// Recusively calculate how many buttons are required
	sum := 0
	for i, r := range path {
		start := 'A'
		if i != 0 {
			start = rune(path[i-1])
		}
		end := r

		options := keypadPath(keypad, dirLookup, start, end)
		// fmt.Printf("%s   Paths %c -> %c: %+v\n", indent(keypads), start, end, options)

		shortest := math.MaxInt
		for option := range options {
			length := recursivePath(keypad, dirLookup, option, keypads-1, cache)
			if length < shortest {
				shortest = length
			}
		}
		// fmt.Printf("%s   Shortest: %d\n", indent(keypads), shortest)
		sum += shortest
	}
	// fmt.Printf("%s  Length: %d\n", indent(keypads), sum)
	cache[Cache{pads: keypads, path: path}] = sum
	return sum
}

// For debug: indents start of log statements
func indent(n int) string {
	return fmt.Sprintf("[%"+strconv.Itoa(25-n)+"d]", n)
}

// Keys the direction instructions needed to type the two digits
// Returns all shortest paths
func keypadPath(keypad map[rune]Point, dirLookup map[Point]rune, start rune, end rune) map[string]bool {

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
		// Add A to end of all paths
		paths[e.path+"A"] = true

		if e.pos != endPoint {
			panic("Found path that did not finish")
		}
	}
	return paths
}
