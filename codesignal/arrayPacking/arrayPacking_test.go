package arraypacking

import "testing"

func Test_arrayPackin(t *testing.T) {
	tests := []struct {
		array []int
		ans   int
	}{
		{[]int{24, 85, 0}, 21784},
		{[]int{23, 45, 39}, 2567447},
	}
	for i, test := range tests {
		got := arrayPacking(test.array)
		if got != test.ans {
			t.Errorf("[Q%v] Expect %v, got %v \n", i, test.ans, got)
		}
	}
}
