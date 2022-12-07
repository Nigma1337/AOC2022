package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Basic idea: Create binary tree from directory structure
type node struct {
	Size     uint64
	Name     string
	Children []*node
	Parent   *node
}

func main() {
	var name string
	root := node{Name: "/"}
	current := &root
	fname := "input.txt"
	readFile, err := os.Open(fname)
	if err != nil {
		panic("Failed to open file")
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		text := fileScanner.Text()
		if text == "$ ls" {
			continue
		}
		sp := strings.Split(text, " ")
		start := sp[0]
		if start == "dir" {
			name = sp[1]
			index := dirInArray(name, current.Children)
			// if -1, its not in array
			if index == -1 {
				new := root.New(name, current)
				current.Children = append(current.Children, new)
			}
		} else if start == "$" {
			name = sp[2]
			if name == ".." {
				current = current.Parent
				continue
			} else if name == "/" {
				current = &root
				continue
			}
			index := dirInArray(name, current.Children)
			// if -1, its not in array
			if index == -1 {
				new := root.New(name, current)
				current.Children = append(current.Children, new)
				current = new
			} else {
				current = current.Children[index]
			}
		} else {
			// file
			length, err := strconv.ParseUint(start, 10, 64)
			if err != nil {
				panic("bad")
			}
			current.Size += length
		}
	}
	dir_len(&root)

	//Get max uint64 (18,446,744,073,709,551,615)
	res, resTwo := uint64(0), ^uint64(0)

	calculatePartOne(&root, uint64(100000), &res)
	calculatePartTwo(&root, 8381165, &resTwo)
	fmt.Printf("Part 1: %d\n", res)
	fmt.Printf("Part 2: %d\n", resTwo)
}

func (n *node) New(name string, parent *node) *node {
	return &node{Name: name, Parent: parent}
}

func (n *node) Print(push int) {
	fmt.Printf("%*s%s %d\n", push, "", n.Name, n.Size)
	for _, child := range n.Children {
		child.Print(push + 1)
	}
}

func dirInArray(name string, children []*node) int {
	for i, child := range children {
		if child.Name == name {
			return i
		}
	}
	return -1
}

func dir_len(root *node) uint64 {
	for _, child := range root.Children {
		child.Size = dir_len(child)
		root.Size += child.Size
	}
	return root.Size
}

func calculatePartOne(root *node, than uint64, result *uint64) {
	for _, child := range root.Children {
		calculatePartOne(child, than, result)
	}
	if root.Size < than {
		*result = *result + root.Size
	}
}

func calculatePartTwo(root *node, than uint64, result *uint64) {
	for _, child := range root.Children {
		calculatePartTwo(child, than, result)
	}
	if root.Size > than && root.Size < *result {
		*result = root.Size
	}
}
