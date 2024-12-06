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
	if v := part2(example); v != 6 {
		log.Fatalf("part2 failed: %v != 6", v)
	}
	rawData, err := os.ReadFile("./data/day06.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part1 -> %d\n", part1(string(rawData)))
	fmt.Printf("Part2 -> %d\n", part2(string(rawData)))
}

func part2(s string) int {
	grid := stringToGrid(s)
	r, c := find_guard(grid)
	paths := collectPath(grid, r, c)[1:]
	format(grid)
	var options int
	for _, p := range paths {
		if canMakeLoop(grid, r, c, p[0], p[1]) {
			options++
		}
	}
	return options
}

func format(grid [][]byte) {
	for i, row := range grid {
		for j, c := range row {
			if c == '.' {
				grid[i][j] = 0
			} else {
				grid[i][j] = 16
			}
		}
	}
}

func canMakeLoop(src [][]byte, r, c, a, b int) bool {
	grid := make([][]byte, 0, len(src))
	for _, srcRow := range src {
		row := make([]byte, len(srcRow))
		copy(row, srcRow)
		grid = append(grid, row)
	}
	grid[r][c] = 0
	grid[a][b] = 16
	m, n := len(grid), len(grid[0])
	for {
		// 1: up, 2: right, 4: down, 8: left
		for _, delta := range [4][3]int{{-1, 0, 1}, {0, 1, 2}, {1, 0, 4}, {0, -1, 8}} {
			for {
				nr, nc := r+delta[0], c+delta[1]
				if nr < 0 || nr >= m || nc < 0 || nc >= n {
					// Out of the grid
					return false
				}
				if grid[nr][nc] == 16 { // Change the direction
					break
				}
				r, c = nr, nc
				bit := byte(delta[2])
				if grid[r][c]&bit == bit {
					// Found the loop!
					return true
				} else {
					grid[r][c] = grid[r][c] | bit
				}

			}
		}
	}
}

func part1(s string) int {
	grid := stringToGrid(s)
	r, c := find_guard(grid)
	paths := collectPath(grid, r, c)
	return len(paths)
}

func collectPath(src [][]byte, r, c int) [][2]int {
	grid := make([][]byte, 0, len(src))
	for _, srcRow := range src {
		row := make([]byte, len(srcRow))
		copy(row, srcRow)
		grid = append(grid, row)
	}
	m, n := len(grid), len(grid[0])
	grid[r][c] = 'X'
	var paths [][2]int = [][2]int{{r, c}}
	for {
		for _, delta := range [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			for {
				nr, nc := r+delta[0], c+delta[1]
				if nr < 0 || nr >= m || nc < 0 || nc >= n {
					// Out of the grid
					return paths
				}
				if grid[nr][nc] == '#' {
					// Change the direction
					break
				}
				r, c = nr, nc
				if grid[r][c] != 'X' {
					grid[r][c] = 'X'
					paths = append(paths, [2]int{r, c})
				}
			}
		}
	}
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
		// fmt.Println(row)
		for _, c := range row {
			fmt.Printf("%02d ", c)
		}
		fmt.Println("")
	}
	fmt.Println("")
}
