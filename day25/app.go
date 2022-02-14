package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const FileName = "input.txt"

type Field struct {
	Cells [][]int
	H, W  int
}

var field = LoadInput()

func main() {
	start := time.Now()

	// fmt.Printf("%+v\n", field)
	// print(field)

	step := 0

	for {
		step++
		next, moves := runStep(field)

		fmt.Println("\nStep: ", step)
		// print(next)

		if moves == 0 || step > 1000 {
			break
		}

		field = next
	}

	elapsed := time.Since(start)
	fmt.Println("done", elapsed)
}

func runStep(f Field) (Field, int) {
	result := Field{W: f.W, H: f.H}

	result.Cells = make([][]int, f.H)

	for line := 0; line < f.H; line++ {
		result.Cells[line] = make([]int, f.W)
	}

	moves := 0

	for line := 0; line < f.H; line++ {
		for x := 0; x < f.W; x++ {
			if f.Cells[line][x] == 2 && x < f.W-1 {
				if f.Cells[line][x+1] == 0 {
					result.Cells[line][x+1] = 2
					moves++
				} else {
					result.Cells[line][x] = 2
				}
			}

			if f.Cells[line][x] == 2 && x == f.W-1 {
				if f.Cells[line][0] == 0 {
					result.Cells[line][0] = 2
					moves++
				} else {
					result.Cells[line][x] = 2
				}
			}
		}
	}

	// print(result)

	for line := 0; line < f.H; line++ {
		for x := 0; x < f.W; x++ {
			if f.Cells[line][x] == 1 && line < f.H-1 {
				if (f.Cells[line+1][x] == 0 || f.Cells[line+1][x] == 2) && result.Cells[line+1][x] == 0 {
					result.Cells[line+1][x] = 1
					moves++
				} else {
					result.Cells[line][x] = 1
				}
			}

			if f.Cells[line][x] == 1 && line == f.H-1 {
				if (f.Cells[0][x] == 0 || f.Cells[0][x] == 2) && result.Cells[0][x] == 0 {
					result.Cells[0][x] = 1
					moves++
				} else {
					result.Cells[line][x] = 1
				}
			}
		}
	}

	return result, moves
}

func print(f Field) {
	fmt.Printf("F=%d x %d\n", f.W, f.H)

	for y := 0; y < f.H; y++ {
		for x := 0; x < f.W; x++ {

			if f.Cells[y][x] == 2 {
				fmt.Print(">")
			}
			if f.Cells[y][x] == 1 {
				fmt.Print("v")
			}
			if f.Cells[y][x] == 0 {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func LoadInput() Field {
	result := Field{}

	file, _ := os.Open(FileName)
	scanner := bufio.NewScanner(file)

	index := 0
	for scanner.Scan() {
		index++
		text := scanner.Text()

		parts := strings.Split(text, "")

		if result.W == 0 {
			result.W = len(parts)
		}

		temp := make([]int, result.W)

		for i, part := range parts {
			if part == ">" {
				temp[i] = 2
			}
			if part == "v" {
				temp[i] = 1
			}
		}

		result.Cells = append(result.Cells, temp)
	}

	result.H = index

	return result
}
