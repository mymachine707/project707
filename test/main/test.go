package test

import "fmt"

func main() {
	var answer string = Minus()
	fmt.Println(answer)
}

func Minus(a, b []string) string {
	var z []string

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

	var answers string
	for i := 0; i < len(z); i++ {
		answers += z[i]
	}

	//==========================
	return answers

}
