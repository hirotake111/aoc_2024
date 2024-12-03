package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	// fileName = "./day02/test.txt"
	fileName = "./day02/input.txt"
)

func main() {
	reports, err := getReports()
	if err != nil {
		log.Fatal(err)
	}
	var answer int
	for _, r := range reports {
		if validateReport(r) {
			answer++
		}
	}
	fmt.Printf("part 1 -> %d\n", answer)

	var answer2 int
	for _, r := range reports {
		if validateReportWithTorelant(r) {
			answer2++
		}
	}
	fmt.Printf("part 2 -> %d\n", answer2)
}

func validateReport(report []int) bool {
	increasing := isIncreasing(report)
	prev := -1
	for _, cur := range report {
		if prev == -1 {
			prev = cur
		} else if !safePair(prev, cur, increasing) {
			return false
		}
		prev = cur
	}
	return true
}

func validateReportWithTorelant(report []int) bool {
	for skip := 0; skip < len(report); skip++ {
		arr := make([]int, 0, len(report)-1)
		for i, n := range report {
			if i != skip {
				arr = append(arr, n)
			}
		}
		if validateReport(arr) {
			return true
		}
	}
	return false
}

func isIncreasing(report []int) bool {
	prev := report[0]
	inc, dec := 0, 0
	for _, cur := range report {
		if prev < cur {
			inc++
		} else if prev > cur {
			dec++
		}
		prev = cur
	}
	return inc > dec
}

func safePair(a, b int, increasing bool) bool {
	if !increasing {
		a, b = b, a
	}
	return a < b && b-a <= 3
}

func getReports() ([][]int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return [][]int{}, err
	}
	reader := bufio.NewReader(f)
	result := make([][]int, 0)
	for {
		l, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return [][]int{}, err
		}
		var arr []int
		for _, s := range strings.Split(string(l), " ") {
			n, err := strconv.Atoi(s)
			if err != nil {
				return [][]int{}, err
			}
			arr = append(arr, n)
		}
		result = append(result, arr)
	}
	return result, nil
}
