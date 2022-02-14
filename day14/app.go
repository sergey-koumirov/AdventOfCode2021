package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type LR = struct {
	L string
	R string
}

type Pairs map[string]LR

func main() {
	start := time.Now()

	chain, data := LoadInput()
	fmt.Println(data, len(data))

	freq2 := make(map[string]int)
	fmt.Println(chain)

	for i := 1; i < len(chain); i++ {
		ss := chain[i-1 : i+1]
		freq2[ss]++
	}

	for i := 0; i < 40; i++ {
		newFreq := make(map[string]int)

		for k, v := range freq2 {
			lr, ex := data[k]
			if ex {
				newFreq[lr.L] = newFreq[lr.L] + v
				newFreq[lr.R] = newFreq[lr.R] + v
			} else {
				newFreq[k] = newFreq[k] + v
			}
		}

		freq2 = newFreq

		fmt.Println(freq2)
	}

	freq := make(map[string]int)
	for k, v := range freq2 {
		freq[k[0:1]] = freq[k[0:1]] + v
		freq[k[1:2]] = freq[k[1:2]] + v
	}

	for k, v := range freq {
		if k == "O" {
			freq[k] = (v-2)/2 + 2
		} else {
			freq[k] = v / 2
		}
	}

	fmt.Println(freq)

	max := -1
	min := -1
	for _, v := range freq {
		if max == -1 || max < v {
			max = v
		}
		if min == -1 || min > v {
			min = v
		}
	}

	fmt.Println("delta", max, min, max-min)

	elapsed := time.Since(start)
	fmt.Println("time", elapsed)
}

func countFreq(chain string) map[string]int {
	result := make(map[string]int)
	for i := 0; i < len(chain); i++ {
		s := chain[i : i+1]
		result[s]++
	}
	return result
}

func LoadInput() (string, Pairs) {
	result := make(Pairs)

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	initial := scanner.Text()
	scanner.Scan()

	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, " -> ")

		s1 := parts[0][0:1]
		s2 := parts[0][1:2]
		s3 := parts[1]

		result[parts[0]] = LR{L: s1 + s3, R: s3 + s2}
	}
	return initial, result
}
