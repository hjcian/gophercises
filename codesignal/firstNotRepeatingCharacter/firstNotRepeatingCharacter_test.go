package firstnotrepeatingcharacter

import "testing"

type suit struct {
	given string
	ans   string
}

func Test_firstNotRepeatingCharacter(t *testing.T) {
	testSuits := []suit{
		{"abacabad", "c"},
		{"abacabaabacaba", "_"},
		{"z", "z"},
		{"bcb", "c"},
	}
	for idx, testSuit := range testSuits {
		if got := firstNotRepeatingCharacter(testSuit.given); got != testSuit.ans {
			t.Errorf("#%v- func('%v'), want '%v', got '%v' \n", idx, testSuit.given, testSuit.ans, got)
		}
	}
}
