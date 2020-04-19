package arraypacking

// https://app.codesignal.com/challenge/7YX8Fcyhxo62wBXoM
func arrayPacking(a []int) int {
	// num := 2
	// fmt.Println(strconv.FormatInt(int64(num), 2))
	total := 0
	for shift, v := range a {
		total += v << (8 * shift)
	}
	return total
}
