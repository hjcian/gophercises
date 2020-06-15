package insertbits

import "testing"

type Input struct {
	n int
	a int
	b int
	k int
}

func Test_insertBits(t *testing.T) {
	tests := []struct {
		name string
		Input
		expected int
	}{
		{
			"Test1",
			Input{n: 1024, a: 1, b: 6, k: 17},
			1058,
		},
	}

	for _, test := range tests {
		got := insertBits(test.n, test.a, test.b, test.k)
		if got != test.expected {
			t.Errorf("got %v, expected %v", got, test.expected)
		}
	}
}
