package main

import "fmt"

func functionA() {
	add(20, 30)
	sum, mul := getNumber(20, 45)
	fmt.Println(sum)
	fmt.Println(mul)
}

// Note: Normal function
func add(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println(sum)
}

// Note: Function with return type
func getNumber(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}
// Note: Function with return type and named return value
func getNumber1(num1 int, num2 int) (sum int, mul int) {
	sum = num1 + num2
	mul = num1 * num2
	return
}
