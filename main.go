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
}
