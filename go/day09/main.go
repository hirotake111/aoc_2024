package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DiskMap = []int

func main() {
	input, err := os.ReadFile("./data/day09.txt")
	if err != nil {
		panic(err)
	}
	dm := parseDiskMap(string(input))
	fmt.Printf("Part1 -> %d\n", part1(dm))
	fmt.Printf("Part2 -> %d\n", part2(dm))
	day9_2(string(input))
}

func parseDiskMap(input string) DiskMap {
	input = strings.Trim(input, "\n")
	dm := make(DiskMap, 0)
	for i := 0; i < len(input); i++ {
		id := i / 2
		n := int(input[i] - '0')
		for j := 0; j < n; j++ {
			if i%2 == 0 { // handle the number as block
				dm = append(dm, id)
			} else { // handle  the number as freespace
				dm = append(dm, -1)
			}
		}
	}
	return dm
}

func part1(diskMap DiskMap) int {
	dm := moveDisk(diskMap)
	return calcCheckSum(dm)
}

func part2(diskMap DiskMap) int {
	// fmt.Println(diskMap)
	blocks, spaces := getBlocksAndSpaces(diskMap)
	dm := moveDisk2(diskMap, blocks, spaces)
	fmt.Println(dm)
	return calcCheckSum(dm)
}

func moveDisk(diskMap DiskMap) DiskMap {
	dm := make(DiskMap, len(diskMap))
	copy(dm, diskMap)
	l, r := 0, len(dm)-1
	for l < r {
		for dm[l] != -1 {
			l++
		}
		for dm[r] == -1 {
			r--
		}
		if l >= r {
			break
		}
		dm[l], dm[r] = dm[r], dm[l]
		l++
		r--
	}
	return dm
}

func moveDisk2(diskMap DiskMap, blocks, spaces [][2]int) DiskMap {
	dm := make(DiskMap, len(diskMap))
	copy(dm, diskMap)
	rest := make([][2]int, 0)
	for blockId := len(blocks) - 1; blockId >= 0; blockId-- {
		rest = nil
		// fmt.Printf("block %v\n", blocks[blockId])
		// fmt.Printf("spaces: %v\n", spaces)
		blockLeft, fileSize := blocks[blockId][0], blocks[blockId][1]
		for i, space := range spaces {
			spaceLeft, freeSpace := space[0], space[1]
			if freeSpace < fileSize {
				// Look into the next free space
				rest = append(rest, space)
				continue
			}
			// Move file
			// fmt.Printf("Moving block ID %d\n", blockId)
			for k := 0; k < fileSize; k++ {
				dm[spaceLeft+k] = blockId
				dm[blockLeft+k] = -2
				space[0]++
				space[1]--
			}
			// fmt.Printf("dm: %v\n", dm)
			// fmt.Printf("space: %v\n", space)
			if space[1] > 0 {
				rest = append(rest, space)
			} else {
				// fmt.Printf("Removing space %v\n", space)
			}
			spaces = append(rest, spaces[i+1:]...)
			break
		}
	}
	fmt.Printf("final spaces: %v\n", spaces)
	return dm
	// moved := make(map[int]struct{}, 0)
	// for ri >= 0 {
	// 	// fmt.Println("check:", dm, ri)
	// 	// Find file blocks
	// 	for ri >= 0 {
	// 		if _, ok := moved[dm[ri]]; ok || dm[ri] < 0 {
	// 			ri--
	// 		} else {
	// 			break
	// 		}
	// 	}
	// 	if ri <= li {
	// 		// fmt.Printf("ri(%d) <= li(%d) -> break\n", ri, li)
	// 		break
	// 	}
	// 	rj := ri
	// 	for rj >= 0 && dm[rj] == dm[ri] {
	// 		rj--
	// 	}
	// 	fileSize := ri - rj
	// 	fmt.Printf("Block detail - ID %d,  ri: %d, rj: %d, fileSize: %d\n", dm[ri], ri, rj, fileSize)
	// 	// Find space that can accomodate fileSize
	// 	for li < ri {
	// 		// fmt.Printf("li(%d), ri(%d)\n", li, ri)
	// 		for li < ri && dm[li] != -1 {
	// 			li++
	// 		}
	// 		if li >= ri {
	// 			fmt.Printf("Skip moving ID %d\n", dm[ri])
	// 			ri = rj
	// 			li = 0
	// 			break
	// 		}
	// 		lj := li
	// 		for lj < ri && dm[lj] == dm[li] {
	// 			lj++
	// 		}
	// 		spaceSize := lj - li
	// 		if fileSize > spaceSize {
	// 			li = lj
	// 			continue
	// 		}
	// 		// Move disk blocks
	// 		block := dm[ri]
	// 		for i := 0; i < fileSize; i++ {
	// 			dm[li] = block
	// 			dm[ri] = -2
	// 			li++
	// 			ri--
	// 		}
	// 		// fmt.Println(dm)
	// 		// fmt.Println("==== end of iteration === ")
	// 		moved[block] = struct{}{}
	// 		break
	// 	}
	// }
	// // fmt.Println(dm)
	// return dm
}

func getBlocksAndSpaces(dm DiskMap) (blocks, spaces [][2]int) {
	i := 0
	for i < len(dm) {
		j := i
		for i < len(dm) && dm[i] == dm[j] {
			i++
		}
		if dm[j] == -1 {
			spaces = append(spaces, [2]int{j, i - j})
		} else { // block
			blocks = append(blocks, [2]int{j, i - j})
		}
	}
	return
}

func calcCheckSum(dm DiskMap) int {
	checkSum := 0
	for i, n := range dm {
		if n > 0 {
			checkSum += i * n
		}
	}
	return checkSum
}

func getCheckSumFromExample(input string) int {
	input = strings.Trim(input, "\n")
	var cs int
	for i := 0; i < len(input); i++ {
		if input[i] != '.' {
			cs += i * int(input[i]-'0')
		}
	}
	return cs
}

func day9_2(input string) {
	blocks := []rune{}
	isBlock := true
	fileId := 0

	for _, numRune := range input {
		num, _ := strconv.Atoi(string(numRune))

		if isBlock {
			for i := 0; i < num; i++ {
				blocks = append(blocks, rune('0'+fileId))
			}
			fileId++
		} else {
			for i := 0; i < num; i++ {
				blocks = append(blocks, '.')
			}
		}
		isBlock = !isBlock
	}

	for currentFile := fileId - 1; currentFile >= 0; currentFile-- {
		fileBlocks := []int{}

		for i, block := range blocks {
			if block == rune('0'+currentFile) {
				fileBlocks = append(fileBlocks, i)
			}
		}

		freeStart := -1
		freeLength := 0

		for i := 0; i < fileBlocks[0]; i++ {
			if blocks[i] != '.' {
				freeStart = -1
				freeLength = 0
				continue
			}

			if freeStart == -1 {
				freeStart = i
			}

			freeLength++

			if freeLength == len(fileBlocks) {
				break
			}
		}

		if freeLength == len(fileBlocks) {
			for i := 0; i < freeLength; i++ {
				blocks[fileBlocks[i]] = '.'
				blocks[freeStart+i] = rune('0' + currentFile)
			}
		}
	}

	output := 0

	for blockId := 0; blockId < len(blocks); blockId++ {
		// fmt.Printf("%c", blocks[blockId])
		if blocks[blockId] != '.' {
			output += blockId * int(blocks[blockId]-'0')
		}
	}

	// fmt.Println("")
	fmt.Println("Output Day 9 Part 2", output)
}
