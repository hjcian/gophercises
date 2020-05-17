package word2num

import "testing"

func Test_word2num(t *testing.T) {
	tests := []struct {
		name          string
		given         string
		expectResults int
	}{
		{
			name:          "test 1",
			given:         "one",
			expectResults: 1,
		},
		{
			name:          "test 2",
			given:         "One hundred Twenty three thousand, Four hundred and fifty-six",
			expectResults: 123456,
		},
		{
			name:          "test 3",
			given:         "One thousand, nine hundred and seventy-nine",
			expectResults: 1979,
		},
		{
			name:          "test 4",
			given:         "Fourteen",
			expectResults: 14,
		},
		{
			name:          "test 5",
			given:         "Two hundred and thirty-seven",
			expectResults: 237,
		},
		{
			name:          "test 6",
			given:         "Eleven",
			expectResults: 11,
		},
		{
			name:          "test 8",
			given:         "Nine hundred eighty seven million six hundred fifty-four thousand three hundred and twenty one",
			expectResults: 987654321,
		},
		{
			name:          "test 9",
			given:         "twenty one hundred",
			expectResults: 2100,
		},
		{
			name:          "test 10",
			given:         "one billion",
			expectResults: 1000000000,
		},
		{
			name:          "hidden 1",
			given:         "million",
			expectResults: 1000000,
		},
		{
			name:          "edge 1",
			given:         "one Hundred, seventY-four million, aNd twenty one HUNDRED twelve",
			expectResults: 174002112,
		},
		{
			name:          "edge 2",
			given:         "thirty one hundred thousand",
			expectResults: 3100000,
		},
		{
			name:          "edge 3",
			given:         "three million and one hundred thousand",
			expectResults: 3100000,
		},
	}

	for _, test := range tests {
		got := word2Num(test.given)
		if got != test.expectResults {
			t.Errorf("[%v] got %v, expect %v \n",
				test.name,
				got,
				test.expectResults,
			)
		}
	}
}
