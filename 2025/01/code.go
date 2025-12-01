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

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {

	// int variables
	dial := 50
	countZero := 0

	// Part 1
	if !part2 {
		for _, line := range strings.Split(strings.TrimSpace(input), "\n") {

			// fmt.Print("Part 1 Line : ", line, "  ->  ", " dial before ", dial)

			// Parse number of steps to int
			steps, err := strconv.Atoi(line[1:])
			if err != nil {
				// ... handle error
				panic(err)
			}

			// choose direction und turn dial the number of steps
			if line[0:1] == "L" {
				dial -= steps
			} else {
				dial += steps
			}

			// number of steps can get the dial above 99 (which means it starts with 0 again)
			for dial > 99 {
				dial = dial - 100
			}
			// number of steps can get the dial below 0 (which means it starts with 99 again)
			for dial < 0 {
				dial = dial + 100
			}

			// count the zeros
			if dial == 0 {
				countZero++
			}
			// fmt.Println(" dial after  ", dial, "CountZero", countZero)
		}
		fmt.Println("Part 1 Dial : ", dial, "CountZero", countZero)
		return countZero
	}

	// int variables (again)
	dial = 50
	countZero = 0

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {

		for _, line := range strings.Split(strings.TrimSpace(input), "\n") {

			// fmt.Print("Part 2 Line : ", line, "  ->  ", " dial before ", dial)

			// Parse number of steps to int
			steps, err := strconv.Atoi(line[1:])
			if err != nil {
				// ... handle error
				panic(err)
			}

			// different approach for 2. task
			// we loop much more and go only one step each loop - to make sure to not miss a '0'
			for steps > 0 {
				// depending on the direction, go one step up or down
				if line[0:1] == "L" {
					dial--
				} else {
					dial++
				}
				// if we go above 99 the last step (can actually only be 100)
				// reset dial to 0
				if dial > 99 {
					dial = dial - 100
				}
				// if we go below 0 with the last stepp (can only be -1)
				// reset dial to 99
				if dial < 0 {
					dial = dial + 100
				}
				// count the '0'
				if dial == 0 {
					countZero++
				}
				// one step each loop
				steps--
			}
			// fmt.Println(" dial after  ", dial, "CountZero", countZero)
		}
		fmt.Println("PART 2 Dial : ", dial, "CountZero", countZero)
	}

	return countZero
}
