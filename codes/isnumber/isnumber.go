package isnumber

import "unicode"

// entered slice and validate in element
func Isnumber(num []string) bool {
	var answer bool
	for i := 0; i < len(num); i++ {
		if IsnumberStr(num[i]) {
			answer = true
		} else if !IsnumberStr(num[i]) {
			panic("This string is not a number!")
		}
	}
	return answer
}

// entered string and validate
func IsnumberStr(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
