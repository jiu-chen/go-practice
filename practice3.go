package main

import "fmt"

// IsPalindromic ... 回文数判断
func IsPalindromic(num int) bool {
	var temp int = num
	var res int = 0
	var d int
	for {
		d = num % 10
		num = num / 10
		res = res*10 + d
		if num == 0 {
			break
		}
	}
	fmt.Println("res is:", res)
	return temp == res
}

func main() {
	a, b := 12321, 112234
	fmt.Printf("a is palindromic: %v\n", IsPalindromic(a))
	fmt.Printf("b is palindromic: %v\n", IsPalindromic(b))
}
