// package main

// import "fmt"

// func main() {
// 	// Note: Many way to variable calling
// 	// New:  var name string = "zia"
// 	// New: name := "zia"
// 	/* New:
// 	var name string
// 	name = "Zia"
// 	*/

// 	// New: Group variable declaration
// 	var (
// 		name string = "zia"
// 		age  int    = 25
// 	)
// 	// New: Multiple value declaration
// 	var x, y int
// 	x = 10
// 	y = 20

// 	fmt.Println(x, y)
// 	fmt.Println(name)
// 	fmt.Println(age)
// }

package main

import "fmt"

func main() {

	// variable()
	// initialValue()
	// MultipleValue()
	// BlockDeclaration()
	// Constants()
	// multiConst()
	outputFunc()
}
func syntax() {
	// Syntax: with the var keyword
	// var variableName type = value
	var name string = "jim"

	// Syntax: with the := sign
	// variableName := value
	age := 25

	// Syntax:

	fmt.Println("Hello Jim")
	fmt.Println(name, age)
}
func variable() {
	var Animal1 string = "cat"
	var Animal2 = "Fish" //type is inferred
	count := 2

	fmt.Println(Animal1, Animal2, count)
}
func initialValue() {
	var a string
	var b int
	var c bool
	a = "jim"
	b = 5
	c = true
	fmt.Println(a, b, c)
}
func MultipleValue() {
	var a, b, c int = 1, 2, 3
	var d, e = 3, "jim"

	fmt.Println(a, b, c)
	fmt.Println(d, e)
}
func BlockDeclaration() {
	var (
		a int
		b int    = 1
		c string = "mim"
	)
	fmt.Println(a, b, c)
}
func Constants() {
	// Syntax: const ConstName type = value
	// Syntax: const constName = value

	const A string = "jim"
	const a = "mim"
	fmt.Println(A, a)
}
func multiConst() {
	const (
		a        = 1
		b string = "cat"
		c        = "mog"
	)
	fmt.Println(a, b, c)
}
func outputFunc() {
	var i, j string = "jim", "mim"

	// Syntax: fmt.Print() not spacing in the value and new line not generate
	fmt.Print(i, j)
	fmt.Print("\n", i, " ", j, "\n")

	// Syntax: fmt.Println() use per value spacing and new line generate
	fmt.Println(i, j)

	// Syntax: fmt.Printf() use formatting verbs
	var (
		name  = "jim"
		score = 5
	)
	fmt.Printf("Name: %v, Type: %T\n", name, name)
	fmt.Printf("Score: %v, Type: %T\n", score, score)
}
