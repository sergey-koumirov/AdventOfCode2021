package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y, N int64
}

type Line struct {
	X1, Y1, X2, Y2 int64
}

type Input struct {
	Lines []Line
}

func LoadInput() Input {
	var result Input

	result.Lines = make([]Line, 0, 500)

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Replace(text, " -> ", ",", -1)
		nums := SplitToInt64(parts, ",")

		temp := Line{X1: nums[0], Y1: nums[1], X2: nums[2], Y2: nums[3]}

		result.Lines = append(result.Lines, temp)
	}

	return result
}

func SplitToInt64(line string, s string) []int64 {
	parts := strings.Split(line, s)

	result := make([]int64, len(parts))

	for i, part := range parts {
		result[i], _ = strconv.ParseInt(part, 10, 64)
	}

	return result
}
