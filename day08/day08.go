package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Line struct {
	Digits []string
	Codes  []string
}

type Use struct {
	sym string
	n   int64
}

func main() {
	start := time.Now()

	data := LoadInput()
	// fmt.Println(data)

	total := int64(0)

	for i, line := range data {
		d3 := findFirstByLen(line.Digits, 3)
		d2 := findFirstByLen(line.Digits, 2)
		uses := countUses(line.Digits)

		diff := minus(d3, d2)
		s4 := findByUse(uses, 4)
		s6 := findByUse(uses, 6)
		s9 := findByUse(uses, 9)

		s8 := findByUse(uses, 8)
		s8m := removeEl(s8, diff)

		d4 := findFirstByLen(line.Digits, 4)
		d4m1 := minus(d4, s8m[0]+s9[0]+s6[0])

		last := minus("abcdefg", diff+s8m[0]+s9[0]+s4[0]+s6[0]+d4m1)

		replace(&data[i], diff, "1")
		replace(&data[i], s8m[0], "2")
		replace(&data[i], s9[0], "3")
		replace(&data[i], last, "4")
		replace(&data[i], s4[0], "5")
		replace(&data[i], s6[0], "6")
		replace(&data[i], d4m1, "7")

		fmt.Println(data[i])

		decoded := [4]int64{}
		for i, code := range line.Codes {
			n := codeToInt(code)
			decoded[i] = n
		}

		total = total + decoded[0]*1000 + decoded[1]*100 + decoded[2]*10 + decoded[3]
	}
	fmt.Println("total", total)

	elapsed := time.Since(start)
	fmt.Println("time", elapsed)
}

func codeToInt(code string) int64 {
	parts := strings.Split(code, "")
	sort.Strings(parts)
	js := strings.Join(parts, "")

	if js == "123456" {
		return 0
	}

	if js == "23" {
		return 1
	}

	if js == "12457" {
		return 2
	}

	if js == "12347" {
		return 3
	}

	if js == "2367" {
		return 4
	}

	if js == "13467" {
		return 5
	}

	if js == "134567" {
		return 6
	}

	if js == "123" {
		return 7
	}

	if js == "1234567" {
		return 8
	}

	if js == "123467" {
		return 9
	}

	return -1
}

func removeEl(ss []string, el string) []string {
	result := make([]string, 0)

	for _, s := range ss {
		if s != el {
			result = append(result, s)
		}
	}

	return result
}

func findByUse(uses map[string]int64, n int64) []string {
	result := make([]string, 0)
	for s, cnt := range uses {
		if n == cnt {
			result = append(result, s)
		}
	}
	return result
}

func countUses(dd []string) map[string]int64 {
	result := make(map[string]int64)
	result["a"] = 0
	result["b"] = 0
	result["c"] = 0
	result["d"] = 0
	result["e"] = 0
	result["f"] = 0
	result["g"] = 0

	for _, d := range dd {
		for i := 0; i < len(d); i++ {
			s := string(d[i])
			result[s] = result[s] + 1
		}
	}

	return result
}

func replace(l *Line, what string, with string) {
	for i := 0; i < len(l.Digits); i++ {
		l.Digits[i] = strings.ReplaceAll(l.Digits[i], what, with)
	}

	for i := 0; i < len(l.Codes); i++ {
		l.Codes[i] = strings.ReplaceAll(l.Codes[i], what, with)
	}
}

func minus(v1, v2 string) string {
	result := ""

	for i := 0; i < len(v1); i++ {
		found := false
		for j := 0; j < len(v2); j++ {
			if v1[i] == v2[j] {
				found = true
			}
		}
		if !found {
			result = result + string(v1[i])
		}
	}

	return result
}

func findFirstByLen(dd []string, n int) string {
	for _, d := range dd {
		if len(d) == n {
			return d
		}
	}
	return ""
}

func LoadInput() []Line {
	result := make([]Line, 0, 200)
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, " | ")
		temp := Line{
			Digits: strings.Split(parts[0], " "),
			Codes:  strings.Split(parts[1], " "),
		}
		result = append(result, temp)
	}
	return result
}
