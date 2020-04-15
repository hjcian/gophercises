package kpalindrome

func kpalindrome(s string, k int) bool {
	if k >= len(s)-1 {
		return true
	}

	if len(s) <= 1 {
		return true
	}

	if s[0] == s[len(s)-1] {
		return kpalindrome(s[1:len(s)-1], k) // no need to waste k, reduce problem size
	}

	if k == 0 {
		return false
	}

	return kpalindrome(s[0:len(s)-1], k-1) || kpalindrome(s[1:len(s)], k-1) // need to use k
}
