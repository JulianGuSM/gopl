package main

/**
 * @Author: sm.gu
 * @Author: sm.gu@aftership.com
 * @Date: 2021/7/19 2:04 下午
 * @Desc:
 */

func main() {

}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
