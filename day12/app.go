package main

import (
	"advent/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Cave struct {
	Name string
	Once bool
}

type Links map[string][]Cave

func main() {
	start := time.Now()

	data := LoadInput()
	fmt.Println(data)

	totals := make([]int64, len(data)+1)

	totals[0] = calc(&data, "")
	index := 1
	for k := range data {
		if k == "start" || k == "end" || !utils.IsLower(k) {
			totals[index] = totals[0]
		} else {
			totals[index] = calc(&data, k)
		}
		index++
	}

	total := totals[0]
	for i := 1; i <= len(data); i++ {
		total = total + (totals[i] - totals[0])
	}

	fmt.Println("total", total, len(data), totals)

	elapsed := time.Since(start)
	fmt.Println("time", elapsed)
}

func calc(data *Links, special string) int64 {
	total := int64(0)
	path := make([]string, 3*len(*data))
	path[0] = "start"
	current := 0
	used := make(map[string]bool)
	used["start"] = true
	specialUses := 0
	deepCheck("start", data, &path, current, &total, used, special, specialUses)

	return total
}

func deepCheck(key string, data *Links, path *[]string, current int, total *int64, used map[string]bool, special string, specialUses int) {
	if key == "end" {
		*total++
		// print(path, current) 3230
		return
	}

	current++
	for _, nextKey := range (*data)[key] {
		u, ex := used[nextKey.Name]
		if !ex || !u {
			if nextKey.Once && nextKey.Name != special {
				used[nextKey.Name] = true
			} else if nextKey.Once && nextKey.Name == special {
				specialUses++
				if specialUses == 2 {
					used[nextKey.Name] = true
				}
			}

			(*path)[current] = nextKey.Name
			deepCheck(nextKey.Name, data, path, current, total, used, special, specialUses)

			if nextKey.Once && nextKey.Name != special {
				used[nextKey.Name] = false
			} else if nextKey.Once && nextKey.Name == special {
				specialUses--
				used[nextKey.Name] = false
			}

		}
	}
}

func print(path *[]string, current int) {
	for i := 0; i <= current; i++ {
		if i > 0 {
			fmt.Print(" - ")
		}
		fmt.Print((*path)[i])
	}
	fmt.Println()
}

func LoadInput() Links {
	result := make(Links)

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, "-")

		link0, ex0 := result[parts[0]]
		if !ex0 {
			link0 = make([]Cave, 0)
		}
		link0 = append(link0, Cave{Name: parts[1], Once: utils.IsLower(parts[1])})
		result[parts[0]] = link0

		link1, ex1 := result[parts[1]]
		if !ex1 {
			link1 = make([]Cave, 0)
		}
		link1 = append(link1, Cave{Name: parts[0], Once: utils.IsLower(parts[0])})
		result[parts[1]] = link1
	}
	return result
}
