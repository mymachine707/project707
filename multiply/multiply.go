package multiply

import (
	"project/add"
	"project/converter"
	"strings"
)

// converter.Int_str
// converter.Str_Int

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
