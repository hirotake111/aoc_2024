package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`

func TestMain(t *testing.T) {
	grid := getGrid(example)
	assert.Equal(t, 36, part1(grid))
	assert.Equal(t, 81, part2(grid))
}
