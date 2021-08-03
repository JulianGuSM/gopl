package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
 * @Author: sm.gu
 * @Email: sm.gu@aftership.com
 * @Date: 2021/8/2 6:01 下午
 * @Desc:
 */

func main() {
	resp, err := http.Get("http://localhost:9080/test")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close() //ok, most of the time :-)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
