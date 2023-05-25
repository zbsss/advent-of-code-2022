package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FsNode struct {
	name     string
	isDir    bool
	size     int
	parent   *FsNode
	children map[string]*FsNode
}

func NewFsNode(name string, isDir bool) *FsNode {
	node := FsNode{
		name:  name,
		isDir: isDir,
	}

	if isDir {
		node.children = map[string]*FsNode{}
	}

	return &node
}

func parseFsNode(parent *FsNode, line string) *FsNode {
	fields := strings.Fields(line)

	size, err := strconv.Atoi(fields[0])
	isDir := err != nil

	node := NewFsNode(fields[1], isDir)

	node.parent = parent
	node.size = size

	return node
}

func updateDirSizes(node *FsNode) {
	for _, child := range node.children {
		updateDirSizes(child)
		node.size += child.size
	}
}

func a(node *FsNode) int {
	acc := 0

	if node.isDir && node.size < 100000 {
		acc += node.size
	}

	for _, child := range node.children {
		acc += a(child)
	}

	return acc
}

func b(missing int, node *FsNode) int {
	smallest := node.size

	for _, child := range node.children {
		if child.isDir && child.size > missing {
			if res := b(missing, child); res < smallest {
				smallest = res
			}
		}
	}

	return smallest
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)

	// skip first line with "cd /"
	sc.Scan()

	root := NewFsNode("/", true)
	current := root

	for sc.Scan() {
		line := sc.Text()

		if line[0] == '$' {
			fields := strings.Fields(line)

			if fields[1] == "cd" {
				if fields[2] == "/" {
					current = root
				} else if fields[2] == ".." {
					current = current.parent
				} else {
					current = current.children[fields[2]]
				}
			}
		} else {
			node := parseFsNode(current, line)
			current.children[node.name] = node
		}
	}

	updateDirSizes(root)

	solutionA := a(root)
	fmt.Println(solutionA)

	free := 70000000 - root.size
	missing := 30000000 - free

	solutionB := b(missing, root)
	fmt.Println(solutionB)
}
