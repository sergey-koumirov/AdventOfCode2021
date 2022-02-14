package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Data struct {
	Alg   []int
	Image [][]int
	S     int
}

func main() {
	start := time.Now()

	data := LoadInput()
	fmt.Println(data.Alg)
	fmt.Println("000 000 000 =>", data.Alg[0])
	fmt.Println("111 111 111 =>", data.Alg[511])

	// fmt.Println(data.Image)

	// print(data.Image)

	fmt.Println("--------------------------------------------------------------")

	for i := 0; i < 50; i++ {
		fmt.Println("x", i)
		enchance(&data, i)
	}

	fmt.Println("Result:", countLights(data.Image))

	elapsed := time.Since(start)
	fmt.Println("done", elapsed)
}

func enchance(data *Data, index int) {
	temp, framed := makeCopy(data, index)

	for i := 1; i < data.S-1; i++ {
		for j := 1; j < data.S-1; j++ {
			bitIndex := 0
			for k := -1; k <= 1; k++ {
				for n := -1; n <= 1; n++ {
					bitIndex = (bitIndex << 1) | framed[i+k][j+n]
				}
			}
			temp[i][j] = data.Alg[bitIndex]
		}
	}

	data.Image = temp
	// print(framed)
	// print(temp)
}

func countLights(dd [][]int) int {
	result := 0
	for i := 0; i < len(dd); i++ {
		for j := 0; j < len(dd); j++ {
			if dd[i][j] == 1 {
				result++
			}
		}
	}
	return result
}

func print(dd [][]int) {
	fmt.Println()
	for i := 0; i < len(dd); i++ {
		for j := 0; j < len(dd); j++ {
			if dd[i][j] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("Count:", countLights(dd))
}

func makeCopy(data *Data, index int) ([][]int, [][]int) {
	data.S = data.S + 4

	var frameBit, nextFrameBit int
	isBlinking := data.Alg[0] == 1 && data.Alg[511] == 0

	if index%2 == 1 && isBlinking {
		frameBit = 1
		nextFrameBit = 0
	} else if index%2 == 0 && isBlinking {
		frameBit = 0
		nextFrameBit = 1
	}

	framed := make([][]int, data.S)
	temp := make([][]int, data.S)

	for i := 0; i < data.S; i++ {
		temp[i] = make([]int, data.S)
		framed[i] = make([]int, data.S)

		if i == 0 || i == data.S-1 {
			for j := 0; j < data.S; j++ {
				temp[i][j] = nextFrameBit
				framed[i][j] = frameBit
			}
		}
		if i == 1 || i == data.S-2 {
			for j := 0; j < data.S; j++ {
				framed[i][j] = frameBit
			}
		}

		temp[i][0] = nextFrameBit
		temp[i][data.S-1] = nextFrameBit

		framed[i][0] = frameBit
		framed[i][1] = frameBit
		framed[i][data.S-2] = frameBit
		framed[i][data.S-1] = frameBit

		if i > 1 && i < data.S-2 {
			for j := 2; j < data.S-2; j++ {
				framed[i][j] = data.Image[i-2][j-2]
			}
		}
	}

	return temp, framed
}

func LoadInput() Data {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	text := scanner.Text()

	result := Data{}
	result.Alg = make([]int, len(text))

	for i := 0; i < len(text); i++ {
		if text[i:i+1] == "#" {
			result.Alg[i] = 1
		} else {
			result.Alg[i] = 0
		}
	}

	scanner.Scan()

	for scanner.Scan() {
		imgText := scanner.Text()

		temp := make([]int, len(imgText))

		for i := 0; i < len(imgText); i++ {
			if imgText[i:i+1] == "#" {
				temp[i] = 1
			} else {
				temp[i] = 0
			}
		}

		result.Image = append(result.Image, temp)
	}

	result.S = len(result.Image)

	return result
}
