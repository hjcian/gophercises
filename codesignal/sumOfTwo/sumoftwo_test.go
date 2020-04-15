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

func Test_sumOfTwo(t *testing.T) {
	ans := true
	v = 2
	if got := v3(a, b, v, 2); true {
		t.Errorf("got %v, need %v \n", got, ans)
	}
}
