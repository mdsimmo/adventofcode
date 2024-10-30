package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func aoc_15_run() {
	file, err := os.Open("aoc_15_in.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	ingr := make([][]int, 0)
	calories := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		p1 := parts[2]
		p2 := parts[4]
		p3 := parts[6]
		p4 := parts[8]
		p5 := parts[10]
		points := make([]int64, 4)
		points[0], _ = strconv.ParseInt(p1[:len(p1)-1], 10, 32)
		points[1], _ = strconv.ParseInt(p2[:len(p2)-1], 10, 32)
		points[2], _ = strconv.ParseInt(p3[:len(p3)-1], 10, 32)
		points[3], _ = strconv.ParseInt(p4[:len(p4)-1], 10, 32)
		cals, _ := strconv.ParseInt(p5, 10, 32)
		prt2 := make([]int, 4)
		for i := 0; i < 4; i++ {
			prt2[i] = int(points[i])
		}
		ingr = append(ingr, prt2)
		calories = append(calories, int(cals))
	}

	fmt.Printf("Ingredients: %+v\n", ingr)

	quanties := make([]int, len(ingr))
	for i := 0; i < len(quanties); i++ {
		quanties[i] = 100/len(quanties)
	}
	best := 0
	out:
	for true {
		for i := 0; i < len(quanties); i++ {
			for j := 0; j < len(quanties); j++ {
				if j == i {
					continue
				}
				trial := make([]int, len(quanties))
				copy(trial, quanties)
				trial[i]++
				trial[j]--
			
				score := score(trial, ingr)
				fmt.Printf("Testing: %+v, %d\n", trial, score)
				if score > best {
					best = score
					quanties = trial
					continue out
				}
			}
		}
		break
	}
	fmt.Printf("Best: %d, %+v\n", best, quanties)

	// search for a nearby 500 cal cookie
	rng := 100
	best = 0
	newquant := make([]int, len(quanties))
	for i1 := -rng; i1 <=rng; i1++ {
		for i2 := -rng; i2 <=rng; i2++ {
			for i3 := -rng; i3 <=rng; i3++ {
				for i4 := -rng; i4 <=rng; i4++ {
					if i1 + i2 + i3 + i4 != 0 {
						continue
					}
					trial := make([]int, len(quanties))
					copy(trial, quanties)
					trial[0] += i1
					trial[1] += i2
					trial[2] += i3
					trial[3] += i4
					cls := cals(trial, calories)
					if cls != 500 {
						continue
					}
					score := score(trial, ingr)
					if score > best {
						best = score
						newquant = trial
					}
				}
			}
		}
	}

	fmt.Printf("500: %d, %+v", best, newquant)

}

func score(trial []int, ingr [][]int) int {
	score := 1
	for flav := 0; flav < len(ingr[0]); flav++ {
		flav_score := 0
		for ing := 0; ing < len(trial); ing++ {
			if trial[ing] <= 0 {
				return 0
			}
			flav_score += trial[ing]*ingr[ing][flav]
		}
		if flav_score <= 0 {
			return 0
		}
		fmt.Printf("  Flav score: %d, %d\n", flav, flav_score)
		score = score * flav_score
	}
	return score
}

func cals(trial []int, cals []int) int {
	total := 0
	for ing := 0; ing < len(trial); ing++ {
		total += trial[ing] * cals[ing]
	}
	return total
}
