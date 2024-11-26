package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileData, err := os.ReadFile("in.txt")
	if err != nil {
		panic(err)
	}
	parts := strings.Split(string(fileData), "\t")
	values := make([]int, len(parts))
	for i, part := range(parts) {
		v, e := strconv.Atoi(strings.Trim(part, "\n"))
		if e != nil {
			panic(e)
		}
		values[i] = v
	}
	
	fmt.Printf("%+v\n", values)
	searched := map[string]int {mapToInt(values): 0}
	steps := 0
	for {
		// Find max value
		maxVal := 0
		maxIndex := 0
		for i := 0; i < len(values); i++ {
			if values[i] > maxVal {
				maxVal = values[i]
				maxIndex = i
			}
		}

		// Redistriute
		values[maxIndex] = 0
		for i := 0; i < maxVal; i++ {
			values[(maxIndex+i+1)%len(values)]++
		}

		steps++
		fmt.Printf("[%d] %+v\n", steps, values)
		
		// Check if aleady done
		mapped := mapToInt(values)
		if searched[mapped] > 0 {
			fmt.Printf("Steps: %d, Cycle: %d\n", steps, steps - searched[mapped])
			return
		}
		searched[mapped] = steps
	}

}

func mapToInt(in []int) string {
	return fmt.Sprintf("%+v", in)
}
