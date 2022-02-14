package main

import (
	"advent/utils"
	"bufio"
	"fmt"
	"os"
	"time"
)

type Field [10][]int64

func main() {
	start := time.Now()

	data := LoadInput()
	fmt.Println(data)

	fmt.Println("======================")

	// total := int64(0)

	i := 1
	for {
		flashes := step(&data)
		print(&data, i)
		i++
		if flashes == 100 || i > 10000 {
			break
		}
	}

	// for i := 0; i < 100; i++ {
	// 	flashes := step(&data)
	// 	print(&data, i)
	// 	total = total + flashes
	// }

	// fmt.Println("======================", total)

	elapsed := time.Since(start)
	fmt.Println("time", elapsed)
}

func print(data *Field, step int) {
	fmt.Println("---", step)
	for i := 0; i < 10; i++ {
		fmt.Println(data[i])
	}
}

func step(data *Field) int64 {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			data[i][j]++
			if data[i][j] == 10 {
				flash(data, i, j)
			}
		}
	}

	result := int64(0)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if data[i][j] > 9 {
				result++
				data[i][j] = 0
			}
		}
	}
	return result
}

func flash(data *Field, i int, j int) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			newI := i + dx
			newJ := j + dy
			if newI >= 0 && newI <= 9 && newJ >= 0 && newJ <= 9 && !(dx == 0 && dy == 0) {
				data[newI][newJ]++
				if data[newI][newJ] == 10 {
					flash(data, newI, newJ)
				}
			}
		}
	}
}

func LoadInput() Field {
	result := Field{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	index := int64(0)
	for scanner.Scan() {
		text := scanner.Text()
		result[index] = utils.SplitToInt64(text, "")
		index++
	}
	return result
}
