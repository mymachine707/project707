package converter

import "strconv"

// stringdan intga o'tqizadi
func Str_Int(str string) int {
	intVar, err := strconv.Atoi(str)
	if err != nil {
		return 0 // sho'tda tanish bilishchili bo'b ketti err bersa 0 qayatrmasligi kere boshqa javob so'ra
	}
	return intVar
}

// intdan stringga o'tkizadi

func Int_str(num int) string {
	str := strconv.Itoa(num)
	return str
}
