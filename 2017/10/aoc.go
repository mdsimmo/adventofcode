package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part2() {
	lengths, err := os.ReadFile("in.txt")
	if err != nil {
		panic(err)
	}
	lengths = []byte(strings.Trim(string(lengths), "\n"))
	data := make([]byte, 256)
	for i := 0; i < len(data); i++ {
		data[i] = byte(i)
	}
	lengths = append(lengths, []byte { 17, 31, 73, 47, 23 }...)
	
	position := 0
	skip := 0
	for i := 0; i < 64; i++ {
		for _, length := range(lengths) {
			subData := make([]byte, length)
			for i := 0; i < int(length); i++ {
				subData[i] = data[(position+i)%len(data)]
			}
			slices.Reverse(subData)
			for i := 0; i < int(length); i++ {
				data[(position+i)%len(data)] = subData[i]
			}
			position += int(length) + skip 
			skip++
		}
	}

	dense := make([]byte, len(data)/16)
	for i := 0; i < len(dense); i++ {
		dense[i] = 0
		for j := 0; j < 16; j++ {
			dense[i] ^= data[i*16+j]
		}
	}

	for _, d := range(dense) {
		fmt.Printf("%02x", d)
	}
}

func part1() {
	file, err := os.ReadFile("in.txt")
	if err != nil {
		panic(err)
	}
	contents := string(file)
	parts := strings.Split(contents, ",")
	lengths := make([]int, len(parts))
	for i := 0; i < len(parts); i++ {
		value, e := strconv.Atoi(strings.Trim(parts[i], "\n"))
		if e != nil {
			panic(e)
		}
		lengths[i] = value
	}
	data := make([]int, 256)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}
	
	position := 0
	for skip, length := range(lengths) {
		subData := make([]int, length)
		for i := 0; i < length; i++ {
			subData[i] = data[(position+i)%len(data)]
		}
		slices.Reverse(subData)
		for i := 0; i < length; i++ {
			data[(position+i)%len(data)] = subData[i]
		}
		position += length + skip 
	}
	fmt.Printf("First two mult: %d\n", data[0]*data[1])
}
