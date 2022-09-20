package main

import (
	"fmt"
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

	fmt.Printf("%s", "\tEntered sign to calculate: ")
	fmt.Scanf("%s", &signs)

	// step-1: validation

	// step-2: validation sign
	//fmt.Println(sign.SignValidation(signs)) ishlavotti
	//=====//step-2-1: sign numbers
	switch signs {
	case "+":
		fmt.Println("+")
		// step-3: yig'indini hisoblash
		//=====//step-3-1: a+b=a+b
		//=====//step-3-2: a+(-b)=a-b  // kattasini aniqlab keyin ayiriladi
		//=====//step-3-3: -a+b=b-a  // kattasini aniqlab keyin ayiriladi
		//=====//step-3-4: -a+(-b)=-(a+b)
	case "-":
		fmt.Println("-")
		// step-4: ayirmani hisoblash
		//=====//step-4-1: a-b=a-b // kattasini aniqlab keyin ayiriladi
		//=====//step-4-2: a-(-b)=a+b
		//=====//step-4-3: -a-b=-(a+b)
		//=====//step-4-4: -a-(-b)=-b-a // kattasini aniqlab keyin ayiriladi
	case "*":
		fmt.Println("*")
		// step-5: ko'paytmani hisoblash
		//=====//step-5-1: a*b=a*b // kattasini aniqlab keyin ayiriladi
		//=====//step-5-2: a*(-b)=-(a*b)
		//=====//step-5-3: (-a)*(-b)=a*b
	case "/":
		fmt.Println("/")
		// step-6: bo'linmani hisoblash
	case "%":
		fmt.Println("%")
		// step-7: qoldiqni hisoblash
	}

}
