package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	// fileName = "./day01/sample.txt"
	fileName = "./day01/sample2.txt"
)

func Run() {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(f)
	ls, rs, err := getData(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("sum: %d\n", sortAndSum(ls, rs))
	fmt.Printf("sumilarity score: %d\n", similarityScore(ls, rs))
}

func similarityScore(ls, rs []int) int64 {
	freq := make(map[int]int)
	for _, r := range rs {
		freq[r]++
	}
	var total int64
	for _, l := range ls {
		total += int64(l) * int64(freq[l])
	}
	return total
}

func sortAndSum(ls, rs []int) int64 {
	sort.Ints(ls)
	sort.Ints(rs)
	var total int64
	for i, l := range ls {
		r := rs[i]
		if l > r {
			total += int64(l - r)
		} else {
			total += int64(r - l)
		}
	}
	return total
}

func getData(reader *bufio.Reader) (ls []int, rs []int, err error) {
	for {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		ss := strings.Split(string(bytes), " ")
		l, err := strconv.Atoi(ss[0])
		if err != nil {
			return nil, nil, err
		}
		ls = append(ls, l)
		r, err := strconv.Atoi(ss[len(ss)-1])
		if err != nil {
			return nil, nil, err
		}
		rs = append(rs, r)
	}
	return
}
