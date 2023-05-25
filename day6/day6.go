package main

import (
	"bufio"
	"fmt"
	"os"
)

func findDistinctSubsequence(seqLength int, text string) (int, error) {
	for i := 0; i < len(text)-seqLength; i++ {
		chars := map[byte]bool{}
		for j := 0; j < seqLength; j++ {
			if chars[text[i+j]] {
				break
			}

			chars[text[i+j]] = true
		}

		if len(chars) == seqLength {
			return i + seqLength, nil
		}
	}

	return 0, fmt.Errorf("invalid input")
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := sc.Text()

		solutionA, _ := findDistinctSubsequence(4, line)
		solutionB, _ := findDistinctSubsequence(14, line)

		fmt.Println(solutionA)
		fmt.Println(solutionB)
	}
}
