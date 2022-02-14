package main

import (
	"advent/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Point struct {
	X, Y, Z int64
}

type PointsSet struct {
	Points    []Point
	Distances []int64
}

type Pair struct {
	First, Second int
}

type Data struct {
	Sets  []PointsSet
	Pairs []Pair
}

func main() {
	start := time.Now()

	data := LoadInput()
	// fmt.Println(data)

	// calcPairs(&data)
	// for i := 0; i < len(data.Pairs); i++ {
	// 	fmt.Println("Check pair:", i)
	// 	unitePair(&data, i)
	// }

	scanners := []Point{{X: 0, Y: 0, Z: 0}}
	result := data.Sets[0].Points
	used := map[int]bool{0: true}
	// fmt.Println("start", len(result), result)

	for {
		merged := false

		for i := 1; i < len(data.Sets); i++ {
			if !used[i] {
				has12common, points, tempScanners := unitePair(result, data.Sets[i].Points, scanners)
				if has12common {
					merged = true
					result = points
					used[i] = true
					scanners = tempScanners
					// fmt.Println("has12common", len(result), result)
				}
			}
		}

		if !merged {
			break
		}
	}

	max := findMaxManhattan(scanners)

	fmt.Println("scanners = ", len(scanners), max)
	fmt.Println("result = ", len(result))

	// fmt.Println("Check pair:", 0, 1)
	// data.Pairs = []Pair{{First: 0, Second: 1}}
	// unitePair(&data, 0)

	elapsed := time.Since(start)
	fmt.Println("done", elapsed)
}

func findMaxManhattan(points []Point) int {
	result := -1

	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			a1 := points[i]
			a2 := points[j]
			d := utils.Abs(a1.X-a2.X) + utils.Abs(a1.Y-a2.Y) + utils.Abs(a1.Z-a2.Z)

			if d > int64(result) {
				result = int(d)
			}
		}
	}

	return result
}

func unitePair(first, second []Point, scanners []Point) (bool, []Point, []Point) {

	variants1 := make([][]Point, 0, 100)
	scannersVars := make([][]Point, 0, 100)

	for i1 := 0; i1 < len(first); i1++ {
		firstRel := changeCenter(first, first[i1])
		variants1 = append(variants1, firstRel)

		scannersRel := changeCenter(scanners, first[i1])
		scannersVars = append(scannersVars, scannersRel)
	}

	variants2 := make([][]Point, 0, 100)
	centers := []Point{{X: 0, Y: 0, Z: 0}}
	centersVars := make([][]Point, 0, 100)

	for v2 := 0; v2 < 24; v2++ {
		second24 := convertView(second, v2)
		for i2 := 0; i2 < len(second); i2++ {
			secondRel := changeCenter(second24, second24[i2])
			variants2 = append(variants2, secondRel)

			centersRels := changeCenter(centers, second24[i2])
			centersVars = append(centersVars, centersRels)
		}
	}

	has12, points, v1i, v2i := has12common(variants1, variants2)

	if has12 {
		tempScanners := scannersVars[v1i]
		tempScanners = append(tempScanners, centersVars[v2i][0])
		return true, points, tempScanners
	}

	return false, []Point{}, []Point{}
	// fmt.Println("Variants", len(variants1), len(variants2))
}

func has12common(variants1 [][]Point, variants2 [][]Point) (bool, []Point, int, int) {
	for i := 0; i < len(variants1); i++ {
		for j := 0; j < len(variants2); j++ {
			cnt := countCommon(variants1[i], variants2[j])
			if cnt >= 12 {
				union := uniteVariants(variants1[i], variants2[j])
				return true, union, i, j
			}
		}
	}
	return false, []Point{}, -1, -1
}

func uniteVariants(v1 []Point, v2 []Point) []Point {
	result := make([]Point, 0, 100)

	for i := 0; i < len(v1); i++ {
		result = append(result, v1[i])
	}

	for i := 0; i < len(v2); i++ {
		found := false

		for j := 0; j < len(result); j++ {
			if v2[i].X == result[j].X && v2[i].Y == result[j].Y && v2[i].Z == result[j].Z {
				found = true
			}
		}

		if !found {
			result = append(result, v2[i])
		}
	}
	return result
}

func countCommon(v1 []Point, v2 []Point) int {
	result := 0

	for i := 0; i < len(v1); i++ {
		for j := 0; j < len(v2); j++ {
			if v1[i].X == v2[j].X && v1[i].Y == v2[j].Y && v1[i].Z == v2[j].Z {
				result++
			}
		}
	}

	return result
}

func changeCenter(points []Point, c Point) []Point {
	result := make([]Point, len(points))
	for i := 0; i < len(points); i++ {
		result[i] = Point{
			X: points[i].X - c.X,
			Y: points[i].Y - c.Y,
			Z: points[i].Z - c.Z,
		}
	}
	return result
}

func convertView(points []Point, view int) []Point {
	result := make([]Point, len(points))
	for i := 0; i < len(points); i++ {
		result[i] = convertPoint(points[i], view)
	}
	return result
}

