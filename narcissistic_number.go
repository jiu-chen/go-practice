package main

import (
	"fmt"
	"time"
)

// Power ...
func Power(m, n int) int {
	if n > 0 {
		return m * Power(m, n-1)
	}
	return 1
}


func isNarcissisticNum(num int) bool {
	i := num / 100
	k := (num / 10) % 10
	j := num % 10
	return Power(i, 3)+Power(j, 3)+Power(k, 3) == num
}

func main() {
	fmt.Printf("4的3次方: %d\n", Power(4, 3))
	st := time.Now()
	for i := 100; i < 1000; i++ {
		if isNarcissisticNum(i) {
			fmt.Printf("%d is narcissistic num.\n", i)
		}
	}
	fmt.Printf("time cost: %v\n", time.Since(st))
}
