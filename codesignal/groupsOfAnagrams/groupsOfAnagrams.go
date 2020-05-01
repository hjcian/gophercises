package groupsofanagrams

import (
	"sort"
)

type sortRunes []rune

func (s sortRunes) Len() int {
	return len(s)
}
func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortString(s string) string {
	r := []rune(s) // need this, if just pass string into sortRunes(), the result won't be saved, why?
	sort.Sort(sortRunes(r))
	return string(r)
}

func groupsOfAnagrams(words []string) int {
	set := make(map[string]struct{})
	for _, w := range words {
		sorted := sortString(w)
		set[sorted] = struct{}{}
	}
	return len(set)
}
