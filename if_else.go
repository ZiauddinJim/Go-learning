package main

import "fmt"

func ifElse() {
	// age := 20

	// if age > 18 {
	// 	fmt.Println("you are eligible to be married")
	// } else if age < 18 {
	// 	fmt.Println("you are not eligible to be married, but you can love someone")
	// } else {
	// 	fmt.Println("you are just a tenager, not eligible to be married")
	// }
	// male()
	switchCase()
}

func male() {
	age := 20
	sex := "male"

	if age > 60 || sex == "male" {
		fmt.Println("married")
	}
}

func switchCase() {
	a := 20
	switch a {
	case 10:
		fmt.Println("a is 10")
	default:
		fmt.Println("a is not 10")
	}
}
