package main

import "fmt"

/**
 * @Author: sm.gu
 * @Email: sm.gu@aftership.com
 * @Date: 2021/7/30 3:47 下午
 * @Desc:
 */

type info struct {
	result int
}

func work() (int, error) {
	return 13, nil
}

func main() {
	var data info

	result, err := work() //error
	fmt.Printf("info: %+v\n %v,%v", data, result, err)
}
