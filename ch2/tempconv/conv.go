package main

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2021/7/16 6:16 下午
 * @Desc:
 */

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
