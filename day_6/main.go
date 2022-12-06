package main

import (
	"fmt"
	"os"
)

// Empty structs take up 0 bytes (according to https://linuxhint.com/golang-set/)
type void struct{}

var member void

func main() {
	var m map[byte]void
	var n map[byte]void
	var part_1_done bool
	readFile, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	for i := range readFile {
		m = make(map[byte]void)
		n = make(map[byte]void)
		for _, c := range readFile[i : i+4] {
			m[c] = member
			n[c] = member
		}
		for _, c := range readFile[i+4 : i+14] {
			n[c] = member
		}
		if len(m) == 4 && !part_1_done {
			fmt.Printf("Part 1: %d\n", i+4)
			part_1_done = true
		}
		if len(n) == 14 {
			fmt.Printf("Part 2: %d\n", i+14)
			break
		}
	}
}
