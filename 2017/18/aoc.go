package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main2() {

	c := make(chan int)

	go test(c)

	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", <-c)
		time.Sleep(1 * time.Second)
	}

}

func test(c chan int) {
	for i := 0; ; i++ {
		c <- i
		fmt.Printf("I: %d\n", i)
	}
}

func main() {

	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	insList := []string {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		insList = append(insList, line)
	}

	c0 := make(chan int, 100000)
	c1 := make(chan int, 100000)
	go run(0, insList, c0, c1)
	go run(1, insList, c1, c0)

	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("Wait\n")
	}
}

func run(id int, insList []string, output chan<- int, input <-chan int) {
	registers := make([]int, 26)
	registers['p'-'a'] = id
	sendCount := 0
	for i := 0; i < len(insList); i++ {
		// Interpret the line
		line := insList[i]
		//fmt.Printf("[%d, %d] %s\n", id, i, line)
		parts := strings.Split(line, " ")
		ins := parts[0]
		idxs := make([]int, len(parts)-1)
		vals := make([]int, len(parts)-1)
		for j := 0; j < len(parts)-1; j++ {
			var err error = nil
			vals[j], err = strconv.Atoi(parts[j+1])
			if err != nil {
				idxs[j] = int(parts[j+1][0]) - 'a'
				vals[j] = registers[idxs[j]]
			} else {
				idxs[j] = -1 // not a register
			}
		}

		// Execute the instruction
		switch ins {
		case "snd":
			//fmt.Printf("[%d] Sending %d\n", id, vals[0])
			output <- vals[0]
			sendCount++
			fmt.Printf("[%d] Sent: %d (%d)\n", id, vals[0], sendCount)
		case "set":
			registers[idxs[0]] = vals[1]
		case "add":
			registers[idxs[0]] += vals[1]
		case "mul":
			registers[idxs[0]] *= vals[1]
		case "mod":
			registers[idxs[0]] = vals[0] % vals[1]
		case "rcv":
			//fmt.Printf("[%d] Waiting\n", id)
			registers[idxs[0]] = <- input
			//fmt.Printf("[%d] Recv: %d\n", id, registers[idxs[0]])
		case "jgz":
			if vals[0] > 0 {
				// minus one to account for normal loop increment
				i += vals[1] - 1
			}
		default: 
			panic("Unknown instruction")
		}
	}
	panic("Loop complete")
}
