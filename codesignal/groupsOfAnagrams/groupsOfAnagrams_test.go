package groupsofanagrams

import "testing"

type suit struct {
	name         string
	input        []string
	expectResult int
}

func Test_sortString(t *testing.T) {
	t.Skip()
	sortString("tea")
	t.Errorf("print")
}

func Test_groupsOfAnagrams(t *testing.T) {
	tests := []suit{
		{
			"Test 1 ",
			[]string{"tea", "eat", "apple", "ate", "vaja", "cut", "java", "utc"},
			4,
		},
		{
			"Test 3",
			[]string{"tea", "eat", "eat", "ate", "aet", "eta", "eat", "tae", "tea"},
			1,
		},
	}
	for _, test := range tests {
		ret := groupsOfAnagrams(test.input)
		if ret != test.expectResult {
			t.Errorf("[%v] Got %v, Expect %v",
				test.name,
				ret,
				test.expectResult)
		}
	}
}
