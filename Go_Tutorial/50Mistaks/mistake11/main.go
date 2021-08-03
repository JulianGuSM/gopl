package main

import (
	"fmt"
	"sync"
)

/**
 * @Author: sm.gu
 * @Email: sm.gu@aftership.com
 * @Date: 2021/7/30 5:15 下午
 * @Desc:
 */

type testPtr struct {
	name string
}

type myWaitGroup struct {
	sync.WaitGroup
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1; i++ {
		wg.Add(1)
		println("wg's addr: ", &wg)
		go worker(i+1, &wg)
	}
	var p = testPtr{name: "julian"}
	fmt.Println("p's addr：", &p)
	incr(&p)

	wg.Wait()
	fmt.Println("main exits")
}

func incr(i *testPtr) {
	fmt.Println(i)
	i.name = "sm"
}

func worker(i int, wg *sync.WaitGroup) {
	fmt.Printf("goroutine %d is working...\n", i)

	println("wg: %#v\n", wg)

	//statep, semap := wg.state()

	fmt.Printf("goroutine %d done\n", i)
	defer wg.Done()
}

//type WaitGroupp sync.WaitGroup
//
//func (wg *WaitGroupp) state() (statep *uint64, semap *uint32) {
//	if uintptr(unsafe.Pointer(&wg.state1))%8 == 0 {
//		return (*uint64)(unsafe.Pointer(&wg.state1)), &wg.state1[2]
//	} else {
//		return (*uint64)(unsafe.Pointer(&wg.state1[1])), &wg.state1[0]
//	}
//}
