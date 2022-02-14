package main

import (
	"bufio"
	"day07/utils"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	data := LoadInput()
	fmt.Println(data)

	min, max := utils.MinMax(data)
	fmt.Println("MinMax", min, max)

	minF := int64(-1)
	cnt := len(data)

	for x := min; x <= max; x++ {
		s := int64(0)
		for i := 0; i < cnt; i++ {
			s = s + Fuel(utils.Abs(data[i]-x))
		}

		if minF == -1 || minF > s {
			minF = s
		}
	}

	fmt.Println("Fuel", minF)

	elapsed := time.Since(start)
	fmt.Println("Binomial took", elapsed)
}

func Fuel(n int64) int64 {
	return (n*n + n) / 2
}

func LoadInput() []int64 {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()

	temp := utils.SplitToInt64(text, ",")

	return temp
}
