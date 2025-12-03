package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)

	// Somehow Harness fails :/
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
	// Part 1
	if !part2 {
		sumJolatage := 0
		for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
			joltage := findTwoLargestDigits(line)
			// fmt.Println("Part 1 Line : ", line, "  ->  ", joltage)
			sumJolatage += joltage
		}
		fmt.Println("Part 1 Sum Joltage : ", sumJolatage)
		return sumJolatage
	}
	// Part 2
	if part2 {
		var sumJolatage int64
		for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
			joltage := findTwelveLargestDigits(line)
			// fmt.Println("Part 2 Line : ", line, "  ->  ", joltage)
			sumJolatage += joltage
		}
		fmt.Println("Part 2 Sum Joltage : ", sumJolatage)
		return sumJolatage

	}
	// solve part 1 here
	return 42
}

func findTwoLargestDigits(input string) int {
	line := []rune(input)
	first, second := 0, 0
	pos := 0
	// find the biggest digit between the beginning and one positin before the end of the line
	for i, r := range line[:len(line)-1] {
		digit := int(r - '0')
		if digit > first {
			first = digit
			// remember the position of teh first digit
			pos = i
			// fmt.Println("first : ", first, " i ", i)
		}
	}
	// find the second digit between the remembered postiion and the end
	for _, r := range line[pos+1:] {
		digit := int(r - '0')
		if digit > second {
			second = digit
			// fmt.Println("second : ", second, " i ", i)
		}
	}
	return first*10 + second
}

func findTwelveLargestDigits(input string) int64 {
	line := []rune(input)
	pos := 0
	needed := 12
	// map for result key is the position in result (numbers from 12 to 1)
	joltMap := make(map[int]int)
	// for each needed digit
	for needed > 0 {
		stelle := 0
		stellePos := 0
		// begin from pos, which is 0 for the first digit and the index+1 for each consecutive found digit
		// end has to be 'needed' positions from the end
		// end has to be +1 since needed is always greater than 0
		for i, r := range line[pos : len(line)-needed+1] {
			digit := int(r - '0')
			if digit > stelle {
				stelle = digit
				stellePos = pos + i
			}
			// fmt.Println(" stelle/digit ", stelle, " stellepos ", stellePos)
		}
		// we found the biggest digit in the range, save it in the map
		joltMap[needed] = stelle
		// for the next iteration, set the pos for teh new start-index
		pos = stellePos + 1
		// we found a number, continue until needed is 0
		needed--
	}

	var out int64
	// convert the found digits in the map into a big integer
	for i := 12; i >= 1; i-- {
		out = out*10 + int64(joltMap[i])
	}
	return out
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
