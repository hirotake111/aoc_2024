package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	// fileName = "./day05/sample.txt"
	fileName = "./day05/input.txt"
)

func main() {
	section1, section2, err := getData()
	if err != nil {
		log.Fatal(err)
	}
	if err := calc(section1, section2); err != nil {
		log.Fatal(err)
	}

}

func calc(section1, section2 []string) error {
	m, err := toMap(section1)
	if err != nil {
		return err
	}
	// p(m)
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
			updated := correct_ordering(arr, m)
			corrected += updated[len(updated)/2]
		}
	}
	fmt.Printf("Part1 -> %d\n", total)
	fmt.Printf("Part2 -> %d\n", corrected)
	return nil
}

type Node struct {
	val  int
	next *Node
}

func (n Node) String() string {
	return fmt.Sprintf("%d -> %+v", n.val, n.next)
}

func correct_ordering(arr []int, m map[int]map[int]struct{}) []int {
	root := &Node{}
	for _, n := range arr {
		node := &Node{val: n}
		// fmt.Printf("new node: %v\n", node)
		if root.next == nil {
			// fmt.Printf("connected to root: %+v\n", root)
			root.next = node
			continue
		}
		prev := root
		current := root.next
		for current != nil {
			before := m[current.val]
			if _, ok := before[node.val]; ok {
				// change order
				prev.next = node
				node.next = current
				prev = node
				break
			} else {
				// move on to the next
				prev = current
				current = current.next
			}
		}
		if current == nil {
			prev.next = node
		}
	}
	// fmt.Printf("node: %+v\n", root)
	updated := make([]int, 0)
	node := root.next
	for node != nil {
		updated = append(updated, node.val)
		node = node.next
	}
	// fmt.Printf("updated: %v\n", updated)
	return updated
}

func p(m map[int]map[int]struct{}) {
	for k, v := range m {
		fmt.Printf("%d: %v\n", k, v)
	}
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
		if _, ok := m[l]; !ok {
			m[l] = make(map[int]struct{}, 0)
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
