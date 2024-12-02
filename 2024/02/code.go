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
	if part2 {
		var secureReports []string

		for _, report := range strings.Split(strings.TrimSpace(input), "\n") {
			// fmt.Println(line)

			// Split line into Slice of "Words", splits by "consecutive white space characters"
			level := strings.Fields(report)

			var secure = testSecure(level)

			if secure {
				//fmt.Print(level, " is secure")
				secureReports = append(secureReports, report)
			} else {
				fmt.Println(level, " is NOT secure")

				for i := 0; i < len(level); i++ {
					var sublevel = RemoveIndex(level, i)
					fmt.Println(" Sublevel ", sublevel)
					secure = testSecure(sublevel)
					if secure {
						fmt.Println(" Sublevel ", sublevel, " is secure")
						secureReports = append(secureReports, report)
						break
					}
				}
			}

			// Number of save reports
		}
		return len(secureReports)
	}
	// solve part 1 here

	var secureReports []string

	for _, report := range strings.Split(strings.TrimSpace(input), "\n") {
		// fmt.Println(line)

		// Split line into Slice of "Words", splits by "consecutive white space characters"
		level := strings.Fields(report)

		var secure = testSecure(level)

		if secure {
			//fmt.Print(level, " is secure")
			secureReports = append(secureReports, report)
		} else {
			fmt.Println(level, " is NOT secure")
		}

		// Number of save reports
	}
	return len(secureReports)

}

func testSecure(input []string) bool {

	var secure bool = true
	var up bool = false
	var down bool = false

	for i := 1; i < len(input); i++ {

		i1, err := strconv.Atoi(input[i-1])
		if err != nil {
			// ... handle error
			panic(err)
		}
		i2, err := strconv.Atoi(input[i])
		if err != nil {
			// ... handle error
			panic(err)
		}

		if i == 1 {
			if i1 < i2 {
				up = true
			} else if i1 > i2 {
				down = true
			}
		}

		if i1 < i2 && (i2-i1) <= 3 && (i2-i1) >= 1 && up && secure {
			// go up
			continue
		} else if i1 > i2 && (i1-i2) <= 3 && (i1-i2) >= 1 && down && secure {
			// go down
			continue
		} else {
			// error
			fmt.Println(" ERROR i1 ", i1, " i2 ", i2, " up ", up, " down ", down)
			secure = false
			break
		}
	}
	return secure
}

func RemoveIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	ret = append(ret, s[index+1:]...)
	return ret
}
