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
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	// Creating as int32 lets us skip converting rune to int later, as a rune is just an alias for int32
	var res int32
	var resTwo int32
	var group [3]string
	var i int
	fileScanner := bufio.NewScanner(readFile)
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
		if i >= 2 {
			group[i] = text
			i = 0
			for _, c := range group[0] {
				if strings.ContainsRune(group[1], c) && strings.ContainsRune(group[2], c) {
					resTwo += get_priority(c)
					break
				}
			}
		} else {
			group[i] = text
			i++
		}
	}
	fmt.Printf("Part 1: %d\n", res)
	fmt.Printf("Part 2: %d\n", resTwo)
}

func get_priority(c rune) rune {
	if unicode.IsLower(c) {
		return c - 96
	} else {
		return c - 38
	}
}
