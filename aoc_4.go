package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func aoc_4_run() {
	find_hash("bgvyzdsv", 6)
}

func find_hash(input string, leaders int) {
	for index := 0;; index++ {
		test := input + strconv.Itoa(index)
		hash := md5.Sum([]byte(test))
		ok := true;
		max_index := (leaders+1)/2
		half_index := max_index - 1
		if leaders % 2 == 0 {
			half_index = max_index
		}
		for i := 0; i < max_index; i++ {
			if (
				((i != half_index) && (hash[i] != 0)) ||
				((i == half_index) && ((hash[i] & 0xF0) != 0))) {
				ok = false;
			}
		}
		if ok {
			fmt.Printf("FOUND: %d, %x", index, hash)
			break;
		}
		fmt.Printf("\n%d: %x\n", index, hash)
	}
	
}
