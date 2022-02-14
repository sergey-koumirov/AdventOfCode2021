package main

import (
	"advent/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	X   int64
	Y   int64
	Cnt int
}

type Points []Point

type Instruction struct {
	V   int64
	Dir string
}

type Instructions []Instruction

func main() {
	start := time.Now()

	data, instructions := LoadInput()
	// fmt.Println(data, instructions)

	data1 := data
	for _, ins := range instructions {
		data1 = fold(data1, ins)
		fmt.Println(data1)
	}

	print(data1)

	elapsed := time.Since(start)
	fmt.Println("time", elapsed)
}

func print(pp Points) {

	maxX := 0
	maxY := 0

	for _, p := range pp {
		if maxX < int(p.X) {
			maxX = int(p.X)
		}
		if maxY < int(p.Y) {
			maxY = int(p.Y)
		}
	}

	fmt.Println(maxX, maxY)

	field := make([][]string, maxX+1)
	for i := 0; i < maxX+1; i++ {
		field[i] = make([]string, maxY+1)
	}

	for _, p := range pp {
		field[p.X][p.Y] = "1"
	}

	for j := 0; j < maxY+1; j++ {
		for i := 0; i < maxX+1; i++ {
			if field[i][j] == "" {
				fmt.Print(" ")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
	}

}

func fold(pp Points, ins Instruction) Points {
	if ins.Dir == "x" {
		return foldX(pp, ins.V)
	}
	return foldY(pp, ins.V)
}

func foldX(pp Points, x int64) Points {
	result := make(Points, 0, len(pp))

	for _, p := range pp {
		if p.X < x {
			addPoint(p, &result)
		} else if p.X > x {
			p.X = 2*x - p.X
			addPoint(p, &result)
		}
	}

	return result
}

func foldY(pp Points, y int64) Points {
	result := make(Points, 0, len(pp))

	for _, p := range pp {
		if p.Y < y {
			addPoint(p, &result)
		} else if p.Y > y {
			p.Y = 2*y - p.Y
			addPoint(p, &result)
		}
	}

	return result
}

func addPoint(p Point, result *Points) {
	r := *result
	found := false
	for i := 0; i < len(r); i++ {
		if r[i].X == p.X && r[i].Y == p.Y {
			found = true
			r[i].Cnt++
		}
	}
	if !found {
		*result = append(r, p)
	}
}

func LoadInput() (Points, Instructions) {
	pp := make(Points, 0, 1000)
	ii := make(Instructions, 0, 10)

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		if strings.Index(text, ",") > -1 {
			nums := utils.SplitToInt64(text, ",")
			pp = append(pp, Point{X: nums[0], Y: nums[1], Cnt: 1})
		} else if strings.Index(text, "fold along") > -1 {
			parts := strings.Split(text, "=")

			v, _ := strconv.ParseInt(parts[1], 10, 64)

			var d string
			if strings.Index(text, "fold along x") > -1 {
				d = "x"
			} else {
				d = "y"
			}

			ii = append(ii, Instruction{V: v, Dir: d})
		}

	}
	return pp, ii
}
