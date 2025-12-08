package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Point struct {
	X, Y, Z int
}

// Implementiert fmt.Stringer
func (p Point) String() string {
	return fmt.Sprintf("(%.d, %.d, %.d)", p.X, p.Y, p.Z)
}

type Distance struct {
	point1, point2 Point
	distance       int
}

func run(part2 bool, input string) any {

	lines := strings.Split(strings.TrimSpace(input), "\n")

	points := make([]Point, 0, len(lines))

	for _, line := range lines {
		coordsA := strings.Split(strings.TrimSpace(line), ",")
		coordsI := make([]int, 3)
		for j, coord := range coordsA {
			p := strings.TrimSpace(coord)
			i, _ := strconv.Atoi(p)
			coordsI[j] = i
		}
		p := Point{coordsI[0], coordsI[1], coordsI[2]}
		points = append(points, p)
	}

	// Part 1
	if !part2 {

		distances := CalcNeighbors(points)
		sortDistBySize(distances)

		n := len(distances)
		maxConnections := n
		if n < 1000 {
			// test
			maxConnections = 10
		} else {
			// full task
			maxConnections = 1000
		}

		distances = distances[:maxConnections]

		//for _, d := range distances {
		//	fmt.Println("Part 1 - sorted - Point 1 ", d.point1, " Point 2 ", d.point2, " Distance ", d.distance)
		//}

		circuits := findCircuits(distances)
		sortCircuitsBySize(circuits)
		res := 1
		for i := 0; i < len(circuits[:3]); i++ {
			fmt.Println(" Circuit - Length = ", len(circuits[i]))
			// fmt.Println(circuits[i])
			res *= len(circuits[i])
		}
		return res
	}
	// Part 2
	if part2 {
		distances := CalcNeighbors(points)
		sortDistBySize(distances)
		//for _, d := range distances {
		//	fmt.Println("Part 2 - sorted -  Point 1 ", d.point1, " Point 2 ", d.point2, " Distance ", d.distance)
		//}
		res := findCircuits2(distances, len(points))
		return res
	}
	// solve part 1 here
	return 42
}

// calculate distance between two points
// sqrt is not neccessary for comparing two distanzes
func distance(a, b Point) int {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	dist := dx*dx + dy*dy + dz*dz
	// dist = math.Sqrt(dx*dx + dy*dy + dz*dz)
	return dist
}

// calculate for each point the distance to each other point
// returns a list aof each point with its nearest neighbour and the shortest distance
func CalcNeighbors(points []Point) []Distance {
	numPoints := len(points)
	distances := make([]Distance, 0)

	// i for outer loop of all points
	// j for inner loop of all points
	for i := range numPoints {
		for j := range numPoints {
			if i >= j {
				continue
			}
			dist := distance(points[i], points[j])
			distances = append(distances, Distance{points[i], points[j], dist})
		}
	}
	return distances
}

// distances contains all pathes between points
// find all connected points
func findCircuits(distances []Distance) []map[Point]struct{} {

	circuits := make([]map[Point]struct{}, 0, len(distances))

	for _, dist := range distances {
		// temp Set, to merge all circuits that contain point1 or point2
		temp := make(map[Point]struct{})

		// current points are always included
		temp[dist.point1] = struct{}{}
		temp[dist.point2] = struct{}{}

		// filter slice to keep all circuits that did not match;
		keep := make([]map[Point]struct{}, 0, len(circuits))

		// loop trough all circuits
		for _, setOfPoints := range circuits {
			// if point1 or point 2 is part of the circiut
			_, has1 := setOfPoints[dist.point1]
			_, has2 := setOfPoints[dist.point2]
			if has1 || has2 {
				// merge this circuit into temp
				for p := range setOfPoints {
					temp[p] = struct{}{}
				}
			} else {
				// keep circuit as-is
				keep = append(keep, setOfPoints)
			}

		}
		//add temp to circuits
		circuits = append(keep, temp)
	}
	return circuits
}

// distances contains all pathes between points
// connect all points
// return the product of the x-coordinates of the last connection
func findCircuits2(distances []Distance, anzPoints int) int {

	circuits := make([]map[Point]struct{}, 0, len(distances))
	// going through all distances / list is much longer than in part 1
	for _, dist := range distances {

		// temp Set, to merge all circuits that contain point1 or point2
		temp := make(map[Point]struct{})

		// current points are always included
		temp[dist.point1] = struct{}{}
		temp[dist.point2] = struct{}{}

		// filter slice to keep all circuits that did not match;
		keep := make([]map[Point]struct{}, 0, len(circuits))

		// loop trough all circuits
		for _, setOfPoints := range circuits {
			// if point1 or point 2 is part of the circiut
			_, has1 := setOfPoints[dist.point1]
			_, has2 := setOfPoints[dist.point2]
			if has1 || has2 {
				// merge this circuit into temp
				for p := range setOfPoints {
					temp[p] = struct{}{}
				}
			} else {
				// keep circuit as-is
				keep = append(keep, setOfPoints)
			}
		}
		//add temp to circuits
		circuits = append(keep, temp)

		// as soon as wie have on circuit which has all points, we can stop
		if len(circuits) == 1 && len(circuits[0]) == anzPoints {
			fmt.Println(" Point1 ", dist.point1, " Point2 ", dist.point2)
			// return the product if the last two x-coordinates
			return dist.point1.X * dist.point2.X
		}
	}

	return -1
}

// helper to sort distances
func sortDistBySize(distances []Distance) {
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})
}

// helper to sort circuits
func sortCircuitsBySize(circuits []map[Point]struct{}) {
	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})
}
