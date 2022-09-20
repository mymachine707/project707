package sign

import (
	"project/isnumber"
	"strings"
)

func Signnumbers(num string) bool {
	a := strings.Split(num, "")
	//
	var negative bool

	if a[0] == "+" {
		negative = false
	} else if a[0] == "-" {
		negative = true
	} else if isnumber.Isnumber((a[0])) {
		negative = false
	} else {
		panic("This string is not a number!")
	}

	return negative
}

func SignValidation(sign string) string {
	var answer string
	switch sign {
	case "+":
		answer = sign
	case "-":
		answer = sign
	case "*":
		answer = sign
	case "/":
		answer = sign
	case "%":
		answer = sign
	default:
		panic("This sign is not defined!")
	}
	return answer
}
