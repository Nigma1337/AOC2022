package main

import (
	"fmt"
	"os"
)

func main() {
	var m map[byte]int
	var n map[byte]int
	var part_1_done bool
	readFile, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	for i := range readFile {
		m = make(map[byte]int)
		n = make(map[byte]int)
		for _, c := range readFile[i : i+4] {
			m[c] = 0
			n[c] = 0
		}
		for _, c := range readFile[i+4 : i+14] {
			n[c] = 0
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
