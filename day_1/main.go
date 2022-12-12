package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var max int
	var cur int
	var elves int
	sp := strings.Split(input, "\n")
	for _, text := range sp {
		text := strings.ReplaceAll(text, "\n", "")
		if text == "" {
			elves++
			if max < cur {
				max = cur
			}
			cur = 0
		} else {
			num, err := strconv.Atoi(text)
			if err != nil {
				panic(err)
			}
			cur += num
		}
	}
	fmt.Printf("Part 1: %d\n", max)
	fmt.Printf("Part 2: %d\n", part2(elves, sp))
}

func part2(elves int, sp []string) int {
	cur := 0
	i := 0
	result := make([]int, elves)
	for _, text := range sp {
		text := strings.ReplaceAll(text, "\n", "")
		if text == "" {
			result[i] = cur
			i++
			cur = 0
		} else {
			num, err := strconv.Atoi(text)
			if err != nil {
				panic(err)
			}
			cur += num
		}
	}
	sort.Ints(result)
	res := result[elves-1] + result[elves-2] + result[elves-3]
	return res
}
