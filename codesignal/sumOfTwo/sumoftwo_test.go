package sumoftwo

import (
	"math/rand"
	"testing"
)

var length = 1000

func makeRandIntSlice() []int {
	a := make([]int, length, length)
	for i := 0; i < length; i++ {
		a = append(a, rand.Intn(length))
	}
	return a
}

var a = makeRandIntSlice()
var b = makeRandIntSlice()
var v = -1

func BenchmarkV1(gb *testing.B) {
	gb.SkipNow()
	a := makeRandIntSlice()
	b := makeRandIntSlice()
	v := -1
	for i := 0; i < gb.N; i++ {
		v1(a, b, v)
	}
}

func BenchmarkV2(gb *testing.B) {
	for i := 0; i < gb.N; i++ {
		v2(a, b, v)
	}
}

func BenchmarkV3N2(gb *testing.B) {
	for i := 0; i < gb.N; i++ {
		v3(a, b, v, 2)
	}
}

func BenchmarkV3N4(gb *testing.B) {
	for i := 0; i < gb.N; i++ {
		v3(a, b, v, 4)
	}
}

func BenchmarkV3N8(gb *testing.B) {
	for i := 0; i < gb.N; i++ {
		v3(a, b, v, 8)
	}
}

func BenchmarkV3N16(gb *testing.B) {
	for i := 0; i < gb.N; i++ {
		v3(a, b, v, 16)
	}
}

func Test_sumOfTwo(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	tests := []struct {
		v   int
		ans bool
	}{
		{9, true},
		{10, false},
	}
	for i, test := range tests {
		if got := v2(a, b, test.v); got != test.ans {
			t.Errorf("[test-%v] v2: got %v, need %v \n", i, got, test.ans)
		}
		if got := v3(a, b, test.v, 2); got != test.ans {
			t.Errorf("[test-%v] v3: got %v, need %v \n", i, got, test.ans)
		}
	}

}
