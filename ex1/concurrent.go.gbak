package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
 * @Author: sm.gu
 * @Email: sm.gu@aftership.com
 * @Date:  2021/7/21 10:38 上午
 * @Desc:
 */

// Product 产品类
type Product struct {
	name  int
	value int
}

// producer 生产者
func producer(buffer *int, mu *sync.Mutex, name int, stop *bool) {
	for !*stop {
		product := Product{
			name:  name,
			value: rand.Int(),
		}

		mu.Lock()
		if *buffer > 0 {

			*buffer--
			fmt.Printf("producer %v produce a product,product:%v, buffer size: %v\n", name, product, *buffer)

		}
		mu.Unlock()
	}
}

// consumer 消费者
func consumer(buffer *int, mu *sync.Mutex, name int, limit int) {
	mu.Lock()
	if *buffer < limit {

		*buffer++
		fmt.Printf("consumer %v consume a product, buffer size: %v\n", name, *buffer)

	}
	mu.Unlock()
}

func main() {
	mu := &sync.Mutex{}
	stop := false
	buffer := 10

	// 创建 5 个生产者和 5 个消费者
	for i := 0; i < 5; i++ {
		go consumer(&buffer, mu, i, buffer)
		go producer(&buffer, mu, i, &stop)

	}

	time.Sleep(time.Duration(1) * time.Second)
	stop = true // 设置生产者终止信号

}
