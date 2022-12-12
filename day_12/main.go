package main

import (
	"container/list"
	_ "embed"
	"fmt"
	"strings"
)

type location struct {
	x     int
	y     int
	value rune
	depth int
}

type startLoc struct {
	location location
	list     *list.List
}

type land struct {
	visited   map[string]bool
	landscape [][]rune
	colc      int
	rowc      int
}

//go:embed input.txt
var input string

func main() {
	//var st int
	var startLocation startLoc
	var endLocation startLoc
	//var starts []startLoc
	//var wg sync.WaitGroup
	// Get possible starts for part 2 (+1 is needed as the count doesn't count our S start)
	//possible_starts := strings.Count(input, "a")
	//starts = make([]startLoc, possible_starts)
	sp := strings.Split(input, "\n")
	rowc := len(sp[0])
	colc := len(sp)
	visited := make(map[string]bool)
	var landscape = make([][]rune, colc)
	for i := range landscape {
		landscape[i] = make([]rune, rowc)
	}
	for i, c := range sp {
		for j, b := range c {
			if b != 'E' {
				landscape[i][j] = b
			} else {
				endLocation = startLoc{location: location{y: i, x: j, value: 'z' + 1}, list: list.New()}
				landscape[i][j] = 'z' + 1
			}
			if b == 'S' {
				startLocation = startLoc{location: location{y: i, x: j, value: 'a'}, list: list.New()}
				//starts[st] = startLocation
				//st++
				landscape[i][j] = 'a'
			} //else if b == 'a' {
			//	star := startLoc{location: location{y: i, x: j, value: 'a'}, list: list.New()}
			//	starts[st] = star
			//	st++
			//}
		}
	}
	scape := land{landscape: landscape, colc: colc, rowc: rowc, visited: visited}
	res := solve(scape, startLocation.list, startLocation.location)
	visited = make(map[string]bool)
	scape = land{landscape: landscape, colc: colc, rowc: rowc, visited: visited}
	fmt.Printf("Part 1: %d\n", res)
	resTwo := solveTwo(scape, endLocation.list, endLocation.location)
	fmt.Printf("Part 2: %d\n", resTwo)
	//wg.Add(len(starts))
	//for _, a := range starts {
	//	go func(a startLoc) {
	//		defer wg.Done()
	//		visited = make(map[string]bool)
	//		scape = land{landscape: landscape, colc: colc, rowc: rowc, visited: visited}
	//		cool := solve(scape, a.list, a.location)
	//		if cool < res {
	//			res = cool
	//		}
	//	}(a)
	//}
	//wg.Wait()
}

func solve(scape land, fifo *list.List, start location) int {
	scape.get_friends(start, fifo)
	for val := fifo.Front(); val != nil; val = val.Next() {
		loc := location(val.Value.(location))
		if loc.value == 'z'+1 {
			return loc.depth
		}
		scape.get_friends(loc, fifo)
	}
	return 420
}

func solveTwo(scape land, fifo *list.List, start location) int {
	scape.get_enemies(start, fifo)
	for val := fifo.Front(); val != nil; val = val.Next() {
		loc := location(val.Value.(location))
		if loc.value == 'a' {
			return loc.depth
		}
		scape.get_enemies(loc, fifo)
	}
	return 420
}

func (landscape land) get_friends(loc location, fifo *list.List) {
	// Check "edge" cases (get it, like edge cases, but its also literally checking the edge)
	if loc.x == landscape.rowc-1 {
		landscape.checkBehind(loc, fifo)
	} else if loc.x == 0 {
		landscape.checkFront(loc, fifo)
	} else {
		landscape.checkFront(loc, fifo)
		landscape.checkBehind(loc, fifo)
	}
	if loc.y == landscape.colc-1 {
		landscape.checkAbove(loc, fifo)
	} else if loc.y == 0 {
		landscape.checkBelow(loc, fifo)
	} else {
		landscape.checkAbove(loc, fifo)
		landscape.checkBelow(loc, fifo)
	}
}

func (landscape land) get_enemies(loc location, fifo *list.List) {
	// Check "edge" cases (get it, like edge cases, but its also literally checking the edge)
	if loc.x == landscape.rowc-1 {
		landscape.checkBehindRev(loc, fifo)
	} else if loc.x == 0 {
		landscape.checkFrontRev(loc, fifo)
	} else {
		landscape.checkFrontRev(loc, fifo)
		landscape.checkBehindRev(loc, fifo)
	}
	if loc.y == landscape.colc-1 {
		landscape.checkAboveRev(loc, fifo)
	} else if loc.y == 0 {
		landscape.checkBelowRev(loc, fifo)
	} else {
		landscape.checkAboveRev(loc, fifo)
		landscape.checkBelowRev(loc, fifo)
	}
}

