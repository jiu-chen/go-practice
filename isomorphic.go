package main

import (
	"fmt"
	"strconv"
)

// 判断一个数 是否是同构数
// 正整数n若是它平方数的尾部，则称n为同构数。

// IsomorphicNum ...
func IsomorphicNum(num int) bool {
	lens := len(strconv.Itoa(num))
	r := Power(num, 2)
	res := r % (Power(10, lens))
	return res == num
}

func main() {
	for i := 0; i < 10000; i++ {
		if IsomorphicNum(i) {
			fmt.Printf("%d 是同构数.", i)
		}
	}
	fmt.Println()
}
