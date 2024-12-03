package main

import (
	"fmt"
	"regexp"
	"strconv"

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

	if part2 {
		var sum int = 0
		// keep the state if we have ay do or a don't
		var do bool = true
		// find all commands mul or do or don't
		pattern := regexp.MustCompile(`mul\((\d{1,3})\,(\d{1,3})\)|do\(\)|don\'t\(\)`)
		fmt.Println(pattern)

		// For each match of the regex in the content.
		for _, matches := range pattern.FindAllStringSubmatch(input, -1) {
			// check whick command
			if matches[0] == `do()` {
				do = true
			} else if matches[0] == `don't()` {
				do = false
			} else if do {
				// same as 1st solution, get numbers and calculate sum
				fmt.Println(matches)
				i1, err := strconv.Atoi(matches[1])
				if err != nil {
					// ... handle error
					panic(err)
				}
				i2, err := strconv.Atoi(matches[2])
				if err != nil {
					// ... handle error
					panic(err)
				}
				sum += i1 * i2
			}
		}
		return sum
	}

	// solve part 1 here

	var sum int = 0

	// split for lines not needed, just apply the regex to the whole input, the '\n' don't matter
	//for _, line := range strings.Split(strings.TrimSpace(input), "\n") {

	pattern := regexp.MustCompile(`mul\((\d{1,3})\,(\d{1,3})\)`)

	// For each match of the regex in the content.
	//for _, matches := range pattern.FindAllStringSubmatch(line, -1) {
	for _, matches := range pattern.FindAllStringSubmatch(input, -1) {

		//fmt.Println(matches)
		// if we have a match, than the match itself is the first element,
		// the first number --- (group) are called submatch in go
		i1, err := strconv.Atoi(matches[1])
		if err != nil {
			// ... handle error
			panic(err)
		}
		// the second number --- (group) are called submatch in go
		i2, err := strconv.Atoi(matches[2])
		if err != nil {
			// ... handle error
			panic(err)
		}
		sum += i1 * i2
	}
	//}

	return sum
}
