package main

import (
	"fmt"
)

func Complement(arr []int) (ret int) {
	for i, v := range arr {
		// fmt.Println(i, ^i)
		ret -= ^i + v
	}
	return
}

func Trapezoid(arr []int) int {
	height := len(arr)
	area := (1 + height) * height / 2
	for _, v := range arr {
		area -= v
	}
	return area
}

func missingNumber(arr []int) int {
	bucket := make([]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]] = 1
	}
	for i := 0; i < len(bucket); i++ {
		if bucket[i] == 0 {
			return i
		}
	}
	return 0
}

func main() {
	arr := []int{8, 6, 7, 0, 2, 5, 4, 3}
	v := missingNumber(arr)
	fmt.Println(arr)
	fmt.Println(v)
}
