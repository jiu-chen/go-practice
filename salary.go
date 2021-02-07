package main

import "fmt"

func main() {
	var n, salary, sum, a int
	var err error
	fmt.Print("输入工龄: ")
	fmt.Scanln(&n)
	fmt.Print("输入薪资: ")
	_, err = fmt.Scanln(&salary)
	if err != nil {
		fmt.Println(err)
	}
	if n > 0 && n < 1 {
		a = 200
	} else if n >= 1 && n < 3 {
		a = 500
	} else if n >= 3 && n < 5 {
		a = 1000
	} else if n >= 5 && n < 10 {
		a = 2500
	} else if n >= 10 && n < 15 {
		a = 5000
	}
	sum = salary + a
	fmt.Printf("目前工作了%d年， 基本工资为%d, 应涨工资%d, 涨后工资%d.\n", n, salary, a, sum)

}
