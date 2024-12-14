package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const ex1 = `
AAAA
BBCD
BBCC
EEEC`

const ex2 = `
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

const ex3 = `
EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`

const ex4 = `
AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`

func TestExample1(t *testing.T) {
	g1 := getGarden(ex1)
	p(g1)
	assert.Equal(t, 140, part1(g1))
	assert.Equal(t, 80, part2(g1))
}

func TestExample2(t *testing.T) {
	g2 := getGarden(ex2)
	p(g2)
	assert.Equal(t, 1930, part1(g2))
	assert.Equal(t, 1206, part2(g2))
}
func TestExample3(t *testing.T) {
	g3 := getGarden(ex3)
	p(g3)
	assert.Equal(t, 236, part2(g3))
}
func TestExample4(t *testing.T) {
	g4 := getGarden(ex4)
	p(g4)
	assert.Equal(t, 368, part2(g4))
}
