package main

import (
	"fmt"
)

func factorialMod(n int64) int64 {
	if n == 1 || n == 0 {
		return 1
	}
	ret := int64(1)
	for i := int64(2); i <= n; i++ {
		ret = ret * i
		for {
			if ret%10 == 0 {
				ret = ret / 10
			} else {
				ret = ret % 10000
				break
			}
		}
	}
	return ret
}

func lastDigitDiffZero(n int64) int {
	v := factorialMod(n)
	return int(v) % 10
}

func factorial(n int64) int64 {
	if n == 1 || n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("Hello, playground")
	for i := 0; i < 20; i++ {
		v := factorial(int64(i))
		v2 := lastDigitDiffZero(int64(i))
		fmt.Printf("%v! = %v vs. %v\n", i, v, v2)
	}
}
