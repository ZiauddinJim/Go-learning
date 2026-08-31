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
	// outputFunc()
	// generalFormattingVerbs()
	// integerFormattingVerbs()
	// stringFormattingVerbs()
	// boolFormattingVerbs()
	// floatFormattingVerbs()
	// dataTypes()
	// Note: array.go file function
	Array()
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
func generalFormattingVerbs() {
	value := 42.5
	fmt.Printf("Default: %v\n", value)
	fmt.Printf("Go-syntax: %#v\n", value)
	fmt.Printf("Type: %T\n", value)
	fmt.Printf("Percentage: %v%%\n", value)
}
func integerFormattingVerbs() {
	i := 20
	fmt.Printf("Binary: %b\n", i)
	fmt.Printf("Decimal: %d\n", i)
	fmt.Printf("sign & Decimal: %+d\n", i)
	fmt.Printf("Octal: %o\n", i)
	fmt.Printf("prefix Octal: %O\n", i)
	fmt.Printf("Hexadecimal(x OR X): %x\n", i)
	fmt.Printf("prefix Hexadecimal: %#x\n", i)
	fmt.Printf("ped with space(justified right): %4d\n", i)
	fmt.Printf("ped with space(justified lef): %-4d\n", i)
	fmt.Printf("ped with zeros: %04d\n", i)
}
func stringFormattingVerbs() {
	text := "Hello"
	fmt.Printf("Plain text: %s\n", text)
	fmt.Printf("Double quotation: %q\n", text)
	fmt.Printf("plain string(w-8 justify right): %8s\n", text)
	fmt.Printf("plain string(w-8 justify left): %-8s\n", text)
	fmt.Printf("hex dump of bytes value: %x\n", text)
	fmt.Printf("hex dump of bytes value space: % x\n", text)
}
func boolFormattingVerbs() {
	boolean := true
	fmt.Printf("Bool verify: %t\n", boolean)
}
func floatFormattingVerbs() {
	i := 3.14159265
	fmt.Printf("Default: %f\n", i)
	fmt.Printf("Precision 2: %.2f\n", i)
	fmt.Printf("Width-6 Precision 2: %6.2f\n", i)
	fmt.Printf("Scientific: %e\n", i)
	fmt.Printf("necessary digits: %g\n", i)
}
func dataTypes() {
	var (
		// Boolean
		a bool = true
		// Syntax: Numeric data type. 3 categories
		// Integer
		b int = 5
		// Float
		c float32 = 5.5
		// Complex types
		d complex64 = 3 + 4i
		// String
		e string = "Hello"
	)
	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float: ", c)
	fmt.Println("Complex: ", d)
	fmt.Println("string: ", e)
}
