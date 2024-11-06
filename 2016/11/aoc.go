package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Entry struct {
	layout [][]int
	floor int
	steps int
}

func (e Entry) Equal(other any) bool {
	o, ok := other.(Entry)
	if !ok {
		return false
	}
	if o.floor != e.floor {
		return false
	}
	// do not include steps in equality for speed
	// should always search entries with lowest steps first, so no point re-adding the same layouts
	//if v_e.steps != e.steps {
	//	return false
	//}
	if len(e.layout) != len(o.layout) {
		return false
	}
	for i, items := range(e.layout) {
		o_items := o.layout[i]
		if len(items) != len(o_items) {
			return false
		}
		for j, item := range(e.layout[i]) {
			if item != o_items[j] {
				return false
			}
		}
	}
	return true
}

func (e Entry) Hash() int64 {
	var hash int64 = 17
	hash = 31 * hash + int64(e.floor)
	//hash = 31 * hash + e.steps
	for _, items := range(e.layout) {
		for _, item := range(items) {
			hash = 31 * hash + int64(item)
		}
	}
	return hash
}

type HashEntry[K Hashable, V any] struct {
	key K
	value V
}

type Hashable interface {
	Hash() int64
	Equal(any) bool
}

type HashMap[K Hashable, V any] struct {
	data map[int64][]HashEntry[K, V]
}

func NewHashMap[K Hashable, V any]() HashMap[K, V] {
	m := HashMap[K, V] {
		data: make(map[int64][]HashEntry[K, V]),
	}
	return m
}

func (m *HashMap[K, V]) Insert(key K, value V) (V, bool) {
	hash := key.Hash()
	hashArray := m.data[hash]

	// insert new hash array
	if hashArray == nil {
		m.data[hash] = []HashEntry[K, V] {{
			key: key,
			value: value,
		}}
		var result V
		return result, false
	}

	// Check if existing value needs to be displaced
	for i, entry := range(hashArray) {
		if key.Equal(entry.key) {
			result := entry.value
			m.data[hash][i].value = value
			return result, true
		}
	}

	// Add to end of array
	m.data[hash] = append(m.data[hash], HashEntry[K, V]{
		key: key,
		value: value,
	})
	var result V
	return result, false
}

func (m *HashMap[K, V]) Delete(key K) (V, bool) {
	hash := key.Hash()
	hashArray := m.data[hash]
	if hashArray == nil {
		var result V
		return result, false
	}

	// Check value exists in array
	for i, entry := range(hashArray) {
		if key.Equal(entry.key) {
			result := entry.value
			if len(hashArray) > 1 {
				hashArray[i] = hashArray[len(hashArray)-1]
				hashArray = hashArray[:len(hashArray)-1]
				m.data[hash] = hashArray
			} else {
				delete(m.data, hash)
			}
			return result, true
		}
	}
	
	// nothing to delete
	var result V
	return result, false
}

func (m *HashMap[K, V]) Get(key K) (V, bool) {
	hash := key.Hash()
	hashArray := m.data[hash]
	if hashArray == nil {
		var result V
		return result, false
	}

	// Check value exists in array
	for _, entry := range(hashArray) {
		if key.Equal(entry.key) {
			return entry.value, true
		}
	}
	
	// could not find it 
	var result V
	return result, false
}

func (m *HashMap[K, V]) Contains(key K) bool {
	_, contains := m.Get(key)
	return contains
}

func (m *HashMap[K, V]) AnyKey() (K, bool) {
	for _, entryArray := range(m.data) {
		if len(entryArray) == 0 {
			panic("Should never have an empty array")
		}
		return entryArray[0].key, true
	}
	// Nothing
	var key K
	return key, false
}

func (m *HashMap[K, V]) IsEmpty() bool {
	return len(m.data) == 0
}

func main() {
	layout := loadLayout()
	diffLayout := layoutToDiffLayout(layout)
	processLayout(diffLayout)
}

func loadLayout() [][]string {
	file, err := os.Open("in2.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	floor := 0
	layout := [][]string {}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' '
		})
		layout = append(layout, make([]string, 0))
		for i := 0; i < len(parts); i++ {
			if strings.HasPrefix(parts[i], "gener") {
				layout[floor] = append(layout[floor], parts[i-1])
			}
			if strings.Contains(parts[i], "compa") {
				layout[floor] = append(layout[floor], parts[i])
			}
		}
		floor++
	}

	fmt.Printf("Layout: %+v\n", layout)
	return layout
}

