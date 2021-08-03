package main

import "fmt"

/**
 * @Author: sm.gu
 * @Email: sm.gu@aftership.com
 * @Date: 2021/7/30 3:55 下午
 * @Desc:
 */

func main() {
	x := 1
	fmt.Println(x) //prints 1
	{
		fmt.Println(x) //prints 1

		x := 2
		fmt.Println(x) //prints 2
	}
	fmt.Println(x) //prints 1!
}
