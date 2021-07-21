package main

import (
	"fmt"
	greetings "github.com/JulianGuSM/gopl/Go_Tutorial/CreateModule"
	"log"
)

/**
 * @Author: sm.gu
 * @Author: sm.gu@aftership.com
 * @Date: 2021/7/20 3:05 下午
 * @Desc:
 */

func main() {
	fmt.Println("Hello, World!")

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//msg, err := greetings.Hello("Julian")
	msg, err := greetings.Hello("")
	if err != nil {

	}

	fmt.Println(msg)
}
