package main

import "fmt"

func reverse(s []int) []int {
	// lens := len(s)
	// for i := 0; i < lens/2; i++ {
	// 	s[i], s[lens-i-1] = s[lens-i-1], s[i]
	// }
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
	return s
}

func main() {
	var s []int = []int{1, 2, 3, 4, 5, 6, 7}
	newS := reverse(s)
	fmt.Println("new slice: ", newS)
}
