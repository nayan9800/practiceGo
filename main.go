package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/nayan9800/practiceGo/pkg/handerrors"
	//to import the local package use module name with path of package
	//<module name>/<path of package folder>
)

func main() {
	fmt.Println(color.GreenString("Practice Go"))

	/*To run basic exmples*/
	//basics.Basics()

	//package and modules
	/*fmt.Println(arith.Sum(1, 2, 3, 4, 5, 6))
	fmt.Println(arith.Add(4, 5))
	fmt.Println(arith.Abs(-4))*/
	//fmt.Println(color.BlueString("Hi this is from remote package go get github.com/fatih/color"))

	/*handling errors*/
	handerrors.RunHandleError()
}
