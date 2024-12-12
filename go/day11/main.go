package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./data/day11.txt")
	if err != nil {
		panic(err)
	}
	input, err := getInput(string(data))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part1 -> %d\n", len(part1(input, 25)))
	fmt.Printf("Part2 -> %d\n", part2(input, 75))
}

func part1(input []int, n int) []int {
	for i := 0; i < n; i++ {
		// fmt.Printf("Iteration: %d, length: %d\n", i, len(input))
		arr := make([]int, 0)
		for _, v := range input {
			if v == 0 {
				// fmt.Printf("1st rule: %d\n", v)
				arr = append(arr, 1)
			} else if len(strconv.Itoa(v))%2 == 0 {
				// fmt.Printf("2nd rule: %d\n", v)
				for _, a := range divide(v) {
					arr = append(arr, a)
				}
			} else {
				// fmt.Printf("3rd rule: %d\n", v)
				arr = append(arr, v*2024)
			}
		}
		input = arr
	}
	return input
}

func part2(input []int, steps int) int {
	cache := make(map[[2]int]int, 0)
	var total int
	for _, n := range input {
		total += dfs(n, steps, cache)
	}
	return total
}

func dfs(stone, steps int, cache map[[2]int]int) int {
	if steps == 0 {
		return 1
	}
	key := [2]int{stone, steps}
	if v, ok := cache[key]; ok {
		return v
	}
	if stone == 0 {
		cache[key] = dfs(1, steps-1, cache)
		return cache[key]
	}
	s := strconv.Itoa(stone)
	l := len(s)
	if l%2 == 0 {
		pair := divide(stone)
		cache[key] = dfs(pair[0], steps-1, cache) + dfs(pair[1], steps-1, cache)
		return cache[key]
	}
	cache[key] = dfs(stone*2024, steps-1, cache)
	return cache[key]
}

func divide(n int) [2]int {
	a := len(strconv.Itoa(n)) / 2
	b := 1
	for i := 0; i < a; i++ {
		b *= 10
	}
	return [2]int{n / b, n % b}
}

func getInput(s string) ([]int, error) {
	result := make([]int, 0)
	for _, ss := range strings.Split(strings.Trim(s, "\n"), " ") {
		n, err := strconv.Atoi(ss)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}
	return result, nil
}
