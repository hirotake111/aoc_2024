package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay07(t *testing.T) {
	res, err := part1(example)
	assert.Nil(t, err)
	assert.Equal(t, 3749, res)
	fmt.Println("====")
	res, err = part2(example)
	assert.Nil(t, err)
	assert.Equal(t, 11387, res)
}
