package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	score := 0

	// X means you need to lose,
	// Y means you need to end the round in a draw
	// Z means you need to win.

	scores := map[string]int{"B X": 1, "C X": 2, "A X": 3, "A Y": 4, "B Y": 5, "C Y": 6, "C Z": 7, "A Z": 8, "B Z": 9}

	for sc.Scan() {
		score += scores[sc.Text()]
	}

	fmt.Println("Result: ", score)
}
