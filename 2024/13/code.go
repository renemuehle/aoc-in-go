package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type arcade struct {
	number         int
	buttonAX       int
	buttonAY       int
	buttonBX       int
	buttonBY       int
	prizeX, prizeY int
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block

	zeilen := strings.Split(input, "\n")
	anzZeilen := len(zeilen)

	//r1 := regexp.MustCompile(`Button\s(P<Button>\w):\sX\+(P<X>\d*),\sY\+(P<Y>\d*)`)
	r1 := regexp.MustCompile(`Button\s(?P<Button>\w):\sX\+(?P<X>\d*),\sY\+(?P<Y>\d*)`)
	r2 := regexp.MustCompile(`Prize:\sX=(?P<X>\d*),\sY=(?P<Y>\d*)`)

	for i := 0; i < anzZeilen; i = i + 4 {
		fmt.Println("Zeile: ", zeilen[i])
		res1 := r1.FindStringSubmatch(zeilen[i])
		fmt.Printf("%#v\n", res1)

		fmt.Println("Zeile: ", zeilen[i+1])
		res2 := r1.FindStringSubmatch(zeilen[i+1])
		fmt.Printf("%#v\n", res2)

		fmt.Println("Zeile: ", zeilen[i+2])
		res3 := r2.FindStringSubmatch(zeilen[i+2])
		fmt.Printf("%#v\n", res3)
		fmt.Printf("%#v\n", r2.FindAllStringSubmatch(zeilen[i+2], -1))

		arcade{i, res1[2], res1[3], res2[2], res2[3], res3[2], res3[3]}

		fmt.Println(" ", arcade)
	}

	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	return 42
}
