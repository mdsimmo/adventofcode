package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("in.txt")
	data := strings.Trim(string(file), " \n\r")
	length := decompress(data)
	fmt.Printf("Length: %d\n", length)
}

func decompress(data string) int {
	length := 0
	for i := 0; i < len(data); i++ {
		r := rune(data[i])
		if r != '(' {
			length++
			continue
		}
		chars := 0
		for j := i+1; ; j++ {
			r = rune(data[j])
			if r == 'x' {
				var err error
				chars, err = strconv.Atoi(data[i+1:j])
				if err != nil {
					panic(err)
				}
				i = j
				break
			}
		}
		repeats := 0
		for j := i+1; ; j++ {
			r = rune(data[j])
			if r == ')' {
				var err error
				repeats, err = strconv.Atoi(data[i+1:j])
				if err != nil {
					panic(err)
				}
				i = j
				break
			}
		}
		substring := data[i+1:i+1+chars]
		length += decompress(substring) * repeats
		i += chars
	}
	return length	
}
