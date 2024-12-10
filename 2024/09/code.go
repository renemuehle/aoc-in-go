package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block

	blockmap := make([]string, 0)
	filemap := make(map[int][]int)

	file := true

	fileId := 0

	for _, v := range strings.Split(input, "") {
		fileSize, _ := strconv.Atoi(v)
		if fileSize > 0 {
			s1 := "."
			if file {
				s1 = strconv.Itoa(fileId)
			}
			s2 := make([]string, fileSize)
			for i := range s2 {
				s2[i] = s1
			}
			blockmap = append(blockmap, s2...)

			if file {
				fmt.Println("File ", fileId, " Size ", fileSize, " Pos ", len(blockmap)-fileSize)
				filemap[fileId] = append(filemap[fileId], fileSize)
				filemap[fileId] = append(filemap[fileId], len(blockmap)-fileSize)
				fileId++
			}
		}
		file = !file
	}
	// the current filesystem
	fmt.Println(blockmap)
	// file-informations for easier access
	fmt.Println(filemap)

	fileIds := make([]int, 0, len(filemap))
	for fId := range filemap {
		fileIds = append(fileIds, fId)
	}

	slices.Sort(fileIds)
	slices.Reverse(fileIds)

	if part2 {
		// for each file (slice is already ordered from high to low)
		for _, fId := range fileIds {

			fSize := filemap[fId][0]
			fPos := filemap[fId][1]
			// fmt.Println("File ", fId, " Size ", fSize, " Pos ", fPos)
			gSize := 0
			gPos := 0
			// from beginning to end, scan for am empty blocj where the current file fits in

			for i, v := range blockmap {
				if v == "." {
					// size of gap
					for l := i; l < len(blockmap); l++ {
						if blockmap[l] != "." {
							gSize = l - i
							break
						}
					}
					fmt.Println("Gap at Pos ", i, " Size of Gap ", gSize)
					if fSize <= gSize {
						gPos = i
						break
					}
				}
			}

			if fSize <= gSize && fPos > gPos {
				// move file
				for n := 0; n < fSize; n++ {
					fmt.Println("move file ", blockmap[fPos+n], " from ", fPos+n, " to ", gPos+n)
					blockmap[gPos+n] = blockmap[fPos+n]
					blockmap[fPos+n] = "."
				}
				filemap[fId][1] = gPos
			}
		}

		fmt.Println(blockmap)

		checksum := 0
		for k, v := range blockmap {
			if v != "." {
				v1, _ := strconv.Atoi(v)
				checksum += k * v1
			}
		}

		return checksum
	}

	// solve part 1 here

	for i, v := range blockmap {
		if v == "." {
			// j = Index of file
			j := len(blockmap) - 1
			// go backward from end to find a file
			for k := j; k >= 0; k-- {
				if blockmap[k] != "." {
					j = k
					break
				}
			}
			//fmt.Println("i ", i, "j ", j)
			//fmt.Println(blockmap)
			// if position of file is behind current postion
			if j > i {
				// move file
				blockmap[i] = blockmap[j]
				blockmap[j] = "."
			} else {
				fmt.Println("Break")
				break
			}
		}
	}

	//fmt.Println(blockmap)
	checksum := checksum(blockmap)
	return checksum
}

func checksum(input []string) int {
	checksum := 0
	for k, v := range input {
		if v != "." {
			v1, _ := strconv.Atoi(v)
			checksum += k * v1
		}
	}
	return checksum
}
