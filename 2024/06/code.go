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
	// when you're ready to do part 2, remove this "not implemented" block

	// Split the input into lines and columns
	zeilen := strings.Split(input, "\n")
	anzZeilen := len(zeilen)
	var anzSpalten int
	m := make([][]string, anzZeilen)

	// initial position and direction of the guard
	guardPos := []int{0, 0}
	direction := "U" // "U" / "D" / "L" / "R"
	for i := 0; i < anzZeilen; i++ {
		spalten := strings.Split(zeilen[i], "")
		anzSpalten = len(spalten)
		m[i] = spalten
		for j := 0; j < anzSpalten; j++ {
			if m[i][j] == "^" {
				guardPos[0] = i
				guardPos[1] = j
			}
		}
	}

	// fmt.Println("Zeilen /Spalten ", anzZeilen, " / ", anzSpalten)
	// fmt.Println("Guard Zeile/Spalte:", guardPos[0], " / ", guardPos[1])
	// printArray(m)

	if part2 {
		fmt.Println("===================== PART2 ===================== ")
		loopCount := 0

		// prepare map m2
		// m2 - we only want to test places on the route of the guard
		m2 := deepCopyStringSlice(m)

		// copy the original starting position
		guardPos2 := []int{guardPos[0], guardPos[1]}
		direction := "U" // "U" / "D" / "L" / "R"

		for direction != "X" {
			direction = walkStep(m2, guardPos2, direction, anzZeilen, anzSpalten)
			//fmt.Println("Guard Zeile/Spalte/Direction:", guardPos2[0], " / ", guardPos2[1], " / ", direction)
		}
		// Map to check for possible obstacles
		//fmt.Println("Map for possible obstacles")
		// printArray(m2)
		fmt.Println("Postions from Route of Guard: ", countX(m2))

		// m2 has all positions of the guard marked with "X"
		// test all these positions if they produce a loop
		for i, line := range m2 {
			for j, col := range line {

				if col == "X" {

					//fmt.Println("Position with \"X\" ", i, "/", j)

					// reset the start-position of te guard
					guardPos2 := []int{guardPos[0], guardPos[1]}
					direction := "U" // "U" / "D" / "L" / "R"

					// new empty map
					m3 := deepCopyStringSlice(m)
					// set new obstacle
					m3[i][j] = "#"

					// new map for counting visits
					v := make([][]int, anzZeilen)
					for i := 0; i < anzZeilen; i++ {
						v[i] = make([]int, anzSpalten)
					}

					// walk m3
					for direction != "X" {
						direction = walkStep(m3, guardPos2, direction, anzZeilen, anzSpalten)
						//fmt.Println("Guard Zeile/Spalte/Direction:", guardPos2[0], " / ", guardPos2[1], " / ", direction)
						v[guardPos2[0]][guardPos2[1]] += 1

						// no need to check the whole slice. just check the last incremented postion
						// if we have visited a position more than 4 times it is a loop
						if v[guardPos2[0]][guardPos2[1]] > 4 {
							loopCount++
							//fmt.Println(" -> is loop ", direction)
							break
						}
						// if it is a loop, count and break
						/*
							if isLoop(v) {
								loopCount++
								fmt.Println(" -> is loop")
								break
							}
						*/
					} // end if it not a loop

					// if direction == "X" {
					//  fmt.Println(" -> is not a loop ", direction)
					// }
				}
			}
		}

		fmt.Println("Anzahl Loops", loopCount)
		return loopCount
	}

	// solve part 1 here

	fmt.Println("===================== PART1 ===================== ")

	printArray(m)

	for direction != "X" {
		direction = walkStep(m, guardPos, direction, anzZeilen, anzSpalten)
		fmt.Println("Guard Zeile/Spalte/Direction:", guardPos[0], " / ", guardPos[1], " / ", direction)
		//printArray(m)
	}
	printArray(m)
	var anz = countX(m)
	fmt.Println("Number of visits ", anz)
	return anz
}

/* Walk one Step in the given direction, check for bounds, return the direction of the next step */
func walkStep(m [][]string, gpos []int, direction string, anzZeilen int, anzSpalten int) string {

	// mark current position
	m[gpos[0]][gpos[1]] = "X"

	//current := []int{gpos[0], gpos[1]}

	//fmt.Println("Walk ", direction, " ", gpos[0], "/", gpos[1], " ", anzZeilen, "/", anzSpalten)

	// check if the next move would leave the map
	if (direction == "U" && gpos[0] == 0) || (direction == "D" && gpos[0] >= anzZeilen-1) ||
		(direction == "L" && gpos[1] == 0) || (direction == "R" && gpos[1] >= anzSpalten-1) {
		// fmt.Println(" Leave X 1 ", direction, "->X ", gpos[0], "/", gpos[1], " ", anzZeilen, "/", anzSpalten)
		direction = "X"
		return direction
	}

	clear := false

	//fmt.Print("Turn Direction from ", direction)
	for !clear {
		// check if next move would hit a abstacle -> turn
		// ist possible that 2 turns are needed
		if direction == "U" && m[gpos[0]-1][gpos[1]] == "#" {
			direction = "R"
		} else if direction == "D" && m[gpos[0]+1][gpos[1]] == "#" {
			direction = "L"
		} else if direction == "L" && m[gpos[0]][gpos[1]-1] == "#" {
			direction = "U"
		} else if direction == "R" && m[gpos[0]][gpos[1]+1] == "#" {
			direction = "D"
		} else {
			clear = true
		}
	}
	//fmt.Println(" to Direction ", direction)

	// check if the next move would leave the map
	if (direction == "U" && gpos[0] == 0) || (direction == "D" && gpos[0] >= anzZeilen-1) ||
		(direction == "L" && gpos[1] == 0) || (direction == "R" && gpos[1] >= anzSpalten-1) {
		// fmt.Println(" Leave X 2 ", direction, "->X ", gpos[0], "/", gpos[1], " ", anzZeilen, "/", anzSpalten)
		direction = "X"
		return direction
	}

	// walk
	if direction == "U" {
		gpos[0] -= 1
	} else if direction == "D" {
		gpos[0] += 1
	} else if direction == "L" {
		gpos[1] -= 1
	} else if direction == "R" {
		gpos[1] += 1
	}

	// mark position
	m[gpos[0]][gpos[1]] = "X"

	// return direction
	return direction
}

func countX(input [][]string) int {
	count := 0
	for _, line := range input {
		for _, col := range line {
			if col == "X" {
				count++
			}
		}
	}
	return count
}

func isLoop(input [][]int) bool {
	loop := false
	for _, line := range input {
		for _, col := range line {
			if col > 4 {
				loop = true
				break
			}
		}
	}
	return loop
}

func deepCopyStringSlice(input [][]string) [][]string {
	duplicate := make([][]string, len(input))
	for i := range input {
		duplicate[i] = make([]string, len(input[i]))
		copy(duplicate[i], input[i])
	}
	return duplicate
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

/* print a slice of string */
func printIntArray(input [][]int) {
	fmt.Println("=======================================")
	for _, row := range input {
		for _, colValue := range row {
			fmt.Print(colValue)
		}
		fmt.Println()
	}
}
