package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	locks := [][]int{}
	keys := [][]int{}
	scanner := bufio.NewScanner(file)
	object := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			object = append(object, line)
		} else {
			lock, obj := parseObject(object)
			if lock {
				locks = append(locks, obj)
			} else {
				keys = append(keys, obj)
			}
			object = []string{}
		}
	}
	lock, obj := parseObject(object)
	if lock {
		locks = append(locks, obj)
	} else {
		keys = append(keys, obj)
	}

	fmt.Printf("Keys: %+v\n", keys)
	fmt.Printf("Locks: %+v\n", locks)

	fits := 0
	for _, lock := range locks {
		for _, key := range keys {
			if isMatch(key, lock) {
				fits++
			}
		}
	}
	fmt.Printf("Fits: %d\n", fits)
}

func isMatch(key []int, lock []int) bool {
	height := 6
	for i := 0; i < len(key); i++ {
		if key[i]+lock[i] >= height {
			return false
		}
	}
	return true
}

func parseObject(lines []string) (bool, []int) {
	height := len(lines)
	obj := make([]int, len(lines[0]))
	for i := 0; i < len(lines[0]); i++ {
		sum := 0
		for j := 0; j < height; j++ {
			if lines[j][i] == '#' {
				sum++
			}
		}
		obj[i] = sum - 1
	}
	return lines[0][0] == '#', obj
}
