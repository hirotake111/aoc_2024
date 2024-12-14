package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input, err := os.ReadFile("./data/day12.txt")
	if err != nil {
		panic(err)
	}
	garden := getGarden(string(input))
	fmt.Printf("Part1 -> %d\n", part1(garden))
	fmt.Printf("Part2 -> %d\n", part2(garden))
}

func getGarden(input string) [][]byte {
	garden := make([][]byte, 0)
	for _, s := range strings.Split(strings.Trim(input, "\n"), "\n") {
		garden = append(garden, []byte(s))
	}
	return garden
}

func part2(garden [][]byte) int {
	m, n := len(garden), len(garden[0])
	// Hash set to track visited cell
	seen := make([][]bool, len(garden))
	for i := range garden {
		seen[i] = make([]bool, len(garden[0]))
	}

	var prices int
	for i, row := range garden {
		for j := range row {
			if seen[i][j] {
				continue
			}
			seen[i][j] = true
			fmt.Printf("%c: i: %d, j: %d\n", garden[i][j], i, j)
			plants := 1
			q := Queue{{i, j}}
			sides := make(map[[2]int][][2]int, 0)

			for !q.Empty() {
				v := q.Pop()
				// 1: Right, 2: Down, 3: left, 4: Up
				for _, d := range [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
					r, c := v[0]+d[0], v[1]+d[1]
					if r < 0 || r >= m || c < 0 || c >= n || garden[v[0]][v[1]] != garden[r][c] {
						// fmt.Printf("r: %d c: %d, direction: %v, \n", r, c, d)
						if d[1] == 0 { // Up or down
							sides[d] = append(sides[d], [2]int{r, c})
							// fmt.Printf("sideMap: %+v\n", sides)
						} else {
							sides[d] = append(sides[d], [2]int{r, c})
							// fmt.Printf("sideMap: %+v\n", sides)
						}
					} else {
						if seen[r][c] {
							continue
						}
						seen[r][c] = true
						// Same kind of plant
						q.Push([2]int{r, c})
						plants++
						// fmt.Printf("r: %d, c: %d, Same plant: plants: %d\n", r, c, plants)
					}
				}
			}
			s := calcSides(sides)
			// fmt.Printf("Final value of the sides: %v\n", sides)
			// fmt.Printf("%c: plants: %d, sides: %d -> %d\n", garden[i][j], plants, s, plants*s)
			prices += plants * s
		}
	}
	return prices
}

func calcSides(sides map[[2]int][][2]int) int {
	total := 4
	for d, arr := range sides {
		switch d {
		case [2]int{-1, 0}, [2]int{1, 0}: // up or down
			sort.Slice(arr, func(i, j int) bool {
				if arr[i][0] == arr[j][0] {
					return arr[i][1] < arr[j][1]
				}
				return arr[i][0] < arr[j][0]
			})
			for i := 1; i < len(arr); i++ {
				if arr[i-1][0] != arr[i][0] || arr[i-1][1]+1 != arr[i][1] {
					total++
				}
			}
		case [2]int{0, -1}, [2]int{0, 1}: // left or right
			sort.Slice(arr, func(i, j int) bool {
				if arr[i][1] == arr[j][1] {
					return arr[i][0] < arr[j][0]
				}
				return arr[i][1] < arr[j][1]
			})
			for i := 1; i < len(arr); i++ {
				if arr[i-1][1] != arr[i][1] || arr[i-1][0]+1 != arr[i][0] {
					total++
				}
			}
		default:
			panic(d)
		}
	}
	return total
}

type Queue [][2]int

func (q *Queue) Push(v [2]int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() [2]int {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q Queue) Empty() bool {
	return len(q) == 0
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
