package main

import (
	"fmt"
	"sort"
)

// Factorial ...
func Factorial(count int) map[int]int {
	m := make(map[int]int)
	for i := 0; i <= count; i++ {
		if i == 0 {
			m[i] = 1
		} else {
			m[i] = m[i-1] * i
		}
	}
	arr := make([]int, 0)
	for k := range m {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i], "的阶乘是", m[arr[i]])
	}
	return m
}

func main() {
	num := 20
	Factorial(num)
}
