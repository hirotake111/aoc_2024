package main

import (
	"fmt"
	"os"
	"strings"
)

type Node = [2]int

func main() {
	content, err := os.ReadFile("./data/day08.txt")
	if err != nil {
		panic(err)
	}
	grid, err := getGrid(string(content))
	if err != nil {
		panic(err)
	}
	nodes := part1(grid)
	for _, n := range nodes {
		grid[n[0]][n[1]] = '#'
	}
	// p(grid)
	fmt.Printf("Part1 -> %d\n", len(nodes))
}

func part1(grid [][]byte) []Node {
	antennaGroup := getAntennaGroup(grid)
	posSet := make(map[Node]struct{})
	for _, antenna := range antennaGroup {
		for _, node := range getAntiNodes(antenna, len(grid), len(grid[0])) {
			posSet[node] = struct{}{}
		}
	}
	uniqueNodes := make([]Node, 0)
	for co := range posSet {
		uniqueNodes = append(uniqueNodes, co)
	}
	return uniqueNodes
}

func getAntiNodes(antennas []Node, m, n int) []Node {
	hs := make(map[Node]struct{}, 0)
	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			a, b := getPosition(antennas[i], antennas[j])
			if a[0] >= 0 && a[0] < m && a[1] >= 0 && a[1] < n {
				hs[a] = struct{}{}

			}
			if b[0] >= 0 && b[0] < m && b[1] >= 0 && b[1] < n {
				hs[b] = struct{}{}
			}
		}
	}
	arr := make([]Node, 0)
	for pos := range hs {
		arr = append(arr, pos)
	}
	return arr
}

func getPosition(a, b Node) (Node, Node) {
	ax, ay, bx, by := a[0], a[1], b[0], b[1]
	if ax > bx {
		diff := ax - bx
		ax += diff
		bx -= diff
	} else {
		diff := bx - ax
		ax -= diff
		bx += diff
	}
	if ay > by {
		diff := ay - by
		ay += diff
		by -= diff
	} else {
		diff := by - ay
		ay -= diff
		by += diff
	}
	return Node{ax, ay}, Node{bx, by}
}

func getAntennaGroup(grid [][]byte) map[byte][]Node {
	cm := make(map[byte][]Node, 0)
	for i, row := range grid {
		for j, c := range row {
			if c != '.' {
				if len(cm[c]) == 0 {
					cm[c] = make([]Node, 0)
				}
				cm[c] = append(cm[c], Node{i, j})
			}
		}
	}
	return cm
}

func getGrid(s string) ([][]byte, error) {
	grid := make([][]byte, 0)
	for _, line := range strings.Split(strings.Trim(s, "\n"), "\n") {
		grid = append(grid, []byte(line))
	}
	return grid, nil
}

func p(grid [][]byte) {
	for _, row := range grid {
		for _, c := range row {
			fmt.Printf("%c", c)
		}
		fmt.Println("")
	}
	fmt.Println("")
}
