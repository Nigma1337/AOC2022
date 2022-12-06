package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var stack_count int
	fname := "input.txt"
	readFile, err := os.Open(fname)
	if fname == "input.txt" {
		stack_count = 9
	} else {
		stack_count = 3
	}
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	stacks := make([]string, stack_count)
	stacksTwo := make([]string, stack_count)
	fileScanner.Split(bufio.ScanLines)
	// Create initial stacks
	for fileScanner.Scan() {
		text := fileScanner.Text()
		if text[:2] == " 1" {
			break
		}
		text = " " + text
		for i := 0; i < stack_count; i++ {
			if text[2] != ' ' {
				stacks[i] = string(stacks[i]) + string(text[2])
			}
			text = text[4:]
		}
	}
	// Create new stacks from initial config
	// := doesn't work as we're dealing with a slice, so we need to call copy explicitly
	copy(stacksTwo, stacks)
	// skip emptyline
	fileScanner.Scan()
	for fileScanner.Scan() {
		text := fileScanner.Text()
		sp := strings.Split(text, " ")
		count, _ := strconv.Atoi(sp[1])
		from, _ := strconv.Atoi(sp[3])
		to, _ := strconv.Atoi(sp[5])
		// zero index
		from = from - 1
		to = to - 1
		for i := 0; i < count; i++ {
			// doing :1 instead of 0 returns a string instead of rune, and lets us skip a string() call later
			top := stacks[from][:1]
			stacks[from] = stacks[from][1:]
			stacks[to] = top + string(stacks[to])
		}
		top := stacksTwo[from][:count]
		stacksTwo[from] = stacksTwo[from][count:]
		stacksTwo[to] = top + string(stacksTwo[to])
	}
	fmt.Print("Part 1: ")
	for _, i := range stacks {
		fmt.Print(string(i[0]))
	}
	fmt.Print("\nPart 2: ")
	for _, i := range stacksTwo {
		fmt.Print(string(i[0]))
	}
	fmt.Println()
}
