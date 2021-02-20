package main

import (
	"fmt"
)

func main() {
	var n, m, s int
	s = 0
	fmt.Print("input n: ")
	fmt.Scanln(&n)
	fmt.Print("input m: ")
	_, err := fmt.Scanln(&m)
	if err != nil {
		fmt.Println(err)
	}
	for i := 2; i <= n; i++ {
		s = (s + m) % i
		fmt.Printf("%d\n", s)
	}

	fmt.Printf("The winner is %d\n", s+1)
}
