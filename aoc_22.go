package main

import (
	"fmt"
	"slices"
)

type Effect struct {
	name string
	duration int
	armor int
	damage int
	mana int
}

type Spell struct {
	name string
	cost int
	damage int
	heal int
	effect Effect
}

type Player2 struct {
	health int
	mana int
}

type Boss struct {
	health int
	attack int
}

type State struct {
	player Player2
	boss Boss
	effects []Effect
	manaSpend int
	spellHist []string
	lastState *State
}

func aoc_22_run() {
	spells := []Spell {
		{
			name: "Missile",
			cost: 53,
			damage: 4,
		}, {
			name: "Drain",
			cost: 73,
			damage: 2,
			heal: 2,
		}, {
			name: "Shield",
			cost: 113,
			effect: Effect {
				name: "Shield",
				duration: 6,
				armor: 7,
			},
		}, {
			name: "Poison",
			cost: 173,
			effect: Effect {
				name: "Poisen",
				duration: 6,
				damage: 3,
			},
		}, {
			name: "Recharge",
			cost: 229,
			effect: Effect{
				name: "Recharge",
				duration: 5,
				mana: 101,
			},
		},
	}
	
	to_explore := []State {{
		player: Player2{
			health: 50,
			mana: 500,
		},
		boss: Boss{
			health: 58,
			attack: 9,
		},
		effects: []Effect {},
		manaSpend: 0,
	}}

	bestState := State {
		manaSpend: 10000,
	}

	for len(to_explore) > 0 {
		state := to_explore[0]
		to_explore = to_explore[1:]
		
		fmt.Printf("List size: %d\n", len(to_explore))
		fmt.Printf("Explore: %+v\n", state)

		options := next_options(spells, state)

		fmt.Printf("  Options: %+v\n", options)

		if len(options) == 0 {
			// lose - discard state
			fmt.Printf("  No options\n")
			continue
		}

		//index := rand.Intn(len(options))

		for _, spell := range(options) {//index:index+1]) {
			fmt.Printf("  Check: %s\n", spell.name)
			newstate := apply(state, spell) 

			if newstate.boss.health <= 0 {
				fmt.Printf("    Boss died\n")
				if newstate.manaSpend < bestState.manaSpend {
					bestState = newstate
					fmt.Printf("    NEW BEST: %+v", bestState)
				}
				continue

			}
			if newstate.player.health <= 0 {
				// lose - discard state
				fmt.Printf("    Player died\n")
				continue
			}
			if newstate.manaSpend > bestState.manaSpend {
				fmt.Printf("    Manna spend over, discard\n")
				continue
			}
			
			fmt.Printf("    Game continues: %+v\n", newstate)
			to_explore = append(to_explore, newstate)

		}
	}

	fmt.Printf(" --- BEST GAME ----\n")
	last := &bestState
	for last != nil {
		fmt.Printf("%+v\n", last)
		last = last.lastState
	}

}

func next_options(spells []Spell, state State) []Spell {
	options := []Spell {}
	Outer:
	for _, spell := range(spells) {
		if spell.cost > state.player.mana {
			continue
		}
		for _, effect := range(state.effects) {
			if effect.name == spell.effect.name {
				continue Outer
			}
		}
		options = append(options, spell)
	}
	return options
}
func apply(stateIn State, spell Spell) State {
	state := stateIn

	// duplicate arrays
	state.spellHist = slices.Clone(state.spellHist)
	state.effects = slices.Clone(state.effects)

	// instant spell things
	state.manaSpend += spell.cost

	// part 2
	state.player.health -= 1
	
	state.spellHist = append(state.spellHist, spell.name)
	state.player.mana -= spell.cost
	state.player.health += spell.heal
	state.boss.health -= spell.damage

	if spell.effect.duration > 0 {
		state.effects = append(state.effects, spell.effect)
	}

	// pre boss effects
	state, armor := applyEffects(state)

	if state.boss.health <= 0 {
		return state
	}

	// boss turn
	state.player.health -= state.boss.attack - armor

	if state.player.health <= 0 {
		return state
	}

	// Pre player effects
	state, _ = applyEffects(state)
	state.lastState = &stateIn

	return state
}

func applyEffects(state State) (State, int) {
	armor := 0
	for i := 0; i < len(state.effects); i++ {
		effect := state.effects[i]
		state.player.mana += effect.mana
		state.boss.health -= effect.damage
		armor += effect.armor
		effect.duration--
		if effect.duration <= 0 {
			state.effects = append(state.effects[:i], state.effects[i+1:]...)
			i--
		} else {
			state.effects[i] = effect
		}
	}
	return state, armor
}

