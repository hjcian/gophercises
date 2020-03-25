package main

import (
	"fmt"
)

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
