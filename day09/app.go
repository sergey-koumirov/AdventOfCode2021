package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Field [100][100]int64

type Uses map[int64]map[int64]bool

func main() {
	start := time.Now()

	data := LoadInput()
	fmt.Println(data)

	uses := make(Uses)
	for i := int64(0); i < 100; i++ {
		uses[i] = make(map[int64]bool)
	}

	sizes := []int{}

	for i := int64(0); i < 100; i++ {
		for j := int64(0); j < 100; j++ {
			if !uses[i][j] && isLow(&data, i, j) {
				// total = total + data[i][j] + 1
				uses[i][j] = true
				size := int(0)
				deepCalc(&data, uses, &size, i, j)
				sizes = append(sizes, size)
				fmt.Println("Basin:", size, i, j)
			}
		}
	}

	l := len(sizes)
	sort.Ints(sizes)

	fmt.Println(sizes, sizes[l-1]*sizes[l-2]*sizes[l-3])

	elapsed := time.Since(start)
	fmt.Println("time", elapsed)
}

func deepCalc(data *Field, uses Uses, size *int, i int64, j int64) {
	uses[i][j] = true
	*size++

	if i > 0 && !uses[i-1][j] && data[i-1][j] < 9 {
		deepCalc(data, uses, size, i-1, j)
	}

	if i < 99 && !uses[i+1][j] && data[i+1][j] < 9 {
		deepCalc(data, uses, size, i+1, j)
	}

	if j > 0 && !uses[i][j-1] && data[i][j-1] < 9 {
		deepCalc(data, uses, size, i, j-1)
	}

	if j < 99 && !uses[i][j+1] && data[i][j+1] < 9 {
		deepCalc(data, uses, size, i, j+1)
	}
}

func isLow(data *Field, i int64, j int64) bool {
	v := data[i][j]

	if i > 0 && data[i-1][j] <= v {
		return false
	}

	if i < 99 && data[i+1][j] <= v {
		return false
	}

	if j > 0 && data[i][j-1] <= v {
		return false
	}

	if j < 99 && data[i][j+1] <= v {
		return false
	}

	return true
}

func LoadInput() Field {
	result := Field{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	i := 0

	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, "")
		for j, t := range parts {
			result[i][j], _ = strconv.ParseInt(t, 10, 64)
		}
		i++
	}
	return result
}
