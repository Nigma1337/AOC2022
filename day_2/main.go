package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	// string is possible matchup, int is amount of points
	pos := map[string]int{
		"A Y": 8, // Rock vs Paper
		"A X": 4, // Rock vs Rock
		"A Z": 3, // Rock vs Scissor
		"B Y": 5, // Paper vs Paper
		"B X": 1, // Paper vs Rock
		"B Z": 9, // Paper vs Scissor
		"C Y": 2, // Scissor vs Paper
		"C X": 7, // Scissor vs Rock
		"C Z": 6, // Scissor vs Scissor
	}
	// string is "hand result", int is amount of points
	posTwo := map[string]int{
		"A Y": 4, // Rock vs Rock (draw)
		"A X": 3, // Rock vs Scissor (lose)
		"A Z": 8, // Rock vs Paper (win)
		"B Y": 5, // Paper vs Paper (draw)
		"B X": 1, // Paper vs Rock (lose)
		"B Z": 9, // Paper vs Scissor (win)
		"C Y": 6, // Scissor vs Scissor (draw)
		"C X": 2, // Scissor vs Paper (lose)
		"C Z": 7, // Scissor vs Rock (win)
	}
	score := 0
	scoreTwo := 0
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		text := fileScanner.Text()
		score += pos[text]
		scoreTwo += posTwo[text]
	}
	fmt.Printf("Part 1: %d\n", score)
	fmt.Printf("Part 2: %d\n", scoreTwo)
}
