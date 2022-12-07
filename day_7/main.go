package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Basic idea: Create binary tree from directory structure
type node struct {
	Size     int
	Name     string
	Children []*node
	Parent   *node
}

var result int
var resTwo []int

func main() {
	var name string
	root := node{Name: "/"}
	current := &root
	fname := "input.txt"
	readFile, err := os.Open(fname)
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		text := fileScanner.Text()
		if text[:3] == "dir" {
			name = text[4:]
			index := dirInArray(name, current.Children)
			// if -1, its not in array
			if index == -1 {
				new := root.New(name, current)
				current.Children = append(current.Children, new)
			}
		} else if text[:4] == "$ cd" {
			name = text[5:]
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
		} else if text == "$ ls" {
			continue
		} else {
			// file
			length, err := strconv.Atoi(strings.Split(text, " ")[0])
			if err != nil {
				panic("bad")
			}
			current.Size += length
		}
	}
	root.Print(0)
	dir_len(&root)
	root.Print(0)
	LT(&root, 100000)
	fmt.Printf("Part 1: %d\n", result)
	LTwo(&root, 8381165)
	sort.Ints(resTwo)
	fmt.Printf("Part 2: %d\n", resTwo[0])
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

func dir_len(root *node) int {
	for _, child := range root.Children {
		child.Size = dir_len(child)
		root.Size += child.Size
	}
	return root.Size
}

func LT(root *node, than int) {
	for _, child := range root.Children {
		LT(child, than)
	}
	if root.Size < than {
		result += root.Size
	}
}

func LTwo(root *node, than int) {
	for _, child := range root.Children {
		LTwo(child, than)
	}
	if root.Size > than {
		resTwo = append(resTwo, root.Size)
	}
}
