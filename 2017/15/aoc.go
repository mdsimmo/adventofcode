package main

import "fmt"

func main() {
	genVals := []int { 634, 301 }
	//genVals := []int { 65, 8921 }
	genMult := []int { 16807, 48271 }
	genMods := []int { 4, 8 }
	modulo := 2147483647

	matches := 0
	for i := 0; i < 5000000; i++ {
		for j := 0; j < 2; j++ {
			genVals[j] = (genVals[j] * genMult[j]) % modulo
			for genVals[j] % genMods[j] != 0 {
				genVals[j] = (genVals[j] * genMult[j]) % modulo
			}
		}
		if genVals[0] & 0xffff == genVals[1] & 0xffff {
			matches++
		}
	}
	fmt.Printf("Matches: %d\n", matches)
}
