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

	lines := strings.Split(input, "\n")

	equations := make([][]int, len(lines))

	for i, m := range lines {

		n := strings.Split(m, " ")
		for _, o := range n {

			p := strings.Replace(o, ":", "", 1)
			q, _ := strconv.Atoi(p)
			equations[i] = append(equations[i], q)
		}
	}

	// fmt.Println("Equations ", equations)

	if part2 {
		fmt.Println("===================== PART2 ===================== ")
		sum := 0
		for _, e := range equations {
			sum += e[0] * calcTrueEquations2(e)
		}
		return sum
	}

	// solve part 1 here
	fmt.Println("===================== PART1 ===================== ")
	sum := 0
	for _, e := range equations {
		sum += e[0] * calcTrueEquations(e)
	}
	return sum
}

func calcTrueEquations(e []int) int {
	valid := 0
	result := e[0]
	fmt.Println("Calculation for  ", result, " = ", e[1:])

	calculations := make(map[int]string)
	calculations[e[1]+e[2]] = strconv.Itoa(e[1]) + "+" + strconv.Itoa(e[2])
	calculations[e[1]*e[2]] = strconv.Itoa(e[1]) + "*" + strconv.Itoa(e[2])
	for i := 3; i < len(e); i++ {

		calculations2 := make(map[int]string)
		for k, v := range calculations {
			if k+e[i] <= result {
				calculations2[k+e[i]] = v + "+" + strconv.Itoa(e[i])
			} else {
				//fmt.Println("Is too big - Not added: ", v+"+"+strconv.Itoa(e[i]), " = ", k+e[i])
			}
			if k*e[i] <= result {
				calculations2[k*e[i]] = v + "*" + strconv.Itoa(e[i])
			} else {
				//fmt.Println("Is too big - Not added: ", (v + "*" + strconv.Itoa(e[i])), " = ", k*e[i])
			}
		}
		//fmt.Println("Calculations2  ", calculations2)
		// keep the current iteration
		calculations = calculations2
	}
	// fmt.Println("Calculations  ", calculations)
	for k, v := range calculations {
		if k == result {
			valid++
			fmt.Println(" Valid ", k, "=", v)
		}
	}
	return valid
}

func calcTrueEquations2(e []int) int {
	valid := 0
	result := e[0]
	fmt.Println("Calculation for  ", result, " = ", e[1:])

	calculations := make(map[int]string)
	calculations[e[1]+e[2]] = strconv.Itoa(e[1]) + "+" + strconv.Itoa(e[2])
	calculations[e[1]*e[2]] = strconv.Itoa(e[1]) + "*" + strconv.Itoa(e[2])
	t0, _ := strconv.Atoi(strconv.Itoa(e[1]) + strconv.Itoa(e[2]))
	calculations[t0] = strconv.Itoa(e[1]) + strconv.Itoa(e[2])

	for i := 3; i < len(e); i++ {
		// new map, to calculate the current iteration
		calculations2 := make(map[int]string)
		// for each prevous calculation, calculate/add all 3 new variations
		for k, v := range calculations {
			if k+e[i] <= result {
				calculations2[k+e[i]] = v + "+" + strconv.Itoa(e[i])
			} else {
				//fmt.Println("Is too big - Not added: ", v+"+"+strconv.Itoa(e[i]), " = ", k+e[i])
			}
			if k*e[i] <= result {
				calculations2[k*e[i]] = v + "*" + strconv.Itoa(e[i])
			} else {
				//fmt.Println("Is too big - Not added: ", (v + "*" + strconv.Itoa(e[i])), " = ", k*e[i])
			}

			// t1 = value
			t1, _ := strconv.Atoi(strconv.Itoa(k) + strconv.Itoa(e[i]))
			//fmt.Println("Concatenation ", t1, "=", (v + "||" + strconv.Itoa(e[i])))
			if t1 <= result {
				calculations2[t1] = v + "||" + strconv.Itoa(e[i])
			} else {
				//fmt.Println("Is too big - Not added: ", t1, "=", (v + "||" + strconv.Itoa(e[i])))
			}
		}
		//fmt.Println("Calculations2  ", calculations2)
		// keep the current iteration
		calculations = calculations2
		//fmt.Println("")
	}
	//fmt.Println("Calculations  ", calculations)
	for k, v := range calculations {
		if k == result {
			valid++
			fmt.Println(" Valid ", k, "=", v)
		}
	}
	//fmt.Println("=================================================")
	return valid
}
