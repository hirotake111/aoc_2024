package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const (
	// fileName = "./day03/test.txt"
	fileName = "./day03/input.txt"
)

func main() {
	fp, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	cursor := NewCursor(bufio.NewReader(fp))
	total, err := cursor.getTotal()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 1 -> %d\n", total)
}

type Cursor struct {
	reader *bufio.Reader
}

func NewCursor(reader *bufio.Reader) *Cursor {
	return &Cursor{reader: reader}
}

func (c *Cursor) Peek() (byte, error) {
	b, err := c.reader.Peek(1)
	if err != nil {
		return 0, err
	}
	// fmt.Printf("peek got %c\n", b)
	return b[0], err
}

func (c *Cursor) Seek() (byte, error) {
	return c.reader.ReadByte()
}

func (c *Cursor) FindM() error {
	for {
		b, err := c.Peek()
		if err != nil {
			return err
		}
		if b == 'm' {
			return nil
		}
		c.Seek()
	}
}

func (c *Cursor) Parse() (int, error) {
	// mul(
	for _, b := range []byte{'m', 'u', 'l', '('} {
		t, err := c.Peek()
		if err != nil {
			return 0, err
		}
		if t != b {
			return 0, err
		}
		c.Seek()
	}
	// Get first number
	a, err := c.GetNum()
	if err != nil {
		return 0, err
	}
	fmt.Printf("Got a: %d\n", a)
	t, err := c.Peek()
	if err != nil {
		return 0, nil
	}
	if t != ',' {
		return 0, nil
	}
	c.Seek()
	// Get second number
	b, err := c.GetNum()
	if err != nil {
		return 0, nil
	}
	fmt.Printf("Got b: %d\n", b)
	t, err = c.Peek()
	if err != nil {
		return 0, err
	}
	if t != ')' {
		return 0, err
	}
	c.Seek()
	return a * b, nil
}

var nums = []byte("1234567890")

func (c *Cursor) GetNum() (int, error) {
	var bs []byte
	for {
		b, err := c.Peek()
		if err != nil {
			return 0, nil
		}
		if bytes.Contains(nums, []byte{b}) {
			bs = append(bs, b)
			c.Seek()
		} else {
			break
		}
	}
	a, err := strconv.Atoi(string(bs))
	if err != nil {
		return 0, err
	}
	return a, nil
}

func (c *Cursor) getTotal() (int, error) {
	total := 0
	for {
		if err := c.FindM(); err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("found end of file while searching for 'm'")
				break
			} else {
				log.Fatal(err)
			}
		}
		n, err := c.Parse()
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("found end of file while parsing")
				break
			} else {
				log.Fatal(err)
			}
		}
		fmt.Printf("got result: %d\n", n)
		total += n
	}
	return total, nil
}
