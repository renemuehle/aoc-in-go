package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {

	lines := strings.Split(strings.TrimSpace(input), "\n")

	// helper function, because go doesn't allow to directly
	// securely set a character, watch the bounds
	setByte := func(i, j int, ch byte) {
		if i < 0 || i >= len(lines) {
			return
		}
		if j < 0 || j >= len(lines[i]) {
			return
		}
		b := []byte(lines[i])
		b[j] = ch
		lines[i] = string(b)
	}

	// Part 1
	if !part2 {
		count := 0

		// loops for each line to set the beams
		for i, line := range lines {
			for j, ch := range line {
				// for the start, the beam goes down
				if ch == 'S' {
					setByte(i+1, j, '|')
				}
				// split the beam
				if ch == '^' && lines[i-1][j] == '|' {
					setByte(i, j-1, '|')
					setByte(i, j+1, '|')
					// count a split
					count++
					// continue the split beams, if possible
					if i < len(lines)-1 && lines[i+1][j-1] == '.' {
						setByte(i+1, j-1, '|')
					}
					if i < len(lines)-1 && lines[i+1][j+1] == '.' {
						setByte(i+1, j+1, '|')
					}
				}
				// continue the current beam
				if ch == '|' {
					// if possible
					if i < len(lines)-1 && lines[i+1][j] == '.' {
						setByte(i+1, j, '|')
					}
				}
			}

			// Debug-Ausgabe
			//for _, l := range lines {
			//	fmt.Println(l)
			//}
			//fmt.Println("")
		}

		// count
		fmt.Println(" Splits: ", count)
		return count
	}
	// Part 2
	if part2 {

		colStart := 0

		// like part 1, set all beams
		for i, line := range lines {
			for j, ch := range line {
				if ch == 'S' {
					setByte(i+1, j, '|')
					colStart = j
				}
				if ch == '^' && lines[i-1][j] == '|' {
					setByte(i, j-1, '|')
					setByte(i, j+1, '|')
					if i < len(lines)-1 && lines[i+1][j-1] == '.' {
						setByte(i+1, j-1, '|')
					}
					if i < len(lines)-1 && lines[i+1][j+1] == '.' {
						setByte(i+1, j+1, '|')
					}
				}
				if ch == '|' {
					if i < len(lines)-1 && lines[i+1][j] == '.' {
						setByte(i+1, j, '|')
					}
				}
			}
		}

		cache := make(map[string]int64)
		// recursive function to find all path
		// because of the big number of paths a recursion with cache is needed, otherwise calculation takes ages ...
		count := countPath(lines, 0, colStart, cache)
		// count
		fmt.Println(" Paths: ", count)
		return count
	}
	// solve part 1 here
	return 42
}

func countPath(lines []string, i, j int, cache map[string]int64) int64 {

	// key for current position
	pos := strconv.Itoa(i) + ":" + strconv.Itoa(j)

	var anz int64 = 0

	// if we already know the value, return it
	anz, exists := cache[pos]
	if exists {
		return anz
	}

	// check the bounds of the input , prevent errors
	if i < 0 || i >= len(lines) || j < 0 || j >= len(lines[i]) {
		return 0
	}

	ch := lines[i][j]
	switch ch {
	case 'S':
		// fmt.Println(" Start S go down ", i+1, "/", j)
		anz = countPath(lines, i+1, j, cache)
	case '^':
		// fmt.Println(" Go Left I/J ", i, "/", j-1)
		anzl := countPath(lines, i, j-1, cache)
		anz += anzl
		// fmt.Println(" Go Right I/J ", i, "/", j+1)
		anzr := countPath(lines, i, j+1, cache)
		anz += anzr
	case '|':
		if i < len(lines)-1 && lines[i+1][j] != '.' {
			// fmt.Println(" Go Down I/J ", i+1, "/", j )
			anz = countPath(lines, i+1, j, cache)
		} else if i == len(lines)-1 {
			// fmt.Println(" End of Recursion I/J ", i, "/", j)
			return 1
		}
	default:
		// Anythin else
		anz = 0
	}
	// cache the result, we don't want to calculate this again
	cache[pos] = anz
	return anz
}
