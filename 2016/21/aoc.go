package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	//descramble := false
	//text := []rune("abcdefgh")
	descramble := true
	text := []rune("fbgdceah") //"abcdefgh")

	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	lines := []string {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if descramble {
			line = lines[len(lines)-i-1]
		}
		fmt.Printf("Process: %s\n", line)
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "swap":
			switch parts[1] {
			case "position":
				a, _ := strconv.Atoi(parts[2])
				b, _ := strconv.Atoi(parts[5])
				c := text[b]
				text[b] = text[a]
				text[a] = c
			case "letter":
				for i, ri := range(text) {
					if ri != rune(parts[2][0]) {
						continue
					}
					for j, rj := range(text) {
						if rj != rune(parts[5][0]) {
							continue
						}
						text[i] = rj
						text[j] = ri
						goto complete
					}
				}
				panic("Should never happen")
				complete: ;
			default: 
				panic("Unknown function")
			}
		case "reverse":
			i1, _ := strconv.Atoi(parts[2])
			i2, _ := strconv.Atoi(parts[4])
			slices.Reverse(text[i1:i2+1])
		case "rotate":
			switch parts[1] {
			case "based":
				r := rune(parts[6][0])
				if descramble {
					// find indexScrambled of new character
					var indexScrambled int
					for i, rc := range(text) {
						if rc == r {
							indexScrambled = i
							break
						}
					}
					// Search for how many rotatesTry were done
					var indexUnscram int
					if indexScrambled % 2 == 0 {
						indexUnscram = (indexScrambled - 2) / 2
						for indexUnscram < 4 {
							indexUnscram += len(text)/2
						}
					} else {
						indexUnscram = (indexScrambled - 1) / 2
					} 
					rotates := indexUnscram
					if indexUnscram >= 4 {
						rotates++
					}
					rotates++
					if (indexUnscram + rotates) % len(text) != indexScrambled {
						panic("Calculated incorrect rotate")
					}
					for i := 0; i < rotates; i++ {
						text = append(text[1:], text[0])
					}
				} else {
					var rotates int
					for i, rc := range(text) {
						if rc == r {
							rotates = i
						}
					}
					if rotates >= 4 {
						rotates++
					}
					rotates++
					for i := 0; i < rotates; i++ {
						text = append([]rune{ text[len(text)-1] }, text[:len(text)-1]...)
					}
				}
			case "right":
				fallthrough
			case "left":
				right := parts[1] == "right"
				if descramble {
					right = !right
				}
				steps, _ := strconv.Atoi(parts[2])
				for i := 0; i < steps; i++ {
					if right {
						text = append([]rune{ text[len(text)-1] }, text[:len(text)-1]...)
					} else {
						text = append(text[1:], text[0])
					}
				}
			default:
				panic("unknown option")
			}
		case "move":
			i1, _ := strconv.Atoi(parts[2])
			i2, _ := strconv.Atoi(parts[5])
			if descramble {
				t := i1
				i1 = i2
				i2 = t
			}
			r := text[i1]
			text = append(text[:i1], text[i1+1:]...)
			var sb strings.Builder
			sb.WriteString(string(text[:i2]))
			sb.WriteRune(r)
			sb.WriteString(string(text[i2:]))
			text = []rune(sb.String())
		default:
			panic("Unknown instruction")
		}
		fmt.Printf("  Result: %s\n", string(text))
	}
}
