package main

import "fmt"

func main() {
	fmt.Print("input days: ")
	var n int
	fmt.Scanln(&n)
	res := 11
	for i := n; i > 0; i-- {
		res = (res*(i+1) + 1) / i
	}
	fmt.Println(res)
}
