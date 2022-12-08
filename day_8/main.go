package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var line int
	var forrest [99][99]int
	//var wg sync.WaitGroup
	readFile, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}
	col := 0
	res := 0
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
		res++
	}
	score := 0
	for i := 1; i < line; i++ {
		for j := 1; j < line; j++ {
			//go func(i int, j int) {
			//	defer wg.Done()
			hidden, sc := hidden(&forrest, i, j)
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

func hidden(forrest *[99][99]int, line int, col int) (bool, int) {
	val := &forrest[line][col]
	down, up, left, right := false, false, false, false
	da, ua, la, ra := 0, 0, 0, 0
	// Check down
	for i := line; i < 99; i++ {
		if i == line {
			continue
		}
		//fmt.Printf("Value %d on line %d col %d is hidden down\n", *val, line, col)
		da++
		if forrest[i][col] >= *val {
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
		if forrest[i][col] >= *val {
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
		if forrest[line][i] >= *val {
			//fmt.Printf("Value %d on line %d col %d is hidden to the left\n", *val, line, col)
			left = true
			break
		}
	}
	//// Check right
	for i := col; i < 99; i++ {
		//fmt.Printf("%d %d %d %d %d\n", *val, line, col, line, i)
		if i == col {
			continue
		}
		ra++
		if forrest[line][i] >= *val {
			//fmt.Printf("Value %d on line %d col %d is hidden to the right\n", *val, line, col)
			right = true
			break
		}
	}
	hidden := left && right && up && down
	//fmt.Printf("Value %d on line %d col %d is hidden: %t\n", *val, line, col, hidden)
	fmt.Println(da*ua*la*ra, da, ua, la, ra)
	return hidden, (da * ua * la * ra)
}
