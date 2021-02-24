package main

import "fmt"

func eat(n int) int {
	if n == 1 {
		return 1
	}
	return 2*eat(n-1) + 2
}

func main() {
	fmt.Print("input days: ")
	var n int
	fmt.Scanln(&n)
	fmt.Printf("第 %d 还有1个，则一共有 %d 个桃子。", n, eat(n))
}
