package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("./data/day10.txt")
	if err != nil {
		panic(err)
	}
	grid := getGrid(string(input))
	fmt.Printf("Part1 -> %d\n", part1(grid))
	fmt.Printf("Part2 -> %d\n", part2(grid))
}

func getGrid(input string) [][]byte {
	grid := make([][]byte, 0)
	for _, s := range strings.Split(strings.Trim(input, "\n"), "\n") {
		row := make([]byte, len(s))
		for i := 0; i < len(s); i++ {
			row[i] = s[i] - '0'
		}
		grid = append(grid, row)
	}
	return grid
}

func part1(grid [][]byte) int {
	var total int
	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				seen := make(map[[2]int]struct{})
				seen[[2]int{i, j}] = struct{}{}
				total += trail(i, j, grid, seen)
			}
		}
	}
	return total
}

func part2(grid [][]byte) int {
	var total int
	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				seen := make(map[[2]int]struct{})
				seen[[2]int{i, j}] = struct{}{}
				total += trail2(i, j, grid)
			}
		}
	}
	return total
}

var directions = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func trail2(i, j int, grid [][]byte) int {
	var total int
	if grid[i][j] == 9 {
		return 1
	}
	for _, dir := range directions {
		r, c := i+dir[0], j+dir[1]
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			continue
		}
		if grid[i][j]+1 == grid[r][c] {
			total += trail2(r, c, grid)
		}
	}
	return total
}

func trail(i, j int, grid [][]byte, seen map[[2]int]struct{}) int {
	var total int
	if grid[i][j] == 9 {
		return 1
	}
	for _, dir := range directions {
		r, c := i+dir[0], j+dir[1]
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			continue
		}
		if _, ok := seen[[2]int{r, c}]; !ok && grid[i][j]+1 == grid[r][c] {
			seen[[2]int{r, c}] = struct{}{}
			total += trail(r, c, grid, seen)
		}
	}
	return total
}

func p(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println("")
}
