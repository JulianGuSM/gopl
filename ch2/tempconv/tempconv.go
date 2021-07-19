package main

import "fmt"

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2021/7/16 6:14 下午
 * @Desc:
 */

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
