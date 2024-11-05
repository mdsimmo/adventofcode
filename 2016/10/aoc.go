package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("in.txt")
	ins := []string {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ins = append(ins, line)
	}

	bots := make(map[int][]int)
	firstLoop := true
	for {
		anyAlter := false
		for _, in := range(ins) {
			parts := strings.Split(in, " ")
			fmt.Printf("Check: %s\n", in)
			if firstLoop && parts[0] == "value" {
				value, _ := strconv.Atoi(parts[1])
				bot, _ := strconv.Atoi(parts[5])
				bots[bot] = append(bots[bot], value)
				fmt.Printf("  Value %d given to %d\n", value, bot)
				anyAlter = true
			}

			if !firstLoop && parts[0] == "bot" {
				botA, _ := strconv.Atoi(parts[1])
				botB, _ := strconv.Atoi(parts[6])
				botC, _ := strconv.Atoi(parts[11])
				botBType := parts[5] == "bot"
				botCType := parts[10] == "bot"
				firstLow := parts[3] == "low"
				
				// store outputs at 1000+ registers
				if !botBType {
					botB += 1000
				}
				if !botCType {
					botC += 1000
				}

				if len(bots[botA]) > 2 {
					panic("unexpected length")
				}
				if len(bots[botA]) == 2 {
					value1 := bots[botA][0]
					value2 := bots[botA][1]
					
					if (value1 == 61 && value2 == 17) || (value1 == 17 && value2 == 61) {
						fmt.Printf("BOT COMPARE: %d\n", botA)
						//return
					}

					if (value2 < value1 && firstLow) || (value2 > value1 && !firstLow) {
						bots[botB] = append(bots[botB], value2)
						bots[botC] = append(bots[botC], value1)
						fmt.Printf("  Bot %4d: %4d -> %4d, %4d -> %4d\n", botA, value2, botB, value1, botC)
					} else {
						bots[botB] = append(bots[botB], value1)
						bots[botC] = append(bots[botC], value2)
						fmt.Printf("  Bot %4d: %4d -> %4d, %4d -> %4d\n", botA, value1, botB, value2, botC)
					}
					bots[botA] = make([]int, 0)
					anyAlter = true
				} else {
					fmt.Printf("  Skip: Bot %d had %d tokens\n", botA, len(bots[botA]))
				}
				
			}
		}
		if firstLoop {
			firstLoop = false
			fmt.Printf("First Loop complete\n")
		} else if anyAlter {
			fmt.Printf("Loop complete. Restart\n")
		} else {
			fmt.Printf("No Change. Stopping\n")
			break
		}
	}
	
	mult := 1
	for i := 1000; i <= 1002; i++ {
		for j := 0; j < len(bots[i]); j++ {
			mult *= bots[i][j]
		}
	}
	fmt.Printf("Multiply: %d\n", mult)
}

