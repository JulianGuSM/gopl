package main

import (
	"encoding/json"
	"fmt"
)

/**
 * @Author: sm.gu
 * @Email: sm.gu@aftership.com
 * @Date: 2021/7/27 10:42 下午
 * @Desc:
 */

type Server struct {
	ServerName string
	ServerIp   string
	ServerPort int
}

func Serialize() {
	server := &Server{
		ServerName: "s1",
		ServerIp:   "0.0.0.0",
		ServerPort: 0,
	}
	byt, err := json.Marshal(server)
	if err != nil {
		fmt.Printf("marshal err: %v\n", err)
		return
	}
	fmt.Printf("serialize result: %s\n", string(byt))
}

func main() {
	Serialize()
}
