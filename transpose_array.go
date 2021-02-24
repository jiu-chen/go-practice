package main

import "fmt"

func transpose(a [][]int) [][]int {
	col := len(a)
	row := len(a[0])
	if row == 0 {
		return nil
	}
	arr := make([][]int, row, col)
	for i := 0; i < row; i++ {
		arr[i] = make([]int, col)
	}
	for m := 0; m < row; m++ {
		for n := 0; n < col; n++ {
			arr[m][n] = a[n][m]
		}
	}
	return arr
}

func display(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	var arr [][]int = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("reversed arr: ", transpose(arr))
	display(transpose(arr))
}
