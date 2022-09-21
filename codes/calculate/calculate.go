package main

import (
	"fmt"
	"project/add"
	"project/clean"
	"project/isnumber"
	"project/multiply"
	"project/sign"
	"project/sub"
	"strings"
)

func main() {

	fmt.Println("Sonlarni string holatida qo'shish")
	var num1 string
	var num2 string
	var signs string

	fmt.Printf("%s", "\tEntered number 1: ")
	fmt.Scanf("%s", &num1)

	fmt.Printf("%s", "\tEntered number 2: ")
	fmt.Scanf("%s", &num2)

	fmt.Printf("%s", "\t\nEntered sign to calculate: ")
	fmt.Scanf("%s", &signs)

	// sliceing is number
	a := strings.Split(num1, "")
	b := strings.Split(num2, "")

	// step-1: validation is sign number  == negative (true)  // positive (false)
	var a1 bool // if number one is negative a1 is true else a1 is false
	var b2 bool // if number two is negative b2 is true else b2 is false
	a1 = sign.Signnumbers(a)
	b2 = sign.Signnumbers(b)
	// clean up sign !!!!!
	num1 = clean.Clean(num1)
	num2 = clean.Clean(num2)

	//=====//step-1-2: validation is number true else panic
	isnumber.Isnumber(a)
	isnumber.Isnumber(b)

	// step-2: validation is sign canculate // Bajarish amalini tekshiriladi +  -  *  / is sign true else panic
	sign.SignValidation(signs)

	var answerFunction string // function answer

	//=====//step-2-1: sign numbers
	switch signs {
	case "+":
		fmt.Println("+")
		// step-3: yig'indini hisoblash
		//=====//step-3-1: a+b=a+b
		if !a1 && !b2 {
			// a+b=a+b
			// a > 0  and b > 0
			answerFunction = add.Add(num1, num2)
		} else if !a1 && b2 {
			//=====//step-3-2: a+(-b)=a-b  // kattasini aniqlab keyin ayiriladi
			// a > 0  and b < 0
			if sign.Findtobig(num1, num2) == 1 { // a>b  == a-b
				answerFunction = sub.Minus(num1, num2)
			} else if sign.Findtobig(num1, num2) == 2 { // a<b == b-a  // answer is negative
				answerFunction = sign.AddtoMinus(sub.Minus(num1, num2)) // add to minus for answer
			}
		} else if a1 && !b2 {
			// a < 0  and b > 0
			//=====//step-3-3: -a+b=b-a  // kattasini aniqlab keyin ayiriladi
			if sign.Findtobig(num1, num2) == 1 { // a>b  == a-b = -c
				answerFunction = sign.AddtoMinus(sub.Minus(num1, num2)) // add to minus for answer
			} else if sign.Findtobig(num1, num2) == 2 { // a<b == b-a
				answerFunction = sub.Minus(num1, num2)
			}
		} else if a1 && b2 {
			//=====//step-3-4: -a+(-b)=-(a+b)
			answerFunction = sign.AddtoMinus(add.Add(num1, num2))
		}

	case "-":
		fmt.Println("-")
		// step-4: ayirmani hisoblash
		//=====//step-4-1: a-b=a-b // kattasini aniqlab keyin ayiriladi
		if !a1 && !b2 {
			if sign.Findtobig(num1, num2) == 1 { // a>b
				answerFunction = sub.Minus(num1, num2)
			} else if sign.Findtobig(num1, num2) == 2 { //b>a
				answerFunction = sign.AddtoMinus(sub.Minus(num1, num2))
			}
		} else if !a1 && b2 {
			//=====//step-4-2: a-(-b)=a+b
			answerFunction = add.Add(num1, num2)
		} else if a1 && !b2 {
			//=====//step-4-3: -a-b=-(a+b)
			answerFunction = sign.AddtoMinus(add.Add(num1, num2))
		} else if a1 && b2 {
			//=====//step-4-4: -a-(-b)=b-a // kattasini aniqlab keyin ayiriladi
			if sign.Findtobig(num1, num2) == 2 { // a>b
				answerFunction = sub.Minus(num1, num2)
			} else if sign.Findtobig(num1, num2) == 1 { //b>a
				answerFunction = sign.AddtoMinus(sub.Minus(num1, num2))
			}
		}
	case "*":
		fmt.Println("*")
		// step-5: ko'paytmani hisoblash
		if (!a1 && !b2) || (a1 && b2) {
			//=====//step-5-1: a*b=a*b // kattasini aniqlab keyin ayiriladi
			//=====//step-5-3: (-a)*(-b)=a*b
			answerFunction = multiply.Multiply(num1, num2)
		} else if (a1 && !b2) || (!a1 && b2) {
			//=====//step-5-2: a*(-b)=-(a*b) \\ // -a*b=-(a*b)
			answerFunction = sign.AddtoMinus(multiply.Multiply(num1, num2))
		}
	case "/":
		fmt.Println("/")
		// step-6: bo'linmani hisoblash
	case "%":
		fmt.Println("%")
		// step-7: qoldiqni hisoblash
	default:
		panic("This sign is not defined")
	}

	fmt.Println(answerFunction)
}

// Berilgan raqamlarni qo'shish
//p := add.Add(num1, num2)   // qo'shish
//q := sub.Minus(num1, num2) // ayirish
//s := multiply.Multiply(num1, num2) // ko'paytirish
//s := clean.Clean(num1) // tozalash
//s := clean.Clean(num1) // tozalash
//m := modul.Set(num1) //modul
