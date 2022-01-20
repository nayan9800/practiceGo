package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/nayan9800/practiceGo/pkg/ds"
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
	//handerrors.RunHandleError()

	list := ds.NewLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(5)
	list.Traverse()
	fmt.Println(list.Search(5))
	fmt.Println(list.Delete(2))
	list.Traverse()

	t := ds.NewBTree()
	t.Add(5)
	t.Add(9)
	t.Add(1)
	t.Display()
}
