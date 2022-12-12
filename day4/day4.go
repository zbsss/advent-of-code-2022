package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type interval struct {
	start int
	end   int
}

func (i interval) isFullyContainedIn(other interval) bool {
	return other.start <= i.start && i.end <= other.end
}

func (i interval) isOverlapping(other interval) bool {
	return i.start <= other.start && other.start <= i.end
}

func extractNumbers(str string) []int {
	re := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)
	matches := re.FindAllStringSubmatch(str, -1)
	var numbers []int
	for _, match := range matches {
		for _, char := range match[1:] {
			num, _ := strconv.Atoi(char)
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func extractIntervals(str string) (a, b interval) {
	a = interval{}
	b = interval{}

	fmt.Sscanf(str, "%d-%d,%d-%d", &a.start, &a.end, &b.start, &b.end)
	return a, b
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	fullyContainedCount := 0
	overlapCount := 0

	for sc.Scan() {
		// nums := extractNumbers(sc.Text())

		// a := interval{nums[0], nums[1]}
		// b := interval{nums[2], nums[3]}
		a, b := extractIntervals(sc.Text())

		if a.isFullyContainedIn(b) || b.isFullyContainedIn(a) {
			fullyContainedCount += 1
		}

		if a.isOverlapping(b) || b.isOverlapping(a) {
			fmt.Println(a, b)
			overlapCount += 1
		}
	}

	fmt.Println("Fully contained count: ", fullyContainedCount)
	fmt.Println("All overlaps count: ", overlapCount)
}
