package arrayandslice

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	ret := make([]int, len(slices), len(slices))
	for i, slice := range slices {
		ret[i] = Sum(slice)
	}
	return ret
}
