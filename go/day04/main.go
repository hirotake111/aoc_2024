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
	input, err := getInput()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(input)
	var total int
	for i, row := range input {
		for j, c := range row {
			if c == 'X' {
				// fmt.Printf("Found X: (%d, %d)\n", i, j)
				total += findXMAS(i, j, input)
			}
		}
	}
	fmt.Printf("Part1 -> %d\n", total)
}

var directions = [8][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}

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
