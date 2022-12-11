package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	max := 0
	cur := 0
	elves := 0
	for fileScanner.Scan() {
		text := fileScanner.Text()
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
	readFile.Seek(0, io.SeekStart)
	fmt.Printf("Part 2: %d\n", part2(elves, readFile))
}

func part2(elves int, readFile *os.File) int {
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	cur := 0
	i := 0
	result := make([]int, elves)
	for fileScanner.Scan() {
		text := fileScanner.Text()
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