func layoutToDiffLayout(layout [][]string) [][]int {
	diffLayout := make([][]int, len(layout))
	for i := 0; i < len(diffLayout); i++ {
		diffLayout[i] = make([]int, 0)
	}
	for floor, items := range(layout) {
		itemLoop:
		for _, item := range(items) {
			if !strings.Contains(item, "comp") {
				continue
			}
			genName := strings.Split(item, "-")[0]
			for floor_2, items_2 := range(layout) {
				for _, gen := range(items_2) {
					if gen == genName {
						diffLayout[floor] = append(diffLayout[floor], floor_2-floor)
						continue itemLoop
					}
				}
			}
			panic("Did not find matching item")
		}
	}
	fmt.Printf("Diff Layout: %+v\n", diffLayout)
	return diffLayout
}

func processLayout(startLayout [][]int) {
	startEntry := Entry {
		layout: startLayout,
		floor: 0,
		steps: 0,
	}
	starterMap := NewHashMap[Entry, bool]()
	toScan := map[int]*HashMap[Entry, bool] {
		0: &starterMap,
	}
	toScan[0].Insert(startEntry, true)

	for len(toScan) > 0 {
		bestChance := math.MaxInt 
		for key := range(toScan) {
			if key < bestChance {
				bestChance = key
			}
		}
		scanMap := toScan[bestChance]
		scan, _ := scanMap.AnyKey()
		scanMap.Delete(scan)
		if scanMap.IsEmpty() {
			delete(toScan, bestChance)
		}
		fmt.Printf("Check: %+v\n", scan)

		// See if it is winner (all items on last floor)
		winnerCheck:
		for i, floor := range(scan.layout) {
			if i == len(scan.layout)-1 {
				for _, item := range(scan.layout[i]) {
					if item != 0 {
						break winnerCheck
					}
				}
				fmt.Printf("FOUND SOLUTION: %+v\n", scan)
				return
			}
			if len(floor) > 0 {
				break
			}
		}

		// list of possible items to move (by index)
		permItemMoves := [][][]int {} // list of moves (move = [[]chipmoves[], []genmoves]), chip/gen-move = index of gen/chip move
		genMoves := make([]int, 0) // value of gen items that could move
		chipMoves := make([]int, 0) // value of chip items that could move
		for floor, items := range(scan.layout) {
			for _, item := range(items) {
				// option to move the generators
				if floor + item == scan.floor {
					genMoves = append(genMoves, item)
				}

				// option to move the microchips
				if floor == scan.floor {
					chipMoves = append(chipMoves, item)
				}
			}
		}
		for i := 0; i < len(genMoves) + len(chipMoves); i++ {
			// Put every individual item
			if i >= len(genMoves) {
				permItemMoves = append(permItemMoves, [][]int {{i-len(genMoves)}, {}})
			} else {
				permItemMoves = append(permItemMoves, [][]int {{}, {i}})
			}
			
			// Put every pair
			for j := i+1; j < len(genMoves) + len(chipMoves); j++ {
				if i >= len(genMoves) {
					if j >= len(genMoves) {
						permItemMoves = append(permItemMoves, [][]int {{i-len(genMoves), j-len(genMoves)}, {}})
					} else {
						permItemMoves = append(permItemMoves, [][]int {{i-len(genMoves)}, {j}})
					}
				} else {
					if j >= len(genMoves) {
						permItemMoves = append(permItemMoves, [][]int {{j-len(genMoves)}, {i}})
					} else {
						permItemMoves = append(permItemMoves, [][]int {{}, {i, j}})
					}
				}
			}
		}
		//fmt.Printf("  Options:\n")
		//fmt.Printf("    Gen: %+v\n", genMoves)
		//fmt.Printf("    Chp: %+v\n", chipMoves)
		//fmt.Printf("    Comb: %+v\n", permItemMoves)

		if len(permItemMoves) == 0 {
			panic("Always should be something to move")
		}

		for _, floorDelta := range([]int {-1, 1}) {
			// don't go out of floor levels
			if floorDelta < 0 && scan.floor == 0 {
				continue
			}
			if floorDelta > 0 && scan.floor == len(scan.layout)-1 {
				continue
			}

			for _, perm := range(permItemMoves) {
				//fmt.Printf("  Option: %d, %+v\n", floorDelta, perm)
				// create new layout with selected items moved to new floor
				newLayout := slices.Clone(scan.layout)
				for i := 0; i < len(newLayout); i++ {
					newLayout[i] = slices.Clone(newLayout[i])
				}

				if len(perm) != 2 {
					panic("Expect 2 slots per permutation (gen/chip)")
				}

				newFloor := scan.floor + floorDelta
				
				// move generator and chip together
				moveBoth := false
				for _, iGen := range(perm[1]) {
					for _, iChip := range(perm[0]) {
						if genMoves[iGen] == 0 && chipMoves[iChip] == 0 {
							moveBoth = true
							//fmt.Printf("    Move chip with gen\n")
						}
					}
				}
				for _, iGen := range(perm[1]) {
					if moveBoth {
						continue
					}
					// only need to increment the chip number
					chipFloor := scan.floor - genMoves[iGen]
					//fmt.Printf("    Chip Floor: %d\n", chipFloor)
					for j := 0; j <= len(newLayout[chipFloor]); j++ { // purposefully <= so will panic if not found
						if newLayout[chipFloor][j] == genMoves[iGen] {
							newLayout[chipFloor][j] += floorDelta
							break
						}
					}
				}
				for _, iChip := range(perm[0]) {
					// chip move
					newDist := chipMoves[iChip]-floorDelta
					if moveBoth {
						newDist = 0
					}
					newLayout[newFloor] = append(newLayout[newFloor], newDist) 
					// delete the old chip
					for j := 0; j <= len(newLayout[scan.floor]); j++ { // purposefully <= so will panic if not found
						if newLayout[scan.floor][j] == chipMoves[iChip] {
							newLayout[scan.floor] = append(newLayout[scan.floor][:j], newLayout[scan.floor][j+1:]...)
							break
						}
					}
				}
			
				// sort the floors
				for i := 0; i < len(scan.layout); i++ {
					slices.Sort(scan.layout[i])
				}

				newScan := Entry {
					layout: newLayout,
					floor: scan.floor+floorDelta,
					steps: scan.steps + 1,
				}
			
				if isValidState(newScan) {
					score := guessBestScore(newScan)
					if toScan[score] == nil {
						newMap := NewHashMap[Entry, bool]()
						toScan[score] = &newMap
					}
					_, dup := toScan[score].Insert(newScan, true)
					if dup {
						fmt.Printf("  Dup: (%3d) %+v\n", score, newScan)
					} else {
						fmt.Printf("  Add: (%3d) %+v\n", score, newScan)
					}
				} else {
					fmt.Printf("  Dsc: (   ) %+v\n", newScan)
				}
			}
		}
	}
	panic("Should never happen")
}

