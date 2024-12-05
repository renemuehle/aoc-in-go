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
	// when you're ready to do part 2, remove this "not implemented" block

	parts := strings.Split(input, "\n\n")

	//fmt.Println("Rules ")
	part0 := strings.Split(parts[0], "\n")
	rules := make([][]string, len(part0))
	for i, row := range part0 {
		rules[i] = strings.Split(row, "|")
	}
	//fmt.Println(rules)

	//fmt.Println("Updates")
	part1 := strings.Split(parts[1], "\n")
	updates := make([][]string, len(part1))
	for i, row := range part1 {
		updates[i] = strings.Split(row, ",")
	}
	//fmt.Println(updates)

	if part2 {
		// solve part 2 here
		fmt.Println("===================== PART2 ===================== ")
		sum := 0

		for _, upd := range updates {
			isOk := true
			// j is the first page for checking the rules
			for j := 0; j < len(upd); j++ {
				// k is for all pages following j
				for k := j + 1; k < len(upd); k++ {
					// fmt.Print(" ", upd[j], "->", upd[k])
					// check if any update break the rules
					for _, rule := range rules {
						// if we found two pages which break a rule, reorder the pages
						if rule[0] == upd[k] && rule[1] == upd[j] {

							isOk = false
							// if we found a broken rule, switch pages
							tmp := upd[k]
							upd[k] = upd[j]
							upd[j] = tmp
							// actually restart is not needed, becaus all pages before this possition are already in order
							// and restart to sort this update
							// j = 0
							// k = j + 1
						}
					}
				}
			}
			if !isOk {
				// fmt.Println("Update", upd, " was NOT OK, is ordered now ")
				// find middle page
				mPageNo, _ := strconv.Atoi(upd[(len(upd)-1)/2])
				// fmt.Println("Middle Page No ", mPageNo)
				sum += mPageNo
			} else {
				//fmt.Println("Update", upd, " is OK ")
			}
		}

		return sum
	}

	// solve part 1 here
	fmt.Println("===================== PART1 ===================== ")
	sum := 0

	for _, upd := range updates {
		isOk := true
		// j is the first page for checking the rules
		for j := 0; j < len(upd); j++ {
			// k is for all pages following j
			for k := j + 1; k < len(upd); k++ {
				// fmt.Print(" ", upd[j], "->", upd[k])
				// check if any update break the rules
				for _, rule := range rules {
					// if we found two pages which break a rule, stop the search
					if rule[0] == upd[k] && rule[1] == upd[j] {
						// fmt.Print("BREAK ", rule, " ")
						isOk = false
						break
					}
				}
				if !isOk {
					break
				}
			}
			if !isOk {
				break
			}
		}
		if !isOk {
			// fmt.Println("Update", upd, " is NOT OK ")
		} else {
			// if the update is ok
			// fmt.Println("Update", upd, " is OK ")
			// find middle page
			mPageNo, _ := strconv.Atoi(upd[(len(upd)-1)/2])
			// fmt.Println("Middle Page No ", mPageNo)
			sum += mPageNo
		}
	}

	return sum
}
