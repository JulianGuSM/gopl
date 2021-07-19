package main

import (
	"fmt"
	"github.com/gopl.io/ch2/tempconv"
)

/**
 * @Author: sm.gu
 * @Author: sm.gu@aftership.com
 * @Date: 2021/7/16 6:17 下午
 * @Desc:
 */

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC) // "Brrrr! -273.15°C"
	fmt.Println(tempconv.CToF(tempconv.BoilingC))     // "212°F"
}
