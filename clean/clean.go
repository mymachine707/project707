package clean

import (
	"project/isnumber"
	"strings"
)

// isnumber.Isnumber
func Clean(num string) string {
	a := strings.Split(num, "")
	//

	if a[0] == "+" {
		a = a[1:]
	} else if a[0] == "-" {
		a = a[1:]
	} else if !isnumber.Isnumber(a[0]) {
		panic("This string is not a number!")
	}

	for i := 1; i < len(a); i++ {
		if !isnumber.Isnumber(a[i]) {
			panic("This string is not a number!")
		}
	}

	var b []string
	count := 0

	for i := 0; i < len(a); i++ {

		if a[i] == "0" {
			count++
		} else {
			break
		}
	}
	b = a[count:]

	var answer string

	for i := 0; i < len(b); i++ {
		answer += b[i]
	}

	return answer
} //-end function
