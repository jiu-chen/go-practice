package main

import (
	"fmt"
	"math"
)

// 判断是否是质数
// 一个数除了1和它本身，不可以被其他数整除
func isPrime(i int) bool {
	for j := 2; float64(j) <= math.Sqrt(float64(i)); j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}

func main() {
	var str []int
	j := 0
	for i := 2; i < 1000; i++ {
		if isPrime(i) {
			str = append(str, i)
			j++
		}
	}
	fmt.Printf("质数有: %d 个\n", j)
	fmt.Println(str)
}
