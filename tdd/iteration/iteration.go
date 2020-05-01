package iteration

func Repeat(s string, n int) string {
	var ret string
	for i := 0; i < n; i++ {
		ret += s
	}
	return ret
}
