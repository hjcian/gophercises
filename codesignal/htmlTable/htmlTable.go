package htmltable

import (
	"strings"
)

// Useful constant strings
const (
	NoSuchCell = "No such cell"
)

func maxcian(table string, row int, column int) string {
	// fmt.Println("Given:", table)
	// fmt.Println("===== start =====")
	// defer fmt.Println("=====  end  =====")
	rows := strings.Split(table, "<tr>")
	rows = rows[1:]
	// fmt.Println("Target row:", row, "- split by <tr> we got", len(rows), "row(s):")
	// for i, r := range rows {
	// 	fmt.Println("	", i, r)
	// }
	if row+1 > len(rows) {
		return NoSuchCell
	}

	columns := strings.Split(rows[row], "<td>")
	columns = columns[1:]
	// fmt.Println("Target column:", column, "- split by <td> we got", len(columns), "column(s):")
	// for i, c := range columns {
	// 	fmt.Println("	", i, c)
	// }
	if column+1 > len(columns) {
		return NoSuchCell
	}

	end := strings.Index(columns[column], "</td>")
	return columns[column][:end]
}
