package main

import (
	"fmt"
	"time"
)

var X1 = int64(287)
var X2 = int64(309)
var Y1 = int64(-76)
var Y2 = int64(-48)

func main() {
	start := time.Now()

	// fmt.Print(test, " ")
	// fmt.Print("F")
	// fmt.Println()
	xx := possibleXX()

	fmt.Println("-= X =-")
	fmt.Println(len(xx))
	fmt.Println(xx)

	yy := possibleYY()

	fmt.Println("-= Y =-")
	fmt.Println(len(yy))
	fmt.Println(yy)

	total := 0
	for i := 0; i < len(xx); i++ {
		for j := 0; j < len(yy); j++ {
			if checkSpeed(xx[i], yy[j]) {
				fmt.Println(xx[i], yy[j])
				total++
			}
		}
	}
	fmt.Println("total", total)

	elapsed := time.Since(start)
	fmt.Println("done", elapsed)
}

func checkSpeed(x, y int64) bool {
	testX := int64(0)
	testY := int64(0)
	step := int64(0)
	found := false
	for !found && testY >= Y1 {
		if x-step > 0 {
			testX = testX + (x - step)
		}
		testY = testY + (y - step)
		if X1 <= testX && testX <= X2 && Y1 <= testY && testY <= Y2 {
			return true
		}
		step++
	}
	return false
}

func possibleYY() []int64 {
	yy := make([]int64, 0)
	for y := int64(-76); y <= int64(75); y++ {
		test := int64(0)
		step := int64(0)
		found := false
		for !found && test >= Y1 {
			test = test + (y - step)

			if Y1 <= test && test <= Y2 {
				yy = append(yy, y)
				found = true

			}
			step++
		}

	}
	return yy
}

func possibleXX() []int64 {
	xx := make([]int64, 0)
	for x := int64(20); x < int64(320); x++ {
		test := int64(0)
		step := int64(0)
		found := false
		for !found && test >= 0 && test <= int64(309) {
			test = test + (x - step)

			if X1 <= test && test <= X2 {
				xx = append(xx, x)
				found = true

			}
			step++
		}

	}
	return xx
}
