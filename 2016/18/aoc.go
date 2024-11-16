package main

import (
	"fmt"
)

func main() {
	in := "^^.^..^.....^..^..^^...^^.^....^^^.^.^^....^.^^^...^^^^.^^^^.^..^^^^.^^.^.^.^.^.^^...^^..^^^..^.^^^^"
	//in := ".^^.^.^^^^"

	traps := make([][]bool, 400000)
	traps[0] = make([]bool, len(in))
	for i := 0; i < len(in); i++ {
		traps[0][i] = in[i] == '^'
	}
	for i := 1; i < len(traps); i++ {
		traps[i] = make([]bool, len(in))
		for j := 0; j < len(traps[i]); j++ {
			l := false
			c := traps[i-1][j]
			r := false
			if j > 0 {
				l = traps[i-1][j-1]
			}
			if j < len(in)-1 {
				r = traps[i-1][j+1]
			}
			traps[i][j] = (l && c && !r) || (!l && c && r) || (l && !c && !r) || (!l && !c && r)
		}
	}

	sum := 0
	for _, r := range(traps) {
		for _, t := range(r) {
			if t {
				//fmt.Print("^")
			} else {
				sum++
				//fmt.Print(".")
			}
		}
		//fmt.Println()
	}
	fmt.Printf(
		"Sum: %d\n",
		sum,
	)
}
