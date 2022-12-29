package types

import (
	"eshort/pkg/easylogger"
	"strconv"
)

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func StringArr2Int64Arr(str_arr []string) []int {
	ints := make([]int, len(str_arr))
	for i, v := range str_arr {
		ints[i], _ = strconv.Atoi(v)
	}
	return ints
}

func StringToUint64(str string) uint64 {
	if len(str) == 0 {
		return 0
	}
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		easylogger.LogError(err, "字符串转uint64失败")
	}
	return i
}

// StringToInt 将字符串转换为 int
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		easylogger.LogError(err, "字符串转int失败")
	}
	return i
}

func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}
