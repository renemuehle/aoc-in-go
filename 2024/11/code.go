package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache map[string]int

func main() {
	//aoc.Harness(run)

	// Somehow Harness fails for day 11 :/
	// manually load the file
	//file, _ := FileAsString("input-example.txt")
	file, _ := FileAsString("input-user.txt")
	run(true, file)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block

	// initialize the cache
	cache = make(map[string]int)

	// get the input
	nTmp := strings.Split(input, " ")
	numbers := make([]int, len(nTmp))
	for i, n1 := range nTmp {
		n2, _ := strconv.Atoi(n1)
		numbers[i] = n2
	}

	fmt.Println("Numbers ", numbers)

	if part2 {

		anzStone := 0
		// number of blinks
		b := 75
		for _, n := range numbers {
			anzStone += blink(n, b)
			fmt.Println(" N ", n, " Blink ", b, " Anz ", anzStone)
		}
		fmt.Println("Anz", anzStone)
		return anzStone
	}
	// solve part 1 here

	anzStone := 0
	// number of blinks
	b := 25
	for _, n := range numbers {
		anzStone += blink(n, b)
		fmt.Println(" N ", n, " Blink ", b, " Anz ", anzStone)
		//fmt.Println("Cache ", cache)
	}
	fmt.Println("Anz", anzStone)
	return anzStone
}

func blink(number int, b int) int {
	ret := 0
	// check the cache, if we already have calculated this combination of number and blinks
	// thats a HUGE performance gain
	val, ok := cache[fmt.Sprintf("%d:%d", number, b)]
	// If the key exists
	if ok {
		// no need to calculate it again, just return the result
		ret = val
	} else if b > 1 {
		// more Blinks, we havte to go down recursively
		// apply the rules to the number
		r := rules(number)
		//fmt.Println(" Rule ( ", number, " ) =", r)
		// and than, calculate the blink for each number (could be 1 or 2)
		for _, i := range r {
			r1 := blink(i, b-1)
			// cache the result of this calculation
			cache[fmt.Sprintf("%d:%d", i, (b-1))] = r1
			ret += r1
			//fmt.Println(" Blink (", i, " , ", (b - 1), " ) = ", ret)
		}
	} else {
		// Last Blink, return Number of Stones, according to the rules
		// could be 1 or 2
		ret = len(rules(number))
	}
	return ret
}

func rules(number int) []int {
	ret := make([]int, 0)
	// Rule 1 : 0 => 1
	if number == 0 {
		return append(ret, 1)
	}
	// Rule 2 : even number of digits, split number in middle
	numberS := strconv.Itoa(number)
	laenge := len(numberS)
	if laenge%2 == 0 {
		n1, _ := strconv.Atoi(numberS[0:(laenge / 2)])
		n2, _ := strconv.Atoi(numberS[(laenge / 2):])
		ret = append(ret, n1)
		ret = append(ret, n2)
		return ret
	}
	// Rule 3 : everything else, multiply with 1024
	return append(ret, (number * 2024))
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
