package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "uqwqemis"
	result := make([]rune, 8)
	resultDone := make([]bool, 8)
	j := 0
	for i := 0; i < 8; i++ {
		for {
			test := input + strconv.Itoa(j)
			j++
			hash := md5.Sum([]byte(test))
			hashTxt := fmt.Sprintf("%x", hash)
			if strings.HasPrefix(hashTxt, "00000") {
				fmt.Printf("%s\n", hashTxt)
				ix, _ := strconv.ParseInt(string(rune(hashTxt[5])), 16, 0)
				c := hashTxt[6]
				if ix < 8 && !resultDone[ix] {
					result[ix] = rune(c)
					resultDone[ix] = true;
					fmt.Printf("  %c inserted at %d\n", c, ix)
					break
				} else {
					fmt.Printf("  Skipped (%c -> %d)\n", c, ix)
				}
			}
		}
	}
	fmt.Printf("Result: %s\n", string(result))
}

