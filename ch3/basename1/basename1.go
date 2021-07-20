package main

/**
 * @Author: sm.gu
 * @Author: sm.gu@aftership.com
 * @Date: 2021/7/19 1:42 下午
 * @Desc:
 */

func main() {

}

func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