func convertPoint(p Point, view int) Point {
	switch view {
	case 1:
		return Point{X: -p.Y, Y: p.X, Z: p.Z} // [-y x z]
	case 2:
		return Point{X: -p.X, Y: -p.Y, Z: p.Z} // [-x -y z]
	case 3:
		return Point{X: p.Y, Y: -p.X, Z: p.Z} // [y -x z]
	case 4:
		return Point{X: p.Z, Y: p.Y, Z: -p.X} // [z y -x]
	case 5:
		return Point{X: -p.Y, Y: p.Z, Z: -p.X} // [-y z -x]
	case 6:
		return Point{X: -p.Z, Y: -p.Y, Z: -p.X} // [-z -y -x]
	case 7:
		return Point{X: p.Y, Y: -p.Z, Z: -p.X} // [y -z -x]
	case 8:
		return Point{X: -p.X, Y: p.Y, Z: -p.Z} // [-x y -z]
	case 9:
		return Point{X: -p.Y, Y: -p.X, Z: -p.Z} // [-y -x -z]
	case 10:
		return Point{X: p.X, Y: -p.Y, Z: -p.Z} // [x -y -z]
	case 11:
		return Point{X: p.Y, Y: p.X, Z: -p.Z} // [y x -z]
	case 12:
		return Point{X: -p.Z, Y: p.Y, Z: p.X} // [-z y x]
	case 13:
		return Point{X: -p.Y, Y: -p.Z, Z: p.X} // [-y -z x]
	case 14:
		return Point{X: p.Z, Y: -p.Y, Z: p.X} // [z -y x]
	case 15:
		return Point{X: p.Y, Y: p.Z, Z: p.X} // [y z x]
	case 16:
		return Point{X: p.X, Y: p.Z, Z: -p.Y} // [x z -y]
	case 17:
		return Point{X: -p.Z, Y: p.X, Z: -p.Y} // [-z x -y]
	case 18:
		return Point{X: -p.X, Y: -p.Z, Z: -p.Y} // [-x -z -y]
	case 19:
		return Point{X: p.Z, Y: -p.X, Z: -p.Y} // [z -x -y]
	case 20:
		return Point{X: p.X, Y: -p.Z, Z: p.Y} // [x -z y]
	case 21:
		return Point{X: p.Z, Y: p.X, Z: p.Y} // [z x y]
	case 22:
		return Point{X: -p.X, Y: p.Z, Z: p.Y} // [-x z y]
	case 23:
		return Point{X: -p.Z, Y: -p.X, Z: p.Y} // [-z -x y]]
	default:
		return Point{X: p.X, Y: p.Y, Z: p.Z} //[x y z]
	}
}

func generate4(x, y, z int) [][]int {
	result := make([][]int, 4)
	result[0] = []int{x, y, z}
	result[1] = []int{-y, x, z}
	result[2] = []int{-x, -y, z}
	result[3] = []int{y, -x, z}
	return result
}

func generate24(x, y, z int) [][]int {
	result := make([][]int, 0)
	forward := generate4(x, y, z)
	right := generate4(z, y, -x)
	back := generate4(-x, y, -z)
	left := generate4(-z, y, x)
	up := generate4(x, z, -y)
	down := generate4(x, -z, y)

	result = append(result, forward...)
	result = append(result, right...)
	result = append(result, back...)
	result = append(result, left...)
	result = append(result, up...)
	result = append(result, down...)

	return result
}

func calcPairs(data *Data) {
	for i := 0; i < len(data.Sets); i++ {
		calcDistances(&data.Sets[i])
		// fmt.Println(i, len(data.Sets[i].Points), len(data.Sets[i].Distances))
	}

	for i := 0; i < len(data.Sets)-1; i++ {
		for j := i + 1; j < len(data.Sets); j++ {
			common := commonInt64(data.Sets[i].Distances, data.Sets[j].Distances)
			if len(common) >= 12 {
				data.Pairs = append(data.Pairs, Pair{First: i, Second: j})
			}
		}
	}

	// fmt.Println("Candidates:", len(data.Pairs))
	// fmt.Println(data.Pairs)
}

func calcDistances(set *PointsSet) {
	for i := 0; i < len(set.Points)-1; i++ {
		for j := i + 1; j < len(set.Points); j++ {
			a1 := set.Points[i]
			a2 := set.Points[j]
			d := (a1.X-a2.X)*(a1.X-a2.X) + (a1.Y-a2.Y)*(a1.Y-a2.Y) + (a1.Z-a2.Z)*(a1.Z-a2.Z)
			set.Distances = append(set.Distances, d)
		}
	}
	sort.Slice(set.Distances, func(i, j int) bool { return set.Distances[i] < set.Distances[j] })
}

func commonInt64(a []int64, b []int64) []int64 {
	result := make([]int64, 0)

	indexA := 0
	indexB := 0

	for indexA < len(a) && indexB < len(b) {
		if a[indexA] == b[indexB] {
			result = append(result, a[indexA])
			indexA++
			indexB++
		} else if a[indexA] < b[indexB] {
			indexA++
		} else {
			indexB++
		}
	}

	return result

}

func LoadInput() Data {
	file, _ := os.Open("input.txt")

	result := Data{}
	result.Sets = make([]PointsSet, 0)

	var temp PointsSet

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		if strings.Index(text, "---") > -1 {
			temp = PointsSet{}
			temp.Points = make([]Point, 0)
			temp.Distances = make([]int64, 0)
		} else if text == "" {
			result.Sets = append(result.Sets, temp)
		} else {
			nums := utils.SplitToInt64(text, ",")
			temp.Points = append(temp.Points, Point{X: nums[0], Y: nums[1], Z: nums[2]})
		}
	}

	return result
}
