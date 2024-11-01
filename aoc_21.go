package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	health int
	damage int
	armor int
}

type Item struct {
	name string 
	cost int
	damage int
	armor int
}

func aoc_21_run() {
	boss := Player {
		health: 100,
		damage: 8,
		armor: 2,
	}
	weapons := readItems("aoc_21_in_w.txt")
	armour := readItems("aoc_21_in_a.txt")
	rings := readItems("aoc_21_in_r.txt")

	weapon_perms := getPerms(1, 1, weapons)
	armour_perms := getPerms(0, 1, armour)
	rings_perms := getPerms(0, 2, rings)

	player := Player {
		health: 100,
		damage: 0,
		armor: 0,
	}

	best_cost := math.MaxInt
	worst_cost := 0

	for i := 0; i < len(weapon_perms); i++ {
		for j := 0; j < len(armour_perms); j++ {
			for k := 0; k < len(rings_perms); k++ {
				items := []Item {}
				items = append(items, weapon_perms[i]...)
				items = append(items, armour_perms[j]...)
				items = append(items, rings_perms[k]...)
				
				cost := 0
				test_player := player
				for _, item := range(items) {
					test_player = addItem(test_player, item)
					cost += item.cost
				}

				win := simulate(test_player, boss)

				if win && cost < best_cost {
					best_cost = cost
				}
				if !win && cost > worst_cost {
					worst_cost = cost
				}
			}
		}
	}

	fmt.Printf("Best cost: %d\n", best_cost)
	fmt.Printf("Worst cost: %d\n", worst_cost)
}

func addItem(player Player, item Item) Player {
	player.armor += item.armor
	player.damage += item.damage
	return player
}

func readItems(filename string) []Item {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // skip header row
	items := []Item {}
	for scanner.Scan() {
		line := scanner.Text()
		parts_raw := strings.Split(line, " ")
		//fmt.Printf("parts: %+v\n", parts_raw)
		parts := make([]string, 4)
		index := 0
		for i, part := range(parts_raw) {
			if part == "" {
				continue
			}
			if part[0] == '+' {
				parts[i-1] += part
			} else {
				parts[index] = part
				index++
			}
		}
		if index == 0 {
			continue
		}
		if index != 4 {
			panic("Expected 4")
		}

		name := parts[0]
		cst, err := strconv.Atoi(parts[1])
		dmg, err := strconv.Atoi(parts[2])
		arm, err := strconv.Atoi(parts[3])
		if err != nil {
			panic(err)
		}
		items = append(items, Item {
			name: name,
			cost: cst,
			damage: dmg,
			armor: arm,
		})
	}
	return items
}

func getPerms(min int, max int, items []Item) [][]Item {
	//fmt.Printf("Geting perms: [%d:%d]\n", min, max)

	if max == 0 {
		return [][]Item {{}}
	}

	if min == 0 {
		return append(getPerms(1, max, items), []Item{})
	}
	
	options := [][]Item {}

	for q := min; q <= max; q++ {
		//fmt.Printf(" Finding perms of q: %d\n", q)
		options_q := [][]Item {}
		for i := 0; i < len(items)-q+1; i++ {
			subPerms := getPerms(q-1, q-1, items[i+1:])
			for j := 0; j < len(subPerms); j++ {
				subPerms[j] = append(subPerms[j], items[i])
			}
			options_q = append(options_q, subPerms...)
			
		}
		options = append(options, options_q...)
		//fmt.Printf("  Found: %+v\n", options)
	}
	return options
}


func simulate(you Player, boss Player) bool {
	you_damage := you.damage - boss.armor
	if you_damage <= 1 {
		you_damage = 1
	}
	boss_damage := boss.damage - you.armor
	if boss_damage <= 1 {
		boss_damage = 1
	}
	you_survive := you.health / boss_damage
	boss_survive := boss.health / you_damage
	fmt.Printf("  Boss: %+v\n", boss)
	fmt.Printf("  Player: %+v\n", you)
	fmt.Printf("  Win: %t, %d, %d\n", true, you_survive, boss_survive)
	return you_survive >= boss_survive
}
