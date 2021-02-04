package main

import (
	"fmt"
	"time"
)

func isNarcissisticNum(num int) bool {
	i := num / 100
	k := (num / 10) % 10
	j := num % 10
	return i*i*i+j*j*j+k*k*k == num
}

func main() {
	st := time.Now()
	for i := 100; i < 1000; i++ {
		if isNarcissisticNum(i) {
			fmt.Printf("%d is narcissistic num.\n", i)
		}
	}
	fmt.Printf("time cost: %v\n", time.Since(st))
}
