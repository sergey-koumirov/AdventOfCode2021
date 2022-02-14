package main

import (
	"day04/utils"
	"fmt"
)

func main() {
	data := utils.LoadInput()

	intersections := make([]utils.Point, 0, 100)

	for i := 0; i < len(data.Lines)-2; i++ {
		for j := i + 1; j < len(data.Lines); j++ {
			points := FindIntersections(data.Lines[i], data.Lines[j])

			// if len(points) == 1 {
			// 	fmt.Println("-------------------")
			// 	fmt.Println(data.Lines[i])
			// 	fmt.Println(data.Lines[j])
			// 	fmt.Println(points)
			// }

			intersections = AddToIntersections(points, intersections)
		}
	}

	// fmt.Println(data)
	// fmt.Println(intersections)
	fmt.Println(len(intersections))
}

func FindIntersections(A utils.Line, B utils.Line) []utils.Point {
	result := make([]utils.Point, 0, 100)

	// if A.X1 == A.X2 && B.X1 == B.X2 && A.X1 == B.X1 { // both ||
	// 	ay1, ay2 := sort12(A.Y1, A.Y2)
	// 	by1, by2 := sort12(B.Y1, B.Y2)
	// 	y1, y2 := findRange(ay1, ay2, by1, by2)

	// 	if y1 > -1 && y2 > -1 {
	// 		for y := y1; y <= y2; y++ {
	// 			result = append(result, utils.Point{X: A.X1, Y: y, N: 2})
	// 		}
	// 	}

	// } else if A.Y1 == A.Y2 && B.Y1 == B.Y2 && A.Y1 == B.Y1 { // both =
	// 	ax1, ax2 := sort12(A.X1, A.X2)
	// 	bx1, bx2 := sort12(B.X1, B.X2)
	// 	x1, x2 := findRange(ax1, ax2, bx1, bx2)

	// 	if x1 > -1 && x2 > -1 {
	// 		for x := x1; x <= x2; x++ {
	// 			result = append(result, utils.Point{X: x, Y: A.Y1, N: 2})
	// 		}
	// 	}

	// } else if A.Y1 == A.Y2 && B.X1 == B.X2 { // cross 1 hor + ver
	// 	ax1, ax2 := sort12(A.X1, A.X2)
	// 	by1, by2 := sort12(B.Y1, B.Y2)

	// 	if ax1 <= B.X1 && B.X1 <= ax2 && by1 <= A.Y1 && A.Y1 <= by2 {
	// 		result = append(result, utils.Point{X: B.X1, Y: A.Y1, N: 2})
	// 	}
	// } else if A.X1 == A.X2 && B.Y1 == B.Y2 { // cross 1 ver + hor
	// 	ay1, ay2 := sort12(A.Y1, A.Y2)
	// 	bx1, bx2 := sort12(B.X1, B.X2)
	// 	if ay1 <= B.Y1 && B.Y1 <= ay2 && bx1 <= A.X1 && A.X1 <= bx2 {
	// 		result = append(result, utils.Point{X: A.X1, Y: B.Y1, N: 2})
	// 	}
	// } else {
	// ax1, ax2 := sort12(A.X1, A.X2)
	// ay1, ay2 := sort12(A.Y1, A.Y2)
	// bx1, bx2 := sort12(B.X1, B.X2)
	// by1, by2 := sort12(B.Y1, B.Y2)

	// a1InB := bx1 <= ax1 && ax1 <= bx2 && by1 <= ay1 && ay1 <= by2
	// a2InB := bx1 <= ax2 && ax2 <= bx2 && by1 <= ay2 && ay2 <= by2

	// b1InA := ax1 <= bx1 && bx1 <= ax2 && ay1 <= by1 && by1 <= ay2
	// b2InA := ax1 <= bx2 && bx2 <= ax2 && ay1 <= by2 && by2 <= ay2

	// if a1InB || a2InB || b1InA || b2InA {
	points1 := getPoints(A.X1, A.Y1, A.X2, A.Y2)
	points2 := getPoints(B.X1, B.Y1, B.X2, B.Y2)

	t := commonPoints(points1, points2)

	fmt.Println("==============================")
	fmt.Printf("%+v", A)
	fmt.Printf("%+v", B)
	fmt.Printf("%+v\n", t)

	for _, p := range t {
		result = append(result, p)
	}
	// }

	// }

	return result
}

func commonPoints(point1, point2 []utils.Point) []utils.Point {
	result := make([]utils.Point, 0, 100)

	for _, p1 := range point1 {
		for _, p2 := range point2 {
			if p1.X == p2.X && p1.Y == p2.Y {
				result = append(result, p1)
			}
		}
	}

	return result
}

func getPoints(x1, y1, x2, y2 int64) []utils.Point {
	result := make([]utils.Point, 0, 100)

	dx := sign(x2 - x1)
	dy := sign(y2 - y1)
	cnt := max(abs(x2-x1), abs(y2-y1)) + 1

	for i := int64(0); i < cnt; i++ {
		result = append(result, utils.Point{X: x1 + i*dx, Y: y1 + i*dy, N: 2})
	}

	return result
}

func max(v1 int64, v2 int64) int64 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func abs(v int64) int64 {
	if v > 0 {
		return v
	}
	return v * -1
}

func sign(v int64) int64 {
	if v > 0 {
		return int64(1)
	}
	if v < 0 {
		return int64(-1)
	}
	return 0
}

func sort12(v1, v2 int64) (int64, int64) {
	if v1 < v2 {
		return v1, v2
	}
	return v2, v1
}

func findRange(a1, a2, b1, b2 int64) (int64, int64) {
	v1 := int64(-1)
	v2 := int64(-1)

	if b1 <= a1 && a1 <= b2 {
		v1 = a1
	}
	if a1 <= b1 && b1 <= a2 {
		v1 = b1
	}

	if b1 <= a2 && a2 <= b2 {
		v2 = a2
	}
	if a1 <= b2 && b2 <= a2 {
		v2 = b2
	}

	return v1, v2
}

func AddToIntersections(points []utils.Point, temp []utils.Point) []utils.Point {

	for _, p := range points {
		found := false
		for j := 0; j < len(temp); j++ {
			if p.X == temp[j].X && p.Y == temp[j].Y {
				found = true
				temp[j].N = temp[j].N + 1
			}
		}
		if !found {
			temp = append(temp, p)
		}
	}

	return temp
}
