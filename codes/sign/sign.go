package sign

import (
	"project/isnumber"
	"strings"
)

func Signnumbers(num []string) bool {

	//
	var negative bool

	if num[0] == "+" {
		negative = false
	} else if num[0] == "-" {
		negative = true
	} else if isnumber.IsnumberStr((num[0])) {
		negative = false
	} else {
		panic("This string is not num number!")
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

func AddtoMinus(num string) string {
	var b []string
	b = append(b, "-")
	b = append(b, num)
	var answer string
	for i := 0; i < len(b); i++ {
		answer += b[i]
	}
	return answer

}

func Findtobig(num1, num2 string) int {

	a := strings.Split(num1, "")
	b := strings.Split(num2, "")

	var e int // a katta b dan
	var c int // a kichik b dan
	// qaysi son kattaligini tekshirish kerak
	if len(a) > len(b) {
		// a katta b dan
		e += 1
	} else if len(a) < len(b) {
		// a kichik b dan
		c += 1
	} else if len(a) == len(b) {
		// a teng b ga
		for i := 0; i < len(a); i++ {
			if a[i] > b[i] {
				// a katta b dan
				e += 1
				break
			} else if a[i] < b[i] {
				// a kichik b dan
				c += 1
				break
			}

		}

	}

	var answer int
	if e == 1 {
		answer = 1
	} else if c == 1 {
		answer = 2
	}
	return answer
}
