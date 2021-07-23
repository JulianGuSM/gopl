package main

import (
	"fmt"
	"time"
)

/**
 * @Author: sm.gu
 * @Email: sm.gu@aftership.com
 * @Date: 2021/7/23 11:16 上午
 * @Desc:
 */

// When we use golang write web server, usually go server
// would create a goroutine process business concurrently.
// At the same time, this goroutine maybe creates other goroutine
// to access data from db or RPC services.

// When each our Http request timeout or terminated， we need to close
// all goroutine elegantly and release resource.So, we need a mechanism
// to notify all goroutines this request is canceled.

//For example, when sleepRandom_1() exit, can not notify sleepRandom_2()

func sleeprandom1() {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("this is sleep Random 1: %d\n", i)

		i++
		if i == 5 {
			fmt.Printf("cancel sleep random 1")
			break
		}
	}
}

func sleeprandom2() {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("This is sleep Random 2: %d\n", i)
		i++
	}
}

func main() {
	go sleeprandom1()
	go sleeprandom2()

	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Continue...")

	}
}
