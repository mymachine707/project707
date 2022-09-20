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
