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
	fmt.Printf("Part1 -> %d\n", len(part1(grid)))
	fmt.Printf("Part2 -> %d\n", len(part2(grid)))
}

func part1(grid [][]byte) []Node {
	antennaGroup := getAntennaGroup(grid)
	posSet := make(map[Node]struct{})
	for _, an := range antennaGroup {
		for _, node := range getAntiNodes(an, len(grid), len(grid[0]), getPosition) {
			posSet[node] = struct{}{}
		}
	}
	uniqueNodes := make([]Node, 0)
	for co := range posSet {
		uniqueNodes = append(uniqueNodes, co)
	}
	return uniqueNodes
}

func part2(grid [][]byte) []Node {
	antennaGroup := getAntennaGroup(grid)
	posSet := make(map[Node]struct{})
	for _, an := range antennaGroup {
		for _, node := range getAntiNodes(an, len(grid), len(grid[0]), getPosition2) {
			posSet[node] = struct{}{}
		}
	}
	uniqueNodes := make([]Node, 0)
	for co := range posSet {
		uniqueNodes = append(uniqueNodes, co)
	}
	return uniqueNodes
}

func getAntiNodes(
	antennas []Node,
	m, n int,
	generator func(a, b Node, m, n int) []Node,
) []Node {
	hs := make(map[Node]struct{}, 0)
	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			for _, node := range generator(antennas[i], antennas[j], m, n) {
				hs[node] = struct{}{}
			}
		}
	}
	arr := make([]Node, 0)
	for pos := range hs {
		arr = append(arr, pos)
	}
	return arr
}

func getPosition(a, b Node, m, n int) []Node {
	ax, ay, bx, by := a[0], a[1], b[0], b[1]
	diffx := ax - bx
	diffy := ay - by
	ax += diffx
	ay += diffy
	bx -= diffx
	by -= diffy
	nodes := make([]Node, 0)
	if ax >= 0 && ax < m && ay >= 0 && ay < n {
		nodes = append(nodes, Node{ax, ay})
	}
	if bx >= 0 && bx < m && by >= 0 && by < n {
		nodes = append(nodes, Node{bx, by})
	}
	return nodes
}

func getPosition2(a, b Node, m, n int) []Node {
	ax, ay, bx, by := a[0], a[1], b[0], b[1]
	diffx := ax - bx
	diffy := ay - by
	nodes := make([]Node, 0)
	for {
		nodes = append(nodes, Node{ax, ay})
		ax += diffx
		ay += diffy
		if ax < 0 || ax >= m || ay < 0 || ay >= n {
			break
		}
	}
	for {
		nodes = append(nodes, Node{bx, by})
		bx -= diffx
		by -= diffy
		if bx < 0 || bx >= m || by < 0 || by >= n {
			break
		}
	}
	return nodes
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