func guessBestScore(entry Entry) int {
	// sum items on each floor
	items := make([]int, len(entry.layout))
	lowestFloorWithItems := math.MaxInt
	for floor := 0; floor < len(items); floor++ {
		items[floor] += len(entry.layout[floor])
		for _, item := range(entry.layout[floor]) {
			genFloor := floor + item
			items[genFloor]++
			if genFloor < lowestFloorWithItems {
				lowestFloorWithItems = genFloor
			}
		}
		if items[floor] > 0 && lowestFloorWithItems > floor {
			lowestFloorWithItems = floor
		}
	}

	// complete. Return current step
	if lowestFloorWithItems == len(items)-1 {
		return entry.steps
	}
	
	bestScore := entry.steps
	
	if lowestFloorWithItems > entry.floor {
		panic("Got to a floor below all the items")
	}

	// go to lowest floor
	for floor := entry.floor; floor > lowestFloorWithItems; floor-- {
		items[floor]--
		items[floor-1]++
		bestScore++
	}

	// move all items up, two at a time, but returning one after each move
	for floor := lowestFloorWithItems; floor < len(items)-1; floor++ {
		bestScore += items[floor]*2-3
		items[floor+1] += items[floor]
		items[floor] = 0
	}

	return bestScore
}

func isValidState(entry Entry) bool {
	for floor, items := range(entry.layout) {
		for _, item := range(items) {
			// safe if has generator
			if item == 0 {
				continue
			}

			// Not safe if generator on same level
			for floor2, items2 := range(entry.layout) {
				for _, item2 := range(items2) {
					genFloor := floor2 + item2
					if genFloor == floor {
						return false
					}
				}
			}
		}

	}
	return true
}

