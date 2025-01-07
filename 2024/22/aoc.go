package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	secrets := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		num, e := strconv.Atoi(line)
		if e != nil {
			panic(e)
		}

		secret := num
		secrets = append(secrets, []int{secret})
		for i := 0; i < 2000; i++ {
			secret = (secret ^ (secret * 64)) % 16777216
			secret = (secret ^ (secret / 32)) % 16777216
			secret = (secret ^ (secret * 2048)) % 16777216
			secrets[len(secrets)-1] = append(secrets[len(secrets)-1], secret)
		}
		sum += int(secret)
	}
	fmt.Printf("Sum (Part 1): %d\n", sum)

	prices := make([][]int, len(secrets))
	deltas := make([][]int, len(secrets))
	for i, monkey := range secrets {
		prices[i] = make([]int, len(monkey))
		deltas[i] = make([]int, len(monkey))
		for j, secret := range monkey {
			prices[i][j] = secret % 10
			if j == 0 {
				deltas[i][j] = -100
				continue
			}
			deltas[i][j] = prices[i][j] - prices[i][j-1]
		}
	}

	// fmt.Printf("%+v\n", deltas)

	// This solution is brute forces
	// It takes a couple minute to complete
	best := 0
	// Assume that one of the first three monkeys sells
	for m := 0; m < 3; m++ {
		for i := 1; i < len(deltas[m])-4; i++ {
			window := deltas[m][i : i+4]
			fmt.Printf("[%d, %d] Check (%d): %+v\n", m, i, best, window)

			sum := 0
			for monkey := range secrets {
				sum += bananaPrice(prices[monkey], deltas[monkey], window)
			}
			if sum > best {
				best = sum
			}
		}
	}
	fmt.Printf("Best (Part 2): %d\n", best)
}

func cacheInt(window []int) int {
	val := 1
	for _, v := range window {
		val *= 100
		val += v
	}
	return val
}

func bananaPrice(price []int, delta []int, window []int) int {
outer:
	for j := len(window); j < len(delta); j++ {
		for k := 0; k < 4; k++ {
			if delta[j-3+k] != window[k] {
				continue outer
			}
		}
		return price[j]
	}
	return 0
}
