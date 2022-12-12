package main

import (
	"bufio"
	"fmt"
	"os"
)

func getPriority(char rune) int {
	ascii := int(char)

	// a-z
	if ascii >= 97 {
		return ascii - 97 + 1
	}
	// A-Z
	return ascii - 65 + 27
}

func makeSet(items string) map[rune]bool {
	set := make(map[rune]bool)
	for _, item := range items {
		set[item] = true
	}
	return set
}

func a() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	total := 0

	for sc.Scan() {
		size := len(sc.Text())

		first := make(map[byte]bool)
		duplicates := make(map[byte]bool)

		for i := 0; i < size/2; i++ {
			first[sc.Text()[i]] = true
		}

		for i := size / 2; i < size; i++ {
			char := sc.Text()[i]

			if first[char] && !duplicates[char] {
				duplicates[char] = true
				total += getPriority(rune(char))
			}
		}
	}

	fmt.Println("Result: ", total)
}

func b() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	total := 0

	for sc.Scan() {
		first := makeSet(sc.Text())
		sc.Scan()
		second := makeSet(sc.Text())
		sc.Scan()

		for _, item := range sc.Text() {
			if first[item] && second[item] {
				total += getPriority(item)
				break
			}
		}
	}

	fmt.Println("Result: ", total)
}

func main() {
	a()
	b()
}
