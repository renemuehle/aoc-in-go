package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

var plants map[string]int = make(map[string]int)
var fences map[string]int = make(map[string]int)

// var fences2 map[string][]string = make(map[string][]string)
var fences2H map[string][][]int = make(map[string][][]int)
var fences2V map[string][][]int = make(map[string][][]int)

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

	zeilen := strings.Split(input, "\n")
	anzZeilen := len(zeilen)
	anzSpalten := len(zeilen[0])
	// m for the map of plants
	m := make([][]string, anzZeilen)
	// v for tracking the visits
	v := make([][]int, anzZeilen)

	// initialize both slices
	for i := 0; i < anzZeilen; i++ {
		sm := make([]string, anzSpalten)
		sv := make([]int, anzSpalten)
		for j, s := range zeilen[i] {
			sm[j] = fmt.Sprintf("%c", s)
			sv[j] = 0
		}
		m[i] = sm
		v[i] = sv
	}

	//PrintArray(m)

	key := 1
	for i, zeile := range m {
		for j := range zeile {
			// go trough the map and if there is es field we don't have visited jet
			// visit that field
			if v[i][j] == 0 {
				visitNeighbour(m, v, key, i, j)
				key++

				PrintArray(m)
				PrintArray(v)
				fmt.Println("Key ", key)
				// fmt.Println("Plants ", plants)
				// fmt.Println("Fences ", fences)
				fmt.Println("Fences2H ", fences2H)
				fmt.Println("Fences2V ", fences2V)
			}
		}
	}

	if part2 {
		// when you're ready to do part 2, remove this "not implemented" block
		sum := 0
		for k, f1 := range fences2H {
			// sort the fences horizontal
			sort.SliceStable(f1, func(i, j int) bool {
				if f1[i][0] == f1[j][0] {
					return f1[i][1] < f1[j][1]
				}
				return f1[i][0] < f1[j][0]
			})
			fmt.Println("F1 (horizontal) von ", k, " = ", f1)
			f1 = trimLinesH(f1)
			fmt.Println("F1 (horizontal) von ", k, " = ", f1)

			fences2H[k] = f1

		}

		for k, f1 := range fences2V {

			// sort the fences vertical
			sort.SliceStable(f1, func(i, j int) bool {
				if f1[i][1] == f1[j][1] {
					return f1[i][0] < f1[j][0]
				}
				return f1[i][1] < f1[j][1]
			})
			fmt.Println("F1 (vertical) von ", k, " = ", f1)
			f1 = trimLinesV(f1)
			fmt.Println("F1 (vertical) von ", k, " = ", f1)
			fences2V[k] = f1

		}

		for k := range plants {
			s := (plants[k] * (len(fences2H[k]) + len(fences2V[k])))
			fmt.Println("Sum ", k, " :: ", s, " = ", plants[k], " * ", len(fences2H[k]), "+", len(fences2V[k]))
			sum += s
		}
		return sum
	}

	// solve part 1 here

	//PrintArray(m)
	//fmt.Println(plants)
	//fmt.Println(fences)
	sum := 0
	for k := range plants {
		sum += (plants[k] * fences[k])
	}

	return sum
}

func trimLinesH(f1 [][]int) [][]int {
	laenge := len(f1)

	for x := laenge - 1; x > 0; x-- {
		//fmt.Println("Test ", x, " # ", f1[x][0], "=", f1[x-1][0], " # ", f1[x][1], "-", f1[x-1][1])
		if f1[x][0] == f1[x-1][0] && f1[x][1]-f1[x-1][1] <= 1 {
			if x == laenge-1 {
				// cut last element
				f1 = f1[:x]
				// fmt.Println("Cut Last Elem ", x, " = ", f1)
			} else {
				// cut element in the middle
				f1 = append(f1[:x], f1[x+1:]...)
				// fmt.Println("Cut Middle Elem ", x, " = ", f1)
				x++
			}
			laenge--
		}
	}
	return f1
}

func trimLinesV(f1 [][]int) [][]int {
	laenge := len(f1)

	for x := laenge - 1; x > 0; x-- {
		//fmt.Println("Test ", x, " # ", f1[x][1], "=", f1[x-1][1], " # ", f1[x][0], "-", f1[x-1][0])
		if f1[x][1] == f1[x-1][1] && f1[x][0]-f1[x-1][0] <= 1 {
			if x == laenge-1 {
				// cut last element
				f1 = f1[:x]
				// fmt.Println("Cut Last Elem ", x, " = ", f1)
			} else {
				// cut element in the middle
				f1 = append(f1[:x], f1[x+1:]...)
				// fmt.Println("Cut Middle Elem ", x, " = ", f1)
				x++
			}
			laenge--
		}
	}
	return f1
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

/* print a slice */
func PrintArray[T any](input [][]T) {
	fmt.Println("=======================================")
	for _, row := range input {
		for _, colValue := range row {
			fmt.Print(colValue)
		}
		fmt.Println()
	}
}

func visitNeighbour(m [][]string, v [][]int, key int, i, j int) {

	if v[i][j] == 0 {
		// the plant of the current field for comparison
		plant := m[i][j]
		// track the visit
		v[i][j] += 1

		// create a uniqe key, kombine the letter of the plant with a unique number
		pk := fmt.Sprintf("%s%d", plant, key)
		plants[pk] += 1

		// if above grows the same plant, visit recursivly
		if i > 0 && m[i-1][j] == plant {
			visitNeighbour(m, v, key, (i - 1), j)

		} else {
			// different neighbour means fence
			fmt.Println(" new Fence H between ", i, ":", j, " and ", (i - 1), ":", j)
			fences[pk] += 1
			fences2H[pk] = append(fences2H[pk], []int{(i - 1), j})
		}

		// if right grows the same plant, visit recursivly
		if j < (len(m[i])-1) && m[i][j+1] == plant {
			visitNeighbour(m, v, key, i, (j + 1))

		} else {
			// different neighbour means fence
			fmt.Println(" new Fence V between ", i, ":", j, " and ", i, ":", (j + 1))
			fences[pk] += 1
			fences2V[pk] = append(fences2V[pk], []int{i, (j + 1)})
		}

		// if below grows the same plant, visit recursivly
		if i < (len(m)-1) && m[i+1][j] == plant {
			visitNeighbour(m, v, key, (i + 1), j)

		} else {
			// different neighbour means fence
			fmt.Println(" new Fence H between ", i, ":", j, " and ", (i + 1), ":", j)
			fences[pk] += 1
			fences2H[pk] = append(fences2H[pk], []int{(i + 1), j})
		}

		// if left grows the same plant, visit recursivly
		if j > 0 && m[i][j-1] == plant {
			visitNeighbour(m, v, key, i, (j - 1))
		} else {
			// different neighbour means fence
			fmt.Println(" new Fence V between ", i, ":", j, " and ", i, ":", (j - 11))
			fences[pk] += 1
			fences2V[pk] = append(fences2V[pk], []int{i, (j - 1)})
		}
	}
}
