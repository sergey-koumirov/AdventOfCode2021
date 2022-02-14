package main

import (
	"advent/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

type Field [100][100]int64

type Uses map[int64]map[int64]bool

func main() {
	start := time.Now()

	data := LoadInput()
	// fmt.Println(data)

	scores := []int{}

	for _, line := range data {
		score := check(line)
		if score > 0 {
			scores = append(scores, score)
		}
		// fmt.Println(line)
	}

	sort.Ints(scores)

	fmt.Println("scores", len(scores)/2, scores[23])

	elapsed := time.Since(start)
	fmt.Println("time", elapsed)
}

func check(line string) int {
	stack := utils.Stack{}

	for _, v := range line {
		s1 := string(v)
		if isOpen(s1) {
			stack.Push(s1)
		}
		if isClose(s1) {
			s2 := stack.Pop()
			if !isPair(s2, s1) {
				// fmt.Println("error", i)
				return 0
			}
		}
	}

	score := int(0)
	if stack.Len() > 0 {
		fmt.Println("incomplete", stack)
		for {
			s := stack.Pop()
			if s == "" {
				break
			}
			var point int
			if s == "(" {
				point = 1
			}
			if s == "[" {
				point = 2
			}
			if s == "{" {
				point = 3
			}
			if s == "<" {
				point = 4
			}

			score = score*5 + point

		}
	}

	return score
}

func isPair(s1, s2 string) bool {
	return s1 == "(" && s2 == ")" || s1 == "{" && s2 == "}" || s1 == "[" && s2 == "]" || s1 == "<" && s2 == ">"
}

func isOpen(s string) bool {
	return s == "(" || s == "{" || s == "[" || s == "<"
}

func isClose(s string) bool {
	return s == ")" || s == "}" || s == "]" || s == ">"
}

func LoadInput() []string {
	result := []string{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		result = append(result, text)
	}
	return result
}
