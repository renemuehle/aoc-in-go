package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// 8-neighbors (N, NE, E, SE, S, SW, W, NW)
var dirs = [8][2]int{
	{-1, 0},  // N
	{-1, 1},  // NE
	{0, 1},   // E
	{1, 1},   // SE
	{1, 0},   // S
	{1, -1},  // SW
	{0, -1},  // W
	{-1, -1}, // NW
}

func run(part2 bool, input string) any {

	lines := strings.Split(strings.TrimSpace(input), "\n")

	inBounds := func(r, c int) bool {
		return r >= 0 && r < len(lines) && c >= 0 && c < len(lines[r])
	}

	// Part 1
	if !part2 {
		lessThan4 := 0
		// loop through lines (rows)
		for r, line := range lines {
			// loop through line (cols)
			for c := range line {
				// if no paper,  continue
				if line[c] != '@' {
					continue
				}
				count := 0
				// for all 8 directions
				for _, d := range dirs {
					// coordinates of the neighbour
					nr, nc := r+d[0], c+d[1]
					// check if coordinates are valid, to avoid "out of range" error
					if inBounds(nr, nc) {
						// if paper, then count
						if lines[nr][nc] == '@' {
							count++
						}
					}
				}
				// if less then 4 neighbous
				if count < 4 {
					lessThan4++
				}

			}
		}
		return lessThan4
	}
	// Part 2
	if part2 {
		paperMap := make([][]rune, len(lines))

		for i, line := range lines {
			paperMap[i] = []rune(line)
		}

		// count, how many rolls are removed
		removedTotal := 0
		// loop condition, stop wehen 0
		removed := 1
		// as long as more rolls get removed
		for removed > 0 {

			// reset for each loop
			removed = 0
			// go through the map
			for r, line := range paperMap {
				for c := range line {
					if line[c] != '@' {
						continue
					}
					// found a paper roll
					count := 0
					// check all directions
					for _, d := range dirs {
						nr, nc := r+d[0], c+d[1]
						if inBounds(nr, nc) {
							if paperMap[nr][nc] == '@' {
								count++
							}
						}
					}
					// if the roll has less than 4 neighbours, remove it, and count
					if count < 4 {
						paperMap[r][c] = 'x'
						removed++
						removedTotal++
					}
				}
			}
		}
		// total number of removed rolls
		return removedTotal
	}
	// solve part 1 here
	return 42
}
