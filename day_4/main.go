package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var firsti [2]int
	var seci [2]int
	var resTwo int
	var res int
	readFile, err := os.Open("input.txt")
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		text := fileScanner.Text()
		sp := strings.Split(text, ",")
		checkLen(sp, 2)
		first := sp[0]
		firstsp := strings.Split(first, "-")
		if len(firstsp) != 2 {
			panic("input is messed up")
		}
		checkLen(firstsp, 2)
		for i, c := range firstsp {
			firsti[i], err = strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
		}
		firstStart, firstEnd := firsti[0], firsti[1]
		sec := sp[1]
		secsp := strings.Split(sec, "-")
		checkLen(secsp, 2)
		for i, c := range secsp {
			seci[i], err = strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
		}
		secStart, secEnd := seci[0], seci[1]
		if firstStart <= secStart && firstEnd >= secEnd {
			res++
			resTwo++
		} else if secStart <= firstStart && secEnd >= firstEnd {
			res++
			resTwo++
		} else if firstStart <= secEnd && firstEnd >= secEnd {
			resTwo++
		} else if secStart <= firstEnd && secEnd >= firstEnd {
			resTwo++
		}
	}
	fmt.Printf("Part 1: %d\n", res)
	fmt.Printf("Part 2: %d\n", resTwo)
}

// Because fuck writing this 3 times.
// Method is used to eliminate bound checking
func checkLen(arr []string, wanted int) {
	if len(arr) != wanted {
		panic("input is messed up")
	}
}
