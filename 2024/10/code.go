package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

var cache map[string]int

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

	// initialize the cache
	cache = make(map[string]int)

	zeilen := strings.Split(input, "\n")
	anzZeilen := len(zeilen)
	anzSpalten := len(zeilen[0])
	m := make([][]int, anzZeilen)

	for i := 0; i < anzZeilen; i++ {
		//spalten := strings.Split(zeilen[i], "")
		//anzSpalten = len(spalten)
		//m[i] = spalten
		spalten := make([]int, anzSpalten)

		for j, s := range zeilen[i] {
			spalten[j], _ = strconv.Atoi(string(s))
		}
		m[i] = spalten
	}

	printArray(m)

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		sum := 0

		for i, row := range m {
			for j, colValue := range row {
				if colValue == 0 {
					fmt.Print("Count Path for ", m[i][j], " - ", i, "/", j)
					s := countPath(m, i, j, i, j)
					fmt.Println(" = ", s)
					sum += s
				}
			}
		}

		fmt.Println("Cache ")
		fmt.Println(cache)
		fmt.Println("Summe ", sum)
		return sum
	}
	// solve part 1 here

	sum := 0

	for i, row := range m {
		for j, colValue := range row {
			if colValue == 0 {
				s := countPath(m, i, j, i, j)
				fmt.Print("Count Path for ", m[i][j], " - ", i, "/", j, " = ", s)
				sum += s
			}
		}
	}

	fmt.Println("Cache ")
	fmt.Println(cache)
	fmt.Println("Summe ", sum)

	return len(cache)
}

func countPath(m [][]int, start, end, i, j int) int {
	//fmt.Println("Count Path for ", m[i][j], " - ", i, "/", j, " -- ")

	anz := 0

	if m[i][j] == 9 {
		cache[fmt.Sprintf("%d:%d->%d:%d", start, end, i, j)] += 1
		anz = 1
	} else {
		// Number above is exactly 1 lower
		if i > 0 && m[i][j] == (m[i-1][j]-1) {
			anz += countPath(m, start, end, (i - 1), j)
			// fmt.Println("Above ", m[i][j], " - ", i, "/", j, " = ", anz, " - ")
		}
		// Number right is exactly 1 lower
		if j < (len(m[i])-1) && m[i][j] == (m[i][j+1]-1) {
			anz += countPath(m, start, end, i, (j + 1))
			// fmt.Println("Right ", m[i][j], " - ", i, "/", j, " = ", anz, " - ")
		}
		// Number below is exactly 1 lower
		if i < (len(m)-1) && m[i][j] == (m[i+1][j]-1) {
			anz += countPath(m, start, end, (i + 1), j)
			// fmt.Println("Below ", m[i][j], " - ", i, "/", j, " = ", anz, " - ")
		}
		// Number right is exactly 1 lower
		if j > 0 && m[i][j] == (m[i][j-1]-1) {
			anz += countPath(m, start, end, i, (j - 1))
			// fmt.Println("Left ", m[i][j], " - ", i, "/", j, " = ", anz, " - ")
		}
	}

	//fmt.Println("Path for ", m[i][j], " - ", i, "/", j, " = ", anz)
	return anz
}

/* print a slice of string */
func printArray(input [][]int) {
	fmt.Println("=======================================")
	for _, row := range input {
		for _, colValue := range row {
			fmt.Print(colValue)
		}
		fmt.Println()
	}
}
