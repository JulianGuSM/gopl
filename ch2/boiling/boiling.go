package main

import "fmt"

/**
 * @Author: sm.gu
 * @Author: sm.gu@aftership.com
 * @Date: 2021/7/16 3:44 下午
 * @Desc:
 */

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9

	//%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// Output:
	// boiling point = 212°F or 100°C
}
