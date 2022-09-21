package main

import (
	"fmt"
	"project/add"
	"project/converter"
	"project/isnumber"
	"project/sign"
	"strings"
	"unicode"
)

func main() {
	var num1 string = "984564564564654654654567"
	var num2 string = "1000545646546545641650"
	var sign string = "+"
	var setnum string = "-0000000000002000"

	fmt.Printf("answer is add = %s\n", Add(num1, num2))                                             // add
	fmt.Printf("answer is sub = %s\n", Minus(num1, num2))                                           // sub
	fmt.Printf("answer is multiply = %s\n", Multiply(num1, num2))                                   // multiply
	fmt.Printf("answer is add to minus = %s, second num =%s\n", AddtoMinus(num1), AddtoMinus(num2)) // add minus
	fmt.Printf("answer is sign = %s\n", SignValidation(sign))                                       // sign
	fmt.Printf("answer is set = %s\n", Set(setnum))
	fmt.Printf("answer is sign = %s\n", Clean(setnum))
	fmt.Printf("answer is number = %v\n", IsnumberStr(num2))
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

func Clean(num string) string {
	a := strings.Split(num, "")
	//

	if a[0] == "+" {
		a = a[1:]
	} else if a[0] == "-" {
		a = a[1:]
	} else if !isnumber.IsnumberStr(a[0]) {
		panic("This string is not a number!")
	}
	// o'zgartirish kiritish kerak!
	for i := 1; i < len(a); i++ {
		if !isnumber.IsnumberStr(a[i]) {
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

func Set(num1 string) string {

	a := strings.Split(num1, "")

	var b []string
	if a[0] == "+" {
		b = a[1:]
	} else if a[0] == "-" {
		b = a[1:]
	} else if isnumber.IsnumberStr((a[0])) {
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

func Multiply(a, b string) string {
	var k []string

	var e int // a katta b dan
	var c int // a kichik b dan
	// qaysi son kattaligini tekshirish kerak
	if len(a) >= len(b) {
		// a katta b dan
		e += 1

	} else if len(a) < len(b) {
		// a kichik b dan
		c += 1

	}
	var ans string
	if e == 1 {
		f := strings.Split(b, "")
		for i := len(f); i > 0; i-- {
			ans = mult(a, f[i-1])
			for l := 0; l < len(f)-i; l++ {
				ans += "0"
			}

			k = append(k, ans)
		}

	} else if c == 1 {
		f := strings.Split(a, "")
		for i := len(f); i > 0; i-- {
			ans = mult(b, f[i-1])
			for l := 0; l < len(f)-i; l++ {
				ans += "0"
			}
			k = append(k, ans)

		}

	}

	var answers string

	for i := 0; i < len(k); i++ {
		answers = add.Add(answers, k[i])
	}

	return answers
}

// katta xonali sonni bir xonali songa ko'paytirish

func mult(w, b string) string {

	a := strings.Split(w, "")

	var l []string = a

	var p int
	var m int = 0
	for i := len(a); i > 0; i-- {
		p = converter.Str_Int(a[i-1]) * converter.Str_Int(b)
		if i-1 == 0 {
			p = converter.Str_Int(a[i-1]) * converter.Str_Int(b)
			p += m
			m = 0
			l[i-1] = converter.Int_str(p)
		} else if p < 10 {
			if p+m >= 10 {
				p += m
				m = 0
				q := strings.Split(converter.Int_str(p+m), "")
				l[i-1] = q[1]
				m = converter.Str_Int(q[0])
			} else {
				l[i-1] = converter.Int_str(p)
			}
		} else if p >= 10 {
			p += m
			m = 0
			q := strings.Split(converter.Int_str(p), "")
			l[i-1] = q[1]
			m = converter.Str_Int(q[0])
		}
	}
	answer := ""

	for i := 0; i < len(l); i++ {
		answer += l[i]
	}
	return answer
}

func Minus(num1, num2 string) string {
	a := strings.Split(num1, "")
	b := strings.Split(num2, "")

	// kiruvchi stringlarni uzunligini tenglashtirish======================

	var s []string
	var y int

	var e int // a katta b dan
	var c int // a kichik b dan

	if sign.Findtobig(num1, num2) == 1 {
		e = 1
	} else if sign.Findtobig(num1, num2) == 2 {
		c = 1
	}

	// slicelarni bir biriga tenglashtirish jarayoni bu kod o'zgartirilmasin

	if e == 1 { // a katta b dan
		y = len(a) - len(b)
		for i := 0; i < y; i++ {
			s = append(s, "0")
		}
		s = append(s, b...)
		b = s
		s = nil

	} else if c == 1 { // a kichik b dan
		y = len(b) - len(a)
		for i := 0; i < y; i++ {
			s = append(s, "0")
		}
		s = append(s, a...)
		a = s
		s = nil
	}
	//==========================

	// asosiy shartlar
	var z []string

	for i := 0; i < len(a); i++ {
		z = append(z, a[i])
	}

	//a bilan b berilgan slice holatida
	if e == 1 {
		fmt.Println("\nnumber1 > number 2")
		// a > b bo'lsa:
		for i := len(a); i > 0; i-- {
			if converter.Str_Int(a[i-1])-converter.Str_Int(b[i-1]) > 0 { // Agar joriy indekslar ayirmasi musbat bo'lsa
				z[i-1] = converter.Int_str(converter.Str_Int(a[i-1]) - converter.Str_Int(b[i-1])) // z joriy indeksiga ayirma qiymati beriladi
			} else if converter.Str_Int(a[i-1])-converter.Str_Int(b[i-1]) == 0 { // Agar joriy indekslar ayirmasi 0 ga teng bo'lsa
				z[i-1] = converter.Int_str(0) // z joriy indeksiga 0 qiymat beriladi
			} else if converter.Str_Int(a[i-1])-converter.Str_Int(b[i-1]) < 0 { // agar joriy indekslar ayirmasi 0 dan kichik bo'lsa
				a[i-1] = converter.Int_str(converter.Str_Int(a[i-1]) + 10)                        // katt sonning joriy indeksiga 10 soni qo'shiladi va qaytadan yenglashtiriladi
				z[i-1] = converter.Int_str(converter.Str_Int(a[i-1]) - converter.Str_Int(b[i-1])) // z ning joriy qiymatini sonlarning joriy qiymatlari ayirmasiga tenglashtiriladi va...
				if converter.Str_Int(a[i-2]) > 0 {                                                // agar katta sonning joriy qiymatidan bitta oldingi qiymati 0 dan katta bo'lsa
					a[i-2] = converter.Int_str(converter.Str_Int(a[i-2]) - 1) // katta sonning joriy qiymatidan bitta oldingi sonnini bittaga kamaytiriladi
				} else if converter.Str_Int(a[i-2]) == 0 { // agar katta sonning joriy qiymatidan bitta oldingi qiymati 0 ga teng bo'lsa...
					a[i-2] = "9"           // katta sonning joriy qiymatidan bitta oldingi sonnini 9 soniga tenglashtiriladi
					var boolen bool = true // for loop uchun boolen qiyamtini true holatida ochildi
					var m int = 3          // indekslarni tekshirish uchun m int shaklida ochildi va katta sonning joriy qiymatidan ikkita oldingi qiymatini tekshirish uchun m sonini 3 ga tenglandi
					for boolen {           // for loop ochildi
						if converter.Str_Int(a[i-m]) > 0 { // Agar katta sonning joriy indeksidan ikkita oldingi indeksi 0 dan katta bo'lsa
							a[i-m] = converter.Int_str(converter.Str_Int(a[i-m]) - 1) // katta sonning joriy indeksidan ikkita oldingi indeksidan 1 soni ayiriladi
							boolen = false                                            // boolen false qiymatini oladi
							break                                                     // for loopdan chiqiladi
						} else if converter.Str_Int(a[i-m]) == 0 { // Agar katta sonning joriy indeksidan ikkita oldingi indeksi 0 ga teng bo'lsa
							a[i-m] = "9" // Katta sonning joriy indeksidan ikkita oldingi indeksi "9" ga tenglanadi
							m++          // m soni 1 ga oshiriladi va for loop qaytadan tekshiruvni boshlaydi
						}

					}
				}
			}

		}

	} else if c == 1 {
		fmt.Println("\nnumber1 < number 2")
		// a < b bo'lsa:
		for i := len(b); i > 0; i-- {
			if converter.Str_Int(b[i-1])-converter.Str_Int(a[i-1]) > 0 { // Agar joriy indekslar ayirmasi musbat bo'lsa
				z[i-1] = converter.Int_str(converter.Str_Int(b[i-1]) - converter.Str_Int(a[i-1])) // z joriy indeksiga ayirma qiymati beriladi
			} else if converter.Str_Int(b[i-1])-converter.Str_Int(a[i-1]) == 0 { // Agar joriy indekslar ayirmasi 0 ga teng bo'lsa
				z[i-1] = converter.Int_str(0) // z joriy indeksiga 0 qiymat beriladi
			} else if converter.Str_Int(b[i-1])-converter.Str_Int(a[i-1]) < 0 { // agar joriy indekslar ayirmasi 0 dan kichik bo'lsa
				b[i-1] = converter.Int_str(converter.Str_Int(b[i-1]) + 10)                        // katt sonning joriy indeksiga 10 soni qo'shiladi va qaytadan yenglashtiriladi
				z[i-1] = converter.Int_str(converter.Str_Int(b[i-1]) - converter.Str_Int(a[i-1])) // z ning joriy qiymatini sonlarning joriy qiymatlari ayirmasiga tenglashtiriladi va...
				if converter.Str_Int(b[i-2]) > 0 {                                                // agar katta sonning joriy qiymatidan bitta oldingi qiymati 0 dan katta bo'lsa
					b[i-2] = converter.Int_str(converter.Str_Int(b[i-2]) - 1) // katta sonning joriy qiymatidan bitta oldingi sonnini bittaga kamaytiriladi
				} else if converter.Str_Int(b[i-2]) == 0 { // agar katta sonning joriy qiymatidan bitta oldingi qiymati 0 ga teng bo'lsa...
					b[i-2] = "9"           // katta sonning joriy qiymatidan bitta oldingi sonnini 9 soniga tenglashtiriladi
					var boolen bool = true // for loop uchun boolen qiyamtini true holatida ochildi
					var m int = 3          // indekslarni tekshirish uchun m int shaklida ochildi va katta sonning joriy qiymatidan ikkita oldingi qiymatini tekshirish uchun m sonini 3 ga tenglandi
					for boolen {           // for loop ochildi
						if converter.Str_Int(b[i-m]) > 0 { // Agar katta sonning joriy indeksidan ikkita oldingi indeksi 0 dan katta bo'lsa
							b[i-m] = converter.Int_str(converter.Str_Int(b[i-m]) - 1) // katta sonning joriy indeksidan ikkita oldingi indeksidan 1 soni ayiriladi
							boolen = false                                            // boolen false qiymatini oladi
							break                                                     // for loopdan chiqiladi
						} else if converter.Str_Int(b[i-m]) == 0 { // Agar katta sonning joriy indeksidan ikkita oldingi indeksi 0 ga teng bo'lsa
							b[i-m] = "9" // Katta sonning joriy indeksidan ikkita oldingi indeksi "9" ga tenglanadi
							m++          // m soni 1 ga oshiriladi va for loop qaytadan tekshiruvni boshlaydi
						}

					}
				}
			}

		}
	}
	//natijani birlashtirish
	var answers string
	for i := 0; i < len(z); i++ {
		answers += z[i]
	}

	//==========================
	return answers
}

func Add(num1, num2 string) string {
	a := strings.Split(num1, "")
	b := strings.Split(num2, "")
	// kiruvchi stringlarni uzunligini tenglashtirish======================

	var s []string
	var y int

	if len(a) > len(b) {
		y = len(a) - len(b)
		for i := 0; i < y; i++ {
			s = append(s, "0")
		}
		s = append(s, b...)
		b = s
		s = nil
	} else {
		y = len(b) - len(a)
		for i := 0; i < y; i++ {
			s = append(s, "0")
		}
		s = append(s, a...)
		a = s
		s = nil
	}

	//stringlarni slice holatida har bir elementini qo'shib yangi slice hosil qilish

	var answer string // javobni ansverga biriktirish
	var k []int

	var m int
	for i := 0; i < len(a); i++ {
		z := converter.Str_Int(a[i]) + converter.Str_Int(b[i])
		k = append(k, z)

		for i := len(k); i > 0; i-- {
			if k[i-1] > 10 {
				if i-1 > 0 {
					m = k[i-1] - 10
					k[i-1] = m
					if i-2 >= 0 {
						k[i-2]++
					}
				} else if i-1 == 0 {
					break
				}

			}
		}
	}

	for i := 0; i < len(k); i++ {
		answer += converter.Int_str(k[i])
	}

	return answer
}
