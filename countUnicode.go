package main

import "fmt"

// CountUnicode ...
func CountUnicode(s string) {
	lens := len(s)
	var num, letter, space, other int
	for i := 0; i < lens; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num++
		} else if s[i] >= 'A' && s[i] <= 'z' {
			letter++
		} else if s[i] == ' ' {
			space++
		} else {
			other++
			fmt.Printf("other char: %d\n", s[i])
		}
	}

	fmt.Printf("number count: %d, letter count: %d, other chars count: %d\n", num, letter, other)
}

func main() {
	var s = "hello world 999\ta\n12\n4"
	CountUnicode(s)
}
