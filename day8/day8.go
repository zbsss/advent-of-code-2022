package main

import (
	"bufio"
	"fmt"
	"os"
)

func countVisible(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	visible := make([][]bool, n)
	for i := range visible {
		visible[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		l_max := -1
		r_max := -1

		for j := 0; j < m; j++ {
			l := j
			r := len(visible) - j - 1

			if grid[i][l] > l_max {
				l_max = grid[i][l]
				visible[i][l] = true
			}

			if grid[i][r] > r_max {
				r_max = grid[i][r]
				visible[i][r] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		l_max := -1
		r_max := -1

		for j := 0; j < n; j++ {
			l := j
			r := len(visible) - j - 1

			if grid[l][i] > l_max {
				l_max = grid[l][i]
				visible[l][i] = true
			}

			if grid[r][i] > r_max {
				r_max = grid[r][i]
				visible[r][i] = true
			}
		}
	}

	count := 0
	for i := range visible {
		for j := range visible[i] {
			if visible[i][j] {
				count++
			}
		}
	}
	return count
}

func maxScenicScore(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	result := 0
	moves := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			score := 1

			for _, move := range moves {
				di, dj := move[0], move[1]

				ni := i + di
				nj := j + dj

				startingHeight := grid[i][j]

				visible := 0

				for 0 <= ni && ni < n && 0 <= nj && nj < m {
					visible++

					if grid[ni][nj] >= startingHeight {
						break
					}

					ni += di
					nj += dj
				}

				score *= visible
			}

			if score > result {
				result = score
			}
		}
	}

	return result
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)

	i := 0
	var grid [][]int

	for sc.Scan() {
		line := sc.Text()

		grid = append(grid, make([]int, len(line)))

		for j, raw := range line {
			height := int(raw - '0')
			grid[i][j] = height
		}

		i++
	}

	solutionA := countVisible(grid)
	fmt.Println(solutionA)

	solutionB := maxScenicScore(grid)
	fmt.Println(solutionB)
}
