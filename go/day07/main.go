package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const example = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

type Equation struct {
	testVal int
	nums    []int
}

type Operator func(a, b int) int

func main() {
	input, err := os.ReadFile("./data/day07.txt")
	if err != nil {
		panic(err)
	}
	res, err := part1(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part1 -> %d\n", res)
	res, err = part2(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part2 -> %d\n", res)
}

func part2(s string) (int, error) {
	es, err := parseEquations(s)
	if err != nil {
		return 0, err
	}
	var total int
	for _, e := range es {
		ops := []Operator{
			func(a, b int) int {
				return a + b
			},
			func(a, b int) int {
				return a * b
			},
			func(a, b int) int {
				n := 10
				for n <= b {
					n *= 10
				}
				return a*n + b
			},
		}
		total += getCalibrationResult(e, ops)
	}
	return total, nil
}

func part1(s string) (int, error) {
	es, err := parseEquations(s)
	if err != nil {
		return 0, err
	}
	var total int
	for _, e := range es {
		ops := []Operator{
			func(a, b int) int {
				return a + b
			},
			func(a, b int) int {
				return a * b
			},
		}
		total += getCalibrationResult(e, ops)
	}
	return total, nil
}

func getCalibrationResult(e Equation, ops []Operator) int {
	cand := make([]int, 0)
	for _, n := range e.nums {
		if len(cand) == 0 {
			cand = append(cand, n)
		} else {
			newCand := make([]int, 0)
			for _, a := range cand {
				for _, op := range ops {
					v := op(a, n)
					newCand = append(newCand, v)
				}
			}
			cand = newCand
		}
	}
	for _, n := range cand {
		if n == e.testVal {
			return e.testVal
		}
	}
	return 0
}

func parseEquations(input string) ([]Equation, error) {
	input = strings.Trim(input, "\n")
	es := make([]Equation, 0)
	for _, line := range strings.Split(input, "\n") {
		subLine := strings.Split(line, ": ")
		v, err := strconv.Atoi(subLine[0])
		if err != nil {
			return nil, err
		}
		e := Equation{testVal: v, nums: make([]int, 0)}
		for _, s := range strings.Split(subLine[1], " ") {
			v, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			e.nums = append(e.nums, v)
		}
		es = append(es, e)
	}
	return es, nil
}
