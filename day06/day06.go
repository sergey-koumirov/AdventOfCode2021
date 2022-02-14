package main

import (
	"bufio"
	"day06/utils"
	"fmt"
	"os"
	"time"
)

const DAYS = int64(256)

func main() {
	start := time.Now()

	data := [DAYS + 1][9]int64{}

	data[0] = LoadInput()
	fmt.Println(data)

	for i := int64(1); i <= DAYS; i++ {
		data[i][0] = data[i-1][1]
		data[i][1] = data[i-1][2]
		data[i][2] = data[i-1][3]
		data[i][3] = data[i-1][4]
		data[i][4] = data[i-1][5]
		data[i][5] = data[i-1][6]
		data[i][6] = data[i-1][7] + data[i-1][0]
		data[i][7] = data[i-1][8]
		data[i][8] = data[i-1][0]
	}

	total := int64(0)
	for _, v := range data[DAYS] {
		total = total + v
	}

	fmt.Println(total)

	elapsed := time.Since(start)
	fmt.Println("Binomial took", elapsed)
}

func LoadInput() [9]int64 {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()

	temp := utils.SplitToInt64(text, ",")

	result := [9]int64{}

	for _, v := range temp {
		result[v]++
	}

	return result
}
