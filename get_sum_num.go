package main

import "fmt"

func main() {
	var s []int
	s = make([]int, 50)
	for i := 0; i < 50; i++ {
		s[i] = 2 * i
	}
	fmt.Println("s: ", s)

	var sum = 100
	fmt.Println("sum is 100 pairs: ", GetSumCount(s, sum))
}

func GetSumCount(arr []int, sum int) int {
	var tMap map[int]int = make(map[int]int)
	var count int = 0
	for i := 0; i < len(arr); i++ {
		if _, ok := tMap[sum-arr[i]]; ok {
			count++
		} else {
			tMap[arr[i]] = 1
		}
	}
	return count
}
