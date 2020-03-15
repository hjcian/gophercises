package main

import "fmt"

// 65~90
// 97~122
// Complete the caesarCipher function below.
func caesarCipher(s string, k int32) string {
	var ret string
	k = k % 26
	for _, c := range s {
		if c >= 65 && c <= 90 {
			// Uppercase
			if c+k > 90 {
				ret += string(c + k - 26)
			} else {
				ret += string(c + k)
			}
		} else if c >= 97 && c <= 122 {
			// Lowercase
			if c+k > 122 {
				ret += string(c + k - 26)
			} else {
				ret += string(c + k)
			}
		} else {
			// others
			ret += string(c)
		}
	}
	return ret
}

func main() {
	in := "Alw"
	var k int32 = 5
	out := caesarCipher(in, k)
	fmt.Println(out)
}
