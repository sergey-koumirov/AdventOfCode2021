package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Table [5][5]int64

type Input struct {
	Seq    []int64
	Boards []Table
	Used   []Table
	IsWin  []bool
}

func LoadInput() Input {
	var result Input

	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	seqLine := scanner.Text()
	result.Seq = SplitToInt64(seqLine, ",")

	result.Boards = make([]Table, 0, 50)

	tIndex := int(0)

	buffer := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			buffer = append(buffer, line)
		}

		if len(buffer) == 5 {
			result.Boards = append(result.Boards, Table{})

			for i := 0; i < 5; i++ {
				sNums := strings.Fields(buffer[i])
				for j, sNum := range sNums {
					num, _ := strconv.ParseInt(sNum, 10, 64)
					result.Boards[tIndex][i][j] = num
				}
			}

			tIndex = tIndex + 1
			buffer = make([]string, 0)
		}
	}

	result.Used = make([]Table, len(result.Boards))
	result.IsWin = make([]bool, len(result.Boards))

	return result
}

func SplitToInt64(line string, s string) []int64 {
	parts := strings.Split(line, s)

	result := make([]int64, len(parts))

	for i, part := range parts {
		result[i], _ = strconv.ParseInt(part, 10, 64)
	}

	return result
}
