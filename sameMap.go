package main

import "fmt"

func isSameMap(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xValue := range x {
		if yValue, ok := x[k]; !ok || yValue != xValue {
			return false
		}
	}
	return true
}

func main() {
	x := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
	}
	y := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 5,
	}
	z := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	fmt.Printf("z is equal with x: %v\n", isSameMap(x, z))
	fmt.Printf("y is equal with x: %v\n", isSameMap(x, y))

}
