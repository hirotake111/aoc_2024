package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	fileName = "./day05/sample.txt"
	// fileName = "./day05/input.txt"
)

func main() {
	section1, section2, err := getData()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(section1)
	// fmt.Println(section2)
	m, err := toMap(section1)
	// fmt.Printf("m: %+v\n", m)
	var total, corrected int
	for _, s := range section2 {
		arr, err := toIntArr(s)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("arr: %v\n", arr)
		seen := make(map[int]struct{})
		flag := true
		for _, n := range arr {
			if _, ok := seen[n]; ok {
				flag = false
				break
			}
			if hs, ok := m[n]; ok {
				for k := range hs {
					seen[k] = struct{}{}
				}
			}
		}
		if flag {
			total += arr[len(arr)/2]
		} else {

		}
	}
	fmt.Printf("Part1 -> %d\n", total)
	fmt.Printf("Part2 -> %d\n", total+corrected)

}

func toIntArr(s string) ([]int, error) {
	arr := make([]int, 0)
	for _, ss := range strings.Split(s, ",") {
		n, err := strconv.Atoi(ss)
		if err != nil {
			return nil, err
		}
		arr = append(arr, n)
	}
	return arr, nil
}

func toMap(section []string) (map[int]map[int]struct{}, error) {
	m := make(map[int]map[int]struct{})
	for _, s := range section {
		v := strings.Split(s, "|")
		l, err := strconv.Atoi(v[0])
		if err != nil {
			return nil, err
		}
		r, err := strconv.Atoi(v[1])
		if err != nil {
			return nil, err
		}
		if hs, ok := m[r]; !ok {
			hs := make(map[int]struct{}, 0)
			hs[l] = struct{}{}
			m[r] = hs
		} else {
			hs[l] = struct{}{}
		}
	}
	return m, nil
}

func getData() ([]string, []string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, nil, err
	}
	s := strings.TrimSpace(string(data))
	section1 := make([]string, 0)
	section2 := make([]string, 0)
	section := &section1
	for _, ss := range strings.Split(s, "\n") {
		if ss == "" {
			section = &section2
		} else {
			*section = append(*section, ss)
		}
	}
	return section1, section2, nil
}
