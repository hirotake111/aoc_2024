package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	// fileName = "./day03/test.txt"
	fileName = "./day03/input.txt"
	// fileName2 = "./day03/test2.txt"
	fileName2 = "./day03/input.txt"
)

var NotTargetError = errors.New("not target")
var EOF = errors.New("end of file!")

func main() {
	buf, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	cursor := NewCursor(buf)
	total1, err := cursor.getTotalPt1()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 1 -> %d\n", total1)

	buf, err = os.ReadFile(fileName2)
	if err != nil {
		log.Fatal(err)
	}
	cursor = NewCursor(buf)
	total2, err := cursor.getTotalPt2()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 2 -> %d\n", total2)
}

type Cursor struct {
	buf []byte // Buffered data
	idx int    // index pointing to the current byte
}

func NewCursor(buf []byte) *Cursor {
	return &Cursor{
		buf: buf,
		idx: 0,
	}
}

func (c *Cursor) Peek() (byte, error) {
	if c.idx >= len(c.buf) {
		return 0, EOF
	}
	b := c.buf[c.idx]
	return b, nil
}

func (c *Cursor) Seek() {
	c.idx++
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
	if err := c.parseWord("mul("); err != nil {
		if !errors.Is(err, NotTargetError) {
			return 0, nil
		}
	}
	// Get first number
	a, err := c.GetNum()
	if err != nil {
		return 0, err
	}
	// fmt.Printf("Got a: %d\n", a)
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
	// fmt.Printf("Got b: %d\n", b)
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

func (c *Cursor) getTotalPt1() (int, error) {
	total := 0
	var err error

	for {
		if err = c.FindM(); err != nil {
			break
		}
		n, err := c.Parse()
		if err != nil {
			break
		}
		// fmt.Printf("got result: %d\n", n)
		total += n
	}

	if errors.Is(err, EOF) {
		log.Println("found end of file while searching for 'm'")
		return total, nil
	}
	return 0, err
}

func (c *Cursor) getTotalPt2() (int, error) {
	total := 0
	var err error

	for {
		var b byte
		b, err = c.Peek()
		if err != nil {
			break
		}
		switch b {
		case 'm':
			var n int
			n, err = c.Parse()
			if err != nil {
				break
			}
			// fmt.Printf("got result: %d\n", n)
			total += n
		case 'd':
			if err = c.parseWord("don't()"); err != nil {
				continue
			}
			fmt.Println("Fond don't()")
			// Move cursor until it finds do()
			for {
				b, err = c.Peek()
				if err != nil {
					break
				}
				if b != 'd' {
					c.Seek()
					continue
				}
				err = c.parseWord("do()")
				if errors.Is(err, NotTargetError) {
					// Continue seeking until we find "don't()"
					continue
				}
				break
			}
		default:
			c.Seek()
		}
	}

	if errors.Is(err, EOF) {
		log.Println("found end of file while parsing")
		return total, nil
	}
	return 0, err
}

func (c *Cursor) parseWord(w string) error {
	for _, ch := range []byte(w) {
		b, err := c.Peek()
		if err != nil {
			return err
		}
		if b != ch {
			return NotTargetError
		}
		// Still valid
		c.Seek()
	}
	return nil

}
