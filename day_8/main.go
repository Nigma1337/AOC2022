package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var line int
	var length int
	var score int
	//var wg sync.WaitGroup
	var res int
	readFile, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	for i := range readFile {
		if readFile[i] == '\n' {
			break
		}
		length++
	}
	forrest := make([][]int, length)
	for i := range forrest {
		forrest[i] = make([]int, length)
	}
	col := 0
	for _, c := range readFile {
		if c == '\n' {
			col = 0
			line++
			continue
		}
		height, err := strconv.Atoi(string(c))
		if err != nil {
			panic("bad")
		}
		forrest[line][col] = height
		col++
	}
	res = length * length
	//i32 := int32(res)
	//workers := (length - 2) * (length - 2)
	//wg.Add(workers)
	for i := 1; i < line; i++ {
		for j := 1; j < line; j++ {
			//go func(i int, j int) {
			//defer wg.Done()
			hidden, sc := hidden(forrest, i, j, length)
			if hidden {
				res--
			}
			if sc > score {
				score = sc
			}
			//}(i, j)
		}
	}
	//wg.Wait()
	fmt.Printf("Part 1: %d\n", res)
	fmt.Printf("Part 2: %d\n", score)
}

func hidden(forrest [][]int, line int, col int, length int) (bool, int) {
	val := forrest[line][col]
	down, up, left, right := false, false, false, false
	da, ua, la, ra := 0, 0, 0, 0
	// Check down
	for i := line; i < length; i++ {
		if i == line {
			continue
		}
		//fmt.Printf("Value %d on line %d col %d is hidden down\n", *val, line, col)
		da++
		if forrest[i][col] >= val {
			//fmt.Printf("Value %d on line %d col %d is hidden down!\n", *val, line, col)
			down = true
			break
		}
	}
	////// Check up
	for i := line; i >= 0; i-- {
		if i == line {
			continue
		}
		ua++
		//fmt.Printf("Value %d on line %d col %d is hidden up\n", *val, line, col)
		if forrest[i][col] >= val {
			//fmt.Printf("Value %d on line %d col %d is hidden up!\n", *val, line, col)
			up = true
			break
		}
	}
	//// Check left
	for i := col; i >= 0; i-- {
		if i == col {
			continue
		}
		la++
		if forrest[line][i] >= val {
			//fmt.Printf("Value %d on line %d col %d is hidden to the left\n", *val, line, col)
			left = true
			break
		}
	}
	//// Check right
	for i := col; i < length; i++ {
		//fmt.Printf("%d %d %d %d %d\n", *val, line, col, line, i)
		if i == col {
			continue
		}
		ra++
		if forrest[line][i] >= val {
			//fmt.Printf("Value %d on line %d col %d is hidden to the right\n", *val, line, col)
			right = true
			break
		}
	}
	hidden := left && right && up && down
	//fmt.Printf("Value %d on line %d col %d is hidden: %t\n", *val, line, col, hidden)
	return hidden, (da * ua * la * ra)
}
