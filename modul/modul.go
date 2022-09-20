package modul

import (
	"project/isnumber"
	"strings"
)

func Set(num1 string) string {

	a := strings.Split(num1, "")

	var b []string
	if a[0] == "+" {
		b = a[1:]
	} else if a[0] == "-" {
		b = a[1:]
	} else if isnumber.Isnumber((a[0])) {
		b = a
	} else {
		panic("This string is not a number!")
	}

	var answer string

	for i := 0; i < len(b); i++ {
		answer += b[i]
	}

	return answer

}
