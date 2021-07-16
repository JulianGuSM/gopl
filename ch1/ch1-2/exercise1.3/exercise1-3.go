package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("echo1 cost time: " + strconv.FormatInt(echo1(), 10) + "ns")
	fmt.Println("echo2 cost time: " + strconv.FormatInt(echo2(), 10) + "ns")

}

func echo1() int64 {
	beginTime := time.Now().UnixNano()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	endTime := time.Now().UnixNano()
	return endTime - beginTime
}

func echo2() int64 {
	beginTime := time.Now().UnixNano()
	fmt.Println(strings.Join(os.Args[1:], " "))
	endTime := time.Now().UnixNano()
	return endTime - beginTime
}
