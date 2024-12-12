package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = "125 17"

func TestMain(t *testing.T) {
	input, err := getInput(example)
	fmt.Println(input)
	assert.Nil(t, err)
	assert.Equal(t, 22, len(part1(input, 6)))
	assert.Equal(t, 55312, len(part1(input, 25)))
	assert.Equal(t, 22, part2(input, 6))
	assert.Equal(t, 55312, part2(input, 25))
}
