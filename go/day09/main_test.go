package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = "2333133121414131402"

func TestMain(t *testing.T) {
	diskMap := parseDiskMap(example)
	assert.Equal(t, 1928, part1(diskMap))
	// assert.Equal(t, 2858, getCheckSumFromExample("00992111777.44.333....5555.6666.....8888.."))
	assert.Equal(t, 2858, part2(diskMap))
}
