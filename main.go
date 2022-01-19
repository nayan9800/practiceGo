package main

import (
	"fmt"

	"github.com/fatih/color"

	//to import the local package use module name with path of package
	//<module name>/<path of package folder>
	"github.com/nayan9800/practiceGo/pkg/arith"
)

func main() {
	fmt.Println("Practice go")
	/*To run basic exmples*/
	//basics.Basics()
	fmt.Println(arith.Sum(1, 2, 3, 4, 5, 6))
	fmt.Println(arith.Add(4, 5))
	fmt.Println(arith.Abs(-4))
	fmt.Println(color.BlueString("Hi this is from remote package go get github.com/fatih/color"))
}
