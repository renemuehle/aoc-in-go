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

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {

	var firstNum []int
	var secondNum []int
	var sum int

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		// fmt.Println(line)

		// Split line into Slice of "Words", splits by "consecutive white space characters"
		l := strings.Fields(line)

		// Parse 1. String to int
		i1, err := strconv.Atoi(l[0])
		if err != nil {
			// ... handle error
			panic(err)
		}
		firstNum = append(firstNum, i1)

		// Parse 2. String to int
		i2, err := strconv.Atoi(l[1])
		if err != nil {
			// ... handle error
			panic(err)
		}
		secondNum = append(secondNum, i2)
	}

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {

		// solve part 2 here
		sum = 0

		// create a map/dictionary with each number and the count of that number
		var countNum = countNumbers(secondNum)
		//fmt.Println("countNum:", countNum)

		// calculate "similarity score"
		for i := 0; i < len(firstNum); i++ {
			sum += firstNum[i] * countNum[firstNum[i]]
		}
		return sum
	}

	// solve part 1 here
	sum = 0

	sort.Ints(firstNum)
	//fmt.Println("FirstNum:", firstNum)

	sort.Ints(secondNum)
	//fmt.Println("SecondNum:", secondNum)

	for i := 0; i < len(firstNum); i++ {
		sum += diff(firstNum[i], secondNum[i])
	}

	// fmt.Println("Summe:", sum)
	return sum
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func countNumbers(input []int) map[int]int {
	//Create a   dictionary of values for each element
	countNum := make(map[int]int)
	for _, num := range input {
		countNum[num] = countNum[num] + +1
	}
	return countNum
}
