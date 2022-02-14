package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var aim = int64(0)
	var h = int64(0)
	var d = int64(0)

	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")
		x, _ := strconv.ParseInt(parts[1], 10, 64)

		if parts[0] == "down" {
			aim = aim + x
		}
		if parts[0] == "up" {
			aim = aim - x
		}

		if parts[0] == "forward" {
			h = h + x
			d = d + aim*x
		}

	}

	fmt.Println(h, d, h*d)
}
