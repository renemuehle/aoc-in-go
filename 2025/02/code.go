package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)

	// Somehow Harness fails for day 2 :/
	// manually load the file
	// file, _ := FileAsString("input-example.txt")
	// file, _ := FileAsString("input-user.txt")
	// run(true, file)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if !part2 {

		invalid := 0

		for _, rangeStr := range strings.Split(strings.TrimSpace(input), ",") {

			// String split, get start and end of range
			start, err := strconv.Atoi(strings.Split(rangeStr, "-")[0])
			if err != nil {
				// ... handle error
				panic(err)
			}
			end, err := strconv.Atoi(strings.Split(rangeStr, "-")[1])
			if err != nil {
				// ... handle error
				panic(err)
			}

			// loop trough the range
			for id := start; id <= end; id++ {

				idStr := strconv.Itoa(id)
				// find the middle of the string
				mid := len(idStr) / 2
				// split the string in 2 parts
				left := idStr[:mid]
				right := idStr[mid:]
				// compare parts - if eqal ist an invalid id
				if left == right {
					// fmt.Print(" Invalid ", idStr)
					invalid += id
				}
			}
		}
		fmt.Println("\nResult ", invalid)
		return invalid
	}

	if part2 {

		invalid := 0
		// use a map as replacement for set - we need to make sure to count an invalid id only once
		// if we directly add each invalid id (like in the solution from part 1) to the solution
		// an id like "111111" would be found several times like 1/1/1/1/1/1 or 11/11/11 or 111/111
		// and thus add several times
		invSet := make(map[string]int)

		for _, rangeStr := range strings.Split(strings.TrimSpace(input), ",") {

			// String split, get start and end of range
			start, err := strconv.Atoi(strings.Split(rangeStr, "-")[0])
			if err != nil {
				// ... handle error
				panic(err)
			}
			end, err := strconv.Atoi(strings.Split(rangeStr, "-")[1])
			if err != nil {
				// ... handle error
				panic(err)
			}

			// loop trough the range
			for id := start; id <= end; id++ {

				idStr := strconv.Itoa(id)

				// try to find patterns, split the string in substrings from length 1 to length/2
				for j := 1; j <= len(idStr)/2; j++ {

					// helper-function for the split
					idArr := splitStringInParts(idStr, j)

					// fmt.Println(" idArr : ", idArr)

					// first element of the string array for the compare
					compareLeft := idArr[0]
					inv := true
					for _, part := range idArr {
						// if one of the other elemts is not eqal, the id is not invalid, we can stop searching
						if part != compareLeft {
							inv = false
							break
						}
					}
					// if all parts are eqal, the id is invalid

					if inv {
						// fmt.Print(" Invalid ", idStr)
						// add the id to the "Set", using the idStr as key makes sure each invalid id only added once
						invSet[idStr] = id
					}
				}
			}
		}

		// calculate the sum of all found invalid ids
		for key := range invSet {
			invalid += invSet[key]
		}

		fmt.Println("\nResult ", invalid)
		return invalid
	}
	// solve part 1 here
	return 42
}

func FileAsString(name string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := fmt.Sprintf("%s/%s", cwd, name)
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	fileString := string(file)
	return fileString, nil
}

func splitStringInParts(splitStr string, size int) []string {
	var parts []string
	for i := 0; i < len(splitStr); i += size {
		end := i + size
		if end > len(splitStr) {
			end = len(splitStr)
		}
		parts = append(parts, splitStr[i:end])
	}
	// mindestens 2 Teile?
	if len(parts) >= 2 {
		return parts
	}
	return nil
}
