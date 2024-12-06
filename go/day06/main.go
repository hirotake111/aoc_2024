package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const example = `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

func main() {
	if v := part1(example); v != 41 {
		log.Fatalf("part1 failed: %v != 41", v)
	}
	rawData, err := os.ReadFile("./data/day06.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part1 -> %d\n", part1(string(rawData)))
}

func part1(s string) int {
	grid := stringToGrid(s)
	// p(grid)
	// Identify initial guard position
	m, n := len(grid), len(grid[0])
	r, c := find_guard(grid)
	grid[r][c] = 'X'
	// fmt.Printf("x: %d, y: %d\n", x, y)
	steps := 1
	for r >= 0 && r < m && c >= 0 && c < n {
		for _, delta := range [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			for {
				nr, nc := r+delta[0], c+delta[1]
				if nr < 0 || nr >= m || nc < 0 || nc >= n {
					// Out of the grid
					return steps
				}
				if grid[nr][nc] == '#' {
					// Change the direction
					break
				}
				r, c = nr, nc
				if grid[r][c] != 'X' {
					grid[r][c] = 'X'
					steps++
				}
				// fmt.Println(steps)
				// p(grid)
			}
		}
	}
	return -1
}

func stringToGrid(s string) [][]byte {
	rows := strings.Split(strings.Trim(s, "\n"), "\n")
	grid := make([][]byte, 0, len(rows))
	for _, row := range rows {
		grid = append(grid, []byte(row))
	}
	return grid
}

func find_guard(grid [][]byte) (int, int) {
	for i, row := range grid {
		for j, pos := range row {
			if pos == '^' {
				return i, j
			}
		}
	}
	panic("unreachable")
}

func p(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Println("")
}
