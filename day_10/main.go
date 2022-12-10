package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var cycle int
	var strength int
	x := 1
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	fmt.Printf("Part 2: ")
	for fileScanner.Scan() {
		text := fileScanner.Text()

		if text == "noop" {
			drawPixel(cycle, x)
			cycle++
			if cycle%40 == 20 {
				strength += cycle * x
			}
			continue
		} else {
			sp := strings.Split(text, " ")
			number, err := strconv.Atoi(sp[1])
			if err != nil {
				panic(err)
			}
			for i := 0; i < 2; i++ {
				drawPixel(cycle, x)
				cycle++
				if cycle%40 == 20 {
					strength += cycle * x
				}
			}
			x += number
		}
	}
	fmt.Println()
	fmt.Printf("Part 1: %d\n", strength)
}

func drawPixel(cycle int, x int) {
	var character string
	c := cycle % 40
	if x == c+1 || x == c-1 || x == c {
		character = "#"
	} else {
		character = "."
	}
	if cycle%40 == 0 {
		fmt.Println()
	}
	fmt.Printf("%s", character)
}
