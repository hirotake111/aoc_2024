package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	fileName = "./day05/sample.txt"
)

func main() {
	section1, section2, err := getData()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(section1)
	fmt.Println(section2)
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
