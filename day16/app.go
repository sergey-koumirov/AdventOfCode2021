package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	data := LoadInput()
	fmt.Println(data)
	fmt.Println("-----")

	bindata := decode(data)
	fmt.Println(bindata)
	fmt.Println("-----")

	total := 0
	next, calcRes := parsePacket(bindata, 0, 0, &total)
	fmt.Println("Rest: ", bindata[next:])
	// fmt.Println("Sum: ", total)
	fmt.Println("Calc: ", calcRes)

	elapsed := time.Since(start)
	fmt.Println("done", elapsed)
}

func parsePacket(bindata string, start int, level int, total *int) (int, int64) {
	tab := strings.Repeat(" ", level*2)
	vvv := bindata[start : start+3]
	vvv10, _ := strconv.ParseInt(vvv, 2, 32)
	ttt := bindata[start+3 : start+6]

	// fmt.Println(tab, "Version: ", vvv, vvv10)
	// fmt.Println(tab, "Type:    ", ttt)

	*total = *total + int(vvv10)

	if ttt == "100" {
		next, num := parseType4(start, bindata, tab)
		// fmt.Println(tab, "decimal: ", num)
		return next, num
	} else {
		next, values := parseTypeNot4(start, bindata, tab, level, total)
		var result int64
		if ttt == "000" {
			result = 0
			for _, v := range values {
				result = result + v
			}
			fmt.Println(tab, "sum: ", values, result)
		}
		if ttt == "001" {
			result = 1
			for _, v := range values {
				result = result * v
			}
			fmt.Println(tab, "mult: ", values, result)
		}

		if ttt == "010" {
			result = 99999
			for _, v := range values {
				if result > v {
					result = v
				}
			}
			fmt.Println(tab, "min: ", values, result)
		}

		if ttt == "011" {
			result = -1
			for _, v := range values {
				if result < v {
					result = v
				}
			}
			fmt.Println(tab, "max: ", values, result)
		}

		if ttt == "101" {
			if values[0] > values[1] {
				result = 1
			} else {
				result = 0
			}
			fmt.Println(tab, "> : ", values, result)
		}

		if ttt == "110" {
			if values[0] < values[1] {
				result = 1
			} else {
				result = 0
			}
			fmt.Println(tab, "< : ", values, result)
		}

		if ttt == "111" {
			if values[0] == values[1] {
				result = 1
			} else {
				result = 0
			}
			fmt.Println(tab, "= : ", values, result)
		}
		return next, result
	}
}

func parseTypeNot4(start int, bindata string, tab string, level int, total *int) (int, []int64) {
	result := make([]int64, 0)
	l := bindata[start+6 : start+7]
	// fmt.Println(tab, "L Bit:   ", l)
	if l == "0" {
		sumLen := bindata[start+7 : start+7+15]
		sumLen10, _ := strconv.ParseInt(sumLen, 2, 32)
		// fmt.Println(tab, "Sub-Packets: L=", sumLen, sumLen10)
		next := start + 7 + 15
		var value int64
		for {
			if next >= start+7+15+int(sumLen10) {
				break
			}
			// fmt.Println(tab, "Sub-Packet from ", next)
			next, value = parsePacket(bindata, next, level+1, total)
			result = append(result, value)
		}
		return next, result
	} else {
		n := bindata[start+7 : start+18]
		n10, _ := strconv.ParseInt(n, 2, 32)
		// fmt.Println(tab, "Sub-Packets: N=", n, n10)
		next := start + 7 + 11
		var value int64
		for pi := int64(0); pi < n10; pi++ {
			// fmt.Println(tab, "Sub-Packet #", pi)
			next, value = parsePacket(bindata, next, level+1, total)
			result = append(result, value)
		}
		return next, result
	}
}

func parseType4(start int, bindata string, tab string) (int, int64) {
	result := make([]int64, 0)
	next := start + 6
	for {
		v := bindata[next+1 : next+5]
		v10, _ := strconv.ParseInt(v, 2, 32)
		result = append(result, v10)
		if bindata[next:next+1] == "0" {
			next = next + 5
			// fmt.Println(tab, "Number 0: ", v10)
			break
		}
		next = next + 5
		// fmt.Println(tab, "Number 1: ", v10)
	}

	num := int64(0)
	for i := 0; i < len(result); i++ {
		pw := len(result) - i - 1
		num = num + result[i]*int64(math.Pow(16, float64(pw)))
	}

	// fmt.Println(tab, "decimal: ", num)

	return next, num
}

func decode(data string) string {
	temp := make([]string, len(data))
	for i, s := range data {
		if string(s) == "0" {
			temp[i] = "0000"
		}
		if string(s) == "1" {
			temp[i] = "0001"
		}
		if string(s) == "2" {
			temp[i] = "0010"
		}
		if string(s) == "3" {
			temp[i] = "0011"
		}
		if string(s) == "4" {
			temp[i] = "0100"
		}
		if string(s) == "5" {
			temp[i] = "0101"
		}
		if string(s) == "6" {
			temp[i] = "0110"
		}
		if string(s) == "7" {
			temp[i] = "0111"
		}
		if string(s) == "8" {
			temp[i] = "1000"
		}
		if string(s) == "9" {
			temp[i] = "1001"
		}
		if string(s) == "A" {
			temp[i] = "1010"
		}
		if string(s) == "B" {
			temp[i] = "1011"
		}
		if string(s) == "C" {
			temp[i] = "1100"
		}
		if string(s) == "D" {
			temp[i] = "1101"
		}
		if string(s) == "E" {
			temp[i] = "1110"
		}
		if string(s) == "F" {
			temp[i] = "1111"
		}
	}

	return strings.Join(temp, "")
}

func LoadInput() string {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	return scanner.Text()
}
