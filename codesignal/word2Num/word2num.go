package word2num

import (
	"math"
	"strings"
)

type NumOrder string

const (
	hundred  NumOrder = "hundred"
	thousand NumOrder = "thousand"
	million  NumOrder = "million"
	billion  NumOrder = "billion"
	trillion NumOrder = "trillion"
)

func (n NumOrder) Order() int {
	switch n {
	case thousand:
		return 3
	case million:
		return 6
	case billion:
		return 9
	case trillion:
		return 12
	default:
		return 0
	}
}

var small2num = map[string]int{
	"one":       1,
	"two":       2,
	"three":     3,
	"four":      4,
	"five":      5,
	"six":       6,
	"seven":     7,
	"eight":     8,
	"nine":      9,
	"ten":       10,
	"eleven":    11,
	"twelve":    12,
	"thirteen":  13,
	"fourteen":  14,
	"fifteen":   15,
	"sixteen":   16,
	"seventeen": 17,
	"eighteen":  18,
	"nineeen":   19,
	"twenty":    20,
	"thirty":    30,
	"fourty":    40,
	"fifty":     50,
	"sixty":     60,
	"seventy":   70,
	"eighty":    80,
	"ninety":    90,
}

func compute(str string) int {
	str = strings.TrimSpace(str)
	digits := strings.Split(str, " ")
	sum := 0
	for _, v := range digits {
		v = strings.TrimSpace(v)
		if len(v) > 0 {
			sum += small2num[v]
		}
	}
	return sum
}

func class2num(str string) int {
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, "and", "")
	str = strings.ReplaceAll(str, "-", " ")
	str = strings.ReplaceAll(str, ",", " ")

	hundredPos := strings.Index(str, string(hundred))

	tensDigit := 0
	hundredsDigit := 0
	if hundredPos != -1 {
		hundredsDigit = compute(str[:hundredPos]) * 100
		tensDigit = compute(str[hundredPos+len(hundred):])
	} else {
		tensDigit = compute(str)
	}
	sum := hundredsDigit + tensDigit
	return sum
}

func findNumClass(str string, numorders []NumOrder) int {
	if len(numorders) == 0 {
		return class2num(str)
	}

	numClass := numorders[0] // return trillion or billion or ...

	sum := 0
	classPos := strings.Index(str, string(numClass))
	if classPos != -1 {
		ret := class2num(str[:classPos])
		if ret == 0 {
			ret = 1
		}
		sum = ret * int(math.Pow10(numClass.Order()))
		str = str[len(numClass)+classPos:]
	}
	return sum + findNumClass(str, numorders[1:])
}

func word2Num(str string) int {
	str = strings.ToLower(str)

	sum := findNumClass(str, []NumOrder{
		trillion,
		billion,
		million,
		thousand,
	})
	return sum
}
