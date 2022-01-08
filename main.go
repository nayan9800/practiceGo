package main

import (
	"fmt"
)

/*1. Functions*/

/*a. Function in go can take one or many arguments*/
func add(x int, y int) int {

	return x + y
}
func square(x int) int {

	return x * x
}

/*function parameters can noted type in the last
if they share the same type*/
func sub(a, b int) int {
	return a - b
}

/*func can retrun many values*/
func secToHrMinSec(secs int) (int, int, int) {
	var hr, min, sec int
	hr = secs / 3600
	min = sec % 3600
	sec = secs - ((hr * 3600) + min*60)
	return hr, min, sec
}

/*names return values are used here
a,b are the names return values
which can be retrun by using naked
return keyword*/
func swap(x, y int) (a, b int) {
	a = y
	b = x
	return
}
func main() {

	fmt.Println("Hello world")
	fmt.Println(add(1, 2))
	fmt.Println(square(2))
	fmt.Println(sub(10, 1))
	fmt.Println(secToHrMinSec(3600))
	fmt.Println(swap(1, 2))

	/*2. Variables and types*/
	/*Variables can be intialized using var keyword*/
	var c, python, java bool
	fmt.Println(c, python, java)
	var hello string = "string"
	fmt.Println(hello)

	/*short variable declarations*/
	a := 0
	name := "hello"
	fmt.Println(a, name)

	/*Data types*/

	//uint type have 8,16,32,64 variant
	//default value is 0
	var number uint32
	number = 1<<32 - 1
	fmt.Printf("%d has type %T\n", number, number)

	//int type have 8,16,32,64 variant
	//default value is 0
	var number2 int32
	number2 = 1554
	fmt.Printf("%d has type %T\n", number2, number2)

	//float
	//has variant 32 and 64
	//default value is 0.0
	var f1 float64
	f1 = 64.656442
	fmt.Printf("%f has type %T\n", f1, f1)

	//complex
	//complex is data type  used for complex number operations
	//complex data type has two varients complex128 and complex64
	var comp1 complex64
	comp1 = complex64(1 + 4i)
	fmt.Printf("%v has type %T\n", comp1, comp1)

	//string
	//default value is ""
	msg := "Hello my name is gopher"
	fmt.Printf("%s has type %T\n", msg, msg)

	//Rune
	// Rune is datatype which is alias of uint32
	//is used to store unicode values
	//smiles is silce of the rune
	smiles := []rune{'😃', '😎', '😈'}
	/*here for loop is used to range over the slice of rune*/
	for _, s := range smiles {
		fmt.Printf("character:=%c, Unicode:=%U\n", s, s)
	}

	//constants
	//constants are declared with const keyword
	const pi = 3.14
	fmt.Println(pi)

	//Type Casting or type conversion
	//in go explict conversion is needed to convert data type
	var i int = 45
	fmt.Printf("%d has type %T\n", i, i)
	f2 := float32(i)
	fmt.Printf("%f has type %T\n", f2, f2)
	ui := uint(i)
	fmt.Printf("%d has type %T\n", ui, ui)

	//control and loops
	controlAndLoops()
}
