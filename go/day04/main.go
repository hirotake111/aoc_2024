package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	// fileName = "./day04/test.txt"
	fileName = "./day04/input.txt"
)

func main() {
	grid, err := getInput()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(grid)
	var total int
	for i, row := range grid {
		for j, c := range row {
			if c == 'X' {
				// fmt.Printf("Found X: (%d, %d)\n", i, j)
				total += findXMAS(i, j, grid)
			}
		}
	}
	fmt.Printf("Part1 -> %d\n", total)

	var total2 int
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if grid[i][j] == 'A' {
				total2 += findX_MAS(i, j, grid)
			}
		}
	}
	fmt.Printf("Part2 -> %d\n", total2)
}

var directions = [8][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}

func findX_MAS(i, j int, grid [][]byte) int {
	var total int
	flag := false
	switch [2]byte{grid[i-1][j-1], grid[i+1][j+1]} {
	case [2]byte{'M', 'S'}, [2]byte{'S', 'M'}:
		// p(i, j, grid)
		flag = true
	}
	switch [2]byte{grid[i-1][j+1], grid[i+1][j-1]} {
	case [2]byte{'M', 'S'}, [2]byte{'S', 'M'}:
		// p(i, j, grid)
		flag = flag && true
	default:
		flag = false
	}
	if flag {
		// p(i, j, grid)
		total++
	}
	return total
}

func p(i, j int, grid [][]byte) {
	fmt.Printf("(%d, %d)\n", i, j)
	fmt.Printf("%c%c%c\n", grid[i-1][j-1], grid[i-1][j], grid[i-1][j+1])
	fmt.Printf("%c%c%c\n", grid[i][j-1], grid[i][j], grid[i][j+1])
	fmt.Printf("%c%c%c\n", grid[i+1][j-1], grid[i+1][j], grid[i+1][j+1])
	fmt.Println("")
}

func findXMAS(i, j int, grid [][]byte) int {
	var total int
	for _, d := range directions {
		r, c := i+d[0], j+d[1]
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			continue
		}
		if grid[r][c] != 'M' {
			continue
		}
		r += d[0]
		c += d[1]
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			continue
		}
		if grid[r][c] != 'A' {
			continue
		}
		r += d[0]
		c += d[1]
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			continue
		}
		if grid[r][c] != 'S' {
			continue
		}
		total++
	}
	return total
}

func getInput() ([][]byte, error) {
	raw, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	input := make([][]byte, 0)
	for _, row := range strings.Split(strings.TrimSuffix(string(raw), "\n"), "\n") {
		input = append(input, []byte(row))
	}
	return input, nil
}
