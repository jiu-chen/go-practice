package main

import "fmt"

//是否是同构数
// 正整数n若是它平方数的尾部，则称n为同构数。
// IsomorphicNum ...
func IsomorphicNum(num int) bool {
	a := num % 10
	return Power(a, 2) == num
}
func main() {
	for i := 0; i < 10000; i++ {
		if IsomorphicNum(i) {
			fmt.Printf("%d 是同构数.", i)
		}
	}

}
