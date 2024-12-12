package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const ex1 = `AAAA
BBCD
BBCC
EEEC`

const ex2 = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func TestMain(t *testing.T) {
	var garden [][]byte
	garden = getGarden(ex1)
	p(garden)
	assert.Equal(t, 140, part1(garden))
	garden = getGarden(ex2)
	p(garden)
	assert.Equal(t, 1930, part1(garden))
}
