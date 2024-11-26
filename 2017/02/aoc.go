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
	if err != nil {
		panic(err)
	}
	sum := 0
	sum2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		largest := math.MinInt
		smallest := math.MaxInt
		for i, sn := range(parts) {
			n, e := strconv.Atoi(sn)
			if e != nil {
				panic(e)
			}
			
			// part 1
			if largest < n {
				largest = n
			}
			if smallest > n {
				smallest = n
			}

			// part 2
			for i2, sn2 := range(parts) {
				n2, e := strconv.Atoi(sn2)
				if e != nil {
					panic(e)
				}
				if n % n2 == 0 && i != i2 {
					sum2 += n / n2
				}
			}
		}
		sum += largest - smallest
	}
	fmt.Printf("Checksum: %d\n", sum)
	fmt.Printf("Divide Sum: %d\n", sum2)
}
