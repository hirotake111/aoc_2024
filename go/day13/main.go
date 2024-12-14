package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	raw, err := os.ReadFile("./data/day13.txt")
	if err != nil {
		panic(err)
	}
	inputs, err := getInputs(string(raw))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part1 -> %d\n", part1(inputs))
}

type Input struct {
	A struct {
		X int
		Y int
	}
	B struct {
		X int
		Y int
	}
	Prize struct {
		X int
		Y int
	}
}

func getInputs(s string) ([]Input, error) {
	inputs := make([]Input, 0)
	s = strings.Trim(s, "\n")
	r := regexp.MustCompile("\\d+")
	for _, block := range strings.Split(s, "\n\n") {
		input := Input{}
		var err error
		matched := r.FindAllString(block, -1)
		input.A.X, err = strconv.Atoi(matched[0])
		if err != nil {
			return nil, err
		}
		input.A.Y, err = strconv.Atoi(matched[1])
		if err != nil {
			return nil, err
		}
		input.B.X, err = strconv.Atoi(matched[2])
		if err != nil {
			return nil, err
		}
		input.B.Y, err = strconv.Atoi(matched[3])
		if err != nil {
			return nil, err
		}
		input.Prize.X, err = strconv.Atoi(matched[4])
		if err != nil {
			return nil, err
		}
		input.Prize.Y, err = strconv.Atoi(matched[5])
		if err != nil {
			return nil, err
		}
		inputs = append(inputs, input)
	}
	return inputs, nil
}

func part1(inputs []Input) int {
	var tokens int
	for _, input := range inputs {
		// fmt.Printf("input: %+v\n", input)
		v := getMinTokens(input)
		// fmt.Printf("Game %d: Spend %d tokens\n", i, v)
		tokens += v
	}
	return tokens
}

func getMinTokens(input Input) int {
	minVal := math.MaxInt
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			if input.A.X*i+input.B.X*j == input.Prize.X && input.A.Y*i+input.B.Y*j == input.Prize.Y {
				// fmt.Printf("i: %d, j: %d\n", i, j)
				minVal = min(minVal, i*3+j)
			}
		}
	}
	if minVal == math.MaxInt {
		return 0
	}
	return minVal
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
