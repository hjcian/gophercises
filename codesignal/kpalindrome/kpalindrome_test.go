package kpalindrome

import "testing"

func Test_sumOfTwo(t *testing.T) {
	tests := []struct {
		s   string
		k   int
		ans bool
	}{
		{"pqwuzifwovyddwyvvbu", 16, true},
		{"sjolvybwxxttqogn", 16, true},
		{"glzamzp", 14, true},
		{"abrarbra", 1, true},
		{"xabcdcbay", 2, true},
		{"xazbcdcbay", 2, false},
		{"adbcdbacdb", 2, false},
	}
	for i, test := range tests {
		if got := kpalindrome(test.s, test.k); got != test.ans {
			t.Errorf("[%v] expect %v but got %v \n", i, test.ans, got)
		}
	}
}
