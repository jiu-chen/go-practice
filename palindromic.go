package main

import (
	"fmt"
	"strconv"
)

// IsPalindromic ... 回文数判断
func IsPalindromic(num int) bool {
	temp, res := num, 0
	for num != 0 {
		d := num % 10
		num = num / 10
		res = res*10 + d
	}
	return temp == res
}

// IsPalindromic2 ...
func IsPalindromic2(a string) bool {
	var b bool = true
	lens := len(a) - 1
	for i := 0; i <= lens/2-1; i++ {
		if a[i] != a[lens-i] {
			b = false
			break
		}
	}
	return b
}

func main() {
	var a int
	fmt.Scanln(&a)
	fmt.Print("input a: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("a is palindromic: %v\n", IsPalindromic(a))
	fmt.Printf("a is palindromic2: %v\n", IsPalindromic2(strconv.Itoa(a)))
}
