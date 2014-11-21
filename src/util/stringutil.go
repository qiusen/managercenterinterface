/*
*	字符串操作工具
 */
package util

import (
	"strconv"
)

/*
* 拼接所有参数为字符串
 */
func ToString(args ...interface{}) string {
	result := ""
	for _, args := range args {
		switch val := args.(type) {
		case int:
			result += strconv.Itoa(val)
		case string:
			result += val
		}

	}
	return result
}

/*
* 截取字符串
* start 开始位置， length 截取长度
 */
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}
