package main

import (
	"day04/utils"
	"fmt"
)

func main() {
	data := utils.LoadInput()

	index := int(0)
	for {
		winIndex := checkBoard(&data, index)

		if winIndex > -1 {
			fmt.Println(calcScore(data, winIndex, index))
			break
		}

		index = index + 1

		if index >= len(data.Seq) {
			fmt.Println("No winner")
			break
		}
	}

	fmt.Println(data)
}

func checkBoard(data *utils.Input, index int) int {
	n := data.Seq[index]

	winIndex := -1
	lastWinIndex := -1

	for i := 0; i < len(data.Boards); i++ {
		if !data.IsWin[i] {
			for x := 0; x < 5; x++ {
				for y := 0; y < 5; y++ {
					if data.Boards[i][x][y] == n {
						data.Used[i][x][y] = 1
					}
				}
			}

			for x := 0; x < 5; x++ {
				if data.Used[i][x][0] == 1 && data.Used[i][x][1] == 1 && data.Used[i][x][2] == 1 && data.Used[i][x][3] == 1 && data.Used[i][x][4] == 1 {
					winIndex = i
				}
			}

			for y := 0; y < 5; y++ {
				if data.Used[i][0][y] == 1 && data.Used[i][1][y] == 1 && data.Used[i][2][y] == 1 && data.Used[i][3][y] == 1 && data.Used[i][4][y] == 1 {
					winIndex = i
				}
			}

			if winIndex > -1 {
				data.IsWin[winIndex] = true
				winIndex = -1

				all := true
				for _, v := range data.IsWin {
					if !v {
						all = false
					}
				}

				if all {
					lastWinIndex = i
				}
			}
		}

	}

	return lastWinIndex
}

func calcScore(data utils.Input, winIndex int, index int) int64 {
	s := int64(0)

	n := data.Seq[index]

	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if data.Used[winIndex][x][y] == 0 {
				s = s + data.Boards[winIndex][x][y]
			}
		}
	}

	return s * n
}
