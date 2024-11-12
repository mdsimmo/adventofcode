package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

type Entry struct {
	index int
	c byte
}

func main() {
	input := "zpqevtbw"
	//input := "abc"
	hashes := 2017

	tripStore := []Entry {}
	pentStore := []Entry {}

	count := 0
	IndexLoop:
	for index := 0; count < 64; index++ {
		if len(tripStore) != 0 {
			index = tripStore[0].index
			//fmt.Printf("Skipped to %d\n", index)
		}
		hash := hash_value(input, index, hashes)
		
		// discard all pents less than the current index
		for len(pentStore) > 0 && pentStore[0].index <= index {
			pentStore = pentStore[1:]
		}
		// discard all trips less than the current index
		for len(tripStore) > 0 && tripStore[0].index <= index {
			tripStore = tripStore[1:]
		}

		trippleCheck:
		for i := 0; i < len(hash) - 2; i++ {
			c := hash[i]
			for j := i+1; j < i+3; j++ {
				if hash[j] != c {
					continue trippleCheck
				}
			}
		
			// check values in pent store
			for i := 0; i < len(pentStore); i++ {
				if pentStore[i].c == c {
					count++
					fmt.Printf("  [%d] Remembered pent: index: %d, %d\n", count, index, pentStore[i].index)
					continue IndexLoop
				}
			}

			//fmt.Printf("Found tripple: %d\n", index)
			startIter := index + 1
			if len(tripStore) > 0 {
				startIter = tripStore[len(tripStore)-1].index+1
			}
			for j := startIter; j < index + 1 + 1000; j++ {
				j_hash := hash_value(input, j, hashes)
				
				j_c := byte('-')
				c_count := 0
				for k := 0; k < len(j_hash); k++ {
					if j_hash[k] != j_c {
						c_count = 1
						j_c = j_hash[k]
					} else {
						c_count++
						if c_count == 3 {
							//fmt.Printf("  Added trip store, %d, %c\n", j, j_c)
							tripStore = append(tripStore, Entry{
								index: j,
								c: j_c,
							})
						}
						if c_count == 5 {
							//fmt.Printf("  Added pent store, %d, %c\n", j, j_c)
							pentStore = append(pentStore, Entry {
								index: j,
								c: j_c,
							})
							if j_c == c {
								count++
								fmt.Printf("  [%d] Found pent: %d, %d, %s, %s\n", count, index, j, hash, j_hash)
								continue IndexLoop
							}
						}
					}
				}
			}

			// only consider first tripplet
			break
		}
	}
}


func hash_value(input string, index int, hashes int) string {
	hashString := input + strconv.Itoa(index)
	hashb := md5.Sum([]byte(hashString))
	for i := 0; i < hashes-1; i++ {
		hashb = md5.Sum([]byte(hex.EncodeToString(hashb[:])))
	}
	hash := hex.EncodeToString(hashb[:])
	return hash
}
