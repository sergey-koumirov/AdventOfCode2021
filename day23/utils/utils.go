package utils

import (
	"strconv"
	"strings"
	"unicode"
)

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func MinMax(array []int64) (int64, int64) {
	var max int64 = array[0]
	var min int64 = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func SplitToInt64(line string, s string) []int64 {
	parts := strings.Split(line, s)

	result := make([]int64, len(parts))

	for i, part := range parts {
		result[i], _ = strconv.ParseInt(part, 10, 64)
	}

	return result
}

func Max(v1 int64, v2 int64) int64 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func Abs(v int64) int64 {
	if v > 0 {
		return v
	}
	return v * -1
}

func AbsInt(v int) int {
	if v > 0 {
		return v
	}
	return v * -1
}

func Sign(v int64) int64 {
	if v > 0 {
		return int64(1)
	}
	if v < 0 {
		return int64(-1)
	}
	return 0
}

func Sort12(v1, v2 int64) (int64, int64) {
	if v1 < v2 {
		return v1, v2
	}
	return v2, v1
}

func FindRange(a1, a2, b1, b2 int64) (int64, int64) {
	v1 := int64(-1)
	v2 := int64(-1)

	if b1 <= a1 && a1 <= b2 {
		v1 = a1
	}
	if a1 <= b1 && b1 <= a2 {
		v1 = b1
	}

	if b1 <= a2 && a2 <= b2 {
		v2 = a2
	}
	if a1 <= b2 && b2 <= a2 {
		v2 = b2
	}

	return v1, v2
}
