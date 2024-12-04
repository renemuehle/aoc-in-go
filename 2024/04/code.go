package main

import (
	"fmt"
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

	// Split the input into lines and columns
	zeilen := strings.Split(input, "\n")
	anzZeilen := len(zeilen)
	var anzSpalten int
	m := make([][]string, anzZeilen)

	for i := 0; i < anzZeilen; i++ {
		spalten := strings.Split(zeilen[i], "")
		anzSpalten = len(spalten)
		m[i] = spalten
	}

	//fmt.Println("Zeilen:", anzZeilen)
	//fmt.Println("Spalten:", anzSpalten)

	//printArray(m)

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		var sum int = 0
		// forward
		for i := 0; i < (anzZeilen - 2); i++ {

			// subslice "subi" - 3 rows
			subi := m[i : i+3]
			//fmt.Println("SUB I von ", i)
			//printArray(subi)

			for j := 0; j < (anzSpalten - 2); j++ {

				// go is by reference - a copy is needed
				subj := make([][]string, len(subi))
				copy(subj, subi)
				// subj is a copy, so it contains 3 full rows in each iteration

				//fmt.Println("SUB J von ", j)
				//printArray(subj)

				// in the 4 selected rows, loop through all sub-slices with 3 columns
				for k := range subj {
					subj[k] = subi[k][j : j+3]
				}

				//fmt.Println("SUB von ", i, ":", j)
				//printArray(subj)

				// we have a 3x3 subslice - count all MAS in this subslice
				anz := findMas(subj)
				//fmt.Println("Anzahl von ", i, ":", j, " == ", anz)
				//fmt.Println("=======================================")
				sum += anz

			}
		}

		fmt.Println("\n", "SUMME ", sum)

		return sum
	}

	// solve part 1 here

	var sum int = 0
	// forward
	for i := 0; i < (anzZeilen - 3); i++ {

		// subslice "subi" - 4 rows
		subi := m[i : i+4]
		//fmt.Println("SUB I von ", i)
		//printArray(subi)

		for j := 0; j < (anzSpalten - 3); j++ {

			// go is by reference - a copy is needed
			subj := make([][]string, len(subi))
			copy(subj, subi)
			// subj is a copy, so it contains 4 full rows in each iteration

			//fmt.Println("SUB J von ", j)
			//printArray(subj)

			// in the 4 selected rows, loop through all sub-slices with 4 columns
			for k := range subj {
				subj[k] = subi[k][j : j+4]
			}

			//fmt.Println("SUB von ", i, ":", j)
			//printArray(subj)

			// we have a 4x4 subslice - count all XMAS in this subslice
			anz := findXmas(subj, i, j)
			//fmt.Println("Anzahl von ", i, ":", j, " == ", anz)
			//fmt.Println("=======================================")
			sum += anz

		}
	}

	fmt.Println("\n", "SUMME ", sum)

	return sum
}

func findXmas(input [][]string, zeile int, spalte int) int {

	// extract the 10 possible XMAS
	s := make([]string, 10)
	// 4 rows
	// only for the first line, we need all 4
	if zeile == 0 {
		s[0] = input[0][0] + input[0][1] + input[0][2] + input[0][3]
		s[1] = input[1][0] + input[1][1] + input[1][2] + input[1][3]
		s[2] = input[2][0] + input[2][1] + input[2][2] + input[2][3]
		s[3] = input[3][0] + input[3][1] + input[3][2] + input[3][3]
	} else {
		// we already hat the first 3 lines in the last run
		// so only add/check the last line
		s[3] = input[3][0] + input[3][1] + input[3][2] + input[3][3]
	}
	// 4 columns
	if spalte == 0 {
		s[4] = input[0][0] + input[1][0] + input[2][0] + input[3][0]
		s[5] = input[0][1] + input[1][1] + input[2][1] + input[3][1]
		s[6] = input[0][2] + input[1][2] + input[2][2] + input[3][2]
		s[7] = input[0][3] + input[1][3] + input[2][3] + input[3][3]
	} else {
		// we already hat the first 3 columns in the last run
		// so only add/check the last column
		s[7] = input[0][3] + input[1][3] + input[2][3] + input[3][3]
	}

	// 2 diagonals
	s[8] = input[0][0] + input[1][1] + input[2][2] + input[3][3]
	s[9] = input[0][3] + input[1][2] + input[2][1] + input[3][0]

	// search forward and backward and sum up the occurences
	var sum int = contains(s, "XMAS") + contains(s, "SAMX")

	return sum
}

func findMas(input [][]string) int {

	// extract the 4 possible MAS
	s := make([]string, 1)
	s[0] = input[0][0] + input[0][2] + input[1][1] + input[2][0] + input[2][2]

	// search forward and backward and sum up the occurences
	var sum int = contains(s, "MMASS") + contains(s, "SSAMM") + contains(s, "MSAMS") + contains(s, "SMASM")

	return sum
}

/* print a slice of string */
func printArray(input [][]string) {
	fmt.Println("=======================================")
	for _, row := range input {
		for _, colValue := range row {
			fmt.Print(colValue)
		}
		fmt.Println()
	}
}

/* Count the number of occurences of a string in a slice of strings*/
func contains(s []string, e string) int {
	var anz int = 0
	for _, a := range s {
		if a == e {
			anz++
		}
	}
	//fmt.Println(s, "Anzahl: ", anz)
	return anz
}
