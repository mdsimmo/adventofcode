package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Entry struct {
	path string
	x int
	y int
}

func main() {
	//in := "ihgpwlah"
	in := "vkjiggvb"

	search := []Entry {{
		path: "",
		x: 0,
		y: 0,
	}}

	worst := search[0]

	for len(search) > 0 {
		e := search[0]
		search = search[1:]

		if e.x == 3 && e.y == 3 {
			fmt.Printf("FOUND VAULT: %s\n", e.path)
			
			if len(worst.path) < len(e.path) {
				worst = e
			}

			continue
		}

		hash := md5.Sum([]byte(in + e.path))
		hash_s := hex.EncodeToString(hash[:])
		u := hash_s[0] > 'a'
		d := hash_s[1] > 'a'
		l := hash_s[2] > 'a'
		r := hash_s[3] > 'a'

		if u && e.y > 0{
			search = append(search, Entry{
				path: e.path + "U",
				x: e.x,
				y: e.y - 1,
			})
		}
		if d && e.y < 3 {
			search = append(search, Entry{
				path: e.path + "D",
				x: e.x,
				y: e.y + 1,
			})
		}
		if l && e.x > 0 {
			search = append(search, Entry{
				path: e.path + "L",
				x: e.x - 1,
				y: e.y,
			})
		}
		if r && e.x < 3 {
			search = append(search, Entry{
				path: e.path + "R",
				x: e.x + 1,
				y: e.y,
			})
		}
	}

	fmt.Printf("Worst path: (%d) %s\n", len(worst.path), worst.path)
}
