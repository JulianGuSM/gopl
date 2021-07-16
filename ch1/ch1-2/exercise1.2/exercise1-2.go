package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("idx: %d argument: %s\n", i, os.Args[i])
	}
}