func (landscape land) checkBelow(loc location, fifo *list.List) {
	belowy := fmt.Sprintf("%d.%d", loc.x, loc.y+1)
	if !landscape.visited[belowy] && landscape.landscape[loc.y+1][loc.x] <= loc.value+1 {
		loc := location{x: loc.x, y: loc.y + 1, value: landscape.landscape[loc.y+1][loc.x], depth: loc.depth + 1}
		fifo.PushBack(loc)
		landscape.visited[belowy] = true
	}
}

func (landscape land) checkAbove(loc location, fifo *list.List) {
	abovey := fmt.Sprintf("%d.%d", loc.x, loc.y-1)
	if !landscape.visited[abovey] && landscape.landscape[loc.y-1][loc.x] <= loc.value+1 {
		loc := location{x: loc.x, y: loc.y - 1, value: landscape.landscape[loc.y-1][loc.x], depth: loc.depth + 1}
		fifo.PushBack(loc)
		landscape.visited[abovey] = true
	}
}

func (landscape land) checkFront(loc location, fifo *list.List) {
	frontx := fmt.Sprintf("%d.%d", loc.x+1, loc.y)
	if !landscape.visited[frontx] && landscape.landscape[loc.y][loc.x+1] <= loc.value+1 {
		loc := location{x: loc.x + 1, y: loc.y, value: landscape.landscape[loc.y][loc.x+1], depth: loc.depth + 1}
		fifo.PushBack(loc)
		landscape.visited[frontx] = true
	}
}

func (landscape land) checkBehind(loc location, fifo *list.List) {
	behindx := fmt.Sprintf("%d.%d", loc.x-1, loc.y)
	if !landscape.visited[behindx] && landscape.landscape[loc.y][loc.x-1] <= loc.value+1 {
		loc := location{x: loc.x - 1, y: loc.y, value: landscape.landscape[loc.y][loc.x-1], depth: loc.depth + 1}
		fifo.PushBack(loc)
		landscape.visited[behindx] = true
	}
}

func (landscape land) checkBelowRev(loc location, fifo *list.List) {
	belowy := fmt.Sprintf("%d.%d", loc.x, loc.y+1)
	if !landscape.visited[belowy] && (loc.value-landscape.landscape[loc.y+1][loc.x]) <= 1 {
		loc := location{x: loc.x, y: loc.y + 1, value: landscape.landscape[loc.y+1][loc.x], depth: loc.depth + 1}
		fifo.PushBack(loc)
		landscape.visited[belowy] = true
	}
}

func (landscape land) checkAboveRev(loc location, fifo *list.List) {
	abovey := fmt.Sprintf("%d.%d", loc.x, loc.y-1)
	if !landscape.visited[abovey] && (loc.value-landscape.landscape[loc.y-1][loc.x]) <= 1 {
		loc := location{x: loc.x, y: loc.y - 1, value: landscape.landscape[loc.y-1][loc.x], depth: loc.depth + 1}
		fifo.PushBack(loc)
		landscape.visited[abovey] = true
	}
}

func (landscape land) checkFrontRev(loc location, fifo *list.List) {
	frontx := fmt.Sprintf("%d.%d", loc.x+1, loc.y)
	if !landscape.visited[frontx] && (loc.value-landscape.landscape[loc.y][loc.x+1]) <= 1 {
		loc := location{x: loc.x + 1, y: loc.y, value: landscape.landscape[loc.y][loc.x+1], depth: loc.depth + 1}
		fifo.PushBack(loc)
		landscape.visited[frontx] = true
	}
}

func (landscape land) checkBehindRev(loc location, fifo *list.List) {
	behindx := fmt.Sprintf("%d.%d", loc.x-1, loc.y)
	if !landscape.visited[behindx] && (loc.value-landscape.landscape[loc.y][loc.x-1]) <= 1 {
		loc := location{x: loc.x - 1, y: loc.y, value: landscape.landscape[loc.y][loc.x-1], depth: loc.depth + 1}
		fifo.PushBack(loc)
		landscape.visited[behindx] = true
	}
}
