package main

import (
	"advent/utils"
	"bufio"
	"fmt"
	"os"
	"time"
)

type Field [500][500]int
type Costs [500][500]int

var delta = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func main() {
	start := time.Now()

	field := LoadInput()

	fmt.Println(field)

	costs := Costs{}
	initCosts(&costs)

	print5(&field)
	fmt.Println("-------------")
	printX(&field)

	elapsed := time.Since(start)
	fmt.Println("load", elapsed)

	best := -1
	deep(&field, &costs, 0, 0, &best)

	fmt.Println(costs)
	fmt.Println(costs[499][499])

	elapsed = time.Since(start)
	fmt.Println("done", elapsed)
}

func print5(field *Field) {
	for dy := 0; dy < 5; dy++ {
		for dx := 0; dx < 5; dx++ {
			fmt.Print(field[dx][dy])
		}
		fmt.Println()
	}
}

func printX(field *Field) {
	for dy := 0; dy < 5; dy++ {
		for dx := 0; dx < 5; dx++ {
			fmt.Print(field[dx*100][dy*100+1])
		}
		fmt.Println()
	}

}

func deep(field *Field, costs *Costs, x int, y int, best *int) {
	for i := 0; i < 4; i++ {
		newX := x + delta[i][0]
		newY := y + delta[i][1]
		if newX >= 0 && newX < 500 && newY >= 0 && newY < 500 {
			newCost := costs[x][y] + field[newX][newY]
			currentCost := costs[newX][newY]
			b := *best
			if (b == -1 || b > -1 && b > newCost) && (currentCost == -1 || newCost < currentCost) {
				// fmt.Println(newX, newY, currentCost)

				if newX == 499 && newY == 499 {
					*best = newCost
					fmt.Println("best: ", newCost)
				}
				costs[newX][newY] = newCost
				deep(field, costs, newX, newY, best)
			}
		}
	}
}

func initCosts(costs *Costs) {
	for x := 0; x < 500; x++ {
		for y := 0; y < 500; y++ {
			costs[x][y] = -1
		}
	}
	costs[0][0] = 0
}

func LoadInput() Field {
	result := Field{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	y := 0
	for scanner.Scan() {
		text := scanner.Text()

		temp := utils.SplitToInt64(text, "")

		for x := 0; x < 100; x++ {
			// result[index][i] = int(temp[i])

			for dx := 0; dx < 5; dx++ {
				for dy := 0; dy < 5; dy++ {
					m := dx + dy
					n := int(temp[x])

					result[dx*100+x][dy*100+y] = int((n+m-1)%9 + 1)
				}
			}
		}

		y++
	}
	return result
}
