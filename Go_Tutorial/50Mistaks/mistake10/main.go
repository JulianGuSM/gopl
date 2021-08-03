package main

import (
	"fmt"
	"reflect"
)

/**
 * @Author: sm.gu
 * @Email: sm.gu@aftership.com
 * @Date: 2021/7/30 4:29 下午
 * @Desc:
 */

func main() {
	x := [3]int{1, 2, 3} // x is array
	y := []int{1, 2, 3}  // y is slice

	typeX := reflect.TypeOf(x)
	typeY := reflect.TypeOf(y)

	fmt.Printf("var x's type is %v, var y's type is %v", typeX, typeY)
}
