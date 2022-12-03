package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	res := 0
	resTwo := 0
	i := 0
	var group [3]string
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		text := fileScanner.Text()
		length := len(text) / 2
		one := text[:length]
		two := text[length:]
		for _, c := range one {
			if strings.ContainsRune(two, c) {
				res += get_priority(c)
				break
			}
		}
		group[i] = text
		i++
		if i == 3 {
			i = 0
			for _, c := range group[0] {
				if strings.ContainsRune(group[1], c) && strings.ContainsRune(group[2], c) {
					resTwo += get_priority(c)
					break
				}
			}
		}
	}
	fmt.Printf("Part 1: %d\n", res)
	fmt.Printf("Part 2: %d\n", resTwo)
}

func get_priority(c rune) int {
	if unicode.IsLower(c) {
		return int(c - 96)
	} else {
		return int(c - 38)
	}
}
