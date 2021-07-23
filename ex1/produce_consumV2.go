package main

import (
	"fmt"
	"time"
)

/**
 * @Author: sm.gu
 * @Email: sm.gu@aftership.com
 * @Date: 2021/7/21 3:29 下午
 * @Desc:
 */

func calculateNextInt(prev int) int {
	time.Sleep(1 * time.Second) // pretend this is an expensive operation
	return prev + 1
}

func main() {
	data := make(chan int)
	quit := make(chan bool)

	// producer
	go func() {
		var i = 0
		for {
			i = calculateNextInt(i)
			select {
			case data <- i:
				// do nothing, just send message to data channel, for blocking produce
			case <-quit:
				close(data)
				return
			}

		}
	}()

	// consumer
	for i := range data {
		fmt.Printf("i=%v\n", i)
		if i >= 5 {
			quit <- true
		}
	}
}
