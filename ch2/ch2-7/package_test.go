package main

import (
	"fmt"
	"testing"
)

/**
 * @Author: sm.gu
 * @Author: sm.gu@aftership.com
 * @Date: 2021/7/19 11:03 上午
 * @Desc:
 */

func f() {}

var g = "g"

func Test1(t *testing.T) {
	f := "f"
	fmt.Println(f)
	fmt.Println(g)
	//fmt.Println(h) compile error
}
