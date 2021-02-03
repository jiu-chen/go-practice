package main

import "fmt"

// IsPalindromic ... 回文数判断
func IsPalindromic(num int) bool {
	var temp int = num
	var res, d int = 0, 0
	for num/10 != 0 {
		d = num % 10
		num = num / 10
		res = d*10 + res
	}
	return temp == res
}

func main() {
	a, b := 12321, 112234
	fmt.Printf("a is palindromic: %v\n", IsPalindromic(a))
	fmt.Printf("b is palindromic: %v\n", IsPalindromic(b))
}
