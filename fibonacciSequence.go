package main

import "fmt"

// FibonacciSequence ...
func FibonacciSequence(n int) int {
	if n <= 1 {
		return 1
	}
	return FibonacciSequence(n-1) + FibonacciSequence(n-2)
}

func main() {
	var m, n int
	fmt.Print("input m: ")
	fmt.Scanln(&m)
	fmt.Print("input n: ")
	fmt.Scanln(&n)
	fmt.Printf("第 %d 项是：%d\n", m, FibonacciSequence(m))
	fmt.Printf("第 %d 项是：%d\n", n, FibonacciSequence(n))
}
