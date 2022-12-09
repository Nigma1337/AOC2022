package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type ropeObj struct {
	X    int
	Y    int
	Tail *ropeObj
}

type void struct{}

var member void

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Part 1: %d\n", solve(readFile, 1))
	readFile.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", solve(readFile, 9))
}

func solve(readFile io.Reader, knots int) int {
	visited := make(map[string]void)
	start := ropeObj{}
	cur := &start
	// Initialize list
	for i := 0; i < knots; i++ {
		cur.Tail = &ropeObj{}
		cur = cur.Tail
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		text := fileScanner.Text()
		sp := strings.Split(text, " ")
		amount, err := strconv.Atoi(sp[1])
		if err != nil {
			panic("bad")
		}
		direction := sp[0]
		for i := 0; i < amount; i++ {
			rope := &start
			switch direction {
			case "R":
				rope.Right()
			case "L":
				rope.Left()
			case "U":
				rope.Up()
			case "D":
				rope.Down()
			}
			for i := 0; i < knots; i++ {
				rope.PullTail()
				rope = rope.Tail
				if rope.Tail == nil {
					location := fmt.Sprintf("%d.%d", rope.X, rope.Y)
					visited[location] = member
				}

			}
		}
	}
	return len(visited)
}

func (n *ropeObj) PullTail() {
	//fmt.Printf("values: X: %d, Y: %d, tailX: %d, tailY: %d\n", n.X, n.Y, n.Tail.X, n.Tail.Y)
	// We're apart on the X axis
	if n.X > n.Tail.X+1 {
		//fmt.Printf("Pulled tail right ")
		n.Tail.X++
		if n.Y > n.Tail.Y {
			//fmt.Printf("and up\n")
			n.Tail.Y++
		} else if n.Y < n.Tail.Y {
			//fmt.Printf("and down\n")
			n.Tail.Y--
		}
	} else if n.X < n.Tail.X-1 {
		//fmt.Printf("Pulled tail left ")
		if n.Y > n.Tail.Y {
			//fmt.Printf("and up\n")
			n.Tail.Y++
		} else if n.Y < n.Tail.Y {
			//fmt.Printf("and down\n")
			n.Tail.Y--
		}
		n.Tail.X--
	}
	// We're apart on the Y axis
	if n.Y > n.Tail.Y+1 {
		//fmt.Printf("Pulled tail up ")
		n.Tail.Y++
		if n.X > n.Tail.X {
			//fmt.Printf("and right\n")
			n.Tail.X++
		} else if n.X < n.Tail.X {
			//fmt.Printf("and left\n")
			n.Tail.X--
		}
	} else if n.Y < n.Tail.Y-1 {
		//fmt.Printf("Pulled tail down")
		n.Tail.Y--
		if n.X > n.Tail.X {
			//fmt.Printf("and right\n")
			n.Tail.X++
		} else if n.X < n.Tail.X {
			//fmt.Printf("and left\n")
			n.Tail.X--
		}
	}
	//fmt.Printf("Location: %d,%d\n", n.Tail.X, n.Tail.Y)
}

func (n *ropeObj) Right() {
	n.X++
}

func (n *ropeObj) Left() {
	n.X--
}
func (n *ropeObj) Up() {
	n.Y++
}
func (n *ropeObj) Down() {
	n.Y--
}
