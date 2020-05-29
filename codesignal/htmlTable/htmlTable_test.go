package htmltable

import (
	"fmt"
	"testing"
)

func Test_htmlTable(t *testing.T) {
	tests := []struct {
		name     string
		table    string
		row      int
		col      int
		expected string
	}{
		{
			"Test 1",
			"<table><tr><td>1</td><td>TWO</td></tr><tr><td>three</td><td>FoUr4</td></tr></table>",
			0,
			1,
			"TWO",
		},
		{
			"Test 2",
			"<table><tr><td>1</td><td>TWO</td></tr></table>",
			1,
			0,
			NoSuchCell,
		},
		{
			"Test 3",
			"<table><tr><td>7BusWMJ</td><td>D</td><td>5QPh9o</td></tr><tr><td>4Z</td><td>9Z</td><td>okc3</td></tr><tr><td>7mV88j</td><td>K358zV8</td><td>Y2vE</td></tr></table>",
			1,
			1,
			"9Z",
		},
		{
			"Test 7",
			"<table><tr><th>CIRCUMFERENCE</th><th>1</th><th>2</th><th>3</th><th>4</th><th>5</th><th>6</th></tr><tr><td>BITS</td><td>3</td><td>4</td><td>8</td><td>10</td><td>12</td><td>15</td></tr></table>",
			0,
			6,
			NoSuchCell,
		},
	}
	t.Run("maxcian", func(t *testing.T) {
		for _, test := range tests {
			fmt.Printf("[%v] -->\n", test.name)
			got := maxcian(test.table, test.row, test.col)
			if got != test.expected {
				t.Errorf("[%v] got '%v', expect '%v' \n",
					test.name,
					got,
					test.expected,
				)
			}
			fmt.Println()
		}
	})
}
