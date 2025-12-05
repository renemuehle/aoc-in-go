package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {

	input_parts := strings.Split(strings.TrimSpace(input), "\n\n")
	goodRangesStr := strings.Split(strings.TrimSpace(input_parts[0]), "\n")
	ingredientStr := strings.Split(strings.TrimSpace(input_parts[1]), "\n")

	goodRanges := make([][2]int, len(goodRangesStr))
	for i, f := range goodRangesStr {
		fa := strings.Split(f, "-")

		b, _ := strconv.Atoi(fa[0])
		e, _ := strconv.Atoi(fa[1])

		goodRanges[i] = [2]int{b, e}
	}

	// Part 1
	if !part2 {
		// a 'Set' of all good ingredients
		goodingredients := make(map[int]int)
		for _, f := range ingredientStr {
			i, _ := strconv.Atoi(f)
			// go through all ranges
			for _, goodRange := range goodRanges {
				// if ingredient is in range
				if i >= goodRange[0] && i <= goodRange[1] {
					// add ingredintId to Set
					goodingredients[i] = i
				}
			}
		}
		return len(goodingredients)
	}
	// Part 2
	if part2 {
		// 1. idea brute force -> doesn't work

		// 2. idea
		// go trough the ranges and sort them, so that overlapping ranges are together
		goodRanges2 := sortRanges(goodRanges)
		// compare ranges for overlapping an merge them
		goodRanges3 := mergeRanges(goodRanges2)

		anzGoodIds := 0
		// number of good ids can be calculated from end-begin for each merged range
		for _, g := range goodRanges3 {
			anzGoodIds += (g[1] - g[0] + 1)
		}
		return anzGoodIds
	}

	return 42
}

// sort ranges, first by by 1st id, and if equal by 2nd id
func sortRanges(idRanges [][2]int) [][2]int {
	sort.Slice(idRanges, func(i, j int) bool {
		if idRanges[i][0] == idRanges[j][0] {
			return idRanges[i][1] < idRanges[j][1]
		}
		return idRanges[i][0] < idRanges[j][0]
	})
	return idRanges
}

// the list if ranges is expected to be sorted, only direct neigbours are checked for overlapping
// if ranges are overlapping they get merged to one (bigger) range
func mergeRanges(idRanges [][2]int) [][2]int {

	merged := make([][2]int, 0, len(idRanges))

	cur := [2]int{idRanges[0][0], idRanges[0][1]}

	for i := 1; i < len(idRanges); i++ {
		s, e := idRanges[i][0], idRanges[i][1]
		// if the start of the loop element is lower or equal the last merged,
		// than the ranges are adjacent or overlapping
		if s <= cur[1]+1 {
			// the end is greater than zhe current end
			if e > cur[1] {
				cur[1] = e // erweitern
			}
			// no overlapping for this element, save the merged and the new current
		} else {
			merged = append(merged, cur)
			cur = [2]int{s, e}
		}
	}
	merged = append(merged, cur)

	return merged
}
