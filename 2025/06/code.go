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

func run(part2 bool, input string) any {

	lines := strings.Split(strings.TrimSpace(input), "\n")

	var homework [][]string

	// Part 1
	if !part2 {
		// read the numbers and the operators
		for _, line := range lines {
			numbers := strings.Fields(line)
			homework = append(homework, numbers)
		}
		totalSum := 0
		// for each column
		for i := 0; i < len(homework[0]); i++ {
			// operator from the last element of the column
			op := homework[len(homework)-1][i]
			colValue := 0
			// for each line of the column
			for j := 0; j < len(homework)-1; j++ {
				// get the integer value
				v, _ := strconv.Atoi(homework[j][i])
				// first element, just set it
				if j == 0 {
					colValue = v
					continue
				}
				// depending on the operator
				// calculate value line by line
				switch op {
				case "+":
					colValue = colValue + v
				case "*":
					colValue = colValue * v
				default:
					fmt.Println("Error ", op, v, i, j)
				}
			}
			// add column value to the total
			totalSum += colValue
		}
		fmt.Println(" Total ", totalSum)
		return totalSum
	}
	// Part 2
	if part2 {
		// line with operators
		lineOp := len(lines) - 1
		// set variable for sum
		totalSum := 0
		// colum for last operator ... needed to recocnize a new operator
		colOp := 0
		// numeric value of column / block
		blockNum := 0
		// current operator
		op := string(' ')

		// for each column of input + extra columns needed, because the lines have different length
		// FIXME: better find the longest line and use that value
		for i := 0; i < len(lines[lineOp])+3; i++ {

			// string for concatenating the numbers in that column
			colValue := string(' ')

			// try to get the character in the operator line
			ch := string(' ')
			// prevent index-error
			if i < len(lines[lineOp]) {
				ch = string(lines[lineOp][i])
			}

			// if the operator is not a space
			if ch != string(' ') {
				// new op = new block
				op = string(ch)
				// remember position of new operator / new block
				// needed to reset the blockNum
				colOp = i

			}
			// current column i all rows except the last on (with the operator)
			for r := 0; r < len(lines)-1; r++ {
				// lines have differet length, prevent error
				if i < len(lines[r]) {
					// append the sting-values (numbers or spaces)
					colValue = colValue + string(lines[r][i])
				}
			}
			// remove spaces
			colValue = strings.TrimSpace(colValue)

			// if the colum has content
			if len(colValue) > 0 {
				// numeric value of cloumn
				v, _ := strconv.Atoi(strings.TrimSpace(colValue))
				// if its the beginning of a block
				if i == colOp {
					// no calculation needed, just get the numeric value and start the new block with that number
					blockNum = v
					continue
				}
				// no new block, switch between operation and continue calculation
				switch op {
				case "+":
					blockNum = blockNum + v
				case "*":
					blockNum = blockNum * v
				}
				//fmt.Println(" OP ", op, " i ", i, " colValue ", colValue, " blockNum ", blockNum, " Gesamt ", gesamtSumme)
			} else {
				// column has no content - end of a block
				// fmt.Println("=== End of Block / Add sum ===")
				// fmt.Print(" Old ", totalSum, " BlockNum ", blockNum)
				// add the value from the block (blocknum) to the total value
				totalSum = totalSum + blockNum
				//fmt.Println(" New ", totalSum)
				blockNum = 0

			}
			// End of current column i
		}
		fmt.Println(" Total ", totalSum)
		return totalSum
	}
	// solve part 1 here
	return 42
}
