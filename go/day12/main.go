package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("./data/day12.txt")
	if err != nil {
		panic(err)
	}
	garden := getGarden(string(input))
	fmt.Printf("Part1 -> %d\n", part1(garden))
}

func getGarden(input string) [][]byte {
	garden := make([][]byte, 0)
	for _, s := range strings.Split(strings.Trim(input, "\n"), "\n") {
		garden = append(garden, []byte(s))
	}
	return garden
}

func part1(garden [][]byte) int {
	seen := make([][]bool, len(garden))
	for i := range garden {
		seen[i] = make([]bool, len(garden[0]))
	}

	var prices int
	for i, row := range garden {
		for j := range row {
			if !seen[i][j] {
				seen[i][j] = true
				res := traverse(i, j, garden, seen)
				// fmt.Printf("%c: %d * %d -> %d\n", v, res[0], res[1], res[0]*res[1])
				prices += res[0] * res[1]
			}
		}
	}
	return prices
}

func traverse(i, j int, garden [][]byte, seen [][]bool) [2]int {
	m, n := len(garden), len(garden[0])
	plants, fences := 1, 0
	for _, d := range [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		r, c := i+d[0], j+d[1]
		if r < 0 || r >= m || c < 0 || c >= n || garden[i][j] != garden[r][c] {
			fences++
		} else if seen[r][c] { // senn it before
			continue
		} else { // Same plants
			seen[r][c] = true
			res := traverse(r, c, garden, seen)
			plants += res[0]
			fences += res[1]
		}
	}
	return [2]int{plants, fences}
}

func p(grid [][]byte) {
	for _, row := range grid {
		fmt.Printf("%c\n", row)
	}
	fmt.Println("")
}
