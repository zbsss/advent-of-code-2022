package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseStacks(lines []string) [][]byte {
	width := len(lines[0])
	stacksCount := (width + 1) / 4
	stacks := make([][]byte, stacksCount)

	for l := len(lines) - 2; l >= 0; l-- {
		line := lines[l]
		for i := 1; i < width; i += 4 {
			if line[i] != ' ' {
				stacks[i/4] = append(stacks[i/4], line[i])
			}
		}
	}

	return stacks
}

type Op struct {
	source, dest, count int
}

func parseOp(raw string) int {
	num, err := strconv.Atoi(raw)

	if err != nil {
		log.Fatal("Failed to parse operation!")
	}

	return num
}

func parseOps(lines []string) []Op {
	ops := make([]Op, len(lines))
	for _, line := range lines {
		fields := strings.Fields(line)
		ops = append(ops, Op{
			count:  parseOp(fields[1]),
			source: parseOp(fields[3]) - 1,
			dest:   parseOp(fields[5]) - 1,
		})
	}
	return ops
}

func getResult(stacks [][]byte) string {
	var result []byte
	for _, stack := range stacks {
		result = append(result, stack[len(stack)-1])
	}
	return string(result)
}

func a(stacks [][]byte, ops []Op) string {

	for _, op := range ops {
		chunk := stacks[op.source][len(stacks[op.source])-op.count:]

		for i := len(chunk) - 1; i >= 0; i-- {
			stacks[op.dest] = append(stacks[op.dest], chunk[i])
		}

		stacks[op.source] = stacks[op.source][:len(stacks[op.source])-op.count]
	}

	return getResult(stacks)
}

func b(stacks [][]byte, ops []Op) string {
	for _, op := range ops {
		chunk := stacks[op.source][len(stacks[op.source])-op.count:]

		stacks[op.dest] = append(stacks[op.dest], chunk...)

		stacks[op.source] = stacks[op.source][:len(stacks[op.source])-op.count]
	}

	return getResult(stacks)
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	sc := bufio.NewScanner(input)

	var stackLines []string
	var opLines []string

	isStacks := true

	for sc.Scan() {
		line := sc.Text()

		if line == "" {
			isStacks = false
		} else if isStacks {
			stackLines = append(stackLines, line)
		} else {
			opLines = append(opLines, line)
		}
	}

	stacks := parseStacks(stackLines)
	ops := parseOps(opLines)

	stacksDeepCopy := make([][]byte, len(stacks))
	for i, stack := range stacks {
		stacksDeepCopy[i] = append(stacksDeepCopy[i], stack...)
	}

	solutionA := a(stacksDeepCopy, ops)
	solutionB := b(stacks, ops)

	fmt.Println("A: ", solutionA)
	fmt.Println("B: ", solutionB)
}
