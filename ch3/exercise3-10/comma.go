package main

import (
	"bytes"
)

/**
 * @Author: sm.gu
 * @Author: sm.gu@aftership.com
 * @Date: 2021/7/19 2:05 下午
 * @Desc:
 */

func main() {

}

func comma(s string) string {
	sb := new(bytes.Buffer)

	if len(s) <= 3 {
		return s
	}
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteString(s[i : i+1])
		count++
		if count == 3 {
			count = 0
			sb.WriteString(",")
		}
	}

	return s
}
