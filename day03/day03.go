package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func frf(temp []string) []int {
	fr := make([]int, 12)
	for _, line := range temp {
		for i := 0; i < 12; i++ {
			if line[i] == '1' {
				fr[i] = fr[i] + 1
			}
		}
	}

	return fr
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	temp := make([]string, 1000)
	i := int(0)

	for scanner.Scan() {
		line := scanner.Text()
		temp[i] = line
		i = i + 1
	}

	ci := int(0)

	before := make([]string, 0, 1000)
	for _, line := range temp {
		before = append(before, line)
	}

	for {
		fr := frf(before)

		after := make([]string, 0, 1000)
		lb := len(before)
		for _, line := range before {
			if fr[ci] >= (lb-fr[ci]) && line[ci] == '1' || fr[ci] < (lb-fr[ci]) && line[ci] == '0' {
				after = append(after, line)
			}
		}

		// fmt.Println(ci, after)

		ci = ci + 1
		before = after

		if len(after) <= 1 {
			x, _ := strconv.ParseInt(after[0], 2, 64)
			fmt.Println("oxy", after, x)
			break
		}
	}

	ci = 0
	before = make([]string, 0, 1000)
	for _, line := range temp {
		before = append(before, line)
	}

	for {

		fr := frf(before)

		after := make([]string, 0, 1000)
		lb := len(before)
		for _, line := range before {
			if fr[ci] < (lb-fr[ci]) && line[ci] == '1' || fr[ci] >= (lb-fr[ci]) && line[ci] == '0' {
				after = append(after, line)
			}
		}

		// fmt.Println(ci, after)

		ci = ci + 1
		before = after

		if len(after) <= 1 {
			x, _ := strconv.ParseInt(after[0], 2, 64)
			fmt.Println("co2", after, x)
			break
		}
	}

}
