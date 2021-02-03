package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var i, j int
	for i = 1; i <= 9; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
	tc := time.Since(start)
	fmt.Printf("cost time: %v\n", tc)
}
