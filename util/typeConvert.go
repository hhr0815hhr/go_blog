package util

import (
	"strconv"
)

func String2Int(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

func String2Int64(str string) int64 {
	return int64(String2Int(str))
}

func Int2String(i int) string {
	return strconv.Itoa(i)
}
