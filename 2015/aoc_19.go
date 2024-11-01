package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

type Entry struct {
	steps int
	seq string
	best_possible int
}

type Replace struct {
	from string
	to string
}

func aoc_19_run() {

	file, err := os.Open("aoc_19_in.txt")
	if err != nil {
		panic(err)
	}

	replaces := make([]Replace, 0)
	replaces_xrya := make([][]Replace, 3)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = regexp.MustCompile("Ti").ReplaceAllString(line, "X")
		line = regexp.MustCompile("Al").ReplaceAllString(line, "W")
		line = regexp.MustCompile("Ca").ReplaceAllString(line, "Z")
		parts := strings.Split(line, " ")
		pre := regexp.MustCompile("[a-z]").ReplaceAllString(string(parts[0][0]), "")
		post := regexp.MustCompile("[a-z]").ReplaceAllString(parts[2], "")
		rep := Replace {
			from: post,
			to: pre,
		}
		if strings.Contains(post, "R") {
			yCount := strings.Count(post, "Y")
			replaces_xrya[yCount] = append(replaces_xrya[yCount], rep)
		} else {
			replaces = append(replaces, rep)
		}
	}
	
	slices.SortFunc(replaces, repSort)
	fmt.Printf("Replaces: %+v\n", replaces)

	file2, err := os.Open("aoc_19_in2.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewScanner(file2)
	reader.Scan()
	sequence := reader.Text()
	sequence = regexp.MustCompile("Ti").ReplaceAllString(sequence, "X")
	sequence = regexp.MustCompile("Al").ReplaceAllString(sequence, "W")
	sequence = regexp.MustCompile("Ca").ReplaceAllString(sequence, "Z")
	sequence = regexp.MustCompile("[a-z]").ReplaceAllString(sequence, "")
	fmt.Printf("Sequence: %s\n", sequence)


	steps := 0

	for true {

		any_found := false

		fmt.Printf("Sequence: (%d steps) %s\n", steps, sequence)
		seq, iters := condense_rparts(sequence, replaces)
		steps += iters
		any_found = any_found || iters > 0

		fmt.Printf(" Seq %d steps: %s\n", steps, seq)

		if iters == 0 {

			// final sub
			if len(seq) == 2 {
				for _, rep := range(replaces) {
					if rep.from == seq {
						steps++
						fmt.Printf("COMPLETE: %d steps", steps)
						return
					}
				}
			}

			fmt.Printf(" No subs found - condensing\n")
			seq, iters = condense_before_midr(seq, replaces)
			steps += iters
			any_found = any_found || iters > 0
			fmt.Printf(" Seq %d stepd: %s\n", steps, seq)
		}

		seq, iters = sub_rparts(seq, replaces_xrya)
		steps += iters
		sequence = seq
		any_found = any_found || iters > 0

		fmt.Printf(" Replaced to (%d steps): %s\n", steps, sequence)

		if !any_found {
			fmt.Printf("No replacements made. Quiting\n")
			break
		}
	}
}

func sub_rparts(in string, reps_y [][]Replace) (string, int) {
	steps := 0
	for {
		found := false
		for _, rys := range(reps_y) {
			for _, ry := range(rys) {
				new_in := strings.Replace(in, ry.from, ry.to, 1)
				if new_in != in {
					found = true
					steps++
				}
				in = new_in
			}
		}
		if !found {
			return in, steps
		}
	}
}

func condense_rparts(in string, reps []Replace) (string, int) {
	steps := 0
	rParts := strings.Split(in, "R")

	if len(rParts) == 0 {
		return in, 0
	}

	var newSequence strings.Builder
	first, iter := condense(rParts[0], reps)
	newSequence.WriteString(first)
	steps += iter

	for ri := 1; ri < len(rParts); ri++ { // 1 = skip stuff before first R
		rPart := rParts[ri]
		fmt.Printf("  Inspect %s\n", rPart)

		aParts := strings.Split(rPart, "A")
		if len(aParts) < 2 {
			fmt.Printf("    No A, skipping\n")
			newSequence.WriteString("R")
			newSequence.WriteString(rPart)
			continue
		}
		
		yParts := strings.Split(aParts[0], "Y")
		for y := 0; y < len(yParts); y++ {
			newPart, iter := condense(yParts[y], reps)
			fmt.Printf("    Condensed %s -> %s\n", yParts[y], newPart)
			if len(newPart) != 1 {
				panic("Must condense to 1")
			}
			steps += iter
			yParts[y] = newPart
		}
		var newRPart strings.Builder
		for i := 0; i < len(yParts); i++ {
			newRPart.WriteString(yParts[i])
			if i < len(yParts)-1 {
				newRPart.WriteString("Y")
			}
		}
		newRPart.WriteString("A")
		for i := 1; i < len(aParts); i++ {
			newRPart.WriteString(aParts[i])
			if i < len(aParts)-1 {
				newRPart.WriteString("A")
			}
		}
		fmt.Printf("    Full Condensed %s -> %s\n", rPart, newRPart.String())
		newSequence.WriteString("R")
		newSequence.WriteString(newRPart.String())
	}
	return newSequence.String(), steps
}

func condense(in string, reps []Replace) (string, int) {
	replaces := 0
	for {
		found := false
		for i := 0; i < len(reps); i++ {
			rep := reps[i]
			if rep.to == "" {
				continue
			}
			if index := strings.LastIndex(in, rep.from); index >= 0 {
				var newIn strings.Builder
				newIn.WriteString(in[:index])
				newIn.WriteString(rep.to)
				newIn.WriteString(in[index+len(rep.from):])
				in = newIn.String()
				i = 0
				replaces++
				found = true
			}
		}
		if !found {
			return in, replaces
		}
	}
}


func condense_before_midr(in string, reps []Replace) (string, int) {
	// find r followed by a
	rParts := strings.Split(in, "R")
	for i := 0; i < len(rParts); i++ {
		rPart := rParts[i]

		aParts := strings.Split(rPart, "A")
		if len(aParts) < 2 {
			continue
		}
		
		rPartPre := rParts[i-1]
		lastTwo := rPartPre[len(rPartPre)-2:]

		for _, rep := range(reps) {
			if rep.from == lastTwo {
				var newString strings.Builder
				for j, p := range(rParts) {
					if j == i-1 {
						newString.WriteString(rPartPre[:len(rPartPre)-2])
						newString.WriteString(rep.to)
					} else {
						newString.WriteString(p)
					}
					if j < len(rParts)-1 {
						newString.WriteString("R")
					}
				}
				return newString.String(), 1
			}
		}

		panic("Cannot replace the last two")
	}

	panic("No A following an R")

}

func repSort(a Replace, b Replace) int {
	if len(a.from) > len(b.from) {
		return -1
	} else if len(a.from) < len(b.from) {
		return 1
	} else {
		return 0
	}
}

