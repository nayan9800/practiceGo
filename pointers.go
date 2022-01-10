package main

import "fmt"

func pointers() {
	fmt.Println("Hello Pointers in go")

	//pointers for type T will be denoted as *T
	var no int = 45 //no is int type
	var i *int      // i is *int type which hold pointer of int type
	i = &no         // assigen addres of no to i

	fmt.Println(i)  //print address
	fmt.Println(*i) //after dereferencing print value

}
