package example

import (
	"fmt"
	"strings"
)

/*
	Map example
    Map/Reduce/Filter只是一种控制逻辑，真正的业务逻辑是在传给他们的数据和那个函数来定义的。
*/

func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArr []string
	for _, str := range arr {
		newArr = append(newArr, fn(str))
	}
	return newArr
}

func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArr []int
	for _, str := range arr {
		newArr = append(newArr, fn(str))
	}
	return newArr
}

func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, str := range arr {
		sum += fn(str)
	}
	return sum
}

func Filter(arr []int, fn func(n int) bool) []int {
	var newArr []int
	for _, item := range arr {
		if fn(item) {
			newArr = append(newArr, item)
		}
	}
	return newArr
}

func runMapStrToStr() {
	strInputs := []string{"Hello", "world", "oh", "yes"}

	strOutputs := MapStrToStr(strInputs, func(s string) string {
		return strings.ToUpper(s)
	})

	fmt.Printf("strOutputs = %v\n", strOutputs)
}

func runMapStrToInt() {
	strInputs := []string{"Hello", "world", "oh", "yes"}

	strOutputs := MapStrToInt(strInputs, func(s string) int {
		return len(s)
	})

	fmt.Printf("strOutputs = %v\n", strOutputs)
}

func runReduce() {
	strInputs := []string{"Hello", "world", "oh", "yes"}

	reduceOut := Reduce(strInputs, func(s string) int {
		return len(s)
	})

	fmt.Printf("reduceOut = %v\n", reduceOut)
}

func runFilter() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	output := Filter(input, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("ouput = %v\n", output)
}
