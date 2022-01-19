package basics

import (
	"fmt"
)

/*Structs*/
//a struct in go is collection of fields
type point struct {
	X int
	Y int
}

func moreTypes() {
	fmt.Println("More types in go")
	fmt.Println(point{X: 2, Y: 8})

	p := point{X: 1, Y: 2}
	//Structs fields can be accessed using .
	fmt.Println(p.X)

	//struct field can also be accessed from struct pointer
	//without the explicit dereference.
	v := &p
	v.X = 15
	fmt.Println(p.X)

	/*Arrays*/
	// here nums is array of int with size 6
	//array can not be resized
	nums := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums)
	fmt.Println(nums[0])

	/*Slices*/
	// slices in go are dynamically sized array
	//here the numSlice is slice of int
	numSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numSlice)

	//slice with the struct
	points := []point{
		{0, 0},
		{1, 0},
		{0, 1},
		{1, 1},
		{0, -1},
		{-1, 0},
	}
	fmt.Println(points)

	//len() and cap() function in slices
	fmt.Printf("len:=%d cap:=%d\n", len(numSlice), cap(numSlice))

	//slice without is any values have nil
	var empty []string
	if empty == nil {
		fmt.Printf("nil slice len:=%d cap:=%d\n", len(empty), cap(empty))
	}

	//To build slices dynamatically make() function is used
	//make function takes argumrnts as type ,length of slice
	msg := make([]string, 5)
	fmt.Printf("slice len:=%d cap:=%d\n", len(msg), cap(msg))
	//to specify the third argument is passed
	msg = make([]string, 0, 5)
	fmt.Printf("slice len:=%d cap:=%d\n", len(msg), cap(msg))

	/*slices of slices*/
	box := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	fmt.Println(box[0][1])
	fmt.Println(box[1][1])

	/*appending a slice*/

	//to append the elments in slice append() function is used
	//append function takes slice as first argument and element
	// to add in second argument
	fmt.Printf("box len:=%d cap:=%d\n", len(box), cap(box))
	box = append(box, []int{10, 11, 12})
	fmt.Printf("box len:=%d cap:=%d\n", len(box), cap(box))

	/*Range*/

	//range is used in for loop with map and slice
	//to range over the elements
	//range gives index and value in slice
	fmt.Println("range over the box")
	for i, v1 := range box {
		for j, v2 := range v1 {
			fmt.Printf("\tbox[%d][%d]=%d", i, j, v2)
		}
		fmt.Println()
	}

	/*Maps*/
	// maps are map keys with value
	phonebook := make(map[string]int64)
	phonebook["foo"] = 12545121
	phonebook["alice"] = 789354542
	phonebook["bob"] = 7522221484
	fmt.Println("bob's phone number = ", phonebook["bob"])

	//operatons on map

	//Get value with key
	val := phonebook["foo"]
	fmt.Println(val)

	//get value and check if element is present
	val, ok := phonebook["bob"]
	fmt.Printf("value:=%d, is present %t\n", val, ok)
	val, ok = phonebook["john"]
	fmt.Printf("value:=%d, is present %t\n", val, ok)

	//insert or update value
	phonebook["john"] = 178761352

	//delete value
	delete(phonebook, "foo")

	//range over map
	//using range over map in for loop gives each key
	//value pair in
	for k, v := range phonebook {
		fmt.Printf("name:= %s number:= %d\n", k, v)
	}

	/*Function as values or anonymous functions*/
	//function in go are also values they can be passed
	//around like variables

	sayHello := func(name string) string {

		return fmt.Sprintf("Hello %s", name)
	}

	fmt.Println(sayHello("Alice"))

	//function clousres
	//here the function in sayHellowithPrefix can access the
	//prefix varaible
	prefix := "From go"
	sayHelloWithPrefix := func(name string) string {

		return fmt.Sprintf("%s Hello %s", prefix, name)
	}
	fmt.Println(sayHelloWithPrefix("bob"))
}
