package firstnotrepeatingcharacter

func firstNotRepeatingCharacter(s string) string {
	appearCount := make(map[rune]int)
	for _, c := range s {
		appearCount[c]++ // it is OK if key not found
	}
	for _, c := range s {
		if appearCount[c] == 1 {
			return string(c)
		}
	}
	return "_"
}
