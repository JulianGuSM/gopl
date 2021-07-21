package greetings

import (
	"errors"
	"fmt"
)

/**
 * @Author: sm.gu
 * @Author: sm.gu@aftership.com
 * @Date: 2021/7/20 3:31 下午
 * @Desc:
 */

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty  name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
