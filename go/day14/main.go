package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Input struct {
	px int
	py int
	vx int
	vy int
}

func main() {
	raw, err := os.ReadFile("./data/day14.txt")
	if err != nil {
		panic(err)
	}
	inputs := getInputs(string(raw))
	fmt.Printf("Part1 -> %d\n", part1(inputs, 101, 103, 100))
	part2(inputs, 101, 103)
}

func part2(inputs []Input, width, height int) int {
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}
	for _, in := range inputs {
		grid[in.py][in.px]++
	}
	p(grid)
	for i := 1; i <= 101*103; i++ {
		for j := 0; j < len(inputs); j++ {
			grid[inputs[j].py][inputs[j].px]--
			inputs[j].px = (inputs[j].px + inputs[j].vx + width) % width
			inputs[j].py = (inputs[j].py + inputs[j].vy + height) % height
			grid[inputs[j].py][inputs[j].px]++
		}
		// check if there is no overlaps
		found := true
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if grid[y][x] > 1 {
					found = false
					break
				}
			}
			if !found {
				break
			}
		}
		if found {
			p2(grid)
			return i
		}
	}
	panic("unreachable")
}

func part1(inputs []Input, width, height, seconds int) int {
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}
	// p(grid)
	for _, in := range inputs {
		grid[in.py][in.px]++
	}
	// p(grid)
	for _, in := range inputs {
		// fmt.Printf("input: %+v\n", in)
		grid[in.py][in.px]--
		x := (width*seconds + in.px + in.vx*seconds) % width
		y := (height*seconds + in.py + in.vy*seconds) % height
		// fmt.Printf("#%d - x: %d, y: %d\n", i, x, y)
		grid[y][x]++
	}
	// p(grid)
	hl, vl := height/2, width/2
	// fmt.Printf("hl: %d, vl: %d\n", hl, vl)
	a, b, c, d := 0, 0, 0, 0
	for i, row := range grid {
		for j, n := range row {
			if n == 0 || i == hl || j == vl {
				continue
			}
			if i < hl {
				if j < vl {
					a += n
				} else {
					b += n
				}
			} else {
				if j < vl {
					c += n
				} else {
					d += n
				}
			}
		}
	}
	return a * b * c * d
}

func getInputs(s string) []Input {
	var err error
	s = strings.Trim(s, "\n")
	r := regexp.MustCompile("[\\-]?\\d+")
	inputs := make([]Input, 0)
	for _, line := range strings.Split(s, "\n") {
		m := r.FindAllString(line, -1)
		arr := [4]int{}
		for i := 0; i < 4; i++ {
			arr[i], err = strconv.Atoi(m[i])
			if err != nil {
				panic(err)
			}
		}
		inputs = append(inputs, Input{px: arr[0], py: arr[1], vx: arr[2], vy: arr[3]})
	}
	return inputs
}

func p(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println("")
}

func p2(grid [][]int) {
	for _, row := range grid {
		for _, v := range row {
			if v == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
