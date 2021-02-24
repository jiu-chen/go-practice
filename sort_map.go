package main

import (
	"fmt"
	"sort"
)

// SortMap ...
func SortMap(student map[string]int) []string {
	var names []string
	for name := range student {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s的年龄是%d\n", name, student[name])
	}
	return names
}

func main() {
	student := map[string]int{
		"lisa":  21,
		"tom":   22,
		"jack":  23,
		"stone": 24,
	}
	res := SortMap(student)
	fmt.Println("res: ", res)
}
