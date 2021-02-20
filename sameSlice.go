package main

import "fmt"

func IsSameSlice(m, n []int) bool {
	if len(m) != len(n) {
		return false
	}
	for ix := range m {
		if m[ix] != n[ix] {
			return false
		}
	}
	return true
}

func main() {
	var m []int = []int{1, 2, 3, 4}
	var n []int = []int{1, 2, 4, 3}
	var l []int = []int{1, 2, 4, 3}
	fmt.Printf("m is same with n: %t\n", IsSameSlice(m, n))
	fmt.Printf("n is same with l: %t\n", IsSameSlice(l, n))
}
